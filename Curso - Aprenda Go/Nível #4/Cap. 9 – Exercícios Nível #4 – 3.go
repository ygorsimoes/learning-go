package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
	fmt.Println("Índice:", slice[:3], "- Valor:", slice[:3])

	// Do quinto ao último item do slice (incluindo o último item!)
	fmt.Println("Índice:", slice[4:], "- Valor:", slice[4:])

	// Do segundo ao sétimo item do slice (incluindo o sétimo item!)
	fmt.Println("Índice:", slice[1:7], "- Valor:", slice[1:7])

	// Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	fmt.Println("Índice:", slice[2:8], "- Valor:", slice[2:8])

	// Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
	penultimate := len(slice) - 1
	fmt.Println("Índice:", slice[2:penultimate], "- Valor:", slice[2:penultimate])
}