package services

import (
	"context"
	"errors"
	"fmt"
	"user-management-service/api/dto"
	"user-management-service/api/repository"
	"user-management-service/api/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

func SingupService(ctx context.Context, q *repository.Queries, user dto.SignupDTO) (string, error) {
	//check if email is already present in db, if yes then fail
	existingUser, err := q.GetUserByEmail(ctx, pgtype.Text{String: user.Email, Valid: true})
	if err == nil && existingUser.Email.Valid {
		return "", errors.New("email already registered")
	}
	if err != nil && err.Error() != "no rows in result set" {
		return "", fmt.Errorf("error checking email: %w", err)
	}

	// hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	// create a param to persist in db
	createParams := repository.CreateUserViaEmailParams{
		Email:    pgtype.Text{String: user.Email, Valid: true},
		Username: user.Username,
		Password: pgtype.Text{String: hashedPassword, Valid: true},
		Role:     user.Role,
	}

	_, err = q.CreateUserViaEmail(ctx, createParams)
	if err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}
	return "User created successfully", nil
}
