package main

import (
	"fmt"
	"net/http"
	"os"

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
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}

	connectionString := ":" + port
	s := server.NewServer()
	s.Register(&HelloController{})
	err := s.Run(connectionString)
	if err != nil {
		fmt.Println("Error when running server", err)
	}

}
