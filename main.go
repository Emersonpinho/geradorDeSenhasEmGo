package main

import (
	"fmt"
	"math/rand"
)

func main() {
	senha := gerarSenha(10)
	fmt.Println(senha)
}

func gerarSenha(lenght int) string {
	
	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()+?><:{}[]"

	senha := make([]byte, lenght)

	senha[0] = caracteres[5]
	senha[1] = caracteres[10]
	senha[2] = caracteres[4]
	senha[4] = caracteres[32]

	for i := range senha {
		senha[i] = caracteres[rand.Intn(len(caracteres))]
	}


	return string(senha)
}

// gerar uma senha aleatória: OK!
// gerar uma senha MAS com obrigatoriedade: letras maiusculas, numeros e etc