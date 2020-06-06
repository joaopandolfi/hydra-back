package routes

import (
	"../controllers"
	"../mhandlers"
	"github.com/gorilla/mux"
)

func patient(r *mux.Router) {
	r.HandleFunc("/patient/new", mhandlers.AuthTokenedProtection(controllers.PatientController{}.New)).Methods("POST")
	r.HandleFunc("/patient/get/{id}", mhandlers.AuthTokenedProtection(controllers.PatientController{}.GetByID)).Methods("GET")
}
