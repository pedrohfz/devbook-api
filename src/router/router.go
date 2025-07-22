package router

import "github.com/gorilla/mux"

// Gerar é uma função que vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}