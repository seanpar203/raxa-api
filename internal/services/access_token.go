package services

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/seanpar203/go-api/internal/models"
)

type accessToken struct{}

// CreateToken creates a new access token for the given user.
//
// ctx - The context for the function.
// user - The user for whom the access token is being created.
// Returns the newly created access token and any error encountered.
func (svc *accessToken) CreateToken(ctx context.Context, user *models.User) (*models.AccessToken, error) {
	token := &models.AccessToken{
		UserID: user.ID,
	}

	if err := token.InsertG(ctx, boil.Infer()); err != nil {
		return nil, err
	}

	return token, nil
}

// GetAccessTokenFromRequest retrieves an access token from the given token string.
//
// It takes a context.Context as the first parameter and a token string as the second parameter.
// It returns a pointer to models.AccessToken and an error.
func (svc *accessToken) GetByToken(ctx context.Context, token string) (*models.AccessToken, error) {
	return models.AccessTokens(Load(models.AccessTokenRels.User), models.AccessTokenWhere.Token.EQ(token)).OneG(ctx)
}
