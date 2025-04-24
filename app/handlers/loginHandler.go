package handlers

// Importa os pacotes necessários para o funcionamento do handler
import (
    "fmt"                  // Usado para imprimir mensagens no terminal (debug)
    "servidorHTTP/app/utils" // Importa funções utilitárias, como validação de usuário e criptografia
    "net/http"            // Usado para lidar com requisições e respostas HTTP
    "text/template"       // Usado para renderizar templates HTML
)

// LoginHandler é responsável por processar os dados enviados pelo formulário de login
func LoginHandler(response http.ResponseWriter, request *http.Request) {

    // Verifica se o método da requisição é POST
    if request.Method != http.MethodPost {
        // Retorna um erro caso o método não seja suportado
        http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    // Obtém os valores enviados pelo formulário
    email := request.FormValue("email")       // E-mail do usuário
    password := request.FormValue("password") // Senha do usuário

    // Criptografa a senha para compará-la com a armazenada no banco de dados
    encryptedPassword := utils.Encrypt(password)

    // Verifica se o usuário existe no banco de dados
    isValidUser, err := utils.ValidateUser(email, encryptedPassword)
    if err != nil {
        // Retorna um erro caso ocorra falha ao validar o usuário
        http.Error(response, "Erro ao validar o usuário", http.StatusInternalServerError)
        return
    }

    // Verifica se as credenciais são inválidas
    if !isValidUser {
        // Retorna um erro caso as credenciais estejam incorretas
        http.Error(response, "Credenciais inválidas", http.StatusUnauthorized)
        fmt.Println("Erro") // Mensagem de debug
        return
    }

    // Busca as informações do usuário no banco de dados
    user, err := utils.GetUserByEmail(email)
    if err != nil {
        // Retorna um erro caso ocorra falha ao buscar as informações do usuário
        http.Error(response, "Erro ao buscar informações do usuário", http.StatusInternalServerError)
        return
    }

    // Carrega o template do perfil do usuário
    tmpl, err := template.ParseFiles("static/profile.html")
    if err != nil {
        // Retorna um erro caso ocorra falha ao carregar o template
        http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
        return
    }

    // Renderiza o template com os dados do usuário
    err = tmpl.Execute(response, user)
    if err != nil {
        // Retorna um erro caso ocorra falha ao renderizar o template
        http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
        return
    }

    // Redireciona o usuário para a página de perfil após o sucesso
    http.Redirect(response, request, "/profile.html", http.StatusOK)
}