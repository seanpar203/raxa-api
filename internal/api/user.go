package api

import (
	"context"

	"github.com/seanpar203/go-api/internal/api/oas"
)

func (s *API) V1GetUserByID(ctx context.Context, params oas.V1GetUserByIDParams) (oas.V1GetUserByIDRes, error) {
	return &oas.V1User{}, nil
}
