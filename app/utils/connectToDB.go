package utils

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"          // Driver PostgreSQL
    "github.com/joho/godotenv"     // Para carregar variáveis de ambiente
)

var DB *sql.DB

func ConnectToDB() {
    // Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    // Obtém as variáveis de ambiente
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")

    // String de conexão
    connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        user, password, dbname, host, port)

    // Abre a conexão com o banco de dados
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }

    // Testa a conexão
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
    }

    fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}