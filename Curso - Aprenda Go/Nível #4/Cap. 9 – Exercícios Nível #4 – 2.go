package main

import "fmt"

func main() {

	array := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, v := range array {
		fmt.Println("Índice:", i, "- Valor:", v)
	}

	fmt.Printf("Tipo: %T", array)
}
