package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hash irá receber uma senha do tipo string e irá colocar um hash nela.
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha e um hash e retorna se elas são iguais.
func VerificarSenha(senhaHash, senhaString string) error {

	// ? Mock para testar funcionamento do Hash
	fmt.Println("Hash: ", senhaHash)
	fmt.Println("Senha: ", senhaString)

	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
