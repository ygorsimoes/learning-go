package main

import "fmt"

func main() {

	// Comece com essa slice "x":
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Utilizando slicing e append, crie uma slice "y" que contenha os valores:
	sliceExtra := append(slice[:3], slice[6:]...)

	// Demonstração:
	fmt.Println(sliceExtra)

}
