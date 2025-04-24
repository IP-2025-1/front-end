package utils

// Importa os pacotes necessários para a conexão com o banco de dados
import (
    "database/sql"          // Usado para interagir com o banco de dados
    "fmt"                   // Usado para formatar strings
    "log"                   // Usado para registrar mensagens de log e erros
    "os"                    // Usado para acessar variáveis de ambiente

    _ "github.com/lib/pq"          // Driver PostgreSQL para o pacote database/sql
    "github.com/joho/godotenv"     // Usado para carregar variáveis de ambiente de um arquivo .env
)

// Declara uma variável global para armazenar a conexão com o banco de dados
var DB *sql.DB

// Função responsável por conectar ao banco de dados
func ConnectToDB() {
    // Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        // Encerra o programa caso ocorra um erro ao carregar o arquivo .env
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    // Obtém as variáveis de ambiente necessárias para a conexão
    user := os.Getenv("DB_USER")       // Usuário do banco de dados
    password := os.Getenv("DB_PASSWORD") // Senha do banco de dados
    dbname := os.Getenv("DB_NAME")     // Nome do banco de dados
    host := os.Getenv("DB_HOST")       // Host do banco de dados
    port := os.Getenv("DB_PORT")       // Porta do banco de dados

    // Cria a string de conexão com base nas variáveis de ambiente
    connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        user, password, dbname, host, port)

    // Abre a conexão com o banco de dados usando a string de conexão
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        // Encerra o programa caso ocorra um erro ao abrir a conexão
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }

    // Testa a conexão com o banco de dados
    err = DB.Ping()
    if err != nil {
        // Encerra o programa caso ocorra um erro ao verificar a conexão
        log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
    }

    // Imprime uma mensagem de sucesso no terminal caso a conexão seja estabelecida
    fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}