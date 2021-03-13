package main

import "fmt"

// Declaração de variáveis com package-level scope
var x int
var y string
var z bool

func main() {

	// Imprime as variáveis mostrando seus valores com quebras de linha
	fmt.Printf("%v \n%v \n%v", x, y, z)
}