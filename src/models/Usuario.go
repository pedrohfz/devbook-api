package models

import "time"

type Usuario struct {
	ID       uint64    `json:"id"`
	Nome     string    `json:"nome"`
	Nick     string    `json:"nick"`
	Email    string    `json:"email"`
	Senha    string    `json:"senha"`
	CriadoEm time.Time `json:"CriadoEm"`
}
