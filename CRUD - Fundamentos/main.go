package main

import (
	"net/http"

	"github.com/ygorsimoes/LearningGo/CRUD_-_Fundamentos/routes"
)

func main() {

	routes.CarregaRotas()

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic("[!] Error: " + err.Error())
	}
}
