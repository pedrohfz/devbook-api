package controllers

import (
	"devbook-api/internal/auth"
	"devbook-api/internal/data"
	"devbook-api/internal/repository"
	"devbook-api/internal/security"
	"devbook-api/pkg/models"
	"devbook-api/pkg/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
	}

	if err = usuario.Preparar("cadastro"); err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NovoRepositorioDeUsuarios(db)
	usuario.ID, err = repo.Criar(usuario)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
	}

	utils.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuários salvos no banco de dados.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NovoRepositorioDeUsuarios(db)
	usuarios, err := repo.Buscar(nomeOuNick)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario busca um único usuário salvo no banco de dados.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
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

	repo := repository.NovoRepositorioDeUsuarios(db)
	usuario, err := repo.BuscarPorID(usuarioID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario altera as informações de um único usuário no banco de dados.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
	if err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIDNoToken, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		utils.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIDNoToken {
		utils.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário que não seja o seu"))
		return
	}

	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("edicao"); err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NovoRepositorioDeUsuarios(db)
	if err = repo.Atualizar(usuarioID, usuario); err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}

// DeletarUsuario exclui um único usuário no banco de dados.
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
	if err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIDNoToken, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		utils.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIDNoToken {
		utils.Erro(w, http.StatusForbidden, errors.New("não é possível excluir um usuário que não seja o seu"))
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NovoRepositorioDeUsuarios(db)
	if err = repo.Deletar(usuarioID); err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}

// SeguirUsuario permite que um usuário siga outro.
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		utils.Erro(w, http.StatusUnauthorized, err)
		return
	}

	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
	if err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	if seguidorID == usuarioID {
		utils.Erro(w, http.StatusForbidden, errors.New("não é possível você seguir a sua própria conta"))
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NovoRepositorioDeUsuarios(db)
	if err = repo.Seguir(usuarioID, seguidorID); err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}

// DeixarDeSeguirUsuario permite que um usuário pare de seguir outro.
func DeixarDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		utils.Erro(w, http.StatusUnauthorized, err)
		return
	}

	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
	if err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	if seguidorID == usuarioID {
		utils.Erro(w, http.StatusForbidden, errors.New("não é possível você deixar de seguir a sua própria conta"))
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NovoRepositorioDeUsuarios(db)
	if err = repo.DeixarDeSeguir(usuarioID, seguidorID); err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores traz todos os seguidores de um usuário.
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
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

	repo := repository.NovoRepositorioDeUsuarios(db)
	seguidores, err := repo.BuscarSeguidores(usuarioID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, seguidores)
}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo.
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
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

	repo := repository.NovoRepositorioDeUsuarios(db)
	usuarios, err := repo.BuscarSeguindo(usuarioID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, usuarios)
}

// AtualizarSenha permite o alterar a senha de um usuário.
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDNoToken, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		utils.Erro(w, http.StatusUnauthorized, err)
		return
	}

	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
	if err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	if usuarioIDNoToken != usuarioID {
		utils.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar a senha de um usuário que não seja o seu"))
		return
	}

	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var senha models.Senha
	if err = json.Unmarshal(corpoRequisicao, &senha); err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NovoRepositorioDeUsuarios(db)
	senhaSalvaNoBanco, err := repo.BuscarSenha(usuarioID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerificarSenha(senhaSalvaNoBanco, senha.Atual); err != nil {
		utils.Erro(w, http.StatusUnauthorized, errors.New("a senha atual não condiz com a que está salva no banco"))
		return
	}

	senhaComHash, err := security.Hash(senha.Nova)
	if err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.AtualizarSenha(usuarioID, string(senhaComHash)); err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}
