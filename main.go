package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/a-lukianenko-dev/fem_project/internal/app"
	"github.com/a-lukianenko-dev/fem_project/internal/routes"
)

func main() {
	port := flag.Int("port", 3000, "server port")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Print("Hello from the app!")

	handler := routes.SetupRoutesHandler(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		Handler:      handler,
	}

	app.Logger.Fatal(server.ListenAndServe())

}
