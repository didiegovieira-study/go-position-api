package main

import (
	"github.com/didiegovieira/go-position-api/di"
)

func main() {
	app, cleanup, err := di.InitializeApi()
	if err != nil {
		panic(err)
	}

	err = app.Start()
	if err != nil {
		panic(err)
	}

	cleanup()
}
