package rotas

import (
	"devbook-api/pkg/controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:            "/login",
	Method:         http.MethodPost,
	Function:       controllers.Login,
	Authentication: false,
}
