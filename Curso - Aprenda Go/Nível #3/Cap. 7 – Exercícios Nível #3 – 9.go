package main

import "fmt"

func main() {

	esporteFavorito := "Basquete"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Não gosto de jogar bola!")
	case "baseball":
		fmt.Println("Não sei jogar, hehe!")
	case "basquete":
		fmt.Println("Aceitou! Gosto de fazer cestas.")
	}
}