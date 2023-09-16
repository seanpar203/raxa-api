package api

import (
	"context"

	"github.com/seanpar203/go-api/internal/api/oas"
)

func (s *API) GetUserByID(ctx context.Context, params oas.GetUserByIDParams) (oas.GetUserByIDRes, error) {
	return &oas.User{}, nil
}
