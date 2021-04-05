package main

import "fmt"

type pessoa struct {
	nome              string
	sobrenome         string
	sorvetesFavoritos []string
}

func main() {

	pessoa1 := pessoa{
		nome:              "Renata",
		sobrenome:         "Silva",
		sorvetesFavoritos: []string{"Laranja", "Maracujá", "Acaí"},
	}

	pessoa2 := pessoa{
		nome:              "Júlio",
		sobrenome:         "Almeida",
		sorvetesFavoritos: []string{"Uva", "Maçã", "Cereja"},
	}

	fmt.Print("Sorvetes da Renata: ")
	for _, v := range pessoa1.sorvetesFavoritos {
		fmt.Print(v, " ")
	}

	fmt.Print("\nSorvetes do Júlio: ")
	for _, v := range pessoa2.sorvetesFavoritos {
		fmt.Print(v, " ")
	}
}
