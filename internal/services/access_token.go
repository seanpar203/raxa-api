package services

import (
	"context"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/seanpar203/go-api/internal/common"
	"github.com/seanpar203/go-api/internal/env"
	"github.com/seanpar203/go-api/internal/models"
)

type accessToken struct{}

// CreateToken creates a new access token for the given user.
//
// ctx - The context for the function.
// user - The user for whom the access token is being created.
// Returns the newly created access token and any error encountered.
func (svc *accessToken) CreateToken(ctx context.Context, user *models.User) (*models.AccessToken, error) {
	if err := svc.InvalidateUserAccessTokens(ctx, user); err != nil {
		return nil, err
	}

	token := &models.AccessToken{
		UserID:     user.ID,
		ValidUntil: time.Now().Add(env.ACCESS_TOKEN_EXPIRATION),
	}

	if err := token.InsertG(ctx, boil.Infer()); err != nil {
		return nil, err
	}

	return token, nil
}

// InvalidateUserAccessTokens invalidates all access tokens for a given user.
//
// ctx: The context in which the function is being executed.
// user: The user for whom the access tokens should be invalidated.
// Returns an error if any error occurred during the process.
func (svc *accessToken) InvalidateUserAccessTokens(ctx context.Context, user *models.User) error {

	tokens, err := models.AccessTokens(
		models.AccessTokenWhere.UserID.EQ(user.ID),
		models.AccessTokenWhere.ValidUntil.GTE(time.Now()),
	).AllG(ctx)

	if err != nil {
		return err
	}

	if _, err := tokens.UpdateAllG(ctx, models.M{"valid_until": time.Now()}); err != nil {
		return err
	}

	return nil

}

// GetAccessTokenFromRequest retrieves an access token from the given token string.
//
// It takes a context.Context as the first parameter and a token string as the second parameter.
// It returns a pointer to models.AccessToken and an error.
func (svc *accessToken) GetByToken(ctx context.Context, token string) (*models.AccessToken, error) {
	logger := common.LoggerFromContext(ctx)

	at, err := models.AccessTokens(
		Load(models.AccessTokenRels.User),
		models.AccessTokenWhere.Token.EQ(token)).
		OneG(ctx)

	if err != nil {
		logger.Err(err).Msg("unable to get access token")
		return nil, err
	}

	now := time.Now()

	if now.After(at.ValidUntil) {
		return nil, ErrAccessTokenExpired
	}

	return at, nil
}
