package main

import (
	"fmt"
	"os"
)

func main() {
	numero := 10
	fmt.Print("Insira alguma coisa: ")

	// Entrada de dados
	_, err := fmt.Scan(&numero)

	//Trata o possível erro da entrada de dados
	if err != nil {
		// Sai do programa com o status diferente de 0
		os.Exit(-1)
	}

	// Imprime a variável em Decimal, Binário e Hexadecimal
	fmt.Printf("Decimal: %d\n", numero)
	fmt.Printf("Binário: %b\n", numero)
	fmt.Printf("Hexadecimal: %#x\n", numero)
}