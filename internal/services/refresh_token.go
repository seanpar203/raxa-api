package services

import (
	"context"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/seanpar203/go-api/internal/models"
)

type refreshToken struct{}

// CreateToken creates a refresh token for the given user.
//
// It takes in a context.Context object and a *models.User object as parameters.
// It returns a *models.RefreshToken object and an error.
func (svc *refreshToken) CreateToken(ctx context.Context, user *models.User) (*models.RefreshToken, error) {

	token := &models.RefreshToken{
		UserID:     user.ID,
		ValidUntil: time.Now().Add(time.Hour * 24 * 7),
	}

	if err := token.InsertG(ctx, boil.Infer()); err != nil {
		return nil, err
	}

	return token, nil
}

// GetByToken retrieves a refresh token based on its token value.
//
// ctx: the context.Context object for the function.
// token: the token string to search for.
// Returns a *models.RefreshToken object and an error if any.
func (svc *refreshToken) GetByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	rt, err := models.RefreshTokens(
		Load(models.AccessTokenRels.User),
		models.RefreshTokenWhere.Token.EQ(token)).
		OneG(ctx)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	if now.After(rt.ValidUntil) {
		return nil, ErrRefreshTokenExpired
	}

	return rt, nil
}
