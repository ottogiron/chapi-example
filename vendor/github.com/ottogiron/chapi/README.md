# Chapi Toy Web Framework
### Implemented for learning purposes


### Package server

```go
import "github.com/ottogiron/chapi/server"
```

#### Installation
```bash
go get github.com/ottogiron/chapi
```

#### server.Server

Server is the application container.

**server.NewServer()**

Creates a new server

**Server.Register(plugin Plugin)**

Registers a new server plugin

**Server.Run(add String)**

Runs a server on the specified address

**Server.HandleFunc(path string, f func(http.ResponseWriter, *http.Request))  route**

Registers a new handler function.


***Example***

server.go
```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ottogiron/chapi/server"
)

//Hello Controller Register which is dependent from HelloDepController(another plugin)
type HelloController struct {
	*server.BasePlugin
}

func (helloController *HelloController) Register(server server.Server) {
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


//HelloDep Controller (Register) will run before HelloController since it is a dependency
type HelloDepController struct {
	*server.BasePlugin
}

func (helloDep *HelloDepController) Name() string {
	return "HelloDep"
}

func (helloDep *HelloDepController) Register(server server.Server) {
	server.HandleFunc("/hellodep", handleHelloDep).Methods("GET")
}

func handleHelloDep(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Chapi Dep")
}

func main() {

	connectionString := ":8080"
	s := server.NewServer()
	s.Register(&HelloController{})
	s.Register(&HelloDepController{})
	runError := s.Run(connectionString)
	if runError != nil {
		fmt.Println("Error when running server", runError)
	}

}
```
***Running the server***
```
go run server.go
```
