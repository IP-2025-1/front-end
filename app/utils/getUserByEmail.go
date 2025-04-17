package utils

import (
    "log"
)

type User struct {
    Username string
    Email    string
    BornDate string
}

func GetUserByEmail(email string) (*User, error) {
    query := `SELECT username, email, born_date FROM users WHERE email = $1`
    var user User
    err := DB.QueryRow(query, email).Scan(&user.Username, &user.Email, &user.BornDate)
    if err != nil {
        log.Printf("Erro ao buscar usuário no banco de dados: %v", err)
        return nil, err
    }
    return &user, nil
}