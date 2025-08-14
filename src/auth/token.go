package auth

import (
	"devbook-api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// ValidarToken verifica se o token passado na requisição é válido.
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveVerificacao)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido!")
}

// ExtrairUsuarioID retorna o usuarioID que está salvo no token.
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveVerificacao)
	if err != nil {
		return 0, err
	}

	if perm, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, err := strconv.ParseUint(fmt.Sprintf("%.0f", perm["usuarioID"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return usuarioID, nil
	}

	return 0, errors.New("Token inválido!")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// * Bearer Token = Bearer *seuToken*
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
