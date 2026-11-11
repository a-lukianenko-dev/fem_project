package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-lukianenko-dev/fem_project/internal/app"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Print("Hello from the app!")

	handler := http.NewServeMux()
	handler.HandleFunc("/health", HandleHealth)

	server := &http.Server{
		Addr:         ":3000",
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		Handler:      handler,
	}

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}
