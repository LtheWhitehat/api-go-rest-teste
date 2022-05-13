package rotas

import (
	"api-go-rest-teste/controllers"
	"api-go-rest-teste/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/TodasPersonalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/GetPersonalidade/{id}", controllers.GetPersonalidade).Methods("Get")
	r.HandleFunc("/delPersonalidade/{id}", controllers.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/AddPersonalidade", controllers.AddPersonalidade).Methods("Post")
	r.HandleFunc("/EditPersonalidade/{id}", controllers.EditPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}))(r)))
}
