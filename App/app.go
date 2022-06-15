package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	routing "github.com/Arr-n-D/marvel/Routing"
)

// create a type called Application
type Application struct {
	Router *routing.Router
	Port   string
}

func Bootstrap() *Application {
	a := &Application{
		Router: routing.NewRouter(),
		Port:   ":8080",
	}

	a.BootstrapRoutes()

	return a

}

func (a *Application) BootstrapRoutes() {
	a.Router.Get("/", func() {
		log.Println("Hello World")
	})
}

// create a new server
func (a *Application) Run() {
	log.Println("Listening on port", a.Port)
	// loop over the router's routes and add them with the http.HandleFunc function
	for _, route := range a.Router.Routes {
		// print the route and all its data
		http.HandleFunc(route.GetUri(), func(w http.ResponseWriter, r *http.Request) {

		})
	}

	err := http.ListenAndServe(a.Port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
