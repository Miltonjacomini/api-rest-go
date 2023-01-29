package routes

import (
	"log"
	"net/http"

	"github.com/miltonjacomini/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.Personalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
