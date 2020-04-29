package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaopandolfi/blackwhale/handlers"
	"net/http"
)

// PageController - Interface for controllers that will render pages
type PageController interface {
	Precompile()
}

func precompilePages() {

}

func pages(r *mux.Router) {
	r.HandleFunc("/login", func (w http.ResponseWriter, r *http.Request){
		handlers.RESTResponseError(w, 404)
	}).Methods("GET")
}
