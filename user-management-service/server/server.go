package server

import (
	"context"
	"log"
	"net/http"
	"user-management-service/api/handlers"
	"user-management-service/api/repository"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func StartServer(port string) {
	pool, err := pgxpool.New(context.Background(), "postgres://root:root@127.0.0.1:5432/user_management?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}
	defer pool.Close()
	r := mux.NewRouter()
	repo := repository.New(pool)
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods(("GET"))
	r.HandleFunc("/login", handlers.LoginHandler(context.Background(), repo)).Methods("POST")
	r.HandleFunc("/signup", handlers.SignupHandler(context.Background(), repo)).Methods("POST")
	r.HandleFunc("/resetpassword", handlers.ResetPasswordRequestHandler(context.Background(), repo)).Methods("PUT")

	r.Use(loggingMiddleware)
	log.Print("DB connection established 🔗\n")
	log.Printf("Starting server on port %s 🚀\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Endpoint hit: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
