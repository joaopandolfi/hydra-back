package controllers

import (
	"net/http"

	"../config"
	"../models"
	"github.com/joaopandolfi/blackwhale/configurations"
	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/remotes/mongo"
	"github.com/joaopandolfi/blackwhale/utils"
)

// --- Health ---

// BaseController -
type HealthController struct{}

// Health route
func (cc HealthController) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	handlers.Response(w, true)
}

// Config database
func (cc HealthController) Config(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	hash := handlers.GetHeader(r, "hash")
	utils.Debug("HEADER", hash, configurations.Configuration.ResetHash)
	if hash != configurations.Configuration.ResetHash {
		handlers.Response(w, "No cookies for you")
		return
	}
	configure()
	handlers.Response(w, true)
}

// ResetDatabase route
func (cc HealthController) ResetDatabase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	hash := handlers.GetHeader(r, "hash")
	if hash != configurations.Configuration.ResetHash {
		handlers.Response(w, "Invalid Hash")
		return
	}

	var errors []error
	_, err := mongo.GetSession().GetCollection("traffic").RemoveAll(nil)
	if err != nil {
		errors = append(errors, err)
	}

	//Config
	//configure()

	if len(errors) > 0 {
		handlers.Response(w, errors)
	} else {
		utils.Info("[ResetService] - Traffic-files database RESETED")
		handlers.Response(w, "Reseted")
	}

}

func configure() error {
	userService := NewUserService()
	_, err := userService.NewUser(models.User{
		People: models.People{
			Name: "Heil Hydra",
			CPF:  "00000",
		},
		Email:     "",
		Username:  "hydrante",
		Picture:   "",
		Password:  config.Config.DefaultPassword, // piririco
		Instution: 0,
		Level:     99,
	})
	return err
}
