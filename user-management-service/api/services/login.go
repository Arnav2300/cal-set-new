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

func LoginService(ctx context.Context, q *repository.Queries, credentials dto.LoginDTO) (string, error) {
	user, err := q.GetUserByEmail(ctx, pgtype.Text{String: credentials.Email, Valid: true})
	if err != nil {
		fmt.Printf("Error fetching user by email: %v", err)
		return "", errors.New("incorrect email or password")
	}
	err = utils.VerifyPassword(user.Password.String, credentials.Password)
	if err != nil {
		fmt.Printf("Password mismatch for user: %s", credentials.Email)
		return "", errors.New("incorrect email or password")
	}
	signedToken, err := utils.SignJwtToken(user.Role, credentials.Email, user.Username)
	if err != nil {
		fmt.Printf("Error signing JWT token: %v", err)
		return "", errors.New("we ran into a problem")
	}
	return signedToken, nil
}
