package controllers

import (
	"devbook-api/src/auth"
	"devbook-api/src/data"
	"devbook-api/src/models"
	"devbook-api/src/repository"
	"devbook-api/src/utils"
	"encoding/json"
	"io"
	"net/http"
)

// CriarPublicacao adiciona uma nova publicação no banco de dados.
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		utils.Erro(w, http.StatusUnauthorized, err)
		return
	}

	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao models.Publicacao
	if err = json.Unmarshal(corpoRequisicao, &publicacao); err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	publicacao.AutorID = usuarioID

	if err = publicacao.Preparar(); err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacao.ID, err = repositorio.Criar(publicacao)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusCreated, publicacao)
}

// BuscarPublicacoes traz as publicações que apareceriam no feed do usuário.
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {}

// BuscarPublicacao traz uma única publicação.
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {}

// AtualizarPublicacao altera os dados de uma publicação.
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {}

// DeletarPublicacao exclui os dados de uma publicação.
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {}
