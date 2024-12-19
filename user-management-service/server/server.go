package server

import (
	"log"
	"net/http"
	"user-management-service/api/handlers"

	"github.com/gorilla/mux"
)

func StartServer(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods(("GET"))
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	r.HandleFunc("/resetpassword", handlers.ResetPasswordHandler).Methods("PUT")

	r.Use(loggingMiddleware)

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Endpoint hit: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
