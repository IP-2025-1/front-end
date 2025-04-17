package handlers

import (
	"fmt"
	"front-end/app/utils"
	"net/http"
	"text/template"
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

	user, err := utils.GetUserByEmail(email)
	if err != nil {
		http.Error(response, "Erro ao buscar informações do usuário", http.StatusInternalServerError)
		return
	}

	// Carrega o template do perfil
	tmpl, err := template.ParseFiles("static/profile.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	// Renderiza o template com os dados do usuário
	err = tmpl.Execute(response, user)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, "/profile.html", http.StatusOK)
}
