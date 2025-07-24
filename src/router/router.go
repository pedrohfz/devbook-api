package router

import (
	"devbook-api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar é uma função que vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}