package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatros bool
}

type sedan struct {
	veiculo
	modeloDeLuxo bool
}

func main() {

	carroDoTio := sedan{veiculo{4, "verde"}, true}
	carroDoVo := caminhonete{
		veiculo: veiculo{
			4,
			"azul",
		},
		tracaoNasQuatros: false,
	}

	fmt.Println(carroDoTio)
	fmt.Println(carroDoVo)

	fmt.Println(carroDoTio.cor)
	fmt.Println(carroDoVo.tracaoNasQuatros)

}