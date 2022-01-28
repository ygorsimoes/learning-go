package models

import (
	"database/sql"

	"github.com/ygorsimoes/LearningGo/CRUD_-_Fundamentos/db"
)

type Produtos struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produtos {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic("[!] Error: " + err.Error())
	}

	p := Produtos{}
	var produtos []Produtos

	for selectDeTodosOsProdutos.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic("[!] Error: " + err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("[!] Error: " + err.Error())
		}
	}(db)

	return produtos
}
