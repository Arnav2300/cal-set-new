package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"user-management-service/api/dto"
	"user-management-service/api/repository"
	"user-management-service/api/services"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "healthy",
		"message": "the service is up and running!",
	})
}

func LoginHandler(ctx context.Context, repo *repository.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var requestBody dto.LoginDTO
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Print("Error in login request body: ", err.Error())
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid request"})
			return
		}
		if requestBody.Email == "" || requestBody.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Print("Incomplete request body")
			json.NewEncoder(w).Encode(map[string]string{"message": "incomplete request"})
			return
		}
		token, err := services.LoginService(ctx, repo, requestBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    token,
			HttpOnly: true,
			Secure:   false, // to ensure it's sent only over HTTPS set this to true
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "login successful"})
	}
}

func SignupHandler(ctx context.Context, repo *repository.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var requestBody dto.SignupDTO
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Print("Error in signup request body: ", err.Error())
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid request"})
			return
		}
		if requestBody.Email == "" || requestBody.Password == "" || requestBody.Role == "" || requestBody.Username == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Print("Incomplete request body")
			json.NewEncoder(w).Encode(map[string]string{"message": "incomplete request"})
			return
		}
		message, err := services.SingupService(ctx, repo, requestBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": message})
	}
}

func ResetPasswordRequestHandler(ctx context.Context, repo *repository.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req struct {
			Email string `json:"email"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid request"})
			return
		}
		resetToken, err := services.ResetPasswordRequestService(ctx, repo, req.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
			return
		}
		resetLink := fmt.Sprintf("http://fonrtentlinkblahblah.com/reset-password?token=%s", resetToken)

		// TODO: send above link to user via email service
		fmt.Print(resetLink)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "check your inbox, password reset link will be valid for 30 minutes"})
	}
}

func ResetPasswordHandler(ctx context.Context, repo *repository.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req struct {
			Token    string `json:"token"`
			Password string `json:"new_password"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Token == "" || req.Password == "" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid request"})
			return
		}
		message, err := services.ResetPasswordService(ctx, repo, req.Password, req.Token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": message})
	}
}
