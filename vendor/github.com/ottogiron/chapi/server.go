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

func (helloController *HelloController) Register(server server.Server) {
	fmt.Println("Registering HelloController")
	server.HandleFunc("/", handleHello).Methods("GET")
}

func (helloController *HelloController) Name() string {
	return "HelloController"
}

func (helloController *HelloController) Dependencies() []string {
	return []string{"HelloDep"}
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Chapi")
}

type HelloDep struct {
	*server.BasePlugin
}

func (helloDep *HelloDep) Name() string {
	return "HelloDep"
}

func (helloDep *HelloDep) Register(server server.Server) {
	fmt.Println("Registering HelloDepController")
	server.HandleFunc("/hellodep", handleHelloDep).Methods("GET")
}

func handleHelloDep(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Chapi Dep")
}

func main() {

	connectionString := ":" + os.Getenv("PORT")
	s := server.NewServer()
	s.Register(&HelloController{})
	s.Register(&HelloDep{})
	runError := s.Run(connectionString)
	if runError != nil {
		fmt.Println("Error when running server", runError)
	}

}
