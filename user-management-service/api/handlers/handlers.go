package handlers

import (
	"encoding/json"
	"net/http"
	"user-management-service/api/repository"
	"user-management-service/api/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"login"}`))
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var user repository.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	if user.Email.String == "" || user.Password.String == "" || user.Role == "" || user.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"request is incomplete"}`))
		return
	}
	services.SingupService(user)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"signup"}`))
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message":"reset password"}`))
}
