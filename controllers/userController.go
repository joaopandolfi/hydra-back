package controllers

import (
	"net/http"
	"strconv"

	"../models"
	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/utils"
	"github.com/segmentio/encoding/json"
)

// UserController -
type UserController struct {
}

// NewClientUser - This endpoint create a new user, based on institution
// @rest
func (cc UserController) NewClientUser(w http.ResponseWriter, r *http.Request) {
	var received map[string]string
	userService := NewUserService()
	err := json.NewDecoder(r.Body).Decode(&received)
	if err != nil {
		utils.Error("Erron on get body", err.Error())
		handlers.RESTResponseError(w, "Invalid body "+err.Error())
		return
	}

	instInt, _ := strconv.Atoi(received["institution"])
	level, _ := strconv.Atoi(received["level"])
	cpf := utils.OnlyNumbers(received["cpf"])
	user := models.User{
		People: models.People{
			Name: received["name"],
			CPF:  cpf,
		},
		Email:     received["email"],
		Level:     level,
		Username:  received["username"],
		Picture:   "",
		Password:  received["password"],
		Instution: instInt,
	}

	//result, err := userService.NewUserClient(user)
	result, err := userService.NewUser(user)
	if err != nil {
		utils.Debug("Error on create new user", err.Error())
		handlers.RESTResponseError(w, "Error on create new user")
	} else {
		handlers.RESTResponse(w, result)
	}
}

// SetEspecialty Set in session the speciality is in the question
// @rest
// Store in session with çabel `specialty`
func (cc UserController) SetEspecialty(w http.ResponseWriter, r *http.Request) {
	sess, _ := handlers.GetSession(r)
	vars := handlers.GetVars(r)

	specialty, _ := strconv.Atoi(vars["specialty"])

	if specialty > 0 {
		sess.Values[models.SESSION_VALUE_SPECIALTY] = specialty
		err := sess.Save(r, w)
		if err == nil {
			handlers.RESTResponse(w, true)
			return
		}
	}

	handlers.RESTResponseError(w, "Invalid Specialty")
}
