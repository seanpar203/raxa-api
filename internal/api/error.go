package api

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/common"
)

var (
	ErrRequestTimedOut = errors.New("request timed out")
)

func ErrorCode(err error) int {
	var ve *validate.Error

	switch {
	case errors.Is(err, ErrAuthInvalidToken):
		return http.StatusUnauthorized
	case errors.Is(err, ErrAuthInvalidTokenFormat):
		return http.StatusUnauthorized
	case errors.Is(err, ErrRequestTimedOut):
		return http.StatusRequestTimeout
	case errors.Is(err, context.DeadlineExceeded):
		return http.StatusRequestTimeout
	case errors.Is(err, context.Canceled):
		return http.StatusRequestTimeout
	case errors.As(err, &ve):
		return http.StatusBadRequest
	default:
		return ogenerrors.ErrorCode(err)
	}
}

// NewError creates *V1ErrorResponseStatusCode from error returned by handler.
//
// Used for common default response.
func (api *API) NewError(ctx context.Context, err error) *oas.V1ErrorResponseStatusCode {
	logger := common.LoggerFromContext(ctx)

	logger.Err(err).Msg("error processing request")

	switch {
	case errors.Is(err, ErrAuthInvalidToken):
		return &oas.V1ErrorResponseStatusCode{
			StatusCode: ErrorCode(err),
		}
	default:
		return ResUnableToProcessRequest
	}
}

func ErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	logger := common.LoggerFromContext(ctx)

	code := ErrorCode(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	enc := jx.GetEncoder()

	var ve *validate.Error
	if errors.As(err, &ve) {
		logger.Err(err).Msg("validation error")
		res := V1FieldErrors(ve)
		res.Encode(enc)
	} else {
		logger.Info().Msg("Not a validation error")
	}

	if _, err := enc.WriteTo(w); err != nil {
		logger.Err(err).Msg("unable to write response")
	}
}
