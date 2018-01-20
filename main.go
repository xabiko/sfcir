package main

import (
	"log"

	"github.com/xabiko/sfcir/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
