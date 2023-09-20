package api

import (
	"context"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/common"
)

// HandleBearerAuth handles the bearer authentication for the API.
//
// It takes a context, the operation name, and a bearer authentication token as parameters.
// It returns the modified context and an error.
func (api *API) HandleBearerAuth(ctx context.Context, operationName string, t oas.BearerAuth) (context.Context, error) {

	logger := common.GetAuthLogger().With().Str("access_token", t.Token).Logger()

	user, err := api.Svcs.User.GetUserFromAccessToken(ctx, t.Token)

	if err != nil {
		logger.Error().Err(err).Msg("authentication failure")
	} else {
		logger.Info().Str("user_id", user.ID).Msg("authentication success")
	}

	ctx = common.UserWithContext(ctx, user)

	return ctx, err
}
