package services

import (
	"context"
	"fmt"
	"time"
	"user-management-service/api/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ResetPasswordRequestService(ctx context.Context, q *repository.Queries, email string) (string, error) {
	user, err := q.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("error fetching user by email: %w", err)
	}
	//if email is found in database
	resetToken := uuid.New().String()
	expirationTime := time.Now().Add(30 * time.Minute)
	err = q.CreatePasswordResetToken(ctx, repository.CreatePasswordResetTokenParams{
		UserID:    pgtype.UUID{Bytes: user.ID.Bytes, Valid: true},
		Token:     pgtype.Text{String: resetToken, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: expirationTime, Valid: true},
	})

	if err != nil {
		return "", fmt.Errorf("error creating password reset token: %w", err)
	}

	return resetToken, nil

}
