package main

import "fmt"

// Declara e atribui constantes cujos valores sejam os próximos 4 anos utilizando iota
const (
	_ = 2021 + iota
	primeiroAno
	segundoAno
	terceiroAno
	quartoAno
)

func main() {

	// Imprime as constantes
	fmt.Println(primeiroAno, segundoAno, terceiroAno, quartoAno)
}