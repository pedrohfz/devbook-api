package auth

import (
	"devbook-api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário.
func CriarToken(usuarioID uint64) (string, error) {
	perm := jwt.MapClaims{}
	perm["authorized"] = true
	perm["exp"] = time.Now().Add(time.Hour * 12).Unix()
	perm["usuarioID"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perm)

	return token.SignedString([]byte(config.SecretKey))
}