package main

import (
	"fmt"
	"os"
)

func main() {
	// Declara e atribui o valor 0 a variável
	var primeiroNumero int
	var segundoNumero int

	fmt.Print("Insira o primeiro número: ")
	_, err := fmt.Scan(&primeiroNumero)

	fmt.Print("Insira o segundo número: ")
	_, err = fmt.Scan(&segundoNumero)

	//Trata o possível erro da entrada de dados
	if err != nil {
		// Sai do programa com o status diferente de 0
		os.Exit(-1)
	}

	// Declara e atribui as variáveis com as expressões boolianas
	expressaoIgualdade := primeiroNumero == segundoNumero
	expressaoDiferenca := primeiroNumero != segundoNumero
	expressaoMenor := primeiroNumero < segundoNumero
	expressaoMenorOuIgual := primeiroNumero <= segundoNumero
	expressaoMaior := primeiroNumero > segundoNumero
	expressaoMaiorOuIgual := primeiroNumero >= segundoNumero

	// Quebra de linha
	fmt.Print("\n")

	// Imprime as expressões boolianas
	fmt.Println("(==) Expressão de Igualdade:       ", expressaoIgualdade)
	fmt.Println("(!=) Expressão de Diferença:       ", expressaoDiferenca)
	fmt.Println("(<)  Expressão de Menor:           ", expressaoMenor)
	fmt.Println("(<=) Expressão de Menor ou Igual:  ", expressaoMenorOuIgual)
	fmt.Println("(>)  Expressão de Maior:           ", expressaoMaior)
	fmt.Println("(>=) Expressão de Maior ou Igual:  ", expressaoMaiorOuIgual)
}