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
	r.HandleFunc("/GetPersonalidade/{id}", controllers.GetPersonalidade).Methods("Get")
	r.HandleFunc("/delPersonalidade/{id}", controllers.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/AddPersonalidade", controllers.AddPersonalidade).Methods("Post")
	log.Fatal(http.ListenAndServe(":8080", r))
}
