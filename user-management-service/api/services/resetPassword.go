package services

import (
	"context"
	"fmt"
	"time"
	"user-management-service/api/repository"
	"user-management-service/api/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

func ResetPasswordService(ctx context.Context, q *repository.Queries, password string, token string) (string, error) {
	resetToken, err := q.GetPasswordResetToken(ctx, pgtype.Text{String: token, Valid: true})
	if err != nil || time.Now().After(resetToken.ExpiresAt.Time) {
		return "", fmt.Errorf("invalid or expired token")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", fmt.Errorf("error processing request")
	}

	user, err := q.GetUserById(ctx, resetToken.UserID)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	if err := utils.VerifyPassword(user.Password, password); err == nil {
		return "", fmt.Errorf("passwords cannot be same")
	}

	err = q.UpdateUserById(ctx, repository.UpdateUserByIdParams{
		ID:       pgtype.UUID{Bytes: user.ID.Bytes, Valid: true},
		Username: user.Username,
		Password: hashedPassword,
	})
	if err != nil {
		return "", fmt.Errorf("error updating password: %w", err)
	}
	err = q.DeletePasswordResetToken(ctx, resetToken.UserID)
	if err != nil {
		return "", fmt.Errorf("password could not be updated")
	}
	return "password updated successfully", nil
}
