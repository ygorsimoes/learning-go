package main

import "fmt"

func main() {

	// Slice
	numeros := []int{1, 2, 3, 4, 5}

	fmt.Println(numeros[0:len(numeros)])

	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}
}
