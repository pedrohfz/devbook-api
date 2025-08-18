package models

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação feita por um usuário.
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitzero"`
}

// Preparar vai chamar os métodos para validar e formatar a publicação recebida.
func (p *Publicacao) Preparar() error {
	if err := p.validar(); err != nil {
		return err
	}

	p.formatar()
	return nil
}

func (p *Publicacao) validar() error {
	if p.Titulo == "" {
		return errors.New("O título é obrigatório e não pode estar em branco!")
	}

	if p.Conteudo == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco!")
	}

	return nil
}

func (p *Publicacao) formatar() {
	p.Titulo = strings.TrimSpace(p.Titulo)
	p.Conteudo = strings.TrimSpace(p.Conteudo)
}
