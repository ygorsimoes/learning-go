package main

import "fmt"

func main() {

	// Declara e atribui valores inteiros a variáveis
	anoNascimento := 2000
	anoAtual := 2021

	// Loop que começa em 2000 e termina em 2021
	for anoNascimento <= anoAtual {

		// Imprime o loop
		fmt.Println("Ano:", anoNascimento)

		// Adiciona +1 a variável
		anoNascimento++
	}
}