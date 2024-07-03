package auth

import (
	"context"
	"todo/internal/models"
)

type AuthService interface {
	SignUp(ctx context.Context, name string, password string, limit int) (*models.User, error)
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (string, error)
}
