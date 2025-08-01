package models

import "time"

// Usuario representa os campos de todos os usuários da aplicação.
type Usuario struct {
	ID       uint64    `json:"id"`
	Nome     string    `json:"nome"`
	Nick     string    `json:"nick"`
	Email    string    `json:"email"`
	Senha    string    `json:"senha"`
	CriadoEm time.Time `json:"CriadoEm"`
}
