# Servidor HTTP com GO

## Visão Geral
Este projeto é um servidor HTTP desenvolvido em GoLang que permite a criação, autenticação, atualização e exclusão de contas de usuários. Ele utiliza PostgreSQL como banco de dados e fornece uma interface web simples para interação com os usuários.

---

## Estrutura do Projeto
### Diretórios Principais:
- **`app/`**: Contém a lógica do servidor, incluindo handlers e utilitários.
  - **`handlers/`**: Lida com as requisições HTTP e processa os dados.
  - **`utils/`**: Contém funções auxiliares, como conexão ao banco de dados e manipulação de dados.
- **`static/`**: Contém os arquivos estáticos (HTML e CSS) que compõem o front-end da aplicação.
  - **`forms/`**: Formulários HTML para criar, atualizar, excluir contas e fazer login.
  - **`styles/`**: Arquivos CSS para estilização das páginas.

---

## Configuração do Ambiente

### Pré-requisitos
1. **GoLang**: Certifique-se de que o Go está instalado na sua máquina.
2. **PostgreSQL**: Banco de dados utilizado pelo projeto.
3. **Docker** (opcional): Para facilitar a configuração do banco de dados.

### Passos para Configuração
1. **Clone o repositório**:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd servidorHTTP
   ```

2. **Configure o arquivo `.env`**:
   Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
   ```
   DB_USER=<seu_usuario>
   DB_PASSWORD=<sua_senha>
   DB_NAME=<nome_do_banco>
   DB_HOST=<host_do_banco>
   DB_PORT=<porta_do_banco>
   ```

3. **Configuração do Banco de Dados**:
   - Se estiver usando Docker, utilize o arquivo `docker-compose.yml` para subir o banco de dados:
     ```bash
     docker compose up
     ```
   - Crie a tabela de usuários no banco de dados com o seguinte comando SQL:
     ```sql
     CREATE TABLE users (
         id SERIAL PRIMARY KEY,
         username VARCHAR(100) NOT NULL,
         email VARCHAR(150) NOT NULL UNIQUE,
         born_date DATE NOT NULL,
         password VARCHAR(255) NOT NULL,
         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
     );
     ```

4. **Instale as dependências**:
   Execute o comando abaixo para instalar as dependências do projeto:
   ```bash
   go mod tidy
   ```

---

## Executando o Projeto
1. **Inicie o servidor**:
   ```bash
   go run app/main.go
   ```

2. **Acesse a aplicação**:
   O servidor estará disponível no endereço exibido no terminal, geralmente algo como:
   ```
   http://127.0.0.1:3000/
   ```

---

## Funcionalidades
### Rotas Principais:
- **`/`**: Página inicial com links para os formulários.
- **`/forms/createAccount.html`**: Formulário para criar uma nova conta.
- **`/forms/login.html`**: Formulário para login.
- **`/forms/updateAccount.html`**: Formulário para atualizar os dados da conta.
- **`/forms/deleteAccount.html`**: Formulário para excluir uma conta.

### Handlers:
- **`FormHandler`**: Processa a criação de contas.
- **`LoginHandler`**: Valida as credenciais e exibe o perfil do usuário.
- **`UpdateAccountHandler`**: Atualiza os dados do usuário no banco de dados.
- **`DeleteAccountHandler`**: Remove a conta do usuário.

---

## Estrutura de Pastas
```
.env
docker-compose.yml
go.mod
app/
  main.go
  handlers/
    formHandler.go
    loginHandler.go
    updateAccountHandler.go
    deleteAccountHandler.go
  utils/
    connectToDB.go
    encrypt.go
    validateUser.go
    updateUser.go
    deleteUser.go
    getUserByEmail.go
static/
  index.html
  forms/
    createAccount.html
    login.html
    updateAccount.html
    deleteAccount.html
  styles/
    index.style.css
    createAccount.style.css
    login.style.css
    updateAccount.style.css
    deleteAccount.style.css
```

---

## Observações
- Certifique-se de que o banco de dados está configurado corretamente antes de iniciar o servidor.
- O projeto utiliza o driver `github.com/lib/pq` para conexão com o PostgreSQL.
- Para mais informações, consulte os comentários no código ou entre em contato com o desenvolvedor.
