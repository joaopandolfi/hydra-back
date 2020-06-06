package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaopandolfi/hydra-back/controllers"
	"github.com/joaopandolfi/hydra-back/mhandlers"
)

func patient(r *mux.Router) {
	r.HandleFunc("/patient/new", mhandlers.AuthTokenedProtection(controllers.PatientController{}.New)).Methods("POST")
	r.HandleFunc("/patient/get/{id}", mhandlers.AuthTokenedProtection(controllers.PatientController{}.GetByID)).Methods("GET")
}
