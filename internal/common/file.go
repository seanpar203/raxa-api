package common

import (
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
)

func UniqueFileName(name string) string {
	ext := filepath.Ext(name)
	return fmt.Sprintf("%s%s", uuid.New().String(), ext)
}
