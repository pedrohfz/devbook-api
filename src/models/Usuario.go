package models

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa os campos de todos os usuários da aplicação.
type Usuario struct {
	ID       uint64    `json:"id"`
	Nome     string    `json:"nome"`
	Nick     string    `json:"nick"`
	Email    string    `json:"email"`
	Senha    string    `json:"senha"`
	CriadoEm time.Time `json:"CriadoEm"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido.
func (usuario *Usuario) Preparar() error {
	if err := usuario.validar(); err != nil {
		return err
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco!")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco!")
	}

	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco!")
	}

	if usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco!")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
