package controllers

import (
	"api-go-rest-teste/DAO"
	"api-go-rest-teste/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home Page</h1>")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidadeList []models.Personalidade
	DAO.DB.Find(&personalidadeList)
	json.NewEncoder(w).Encode(personalidadeList)

}

func GetPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	DAO.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func AddPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	DAO.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	DAO.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	DAO.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	DAO.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
