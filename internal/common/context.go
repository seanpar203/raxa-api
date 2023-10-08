package common

import (
	"context"
	"errors"

	"github.com/rs/zerolog"

	"github.com/seanpar203/go-api/internal/models"
)

type ctxUserKey struct{}
type ctxLoggerKey struct{}

// UserFromContext retrieves the User object from the given context.
//
// It takes a context.Context as a parameter and returns a *models.User object
// and an error. If the User object is found in the context, it is returned.
// Otherwise, an error with the message "user not in context" is returned.
func UserFromContext(ctx context.Context) (*models.User, error) {

	if u, ok := ctx.Value(ctxUserKey{}).(*models.User); ok {
		return u, nil
	}

	return nil, errors.New("user not in context")
}

// UserWithContext adds the provided user to the context.
//
// It takes a context and a user as parameters.
// It returns a new context with the user value added.
func UserWithContext(ctx context.Context, u *models.User) context.Context {
	return context.WithValue(ctx, ctxUserKey{}, u)
}

// LoggerFromContext returns the zerolog.Logger instance from the given context.
//
// It takes a context.Context object as a parameter.
// It returns a pointer to a zerolog.Logger object.
func LoggerFromContext(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
