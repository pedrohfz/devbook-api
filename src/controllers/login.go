package controllers

import (
	"devbook-api/src/auth"
	"devbook-api/src/data"
	"devbook-api/src/models"
	"devbook-api/src/repository"
	"devbook-api/src/security"
	"devbook-api/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Login é responsável por autenticar um usuário na API.
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := data.Conectar()
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		utils.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); err != nil {
		utils.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.CriarToken(usuarioSalvoNoBanco.ID)
	// TODO: Função para realizar a autenticação do Token e adicionar Validações de Métodos

	fmt.Println(token)
}
