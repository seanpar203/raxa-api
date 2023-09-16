package api

import (
	"context"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/models"
)

func (s *API) V1GetUserByID(ctx context.Context, params oas.V1GetUserByIDParams) (oas.V1GetUserByIDRes, error) {

	user, err := models.FindUserG(ctx, string(params.ID))

	if err != nil {
		return &oas.V1User{}, nil
	}

	return &oas.V1User{
		ID:   oas.NewOptString(user.ID),
		Name: oas.NewOptString(user.Email.String),
	}, nil
}

func (s *API) V1GetUserList(ctx context.Context) (oas.V1GetUserListRes, error) {
	return &oas.V1GetUserListOKApplicationJSON{}, nil
}
