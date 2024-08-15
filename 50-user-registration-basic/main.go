// main.go
package main

import (
	"log"
	"net/http"

	userhandlers "user-registration-basic/handlers"

	handlers "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//configure cors
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	r.HandleFunc("/users", userhandlers.CreateUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/users/{id}", userhandlers.GetUser).Methods("GET", "OPTIONS")
	r.HandleFunc("/users/{id}", userhandlers.UpdateUser).Methods("PUT", "OPTIONS")
	r.HandleFunc("/users/{id}", userhandlers.DeleteUser).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", cors(r)))
}
