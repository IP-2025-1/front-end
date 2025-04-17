package handlers

import (
    "servidorHTTP/app/utils"
    "net/http"
)

func DeleteAccountHandler(response http.ResponseWriter, request *http.Request) {
    if request.Method != http.MethodPost {
        http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    email := request.FormValue("email")
    password := request.FormValue("password")

    // Criptografa a senha
    encryptedPassword := utils.Encrypt(password)

    // Verifica se o usuário existe no banco de dados
    isValidUser, err := utils.ValidateUser(email, encryptedPassword)
    if err != nil || !isValidUser {
        http.Error(response, "Credenciais inválidas", http.StatusUnauthorized)
        return
    }

    // Remove o usuário do banco de dados
    err = utils.DeleteUser(email)
    if err != nil {
        http.Error(response, "Erro ao apagar a conta", http.StatusInternalServerError)
        return
    }

    // Redireciona para a página inicial após o sucesso
    http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}