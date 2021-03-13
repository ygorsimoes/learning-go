package main

import "fmt"

// Cria um tipo onde o tipo subjacente é int
type exemploTipo int

// Cria uma variável com o tipo criado acima
var x exemploTipo

func main() {

	// Imprime o valor da variável
	fmt.Println(x)

	// Imprime o tipo da variável
	fmt.Printf("%T\n", x)

	// Atribui o valor 42 a variável
	x = 42

	// Imprime novamente o valor da variável
	fmt.Println(x)
}