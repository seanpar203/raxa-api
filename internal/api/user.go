package api

import (
	"context"

	"github.com/seanpar203/go-api/internal/api/oas"
)

// V1CreateSignupUser implements V1_Create_Signup_User operation.
//
// Creates a signup user that will later be converted into an actual user.
//
// POST /v1/signup
func (api *API) V1CreateSignupUser(ctx context.Context, req *oas.V1CreateSignupUserReq) (*oas.V1SignupUser, error) {
	return &oas.V1SignupUser{}, nil
}

// V1GetUserByID implements v1_Get_User_By_ID operation.
//
// Returns a single user.
//
// GET /v1/users/{id}
func (api *API) V1GetUserByID(ctx context.Context, params oas.V1GetUserByIDParams) (oas.V1GetUserByIDRes, error) {
	return &oas.V1GetUserByIDDef{}, nil
}

// V1GetUserList implements v1_Get_User_List operation.
//
// Returns a single user.
//
// GET /v1/users
func (api *API) V1GetUserList(ctx context.Context) (oas.V1GetUserListRes, error) {
	return &oas.V1GetUserListDef{}, nil
}
