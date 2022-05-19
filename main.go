package main

import (
	"api-go-rest-teste/DAO"
	"api-go-rest-teste/rotas"
	"fmt"
)

func main() {

	DAO.ConnectDB()
	fmt.Println("Iniciando o servidor da API Rest...")
	rotas.HandleRequest()
}
