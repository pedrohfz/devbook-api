package controllers

import (
	"devbook-api/src/data"
	"devbook-api/src/models"
	"devbook-api/src/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := data.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuarioID, err := repositorio.Criar(usuario)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioID)))
}

// BuscarUsuarios busca todos os usuários salvos no banco de dados.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

// BuscarUsuario busca um único usuário salvo no banco de dados.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!"))
}

// AtualizarUsuario altera as informações de um único usuário no banco de dados.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando os usuário!"))
}

// DeletarUsuario exclui um único usuário no banco de dados.
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário!"))
}
