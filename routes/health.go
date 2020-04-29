package routes

import (
	"github.com/gorilla/mux"
	"../controllers"
)


func health(r *mux.Router) {
	// Health
	r.HandleFunc("/health", controllers.HealthController{}.Health).Methods("POST", "GET", "HEAD")
	r.HandleFunc("/config", controllers.HealthController{}.Config).Methods("POST", "GET", "HEAD")
	r.HandleFunc("/reset", controllers.HealthController{}.ResetDatabase).Methods("POST", "GET", "HEAD")
}
