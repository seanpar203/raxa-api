package services

import (
	"context"

	"github.com/seanpar203/go-api/internal/backends"
	"github.com/seanpar203/go-api/internal/models"
)

type OTPChannel string

const (
	OTPChannelEmail OTPChannel = "email"
	OTPChannelSMS   OTPChannel = "sms"
)

type otp struct {
	otp backends.OTPBackend
}

func (svc *otp) Send(ctx context.Context, user *models.User, channel OTPChannel) error {

	if channel == OTPChannelEmail && len(user.PhoneNumber.String) == 0 {
		return nil
	}

	if err := svc.otp.SendSMS(ctx, user.PhoneNumber.String, ""); err != nil {
		return err
	}

	return nil
}
