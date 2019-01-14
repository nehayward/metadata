package main

import (
	"log"
	"net/http"

	"github.com/nehayward/metadata/app"
)

const port = "8080"

func main() {
	router := app.NewRouter(app.AllRoutes())
	log.Print("Listening on http://localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
