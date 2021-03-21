package main

import "fmt"

func main() {

	// Declara e atribui valores inteiros a variáveis
	anoNascimento := 2000
	anoAtual := 2021

	// Loop infinito
	for {

		// Condição para checar se o anoNascimento é menor ou igual ao anoAtual
		if anoNascimento <= anoAtual {

			// Imprime o loop
			fmt.Println("Ano:", anoNascimento)

			// Adiciona +1 a variável
			anoNascimento++

		} else {

			// Interrompe o loop caso a condição seja falso
			break
		}
	}
}