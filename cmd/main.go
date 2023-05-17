package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	var database Database
	err := database.initConnection()
	if err != nil {
		log.Fatal(err)
	}

	var controller Controller
	controller.setRepository(&database)

	router := httprouter.New()
	controller.setupRoutes(router)

	const port = 8080

	log.Printf("Serving on localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
