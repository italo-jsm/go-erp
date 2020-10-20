package main

import (
	"github.com/italosm/go-erp/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.CreateRoutes()
	log.Fatal(http.ListenAndServe(":3000", router))
}