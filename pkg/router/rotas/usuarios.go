package rotas

import (
	"devbook-api/pkg/controllers"
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
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}",
		Method:         http.MethodGet,
		Function:       controllers.BuscarUsuario,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}",
		Method:         http.MethodPut,
		Function:       controllers.AtualizarUsuario,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletarUsuario,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}/seguir",
		Method:         http.MethodPost,
		Function:       controllers.SeguirUsuario,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}/deixar-de-seguir",
		Method:         http.MethodPost,
		Function:       controllers.DeixarDeSeguirUsuario,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}/seguidores",
		Method:         http.MethodGet,
		Function:       controllers.BuscarSeguidores,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}/seguindo",
		Method:         http.MethodGet,
		Function:       controllers.BuscarSeguindo,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}/atualizar-senha",
		Method:         http.MethodPost,
		Function:       controllers.AtualizarSenha,
		Authentication: true,
	},
}
