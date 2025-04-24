package handlers

// Importa os pacotes necessários para o funcionamento do handler
import (
    "servidorHTTP/app/utils" // Importa funções utilitárias, como criptografia e inserção no banco de dados
    "net/http"              // Usado para lidar com requisições e respostas HTTP
)

// FormHandler é responsável por processar os dados enviados pelo formulário de criação de conta
func FormHandler(response http.ResponseWriter, request *http.Request) {
    // Verifica se o método da requisição é POST
    if request.Method != http.MethodPost {
        // Retorna um erro caso o método não seja suportado
        http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    // Obtém os valores enviados pelo formulário
    username := request.FormValue("username") // Nome do usuário
    email := request.FormValue("email")       // E-mail do usuário
    bornDate := request.FormValue("bornDate") // Data de nascimento do usuário
    password := request.FormValue("password") // Senha do usuário

    // Criptografa a senha para armazená-la de forma segura no banco de dados
    encryptedPassword := utils.Encrypt(password)

    // Insere os dados do usuário no banco de dados
    err := utils.InsertUser(username, email, bornDate, encryptedPassword)
    if err != nil {
        // Retorna um erro caso ocorra falha ao salvar os dados no banco de dados
        http.Error(response, "Erro ao salvar os dados no banco de dados", http.StatusInternalServerError)
        return
    }

    // Redireciona o usuário para a página inicial após o sucesso
    http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}