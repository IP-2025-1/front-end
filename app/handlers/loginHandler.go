package handlers

import (
	"fmt"
	"front-end/app/utils"
	"net/http"
)

func LoginHandler(response http.ResponseWriter, request *http.Request) {

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
	if err != nil {
		http.Error(response, "Erro ao validar o usuário", http.StatusInternalServerError)
		return
	}

	if !isValidUser {
		http.Error(response, "Credenciais inválidas", http.StatusUnauthorized)
		fmt.Println("Erro")
		return
	}

	http.Redirect(response, request, "/profile.html", http.StatusOK)
}
