package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/miltonjacomini/go-rest-api/controllers"
	"github.com/miltonjacomini/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeInject)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.Personalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadePorId).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades", controllers.EditaPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
