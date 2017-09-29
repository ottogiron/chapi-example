package main

import (
	"fmt"
	"net/http"

	"github.com/ottogiron/chapi/server"
)

type HelloController struct {
	*server.BasePlugin
}

func (c *HelloController) Register(s server.Server) {
	s.HandleFunc("/", handleHello).Methods("GET")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "!!!Hello Chapi")
}

func main() {

	connectionString := ":80"
	s := server.NewServer()
	s.Register(&HelloController{})
	err := s.Run(connectionString)
	if err != nil {
		fmt.Println("Error when running server", err)
	}

}
