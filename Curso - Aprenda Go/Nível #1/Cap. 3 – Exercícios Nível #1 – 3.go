package main

import "fmt"

// Declaração de variáveis com package-level scope
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	/*
		Atribuição de todos os valores em uma variável
		utilizando o operador curto de declaração
	*/
	s := fmt.Sprintf("%v \n%v \n%v", x, y, z)

	// Imprime a variável contendo todos os valores
	fmt.Println(s)
}