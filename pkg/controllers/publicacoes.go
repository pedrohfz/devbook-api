package controllers

import (
	"devbook-api/internal/auth"
	"devbook-api/internal/data"
	"devbook-api/internal/repository"
	"devbook-api/pkg/models"
	"devbook-api/pkg/utils"
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

	repo := repository.NovoRepositorioDePublicacoes(db)
	publicacao.ID, err = repo.Criar(publicacao)
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

	repo := repository.NovoRepositorioDePublicacoes(db)
	publicacoes, err := repo.Buscar(usuarioID)
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

	repo := repository.NovoRepositorioDePublicacoes(db)
	publicacao, err := repo.BuscarPorID(publicacaoID)
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

	repo := repository.NovoRepositorioDePublicacoes(db)
	publicacoesSalvaNoBanco, err := repo.BuscarPorID(publicacaoID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacoesSalvaNoBanco.AutorID != usuarioID {
		utils.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicação que não seja sua"))
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

	if err = repo.Atualizar(publicacaoID, publicacao); err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}

// DeletarPublicacao exclui os dados de uma publicação.
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
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

	repo := repository.NovoRepositorioDePublicacoes(db)
	publicacoesSalvaNoBanco, err := repo.BuscarPorID(publicacaoID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacoesSalvaNoBanco.AutorID != usuarioID {
		utils.Erro(w, http.StatusForbidden, errors.New("não é possível deletar uma publicação que não seja sua"))
		return
	}

	if err = repo.Deletar(publicacaoID); err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}

// BuscarPublicacoesPorUsuario traz todas as publicações de um usuário específico.
func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
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

	repo := repository.NovoRepositorioDePublicacoes(db)
	publicacoes, err := repo.BuscarPorUsuario(usuarioID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, publicacoes)
}
