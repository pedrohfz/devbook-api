package models

// DadosAuthentication{} contém o token e o id do usuário autenticado.
type DadosAuthentication struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
