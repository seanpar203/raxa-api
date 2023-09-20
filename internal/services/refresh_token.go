package services

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/seanpar203/go-api/internal/models"
)

type refreshToken struct{}

func (svc *refreshToken) CreateToken(ctx context.Context, user *models.User) (*models.RefreshToken, error) {

	token := &models.RefreshToken{
		UserID: user.ID,
	}

	if err := token.InsertG(ctx, boil.Infer()); err != nil {
		return nil, err
	}

	return token, nil
}
