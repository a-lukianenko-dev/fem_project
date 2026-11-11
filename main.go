package main

import (
	"net/http"
	"time"

	"github.com/a-lukianenko-dev/fem_project/internal/app"
	"github.com/a-lukianenko-dev/fem_project/internal/routes"
)

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Print("Hello from the app!")

	handler := routes.SetupRoutesHandler(app)

	server := &http.Server{
		Addr:         ":3000",
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		Handler:      handler,
	}

	app.Logger.Fatal(server.ListenAndServe())

}
