package main

import (
	"fmt"

	"github.com/miltonjacomini/go-rest-api/models"
	"github.com/miltonjacomini/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "História 1"},
		{Nome: "Nome 2", Historia: "História 2"},
	}

	fmt.Println("Iniciando o servidor Rest em Go")
	routes.HandleRequest()
}
