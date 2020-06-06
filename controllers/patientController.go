package controllers

import (
	"net/http"

	"github.com/joaopandolfi/blackwhale/handlers"
)

type PatientController struct {
}

func (c PatientController) New(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}

func (c PatientController) GetByID(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}
