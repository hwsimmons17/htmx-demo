package main

import (
	"htmx-demo/pkg/app"
	"log"
)

func main() {
	app := app.NewApp()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
