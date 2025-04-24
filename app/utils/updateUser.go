package utils

// Importa os pacotes necessários para a função
import (
    "fmt"      // Usado para formatar strings
    "log"      // Usado para registrar mensagens de log e erros
    "strings"  // Usado para manipular strings
)

// UpdateUser atualiza os dados de um usuário no banco de dados
func UpdateUser(currentEmail string, updates map[string]string) error {
    // Cria uma lista para armazenar os trechos da query SQL que serão atualizados
    setClauses := []string{}
    // Cria uma lista para armazenar os valores que serão passados como parâmetros na query
    values := []interface{}{}
    // Inicializa um contador para numerar os placeholders ($1, $2, etc.)
    i := 1

    // Itera sobre os campos e valores fornecidos no mapa de atualizações
    for column, value := range updates {
        // Adiciona o trecho "coluna = $i" à lista de trechos da query
        setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, i))
        // Adiciona o valor correspondente à lista de valores
        values = append(values, value)
        // Incrementa o contador
        i++
    }

    // Adiciona o email atual como condição para identificar o usuário a ser atualizado
    values = append(values, currentEmail)
    // Constrói a query SQL completa com os trechos de atualização e a condição
    query := fmt.Sprintf("UPDATE users SET %s WHERE email = $%d", strings.Join(setClauses, ", "), i)

    // Executa a query no banco de dados, passando os valores como parâmetros
    _, err := DB.Exec(query, values...)
    if err != nil {
        // Registra uma mensagem de erro no log caso a execução da query falhe
        log.Printf("Erro ao atualizar usuário no banco de dados: %v", err)
        return err // Retorna o erro para o chamador
    }

    // Registra uma mensagem de sucesso no log caso a atualização seja bem-sucedida
    log.Println("Usuário atualizado com sucesso!")
    return nil // Retorna nil indicando que não houve erros
}