package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miltonjacomini/go-rest-api/database"
	"github.com/miltonjacomini/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}

func Personalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func PersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	if database.DB.Error != nil {
		json.NewEncoder(w).Encode(database.DB.Error.Error())
	} else {
		json.NewEncoder(w).Encode("Personalidade deletada com sucesso!")
	}
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	database.DB.First(personalidade.Id, &personalidade)
	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
