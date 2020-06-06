package controllers

import (
	"net/http"

	"github.com/joaopandolfi/blackwhale/handlers"
)

type HospitalController struct {
}

func (c HospitalController) New(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}

func (c HospitalController) GetByID(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}

func (c HospitalController) List(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}

func (c HospitalController) ListPatients(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}

func (c HospitalController) ListITU(w http.ResponseWriter, r *http.Request) {
	handlers.RESTResponse(w, "")
}
