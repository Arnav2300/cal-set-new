package dto

type ResetPasswordRequest struct {
	Email string
}

type ResetPassword struct {
	Email    string
	Password string
}
