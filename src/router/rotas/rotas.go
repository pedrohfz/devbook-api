package rotas

import "net/http"

// Rota representa todas as rotas da API
type Rota struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}
