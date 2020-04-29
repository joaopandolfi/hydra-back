package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaopandolfi/blackwhale/configurations"
	"github.com/unrolled/secure"
)

// Register routes on server
func Register(r *mux.Router) {
	//r.Use(handlers.LoggedHandler)
	health(r)
	//public(r)
	pages(r)
	rest(r)
}

// Precompile pages
func Precompile() {
	precompilePages()
}

// Handlers default Register
func Handlers(r *mux.Router) {
	secureMiddleware := secure.New(configurations.Configuration.Security.Options)
	r.Use(secureMiddleware.Handler)
}
