### README.md for geradorDeSenhasEmGo

````markdown name=README.md
# Gerador de Senhas em Go

![Go](https://img.shields.io/badge/Go-100%25-blue)

Este é um gerador de senhas simples escrito em Go. Ele permite gerar senhas aleatórias, garantindo que a senha contenha letras maiúsculas, números e caracteres especiais.

## Funcionalidades

- Geração de senhas aleatórias
- Inclusão obrigatória de letras maiúsculas e números
- Possibilidade de incluir caracteres especiais

## Como usar

### Pré-requisitos

- [Go](https://golang.org/dl/) instalado na sua máquina

### Passos

1. Clone o repositório:

```sh
git clone https://github.com/Emersonpinho/geradorDeSenhasEmGo.git
cd geradorDeSenhasEmGo
```

2. Execute o programa:

```sh
go run main.go
```

## Exemplo de Uso

Ao executar o programa, uma senha aleatória será gerada e exibida no console:

```sh
$ go run main.go
A1c!2
```

## Código-Fonte

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	senha := gerarSenha(5)
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
		numeros[rand.Intn(len(numeros))],
	}
	senha := make([]byte, lenght-len(obrigatorio) )

	for i := range senha {
		senha[i] = caractereGrande[rand.Intn(len(caractereGrande))]
	}

	senha = append(senha, obrigatorio...)

	rand.Shuffle(len(senha), func (i, j int)  {
		senha[i], senha[j] = senha[j], senha[i]
	})

	return string(senha)
}
```

## Contribuição

Se você deseja contribuir com este projeto, sinta-se à vontade para fazer um fork do repositório e enviar pull requests.

## Licença

Este projeto não possui uma licença específica.

---

Feito com ❤️ por [Emersonpinho](https://github.com/Emersonpinho)
````

Espero que você goste do README. Sinta-se à vontade para fazer ajustes conforme necessário!
