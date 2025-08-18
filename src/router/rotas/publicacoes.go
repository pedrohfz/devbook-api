package rotas

import (
	"devbook-api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:            "/publicacoes",
		Method:         http.MethodPost,
		Function:       controllers.CriarPublicacao,
		Authentication: true,
	},
	{
		URI:            "/publicacoes",
		Method:         http.MethodGet,
		Function:       controllers.BuscarPublicacoes,
		Authentication: true,
	},
	{
		URI:            "/publicacoes/{publicacaoID}",
		Method:         http.MethodGet,
		Function:       controllers.BuscarPublicacao,
		Authentication: true,
	},
	{
		URI:            "/publicacoes/{publicacaoID}",
		Method:         http.MethodPut,
		Function:       controllers.AtualizarPublicacao,
		Authentication: true,
	},
	{
		URI:            "/publicacoes/{publicacaoID}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletarPublicacao,
		Authentication: true,
	},
}
