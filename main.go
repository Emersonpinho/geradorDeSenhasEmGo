package main

import "fmt"

func main() {
	senha := gerarSenha()
	fmt.Println(senha)
}

func gerarSenha() string {
	
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"

	return "senha"
}

// gerar uma senha aleatória
// gerar uma senha MAS com obrigatoriedade: letras maiusculas, numeros e etc