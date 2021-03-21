package main

import "fmt"

func main() {

	// Loop que percorre do número 33 ao 122
	for i := 33; i <= 122; i++ {

		// Imprime o Loop em formato de ASCII
		fmt.Printf("%#U\n", i)
	}
}