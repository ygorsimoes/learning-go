package main

import "fmt"

func main() {

	mepe := map[string][]string{
		"dasilva_guriaestranhadostrangerthings": {
			"desaparecer", "ser assustadora", "raspar o cabelo",
		},
		"senna_ayrton": {
			"andar de jetski", "ser milionário", "falar com sotaque de paulista mano",
		},
		"pike_rob": {
			"criar linguagens de programação", "usar uns ternos muito loucos",
		},
	}

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
