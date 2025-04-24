package utils

// Importa os pacotes necessários para a função de criptografia
import (
	"crypto/md5"   // Usado para gerar hashes MD5
	"encoding/hex" // Usado para converter o hash em uma string hexadecimal
)

// Encrypt recebe uma string e retorna o hash MD5 correspondente
func Encrypt(data string) (hash string) {
	// Cria um novo objeto hasher para gerar o hash MD5
	hasher := md5.New()
	// Escreve os dados no hasher
	hasher.Write([]byte(data))
	// Gera o hash e o converte para uma string hexadecimal
	hash = hex.EncodeToString(hasher.Sum(nil))
	// Retorna o hash gerado
	return
}
