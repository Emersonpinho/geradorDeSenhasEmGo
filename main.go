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
	
	minusculas := "abcdefghijklmnopqrstuvwxyz"
	maiusculas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numeros:= "0123456789"
	caracteresEspeciais := "!@#$%^&*()+?><:{}[]"
	caractereGrande := minusculas + maiusculas + numeros + caracteresEspeciais

	obrigatorio := []byte{
		maiusculas[rand.Intn(len(maiusculas))],
		numeros,
	}
	senha := make([]byte, lenght)

	for i := range senha {
		senha[i] = caractereGrande[rand.Intn(len(caractereGrande))]
	}

	return string(senha)
}

// gerar uma senha aleatória: OK!
// gerar uma senha MAS com obrigatoriedade: letras maiusculas, numeros e etc