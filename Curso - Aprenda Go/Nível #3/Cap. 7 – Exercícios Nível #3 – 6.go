package main

import (
	"fmt"
	//"os"
	"strings"
)

/*
	JOGO: Acerte a fruta!
*/

func main() {

	/*
		// Declara uma variável do tipo string
		var fruta string

		// Entrada de dados
		_, err := fmt.Scan(&fruta)

		// Trata o erro do método Scan()
		if err != nil {
			// Encerra o programa com o valor diferente de 0
			os.Exit(-1)
		}
	*/

	// Declara e atribui uma string a variável
	fruta := "cereja"

	// Converte a string para caixa baixa (minúsculo)
	fruta = strings.ToLower(fruta)

	// Condição que verifica qual é a fruta da variável.
	if fruta == "maçã" {
		fmt.Println("Errado, maçã não é a fruta! =)")
	} else if fruta == "amora" {
		fmt.Println("Errado, amora não é a fruta! :D")
	} else if fruta == "cereja" {
		fmt.Println("Errado, cereja não é a fruta! :O")
	} else if fruta == "laranja" {
		fmt.Println("Acertou! A fruta é laranja! :)")
	} else {
		fmt.Println("Errado,", fruta, "não é a fruta!")
	}
}