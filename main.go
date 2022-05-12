package main

import (
	"api-go-rest-teste/models"
	"api-go-rest-teste/rotas"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Lilly", Historia: "A hacker dos anos 00s"},
		{Id: 2, Nome: "Joy", Historia: "A Professora revolucionária"},
	}
	fmt.Println("Iniciando o servidor da API Rest...")
	rotas.HandleRequest()
}
