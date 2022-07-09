package app

import (
	"errors"
	"fmt"
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

func w(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get custom function"))
}

func (a *Application) BootstrapRoutes() {
	a.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get organization"))
	})

	a.Router.Get("/foo", w)
}

// create a new server
func (a *Application) Run() {
	err := http.ListenAndServe(a.Port, a.Router)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
