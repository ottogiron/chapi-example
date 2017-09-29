package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

type negroniServer struct {
	*negroni.Negroni
	router *router
	*baseServer
}

func NewServer() Server {
	n := negroni.Classic()
	router := newRouter()
	baseServer := &baseServer{}
	baseServer.plugins = make(map[string]Plugin)
	ns := negroniServer{n, router, baseServer}
	n.UseHandler(router)
	return ns
}

func (server negroniServer) Run(addr string) error {
	processedPlugins := make(map[string]bool)
	serverRegisterError := server.registerPlugins(server.plugins, server, processedPlugins)
	if serverRegisterError != nil {
		return serverRegisterError
	}
	server.Negroni.Run(addr)
	return nil
}

func (server negroniServer) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *route {
	muxRoute := server.router.HandleFunc(path, f)
	route := newRoute(muxRoute)
	return route
}
