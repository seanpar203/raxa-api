package api

import (
	"context"
	"testing"

	"github.com/ogen-go/ogen/middleware"

	"github.com/seanpar203/go-api/internal/common"
)

func TestRecoveryMiddleware(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		inputReq    middleware.Request
		shouldPanic bool
		inputNext   middleware.Next
	}{
		{
			name:     "No panic",
			inputReq: middleware.Request{
				// Initialize request with necessary fields
			},
			shouldPanic: false,
			inputNext: func(req middleware.Request) (middleware.Response, error) {
				// Implement the next function
				return middleware.Response{}, nil
			},
		},
		{
			name:     "Panic occurred",
			inputReq: middleware.Request{
				// Initialize request with necessary fields
			},
			shouldPanic: true,
			inputNext: func(req middleware.Request) (middleware.Response, error) {
				// Implement the next function that causes a panic

				// Simulate a panic
				panic("Something really bad happened...")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock logger
			logger, out := common.GetTestLogger()

			// Set up the request context with the mock logger
			ctx := logger.WithContext(context.Background())
			tt.inputReq.Context = ctx

			// Call the RecoveryMiddleware function with the input parameters
			RecoveryMiddleware()(tt.inputReq, tt.inputNext)

			if tt.shouldPanic && len(out.Bytes()) == 0 {
				t.Error("RecoveryMiddleware did not panic", tt.name)
			}
		})
	}
}
