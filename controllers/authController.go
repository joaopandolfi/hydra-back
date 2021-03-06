package controllers

import (
	"fmt"
	"net/http"

	"github.com/segmentio/encoding/json"

	"github.com/joaopandolfi/hydra-back/models"
	"github.com/joaopandolfi/hydra-back/services"

	"github.com/flosch/pongo2"
	"github.com/joaopandolfi/blackwhale/configurations"
	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/utils"
)

// AuthController -
type AuthController struct{}

// PreCompile all pages used in this controller
// @method
func (cc AuthController) PreCompile() {
	configurations.Configuration.Templates["login"] = utils.GetHbsPage("login.hbs")
}

// ----- PAGE -----

// LoginPage - Render login page
// @page
func (cc AuthController) LoginPage(w http.ResponseWriter, r *http.Request) {
	data := pongo2.Context{"test": "test"}
	err := configurations.Configuration.Templates["login"].ExecuteWriter(data, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ------ REST ------

// Login - endpoint ro receive `username` and `password` to check login
// @rest
// If login is successfull, the data will be stored in `session`
// Session keys: `logged`, `username`, `institution`, `level`, `token`
func (cc AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var userService = services.NewUser()
	var received map[string]string
	err := json.NewDecoder(r.Body).Decode(&received)

	var msg string
	var result map[string]interface{}

	if err != nil {
		utils.CriticalError("[AuthController][REST][Login]- Error on unmarshal body", err.Error())
		handlers.ResponseError(w, "Error on unmarshal body")
		return
	}

	if received["username"] != "" && received["password"] != "" && received["institution"] != "" {
		user, success, err := userService.Login(received["username"], received["password"], received["institution"])
		if err != nil {
			success = false
		}
		if success {
			var sess, err = handlers.GetSession(r)
			if err != nil {
				utils.CriticalError("Error on get session", err.Error())
				handlers.RESTResponseError(w, "Session Error")
				return
			}

			token, err := utils.NewJwtToken(utils.Token{
				ID:          fmt.Sprint(user.ID),
				Institution: fmt.Sprint(user.Instution),
				Permission:  fmt.Sprint(user.Level),
			}, configurations.Configuration.Security.TokenValidity)
			if err != nil {
				token = user.Token
			}

			sess.Values[models.SESSION_VALUE_ID] = user.ID
			sess.Values[models.SESSION_VALUE_NAME] = user.Name
			sess.Values[models.SESSION_VALUE_LEVEL] = user.Level
			sess.Values[models.SESSION_VALUE_TOKEN] = token
			sess.Values[models.SESSION_VALUE_LOGGED] = true
			sess.Values[models.SESSION_VALUE_USERNAME] = user.Username
			sess.Values[models.SESSION_VALUE_INSTITUTION] = user.Instution
			sess.Options = configurations.Configuration.Session.Options
			err = sess.Save(r, w)
			utils.Debug("[AuthController][REST][Login]-Authenticated", sess.Values[models.SESSION_VALUE_USERNAME], err, sess.Values[models.SESSION_VALUE_LOGGED])

			result = make(map[string]interface{})

			result["id"] = user.ID
			result["token"] = token
			result["success"] = true
			result["permission"] = user.Level
			result["institution"] = user.Instution

			handlers.Response(w, result)
			return
		}
		msg = "Invalid credentials"
	} else {
		utils.Error("Wrong parameters ->", received)
		msg = "Wrong parameters"
	}

	handlers.RESTResponseError(w, msg)
}

//Logout Destroy session and logout
// @rest
func (cc AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	var sess, _ = handlers.GetSession(r)
	sess.Options.MaxAge = -1
	err := sess.Save(r, w)

	if err != nil {
		utils.CriticalError("[AuthController][Logout] - Error on destroy session", err.Error())
		handlers.RESTResponseError(w, "Error on logout")
		return
	}
	utils.Debug("[PostLogin]-Authenticated", sess.Values["username"], err, sess.Values["logged"])
	handlers.Response(w, map[string]bool{"success": true})
}

// CheckAuth - acess dabatase and validate token
// TODO: Acessar o banco de dados
// Check authentication
func (cc AuthController) CheckAuth(w http.ResponseWriter, r *http.Request) {
	var sess, _ = handlers.GetSession(r)
	token := handlers.GetHeader(r, "token")
	q := handlers.GetQueryes(r)
	bearer := q.Get("bearer")

	if token != "" || bearer != "" {
		token = token + bearer
		t, err := utils.CheckJwtToken(token)
		if err != nil || !t.Authorized {
			handlers.Response(w, map[string]interface{}{"logged": false, "message": "Invalid Token"})

		} else {
			handlers.Response(w, map[string]interface{}{"logged": true, "data": t})
		}
		return
	}

	if sess == nil || sess.Values == nil || sess.Values["logged"] == nil {
		//ntoken, _ := utils.NewJwtToken("teste", 10)
		ntoken := ""
		handlers.Response(w, map[string]interface{}{"logged": false, "token": ntoken})
		return
	}
	handlers.Response(w, map[string]bool{"logged": sess.Values["logged"].(bool)})
}
