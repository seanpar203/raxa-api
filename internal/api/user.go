package api

import (
	"context"

	"github.com/volatiletech/null/v8"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/common"
	"github.com/seanpar203/go-api/internal/models"
)

func mapUserToV1User(user *models.User) *oas.V1User {
	return &oas.V1User{
		ID:    oas.UUID(user.ID),
		Name:  user.Name.String,
		Email: user.Email,
	}
}

// V1CreateSignupUser implements V1_Create_Signup_User operation.
//
// Creates a signup user that will later be converted into an actual user.
//
// POST /v1/signup
func (api *API) V1UsersCreate(ctx context.Context, req *oas.V1UsersCreateReq) (oas.V1UsersCreateRes, error) {

	logger := common.LoggerFromContext(ctx)

	var errRes = &oas.V1ErrorResponse{Message: "unable to create user"}

	user, err := api.Svcs.User.CreateUser(ctx, req.Email, req.Password)

	if err != nil {
		return errRes, err
	}

	at, err := api.Svcs.AccessToken.CreateToken(ctx, user)

	if err != nil {
		return errRes, err
	}

	rt, err := api.Svcs.RefreshToken.CreateToken(ctx, user)

	if err != nil {
		return errRes, err
	}

	logger.Info().Msg("user created")

	return &oas.V1CreateUserResponse{
		User:         *mapUserToV1User(user),
		AccessToken:  oas.UUID(at.Token),
		RefreshToken: oas.UUID(rt.Token),
	}, nil
}

// V1UsersMe implements V1_Users_Me operation.
//
// Gets the current user.
//
// GET /v1/users/me
func (api *API) V1UsersMe(ctx context.Context) (oas.V1UsersMeRes, error) {
	user, _ := common.UserFromContext(ctx)

	return mapUserToV1User(user), nil
}

// V1UsersMeUpdate implements V1_Users_Me_Update operation.
//
// Updates the user.
//
// PATCH /v1/users/me
func (api *API) V1UsersMeUpdate(ctx context.Context, req oas.OptV1UsersMeUpdateReq) (oas.V1UsersMeUpdateRes, error) {
	user, _ := common.UserFromContext(ctx)

	if !req.IsSet() || (!req.Value.Name.IsSet()) {
		return mapUserToV1User(user), nil
	}

	logger := common.LoggerFromContext(ctx)

	if req.Value.Name.IsSet() {
		user.Name = null.StringFrom(req.Value.Name.Value)
	}

	user, err := api.Svcs.User.UpdateUser(ctx, user)

	if err != nil {
		logger.Err(err).Msg("unable to update user")
	}

	return mapUserToV1User(user), nil
}
