package rotas

import (
	"devbook-api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:            "/usuarios",
		Method:         http.MethodPost,
		Function:       controllers.CriarUsuario,
		Authentication: false,
	},
	{
		URI:            "/usuarios",
		Method:         http.MethodGet,
		Function:       controllers.BuscarUsuarios,
		Authentication: true, // ? Teste Token de Autenticação
	},
	{
		URI:            "/usuarios/{usuarioID}",
		Method:         http.MethodGet,
		Function:       controllers.BuscarUsuario,
		Authentication: false,
	},
	{
		URI:            "/usuarios/{usuarioID}",
		Method:         http.MethodPut,
		Function:       controllers.AtualizarUsuario,
		Authentication: false,
	},
	{
		URI:            "/usuarios/{usuarioID}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletarUsuario,
		Authentication: false,
	},
}
