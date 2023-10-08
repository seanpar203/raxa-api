package api

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/common"
)

var (
	ErrAuthInvalidToken       = errors.New("invalid token")
	ErrAuthInvalidTokenFormat = errors.New("improperly formatted auth token")
)

// HandleBearerAuth handles the bearer authentication for the API.
//
// It takes a context, the operation name, and a bearer authentication token as parameters.
// It returns the modified context and an error.
func (api *API) HandleBearerAuth(ctx context.Context, operationName string, t oas.BearerAuth) (context.Context, error) {
	if _, err := uuid.Parse(t.Token); err != nil {
		return ctx, ErrAuthInvalidTokenFormat
	}

	logger := common.GetAuthLogger().With().Str("access_token", t.Token).Logger()

	user, err := api.Svcs.User.GetUserFromAccessToken(ctx, t.Token)

	if err != nil {
		logger.Error().Err(err).Msg("authentication failure")
		return ctx, ErrAuthInvalidToken
	}

	logger.Info().Str("user_id", user.ID).Msg("authentication success")

	ctx = common.UserWithContext(ctx, user)

	return ctx, err
}

// V1AuthLogin is a function that handles the login process for the API.
//
// It takes a context and a request object as parameters and returns a response object and an error.
func (api *API) V1AuthLogin(ctx context.Context, req *oas.V1AuthLoginReq) (oas.V1AuthLoginRes, error) {
	logger := common.LoggerFromContext(ctx)

	user, err := api.Svcs.User.GetByEmail(ctx, req.Email)

	if err != nil {
		logger.Err(err).Str("email", req.Email).Msg("user not found")
		return ResAuthLoginError, nil
	}

	logger = common.AddUserToLogger(logger, user)

	if !common.PasswordsMatch(user.Password, req.Password) {
		logger.Info().Msg("incorrect password")
		return ResAuthLoginError, nil
	}

	at, err := api.Svcs.AccessToken.CreateToken(ctx, user)

	if err != nil {
		logger.Info().Err(err).Msg("unable to create access token")
		return ResAuthLoginError, nil
	}

	rt, err := api.Svcs.RefreshToken.CreateToken(ctx, user)

	if err != nil {
		logger.Info().Err(err).Msg("unable to create refresh token")
		return ResAuthLoginError, nil
	}

	logger.Info().Msg("user logged in")
	return V1LoginUserResponse(user, at, rt), nil
}

func (api *API) V1AuthRefresh(ctx context.Context, req *oas.V1AuthRefreshReq) (oas.V1AuthRefreshRes, error) {
	logger := common.LoggerFromContext(ctx)

	user, err := api.Svcs.User.GetUserFromRefreshToken(ctx, uuid.UUID(req.RefreshToken).String())

	if err != nil {
		logger.Err(err).Msg("user not found")
		return ResAuthRefreshError, nil
	}

	logger = common.AddUserToLogger(logger, user)

	at, err := api.Svcs.AccessToken.CreateToken(ctx, user)

	if err != nil {
		logger.Info().Err(err).Msg("unable to create access token")
		return ResAuthLoginError, nil
	}

	logger.Info().Msg("user auth refreshed")

	return V1AuthRefreshResponse(at), nil
}
