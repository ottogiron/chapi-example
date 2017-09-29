package server

import "github.com/gorilla/mux"

type route struct {
	*mux.Route
}

func newRoute(muxRoute *mux.Route) *route {

	return &route{muxRoute}
}
