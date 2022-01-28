package controllers

import (
	"html/template"
	"net/http"

	"github.com/ygorsimoes/LearningGo/CRUD_-_Fundamentos/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()

	err := temp.ExecuteTemplate(w, "Index", todosOsProdutos)
	if err != nil {
		panic("[!] Error: " + err.Error())
	}
}
