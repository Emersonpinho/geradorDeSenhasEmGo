package main

import "fmt"

func main() {
	senha := gerarSenha()
	fmt.Println(senha)
}

func gerarSenha() string {
	return "senha"
}

// gerar uma senha aleatória
// gerar uma senha MAS com obrigatoriedade: letras maiusculas, numeros e etc