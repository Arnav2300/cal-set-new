package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-management-service/api/dto"
	"user-management-service/api/repository"
	"user-management-service/api/services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
	ctx := r.Context()
	print(ctx)
	var repo *repository.Queries
	token, err := services.LoginService(ctx, repo, requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
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

func SignupHandler(w http.ResponseWriter, r *http.Request) {
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
	ctx := r.Context()
	print(ctx)
	var repo *repository.Queries
	message, err := services.SingupService(ctx, repo, requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message":"reset password"}`))
}
