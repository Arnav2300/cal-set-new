package main

import (
	"log"
	"net/http"
	"user-management-service/api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/signup", handlers.Signup).Methods("POST")
	r.HandleFunc("/resetpassword", handlers.ResetPassword).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
