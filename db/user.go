package db

import (
	"database/sql"

)

// Usuario representa um usuário no banco de dados
type Usuario struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
}

// TodosUsuarios retorna todos os usuários no banco de dados
func TodosUsuarios(db *sql.DB) ([]Usuario, error) {
	rows, err := db.Query("SELECT id, nome, email, telefone FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Telefone)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}

// NovoUsuario insere um novo usuário no banco de dados
func NovoUsuario(db *sql.DB, nome string, email string, telefone string) error {
	_, err := db.Exec("INSERT INTO usuarios(nome, email, telefone) VALUES(?, ?, ?)",
		nome, email, telefone)
	return err
}
