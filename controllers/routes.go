package controllers

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"Mutants", "POST", "/mutant", Mutants},
	Route{"Statistics", "GET", "/stats", Statistics},
}
