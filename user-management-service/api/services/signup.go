package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"user-management-service/api/dto"
	"user-management-service/api/repository"
	"user-management-service/api/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func SingupService(ctx context.Context, q *repository.Queries, user dto.SignupDTO) (string, error) {
	//check if email is already present in db, if yes then fail
	_, err := q.GetUserByEmail(ctx, user.Email)
	fmt.Print("existing user->", err)
	if err == nil {
		return "", errors.New("email already registered")
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("error checking email: %w", err)
	}

	// hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	// create a param to persist in db
	createParams := repository.CreateUserViaEmailParams{
		ID:       pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
		Role:     user.Role,
	}

	_, err = q.CreateUserViaEmail(ctx, createParams)
	if err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}
	return "User created successfully", nil
}
