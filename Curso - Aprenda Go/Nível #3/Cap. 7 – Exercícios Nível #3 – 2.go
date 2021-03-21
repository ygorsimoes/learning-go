package main

import "fmt"

func main() {

	// Loop que começa em 60 e termina em 90
	for i := 65; i <= 90; i++ {
		// Imprime o número do UNICODE
		fmt.Println(i)

		// Loop que começa em 1 e termina no 3
		for x := 1; x <= 3; x++ {
			// Imprime o UNICODE
			fmt.Printf("%#U\n", i)
		}

		// Quebra de linha
		fmt.Print("\n")
	}
}