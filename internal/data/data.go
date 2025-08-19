package data

import (
	"database/sql"
	"devbook-api/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre e retorna a conexão com o banco de dados.
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
