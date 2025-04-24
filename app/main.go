package main

// Importa os pacotes necessários para o funcionamento do programa
import (
    "fmt"                      // Usado para imprimir mensagens no terminal
    "servidorHTTP/app/handlers" // Importa os handlers definidos na aplicação
    "servidorHTTP/app/utils"    // Importa utilitários como conexão ao banco de dados
    "log"                      // Usado para registrar mensagens de erro ou log
    "net"                      // Usado para obter informações de rede, como IPs
    "net/http"                 // Usado para criar o servidor HTTP
)

func main() {
    // Conecta ao banco de dados utilizando a função definida em utils
    utils.ConnectToDB()

    // Cria um file server para servir arquivos estáticos da pasta "./static"
    fileserver := http.FileServer(http.Dir("./static"))

    // Define a rota raiz ("/") para servir os arquivos estáticos
    http.Handle("/", fileserver)

    // Define a rota "/hello" e associa ao handler HelloHandler
    http.HandleFunc("/hello", handlers.HelloHandler)

    // Define a rota "/form" e associa ao handler FormHandler
    http.HandleFunc("/form", handlers.FormHandler)

    // Define a rota "/login" e associa ao handler LoginHandler
    http.HandleFunc("/login", handlers.LoginHandler)

    // Define a rota "/updateAccount" e associa ao handler UpdateAccountHandler
    http.HandleFunc("/updateAccount", handlers.UpdateAccountHandler)

    // Define a rota "/deleteAccount" e associa ao handler DeleteAccountHandler
    http.HandleFunc("/deleteAccount", handlers.DeleteAccountHandler)

    // Obtém os endereços de rede disponíveis na máquina
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        // Encerra o programa caso ocorra um erro ao obter os endereços
        log.Fatal(err)
    }

    var localIP string
    // Itera sobre os endereços de rede para encontrar o IP local
    for _, addr := range addrs {
        // Verifica se o endereço é um IP válido (não é loopback e é IPv4)
        if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
            // Armazena o IP local encontrado
            localIP = ipNet.IP.String()
            break
        }
    }

    // Define a porta padrão do servidor como "3000"
    port := "3000"

    // Caso nenhum IP local seja encontrado, utiliza "127.0.0.1" como padrão
    if localIP == "" {
        localIP = "127.0.0.1"
    }

    // Imprime no terminal o endereço e a porta onde o servidor está rodando
    fmt.Printf("Servidor rodando em: http://%s:%s/\n", localIP, port)

    // Inicia o servidor HTTP na porta especificada e escuta conexões
    if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
        // Encerra o programa caso ocorra um erro ao iniciar o servidor
        log.Fatal(err)
    }
}