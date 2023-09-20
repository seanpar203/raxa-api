package services

import (
	"context"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/seanpar203/go-api/internal/common"
	"github.com/seanpar203/go-api/internal/models"
)

type user struct{}

var (
	ErrUnableToGetUser      = errors.New("unable to get user")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrUnableToCreateUser   = errors.New("unable to create user")
	ErrInvalidEmailAddress  = errors.New("invalid email address")
	ErrUserPasswordTooShort = errors.New("password must be at least 8 characters")
)

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

// GetUserFromAccessToken retrieves a user using an access token.
//
// The function takes in the context and the access token as parameters.
// It returns a pointer to the User struct and an error if any.
func (svc *user) GetUserFromAccessToken(ctx context.Context, token string) (*models.User, error) {
	at, err := accessTokenSvc.GetByToken(ctx, token)

	if err != nil {
		return nil, ErrUnableToGetUser
	}

	return at.R.User, nil
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
