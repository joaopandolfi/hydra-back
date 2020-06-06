package routes

import (
	"../controllers"
	"../mhandlers"
	"github.com/gorilla/mux"
)

func hospital(r *mux.Router) {
	r.HandleFunc("/hospital/new", mhandlers.AuthTokenedProtection(controllers.HospitalController{}.New)).Methods("POST")
	r.HandleFunc("/hospital/get/{id}", mhandlers.AuthTokenedProtection(controllers.HospitalController{}.GetByID)).Methods("GET")
	r.HandleFunc("/hospital/list", mhandlers.AuthTokenedProtection(controllers.HospitalController{}.List)).Methods("GET")
	r.HandleFunc("/hospital/list/patients", mhandlers.AuthTokenedProtection(controllers.HospitalController{}.ListPatients)).Methods("GET")
	r.HandleFunc("/hospital/list/itu", mhandlers.AuthTokenedProtection(controllers.HospitalController{}.ListITU)).Methods("GET")
}
