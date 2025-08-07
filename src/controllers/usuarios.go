package controllers

import (
	"devbook-api/src/data"
	"devbook-api/src/models"
	"devbook-api/src/repository"
	"devbook-api/src/utils"
	"encoding/json"
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

	if err = usuario.Preparar(); err != nil {
		utils.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuario.ID, err = repositorio.Criar(usuario)
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

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuarios, err := repositorio.Buscar(nomeOuNick)
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

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuario, err := repositorio.BuscarPorID(usuarioID)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario altera as informações de um único usuário no banco de dados.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando os usuário!"))
}

// DeletarUsuario exclui um único usuário no banco de dados.
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário!"))
}
