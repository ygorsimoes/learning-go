package main

import "fmt"

func main() {

	// Começando com a seguinte slice:
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Anexe a ela o valor 52;
	slice = append(slice, 52)

	// Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
	slice = append(slice, 53, 54, 55)

	// Demonstre a slice
	fmt.Println(slice)

	//  Anexe a ela a seguinte slice:
	anexoSlice := []int{56, 57, 58, 59, 60}

	slice = append(slice, anexoSlice...)

	// Demonstre a slice
	fmt.Println(slice)
}
