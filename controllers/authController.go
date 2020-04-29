package controllers

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	"../models"

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
	var userService = NewUserService()
	var sess, _ = handlers.GetSession(r)
	var received map[string]string
	err := json.NewDecoder(r.Body).Decode(&received)

	var msg string
	var result map[string]interface{}

	if err != nil {
		utils.CriticalError("[AuthController][REST][Login]- Error on unmarshal body", err.Error())
		handlers.ResponseError(w, "Error on unmarshal body")
		return
	}

	if received["username"] != "" && received["password"] != "" {

		user, success, err := userService.Login(received["username"], received["password"])
		//( == "joao" &&  == "202cb962ac59075b964b07152d234b70")
		if success {
			sess.Values[models.SESSION_VALUE_LOGGED] = true
			sess.Values[models.SESSION_VALUE_SPECIALTY] = 1 // Default - oftalmo
			sess.Values[models.SESSION_VALUE_USERNAME] = user.Username
			sess.Values[models.SESSION_VALUE_NAME] = user.Name
			sess.Values[models.SESSION_VALUE_INSTITUTION] = user.Instution
			sess.Values[models.SESSION_VALUE_LEVEL] = user.Level
			sess.Values[models.SESSION_VALUE_TOKEN] = user.Token
			sess.Values[models.SESSION_VALUE_ID] = user.ID
			sess.Options = configurations.Configuration.Session.Options
			err = sess.Save(r, w)
			utils.Debug("[AuthController][REST][Login]-Authenticated", sess.Values[models.SESSION_VALUE_USERNAME], err, sess.Values[models.SESSION_VALUE_LOGGED])

			result = make(map[string]interface{})

			result["success"] = true
			result["token"] = user.Token
			result["institution"] = user.Instution
			result["id"] = user.ID
			result["permission"] = user.Level

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
	if sess == nil || sess.Values == nil || sess.Values["logged"] == nil {
		handlers.Response(w, map[string]bool{"logged": false})
		return
	}

	// CHECAR USANDO TOKEN -> CRIADO NA TABELA

	handlers.Response(w, map[string]bool{"logged": sess.Values["logged"].(bool)})
}
