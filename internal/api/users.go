package api

import (
	"context"
	"time"

	"github.com/volatiletech/null/v8"

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

	user, err := api.Svcs.User.CreateUser(ctx, req.Email, req.Password)

	if err != nil {
		logger.Err(err).Msg("unable to create user")
		return ResUnableToCreateUser, nil
	}

	user.Name = null.StringFrom(req.Name)

	user, err = api.Svcs.User.UpdateUser(ctx, user)

	if err != nil {
		logger.Err(err).Msg("unable to set user name")
		return ResUnableToCreateUser, nil
	}

	at, err := api.Svcs.AccessToken.CreateToken(ctx, user)

	if err != nil {
		logger.Err(err).Msg("unable to create access token")
		return ResUnableToCreateUser, nil
	}

	rt, err := api.Svcs.RefreshToken.CreateToken(ctx, user)

	if err != nil {
		logger.Err(err).Msg("unable to create refresh token")
		return ResUnableToCreateUser, nil
	}

	logger.Info().Msg("user created")

	return V1LoginUserResponse(user, at, rt), nil
}

// V1UsersMe implements V1_Users_Me operation.
//
// Gets the current user.
//
// GET /v1/users/me
func (api *API) V1UsersMe(ctx context.Context) (oas.V1UsersMeRes, error) {
	user, _ := common.UserFromContext(ctx)

	return V1User(user), nil
}

// V1UsersMeUpdate implements V1_Users_Me_Update operation.
//
// Updates the user.
//
// PATCH /v1/users/me
func (api *API) V1UsersMeUpdate(ctx context.Context, req oas.OptV1UsersMeUpdateReq) (oas.V1UsersMeUpdateRes, error) {
	var err error
	logger := common.LoggerFromContext(ctx)
	user, _ := common.UserFromContext(ctx)

	if !req.IsSet() || (!req.Value.Name.IsSet() && !req.Value.Birthday.IsSet() && !req.Value.Photo.IsSet()) {
		logger.Info().Msg("empty request body")
		return V1User(user), nil
	}

	if req.Value.Name.IsSet() {
		user.Name = null.StringFrom(req.Value.Name.Value)
	}

	if req.Value.Birthday.IsSet() {
		user.Birthday = null.TimeFrom(time.Time(req.Value.Birthday.Value))
	}

	user, err = api.Svcs.User.UpdateUser(ctx, user)

	if err != nil {
		logger.Err(err).Msg("unable to update user")
		return ResUnableToUpdateUser, nil
	}

	if req.Value.Photo.IsSet() {
		file := req.Value.Photo.Value.File
		name := req.Value.Photo.Value.Name

		user, err = api.Svcs.User.SetPhoto(ctx, user, file, name)

		if err != nil {
			logger.Err(err).Msg("unable to set user photo")
			return ResUnableToUpdateUser, nil
		}
	}

	return V1User(user), nil
}

// V1UsersMeContactsCreate implements v1_Users_Me_Contacts_Create operation.
//
// Creates contacts for the current logged in user OR future.
//
// POST /v1/users/me/contacts
func (api *API) V1UsersMeContactsCreate(ctx context.Context, req oas.Contacts) (oas.V1UsersMeContactsCreateRes, error) {
	return &oas.V1UsersMeContactsCreateNoContent{}, nil
}
