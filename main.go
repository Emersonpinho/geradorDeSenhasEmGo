package main

import (
	"fmt"
	"math/rand"
)

func main() {
	senha := gerarSenha()
	fmt.Println(senha)
}

func gerarSenha() string {
	
	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()+?><:{}[]"

	senha := make([]byte, 7)
	senha[0] = caracteres[5]
	senha[1] = caracteres[10]
	senha[2] = caracteres[4]
	senha[4] = caracteres[32]

	for i := range senha {
		senha[i] = caracteres[rand.Intn(len(caracteres))]
	}


	return string(senha)
}

// gerar uma senha aleatória
// gerar uma senha MAS com obrigatoriedade: letras maiusculas, numeros e etc