package rotas

import (
	"api-go-rest-teste/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/TodasPersonalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/TodasPersonalidades/{id}", controllers.GetPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8080", r))
}
