package handlers

import (
	"fmt"
	"servidorHTTP/app/utils"
	"net/http"
	"text/template"
)

func UpdateAccountHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os valores do formulário
	username := request.FormValue("username")
	email := request.FormValue("email")
	bornDate := request.FormValue("bornDate")
	newPassword := request.FormValue("newPassword")
	currentPassword := request.FormValue("currentPassword")
	currentEmail := request.FormValue("currentEmail")

	fmt.Println(currentEmail, currentPassword)

	// Verifica se a senha atual está correta
	encryptedCurrentPassword := utils.Encrypt(currentPassword)
	isValidUser, err := utils.ValidateUser(currentEmail, encryptedCurrentPassword)
	if err != nil || !isValidUser {
		http.Error(response, "Senha atual ou email inválidos", http.StatusUnauthorized)
		return
	}

	// Cria um mapa para armazenar os campos a serem atualizados
	updates := make(map[string]string)

	if username != "" {
		updates["username"] = username
	}
	if email != "" {
		updates["email"] = email
	}
	if bornDate != "" {
		updates["born_date"] = bornDate
	}
	if newPassword != "" {
		updates["password"] = utils.Encrypt(newPassword) // Criptografa a nova senha
	}

	// Atualiza os campos informados no banco de dados
	err = utils.UpdateUser(currentEmail, updates)
	if err != nil {
		http.Error(response, "Erro ao atualizar os dados no banco de dados", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	var consultationEmail string

	if email == "" {
		consultationEmail = currentEmail
	} else {
		consultationEmail = email
	}

	user, err := utils.GetUserByEmail(consultationEmail)
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

	// Redireciona para a página de perfil ou outra página de sucesso
	http.Redirect(response, request, "/profile.html", http.StatusOK)
}
