package main

import (
	"fmt"
	// os
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

	switch {
	case fruta == "laranja":
		fmt.Println("Essa é a fruta!")
	default:
		fmt.Println("Essa não é a fruta!")
	}
}
