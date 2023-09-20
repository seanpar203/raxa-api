package api

import (
	"fmt"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/services"
)

type API struct {
	Svcs *services.Services
}

func New() (*oas.Server, error) {

	svcs, err := services.New(nil)

	if err != nil {
		return &oas.Server{}, fmt.Errorf("failed to create services: %w", err)
	}

	api := &API{Svcs: svcs}

	return oas.NewServer(api, api, oas.WithMiddleware(LoggerMiddleware()))
}
