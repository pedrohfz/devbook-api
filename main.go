package main

import (
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Gerar()

	fmt.Println("Servidor rodando e escutando na porta :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
