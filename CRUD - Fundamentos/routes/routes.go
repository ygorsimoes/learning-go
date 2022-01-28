package routes

import (
	"net/http"

	"github.com/ygorsimoes/LearningGo/CRUD_-_Fundamentos/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
