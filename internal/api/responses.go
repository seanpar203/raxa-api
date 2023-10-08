package api

import (
	"net/http"

	"github.com/seanpar203/go-api/internal/api/oas"
)

// Our generic responses for our operations.
//
// Internally we're logging what really happened limiting external knowledge.
var (
	ResUnableToCreateUser = &oas.V1ErrorResponseStatusCode{
		StatusCode: http.StatusBadRequest,
		Response:   oas.NewV1ErrorMessageV1ErrorResponse(oas.V1ErrorMessage{Message: "unable to create user"}),
	}

	ResUnableToUpdateUser = &oas.V1ErrorResponseStatusCode{
		StatusCode: http.StatusBadRequest,
		Response:   oas.NewV1ErrorMessageV1ErrorResponse(oas.V1ErrorMessage{Message: "unable to update user"}),
	}
	ResUnableToProcessRequest = &oas.V1ErrorResponseStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: oas.NewV1ErrorMessageV1ErrorResponse(oas.V1ErrorMessage{
			Message: "unable to process request",
		}),
	}
	ResRequestTimout = &oas.V1ErrorResponseStatusCode{
		StatusCode: http.StatusRequestTimeout,
		Response:   oas.NewV1ErrorMessageV1ErrorResponse(oas.V1ErrorMessage{Message: "request timed out"}),
	}

	ResAuthLoginError = &oas.V1ErrorResponseStatusCode{
		StatusCode: http.StatusBadRequest,
		Response:   oas.NewV1ErrorMessageV1ErrorResponse(oas.V1ErrorMessage{Message: "unable to login user"}),
	}

	ResAuthRefreshError = &oas.V1ErrorResponseStatusCode{
		StatusCode: http.StatusBadRequest,
		Response:   oas.NewV1ErrorMessageV1ErrorResponse(oas.V1ErrorMessage{Message: "unable to refresh authentication"}),
	}
)
