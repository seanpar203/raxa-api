package api

import (
	"context"
	"net/url"

	"github.com/google/uuid"
	"github.com/ogen-go/ogen/validate"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/backends"
	"github.com/seanpar203/go-api/internal/models"
)

var (
	fileBackend = backends.GetFileBackend()
)

func Image(path string) oas.Image {

	fp := fileBackend.MustGet(context.Background(), path)

	u, _ := url.Parse(fp)

	return oas.Image(*u)
}

// Maps a user model to an oas.V1User
func V1User(user *models.User) *oas.V1User {
	return &oas.V1User{
		ID:          oas.UUID(uuid.MustParse(user.ID)),
		Name:        user.Name.String,
		Email:       user.Email,
		Birthday:    oas.Date(user.Birthday.Time),
		PhoneNumber: oas.E164PhoneNumber(user.PhoneNumber.String),
		Photo:       Image(user.Photo.String),
	}
}

// Creates a V1 CreateUserResponse
func V1LoginUserResponse(user *models.User, at *models.AccessToken, rt *models.RefreshToken) *oas.V1AuthLoginResponse {
	return &oas.V1AuthLoginResponse{
		User:         *V1User(user),
		AccessToken:  oas.UUID(uuid.MustParse(at.Token)),
		RefreshToken: oas.UUID(uuid.MustParse(rt.Token)),
	}
}

// V1AuthRefreshResponse creates a new V1AuthRefreshResponse object.
func V1AuthRefreshResponse(at *models.AccessToken) *oas.V1AuthRefreshResponse {
	return &oas.V1AuthRefreshResponse{
		AccessToken: oas.UUID(uuid.MustParse(at.Token)),
	}
}

// Builds a V1FieldErrors response from validation errors.
func V1FieldErrors(err *validate.Error) *oas.V1FieldErrors {
	var fieldErrors = &oas.V1FieldErrors{}
	var fieldErrMap = make(map[string]*oas.V1FieldError)

	for _, field := range err.Fields {

		name := field.Name
		err := field.Error.Error()
		fieldErr, ok := fieldErrMap[name]

		if !ok {
			fieldErrMap[name] = &oas.V1FieldError{
				Field:    name,
				Messages: []string{},
			}

			fieldErr = fieldErrMap[name]
		}

		fieldErr.Messages = append(fieldErr.Messages, err)
	}

	for _, field := range fieldErrMap {
		fieldErrors.Errors = append(fieldErrors.Errors, *field)
	}

	return fieldErrors
}
