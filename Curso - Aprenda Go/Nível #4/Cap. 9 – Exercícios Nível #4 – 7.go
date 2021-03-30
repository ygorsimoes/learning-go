package main

import "fmt"

func main() {

	sliceMulti := [][]string{
		{
			"miu",
			"milton",
			"encher o saco",
		},
		{
			"mimi",
			"martha",
			"pedir comida",
		},
		{
			"meus alunos queridos",
			"que estudam bastante",
			"fazer os exercícios ninja",
		},
	}

	for _, v := range sliceMulti {
		fmt.Println(v)
	}

	fmt.Print("\n\n\n")

	for _, v := range sliceMulti {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}
}
