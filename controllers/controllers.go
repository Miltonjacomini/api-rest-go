package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/miltonjacomini/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}

func Personalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
