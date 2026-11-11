package main

import "github.com/a-lukianenko-dev/fem_project/internal/app"

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Print("Hello from the app!")
}
