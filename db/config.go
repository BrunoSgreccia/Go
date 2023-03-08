package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // importa o driver para SQL Server

)

const (
	server   = "localhost"
	port     = 1433
	user     = "usuario"
	password = "senha"
	database = "meubanco"
)

// AbrirConexao retorna uma nova conexão com o banco de dados SQL Server
func AbrirConexao() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	return db, nil
}
