package main

import (
	"fmt"
	"servidorHTTP/app/handlers"
	"servidorHTTP/app/utils"
	"log"
	"net"
	"net/http"
)

func main() {
	// Conecta ao banco de dados
	utils.ConnectToDB()

	// Configuração do servidor de arquivos estáticos
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/updateAccount", handlers.UpdateAccountHandler)
	http.HandleFunc("/deleteAccount", handlers.DeleteAccountHandler)

	// Obtém o IP local da máquina
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	var localIP string
	for _, addr := range addrs {
		// Verifica se o endereço é IPv4 e não é de loopback
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			localIP = ipNet.IP.String()
			break
		}
	}

	port := "3000"
	if localIP == "" {
		localIP = "127.0.0.1" // Fallback para localhost
	}

	// Exibe o IP e a porta no terminal
	fmt.Printf("Servidor rodando em: http://%s:%s/\n", localIP, port)

	// Inicia o servidor na porta 3000
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Fatal(err)
	}
}
