package routes

import (
	"../controllers"
	"../mhandlers"
	"github.com/gorilla/mux"
)

func itu(r *mux.Router) {
	r.HandleFunc("/itu/new", mhandlers.AuthTokenedProtection(controllers.ITUController{}.New)).Methods("POST")
	r.HandleFunc("/itu/get/{id}", mhandlers.AuthTokenedProtection(controllers.ITUController{}.GetByID)).Methods("GET")
	r.HandleFunc("/itu/list", mhandlers.AuthTokenedProtection(controllers.ITUController{}.List)).Methods("GET")
	r.HandleFunc("/itu/set/patient", mhandlers.AuthTokenedProtection(controllers.ITUController{}.SetPatient)).Methods("GET")
}
