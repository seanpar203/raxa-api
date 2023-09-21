package services

import "errors"

var (
	// Refresh Token errors
	ErrRefreshTokenExpired = errors.New("refresh token expired")

	// Access Token errors
	ErrAccessTokenExpired      = errors.New("access token expired")
	ErrAccessTokenDoesNotExist = errors.New("access token does not exist")

	// User related errors
	ErrUnableToGetUser      = errors.New("unable to get user")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrUnableToCreateUser   = errors.New("unable to create user")
	ErrrUnableToUpdateUser  = errors.New("unable to update user")
	ErrInvalidEmailAddress  = errors.New("invalid email address")
	ErrUserPasswordTooShort = errors.New("password must be at least 8 characters")
)
