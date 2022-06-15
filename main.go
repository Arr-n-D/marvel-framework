package main

import (
	"log"

	app "github.com/Arr-n-D/marvel/App"
	routing "github.com/Arr-n-D/marvel/Routing"
)

func main() {
	// print all verbs from Routing\route.go
	log.Println(routing.VERBS)
	app := app.Bootstrap()
	app.Run()

	// add a new route
}
