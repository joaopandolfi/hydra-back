package controllers

import (
	"net/http"

	"github.com/joaopandolfi/blackwhale/handlers"
)

type ITUController struct {
}

func (c ITUController) New(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}

func (c ITUController) GetByID(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}

func (c ITUController) List(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}

func (c ITUController) SetPatient(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}
