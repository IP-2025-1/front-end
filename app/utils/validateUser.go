package utils

import (
	"log"
)

// Valida as credenciais do usuário no banco de dados
func ValidateUser(email, encryptedPassword string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE email = $1 AND password = $2`
	var count int
	err := DB.QueryRow(query, email, encryptedPassword).Scan(&count)
	if err != nil {
		log.Printf("Erro ao validar usuário no banco de dados: %v", err)
		return false, err
	}
	return count > 0, nil
}
