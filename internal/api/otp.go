package api

import (
	"context"

	"github.com/seanpar203/go-api/internal/api/oas"
)

func (api *API) V1OTPCodeSend(ctx context.Context, req *oas.V1OTPCodeSendReq) (oas.V1OTPCodeSendRes, error) {
	// user, _ := common.UserFromContext(ctx)

	return &oas.V1OTPCodeSendNoContent{}, nil
}

func (api *API) V1OTPCodeEnter(ctx context.Context, req *oas.V1OTPCodeEnterReq) (oas.V1OTPCodeEnterRes, error) {
	return &oas.V1OTPCodeEnterNoContent{}, nil
}
