package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaopandolfi/hydra-back/controllers"
	"github.com/joaopandolfi/hydra-back/mhandlers"
)

func rest(r *mux.Router) {
	//Common
	r.HandleFunc("/rest/login", controllers.AuthController{}.Login).Methods("POST")
	r.HandleFunc("/rest/logout", mhandlers.AuthProtection(controllers.AuthController{}.Logout)).Methods("POST", "GET")
	r.HandleFunc("/rest/check/auth", controllers.AuthController{}.CheckAuth).Methods("GET")

	//User
	r.HandleFunc("/rest/user/new", mhandlers.AuthProtection(controllers.UserController{}.NewClientUser)).Methods("POST")
	//r.HandleFunc("/rest/user/new/t", controllers.UserController{}.NewClientUser).Methods("POST")

	//Generic
	r.HandleFunc("/lambda/new", mhandlers.AuthTokenedProtection(controllers.LambdaController{}.Save)).Methods("POST")
	r.HandleFunc("/lambda/tag/new", mhandlers.AuthTokenedProtection(controllers.LambdaController{}.SaveWithTag)).Methods("POST")
	r.HandleFunc("/lambda/get/{id}", mhandlers.AuthTokenedProtection(controllers.LambdaController{}.GetByID)).Methods("GET")

}
