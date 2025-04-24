# Servidor HTTP com GO

## Começando
Para configurar o ambiente de desenvolvimento, deve-se criar um arquivo .env (vide exemplo). No arquivo .env deve ser realizada a conexão com o banco de dados. Como sugestão, pode-se utilizar o docker como sistema para armazenar um banco de dados, para isso, crie um arquivo docker-compose.yml (vide exemplo) e execute o comando:

```docker compose up```

Como banco de dados, deve-se utilizar o PostgreSQL.

### Criando a tabela de usuários

O modelo para a criação da tabela de usuários é regido pelo seguinte SQL

```
    CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    born_date DATE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## Servidor HTTP e Rotas
No arquivo main.go é configurado o servidor HTTP, este servidor abrirá uma porta no IP do roteador utilizado; todos que estiverem conectados na mesma rede que a máquina que está executando o código poderão acessar a aplicação pelo IP informado no terminal. Neste arquivo também são definidas as rotas de processamento de dados do servidor. Ao final do arquivo, pode-se alterar a porta e o IP em que o código será hospedado.

---

## Estrutura de pastas do servidor
### Pasta static
A pasta static serve para armazenar os arquivos estáticos da aplicação, HTML e CSS, ou seja, o front-end da aplicação

### Pasta app
A pasta app serve para armazenar arquivos do servidor, handlers para lidar com formulários e requisições do front-end, utils serve para criar funções de auxílio para as handlers.