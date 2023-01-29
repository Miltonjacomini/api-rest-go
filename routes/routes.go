package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miltonjacomini/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.Personalidades)
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadePorId)
	log.Fatal(http.ListenAndServe(":8000", r))
}
