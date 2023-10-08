package services

import (
	"context"
	"fmt"
	"io"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/seanpar203/go-api/internal/backends"
	"github.com/seanpar203/go-api/internal/common"
	"github.com/seanpar203/go-api/internal/models"
)

var userBlacklistColumns = boil.Blacklist("id", "password")

type user struct {
	fb  backends.FileBackend
	otp backends.OTPBackend
}

// CreateUser creates a new user with the given email and password.
//
// Parameters:
// - ctx: the context.Context object for the request.
// - email: the email address of the user to be created.
// - password: the password for the user to be created.
//
// Returns:
// - *models.User: the newly created user object.
// - error: an error object if there was an issue creating the user.
func (svc *user) CreateUser(ctx context.Context, email string, password string) (*models.User, error) {

	exists, err := svc.DoesEmailExist(ctx, email)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrUserAlreadyExists
	}

	if valid := common.IsValidEmail(email); !valid {
		return nil, ErrInvalidEmailAddress
	}

	if len(password) < 8 {
		return nil, ErrUserPasswordTooShort
	}

	hash, err := common.HashPassword(password)

	if err != nil {
		return nil, ErrUnableToCreateUser
	}

	user := &models.User{
		Email:    email,
		Password: hash,
	}

	if err := user.InsertG(ctx, boil.Infer()); err != nil {
		return nil, ErrUnableToCreateUser
	}

	return user, nil
}

// UpdateUser updates a user in the database.
//
// ctx: the context.Context to be used for the database operation.
// user: the user model to be updated.
// Returns the updated user model and an error, if any.
func (svc *user) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {

	if _, err := user.UpdateG(ctx, userBlacklistColumns); err != nil {
		return user, ErrrUnableToUpdateUser
	}

	return user, nil
}

// Sets the photo of the user and returns the full qualified URL path.
func (svc *user) SetPhoto(ctx context.Context, user *models.User, file io.Reader, name string) (*models.User, error) {
	path := fmt.Sprintf("/users/%s/photos/%s", user.ID, common.UniqueFileName(name))

	if err := svc.fb.Save(ctx, file, path); err != nil {
		return user, err
	}

	user.Photo = null.StringFrom(path)

	if _, err := user.UpdateG(ctx, userBlacklistColumns); err != nil {
		return user, err
	}

	return user, nil
}

func (svc *user) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return models.Users(models.UserWhere.Email.EQ(email)).OneG(ctx)
}

// GetUserFromAccessToken retrieves a user using an access token.
//
// The function takes in the context and the access token as parameters.
// It returns a pointer to the User struct and an error if any.
func (svc *user) GetUserFromAccessToken(ctx context.Context, token string) (*models.User, error) {
	at, err := accessTokenSvc.GetByToken(ctx, token)

	if err != nil {
		return nil, err
	}

	return at.R.User, nil
}

// GetUserFromRefreshToken retrieves a user from a refresh token.
//
// ctx is the context to carry deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.
// token is the refresh token used to retrieve the user.
// *models.User is the user retrieved from the refresh token.
// error is any error that occurred while retrieving the user.
func (svc *user) GetUserFromRefreshToken(ctx context.Context, token string) (*models.User, error) {
	rt, err := refreshTokenSvc.GetByToken(ctx, token)

	if err != nil {
		return nil, err
	}

	return rt.R.User, nil
}

// DoesEmailExist checks if the given email exists in the database.
//
// ctx is the context.Context object for controlling the request lifecycle.
// email is the email address to check.
// The function returns a boolean value indicating whether the email exists or not.
// It also returns an error if there was an issue while querying the database.
func (svc *user) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	return models.Users(models.UserWhere.Email.EQ(email)).ExistsG(ctx)
}
