package controllers

import (
	"devbook-api/src/auth"
	"devbook-api/src/data"
	"devbook-api/src/models"
	"devbook-api/src/repository"
	"devbook-api/src/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacao.ID, err = repositorio.Criar(publicacao)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusCreated, publicacao)
}

// BuscarPublicacoes traz as publicações que apareceriam no feed do usuário.
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		utils.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacoes, err := repositorio.Buscar(usuarioID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, publicacoes)
}

// BuscarPublicacao traz uma única publicação.
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(param["publicacaoID"], 10, 64)
	if err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacao, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, publicacao)
}

// AtualizarPublicacao altera os dados de uma publicação.
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		utils.Erro(w, http.StatusUnauthorized, err)
		return
	}

	param := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(param["publicacaoID"], 10, 64)
	if err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacoesSalvaNoBanco, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacoesSalvaNoBanco.AutorID != usuarioID {
		utils.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não seja sua!"))
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

	if err = publicacao.Preparar(); err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repositorio.Atualizar(publicacaoID, publicacao); err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}

// DeletarPublicacao exclui os dados de uma publicação.
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {}
