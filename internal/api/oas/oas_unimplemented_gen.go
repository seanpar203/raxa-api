// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// V1UsersCreate implements V1_Users_Create operation.
//
// Creates and returns a new user.
//
// POST /v1/users
func (UnimplementedHandler) V1UsersCreate(ctx context.Context, req *V1UsersCreateReq) (r V1UsersCreateRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1UsersMe implements V1_Users_Me operation.
//
// Gets the current user.
//
// GET /v1/users/me
func (UnimplementedHandler) V1UsersMe(ctx context.Context) (r V1UsersMeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1UsersMeUpdate implements V1_Users_Me_Update operation.
//
// Updates the current user.
//
// PATCH /v1/users/me
func (UnimplementedHandler) V1UsersMeUpdate(ctx context.Context, req OptV1UsersMeUpdateReq) (r V1UsersMeUpdateRes, _ error) {
	return r, ht.ErrNotImplemented
}
