package backends

import (
	"context"

	"github.com/seanpar203/go-api/internal/env"
)

type OTPBackend interface {
	SendSMS(ctx context.Context, phone, code string) error
	Validate(ctx context.Context, code string) error
}

type MockOTPBackend struct{}

func (b *MockOTPBackend) SendSMS(ctx context.Context, phone, code string) error {
	return nil
}

func (b *MockOTPBackend) Validate(ctx context.Context, code string) error {
	return nil
}

func GetOTPBackend() OTPBackend {
	switch env.APP_ENV {
	case "test":
		return &MockOTPBackend{}
	case "dev":
		return &MockOTPBackend{}
	default:
		return &MockOTPBackend{}
	}
}
