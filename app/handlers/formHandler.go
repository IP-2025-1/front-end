package handlers

import (
    "servidorHTTP/app/utils"
    "net/http"
)

func FormHandler(response http.ResponseWriter, request *http.Request) {
    if request.Method != http.MethodPost {
        http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    // Obtém os valores do formulário
    username := request.FormValue("username")
    email := request.FormValue("email")
    bornDate := request.FormValue("bornDate")
    password := request.FormValue("password")

    // Criptografa a senha
    encryptedPassword := utils.Encrypt(password)

    // Insere os dados no banco de dados
    err := utils.InsertUser(username, email, bornDate, encryptedPassword)
    if err != nil {
        http.Error(response, "Erro ao salvar os dados no banco de dados", http.StatusInternalServerError)
        return
    }

    // Redireciona para a página inicial após o sucesso
    http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}