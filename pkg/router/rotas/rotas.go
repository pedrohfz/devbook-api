package rotas

import (
	"devbook-api/internal/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API.
type Rota struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

// Configurar posiciona todas as rotas dentro do router.
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {
		if rota.Authentication {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Function))).Methods(rota.Method)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Function)).Methods(rota.Method)
		}
	}

	return r
}
