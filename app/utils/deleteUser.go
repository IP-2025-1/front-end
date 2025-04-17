package utils

import (
    "log"
)

func DeleteUser(email string) error {
    query := `DELETE FROM users WHERE email = $1`
    _, err := DB.Exec(query, email)
    if err != nil {
        log.Printf("Erro ao apagar usuário do banco de dados: %v", err)
        return err
    }
    log.Println("Usuário apagado com sucesso!")
    return nil
}