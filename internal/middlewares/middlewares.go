package middlewares

import (
	"devbook-api/internal/auth"
	"devbook-api/pkg/utils"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal.
func Logger(logs http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n -> %s %s %s", r.Method, r.RequestURI, r.Host)
		logs(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado.
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidarToken(r); err != nil {
			utils.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
