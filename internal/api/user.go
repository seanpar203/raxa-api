package api

import (
	"context"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/common"
)

// V1CreateSignupUser implements V1_Create_Signup_User operation.
//
// Creates a signup user that will later be converted into an actual user.
//
// POST /v1/signup
func (api *API) V1UsersCreate(ctx context.Context, req *oas.V1UsersCreateReq) (oas.V1UsersCreateRes, error) {

	logger := common.LoggerFromContext(ctx)

	var errRes = &oas.V1ErrorResponse{Message: "unable to create user"}

	tx, err := boil.BeginTx(ctx, nil)

	if err != nil {
		return errRes, err
	}

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

	if err := tx.Commit(); err != nil {
		return errRes, err
	}

	logger.Info().Msg("user created")

	return &oas.V1CreateUserResponse{
		AccessToken:  oas.UUID(at.Token),
		RefreshToken: oas.UUID(rt.Token),
		User: oas.V1User{
			ID:    oas.UUID(user.ID),
			Email: user.Email,
		},
	}, nil
}

// V1UsersMe implements V1_Users_Me operation.
//
// Gets the current user.
//
// GET /v1/users/me
func (api *API) V1UsersMe(ctx context.Context) (oas.V1UsersMeRes, error) {
	user, err := common.UserFromContext(ctx)

	if err != nil {
		return &oas.V1ErrorResponse{}, errors.New("unable to authenticate")
	}

	return &oas.V1User{
		ID:    oas.UUID(user.ID),
		Email: user.Email,
	}, nil
}

// V1UsersMeUpdate implements V1_Users_Me_Update operation.
//
// Updates the user.
//
// PATCH /v1/users/me
func (api *API) V1UsersMeUpdate(ctx context.Context) (oas.V1UsersMeUpdateRes, error) {
	return nil, nil
}
