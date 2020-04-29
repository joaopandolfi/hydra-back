package mhandlers

import (
	"fmt"
	"net/http"
	"strconv"

	"../dao"
	"../models"
	"../services"
	"github.com/joaopandolfi/blackwhale/configurations"
	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/utils"
)

var uservice = services.User{
	UserDAO: dao.User{},
}

// TokenHandler -
// @handler
// Intercept all transactions and check if is authenticated by token
func TokenHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sess, _ = handlers.GetSession(r)
		url := r.URL.String()

		if sess != nil {
			l := sess.Values[models.SESSION_VALUE_LOGGED]
			if logged, ok := l.(bool); ok && logged {
				utils.Debug("[TokenHandler]", "Authenticated by session", url)
				next.ServeHTTP(w, r)
				return
			}
		}

		token := handlers.GetHeader(r, "token")
		userIDs := handlers.GetHeader(r, "id")
		userID, _ := strconv.Atoi(userIDs)

		s, err := uservice.CheckToken(userID, token)

		if !s || err != nil {
			utils.Debug("[TokenHandler]", "Auth Error", url)
			handlers.Redirect(r, w, "/login")
			return
		}

		utils.Debug("[TokenHandler]", "Authenticated", url)
		next.ServeHTTP(w, r)
	})
}

// LoggedHandler -
// @handler
// Intercept all transactions and check if is authenticated
func LoggedHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sess, _ = handlers.GetSession(r)
		url := r.URL.String()

		if sess == nil {
			utils.Debug("[LoggedHandler]", "Error on get session", url)
			handlers.Redirect(r, w, "/login")
			return
		}

		l := sess.Values[models.SESSION_VALUE_LOGGED]
		//utils.Debug("[LoggedHandler]", sess.Values[models.SESSION_VALUE_LOGGED])
		if logged, ok := l.(bool); !ok || !logged {
			utils.Debug("[LoggedHandler]", "Auth Error", url)
			handlers.Redirect(r, w, "/login")
			return
		}

		utils.Debug("[LoggedHandler]", "Authenticated", url)
		next.ServeHTTP(w, r)
	})
}

// BlockToClientUserHander -
// @handler
// Intercept all transactions and check if is Client
func BlockToClientUserHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sess, _ = handlers.GetSession(r)
		url := r.URL.String()

		l := sess.Values[models.SESSION_VALUE_LEVEL]
		if l == models.USER_CLIENT {
			utils.Debug("[Permission][BlockToClientUser]", "Permission Denied", url)
			handlers.Redirect(r, w, fmt.Sprintf("%s/forbidden", configurations.Configuration.StaticPath))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// OnlyAdminHandler -
// @handler
// Intercept all transactions and check if user is admin
func OnlyAdminHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sess, _ = handlers.GetSession(r)
		url := r.URL.String()

		l := sess.Values[models.SESSION_VALUE_LEVEL]
		if l != models.USER_ROOT && l != models.USER_ADMIN {
			utils.Debug("[Permission][OnlyAdminHandler]", "Permission Denied", url)
			handlers.Redirect(r, w, fmt.Sprintf("%s/forbidden", configurations.Configuration.StaticPath))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// AuthProtection - Chain Logged handler to protect connections
// @middleware
// Uses session stored value `logged` to make a best gin of the world
func AuthProtection(f http.HandlerFunc) http.HandlerFunc {
	return handlers.Chain(f, LoggedHandler)
}

// AuthTokenedProtection - Chain Logged handler to protect connections
// @middleware
// Uses session stored value `logged` to make a best gin of the world
// If is not connected, check token
func AuthTokenedProtection(f http.HandlerFunc) http.HandlerFunc {
	return handlers.Chain(f, TokenHandler)
}

// BlockForClients - Deny acess to `users` level
// @middleware
// Chain conections to restrict area from Client User
// Uses session stored value `level` to make the magic
func BlockForClients(f http.HandlerFunc) http.HandlerFunc {
	return handlers.Chain(f, LoggedHandler, BlockToClientUserHandler)
}