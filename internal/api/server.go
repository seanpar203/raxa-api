package api

import (
	"github.com/seanpar203/go-api/internal/api/oas"
)

type API struct {
	
}

func New() (*oas.Server, error) {
	return oas.NewServer(&API{})
}
