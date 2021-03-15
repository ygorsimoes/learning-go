package main

import "fmt"

func main() {
	// Declara e atribui valores para as constantes tipadas e não tipadas
	const constanteTipada int = 10
	const constanteNaoTipada = 10

	// Imprime as constantes
	fmt.Println("Constante tipada:", constanteTipada)
	fmt.Println("Constante não tipada:", constanteNaoTipada)
}