package utils

import (
	"log"
)

// Insere um novo usuário no banco de dados
func InsertUser(username, email, bornDate, password string) error {
	query := `INSERT INTO users (username, email, born_date, password) VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(query, username, email, bornDate, password)
	if err != nil {
		log.Printf("Erro ao inserir usuário no banco de dados: %v", err)
		return err
	}
	log.Println("Usuário inserido com sucesso!")
	return nil
}
