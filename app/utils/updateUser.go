package utils

import (
	"fmt"
	"log"
	"strings"
)

func UpdateUser(currentEmail string, updates map[string]string) error {
	// Constrói a query dinamicamente com base nos campos fornecidos
	setClauses := []string{}
	values := []interface{}{}
	i := 1

	for column, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, i))
		values = append(values, value)
		i++
	}

	// Adiciona o email atual como condição
	values = append(values, currentEmail)
	query := fmt.Sprintf("UPDATE users SET %s WHERE email = $%d", strings.Join(setClauses, ", "), i)

	// Executa a query
	_, err := DB.Exec(query, values...)
	if err != nil {
		log.Printf("Erro ao atualizar usuário no banco de dados: %v", err)
		return err
	}
	log.Println("Usuário atualizado com sucesso!")
	return nil
}
