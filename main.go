package main

import (
	"devbook-api/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
