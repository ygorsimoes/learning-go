package main

import "fmt"

func main() {

	// Slice
	numeros := []int{1, 2, 3, 4, 5}

	// Imprime o slice percorrendo os indices setados
	fmt.Println(numeros[0:len(numeros)])

	// Percorre o slice utilizando o for
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}
}