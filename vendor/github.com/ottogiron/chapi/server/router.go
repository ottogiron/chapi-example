package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type router struct {
	*mux.Router
}

func newRouter() *router {
	negroniRouter := mux.NewRouter()
	return &router{negroniRouter}
}

//Vars returns the path variables fron a request
func Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}
