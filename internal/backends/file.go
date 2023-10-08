package backends

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/seanpar203/go-api/internal/common"
	"github.com/seanpar203/go-api/internal/env"
)

type FileBackend interface {

	// Upload uploads a file to the server
	Save(ctx context.Context, file io.Reader, path string) error

	// Returns the qualified URL path to the file
	Get(ctx context.Context, path string) (string, error)

	// MustGet returns the path to the file or panics if the file does not exist
	MustGet(ctx context.Context, path string) string
}

type MockFileBackend struct{}

func (b *MockFileBackend) Save(ctx context.Context, file io.Reader, path string) error {
	logger := common.LoggerFromContext(ctx)
	logger.Info().Msg("file uploaded")
	return nil
}
func (b *MockFileBackend) Get(ctx context.Context, path string) (string, error) {
	logger := common.LoggerFromContext(ctx)
	logger.Info().Msg("file retrieved with Get")
	return "https://george-fx.github.io/apitex/users/01.png", nil
}

func (b *MockFileBackend) MustGet(ctx context.Context, path string) string {
	logger := common.LoggerFromContext(ctx)
	logger.Info().Msg("file retrieved with MustGet")
	return "https://george-fx.github.io/apitex/users/01.png"
}

type LocalStorageBackend struct{}

func (b *LocalStorageBackend) storageFP(path string) string {
	return filepath.Join(env.MEDIA_DIR, path)
}

func (b *LocalStorageBackend) mediaFP(path string) string {

	return fmt.Sprintf("http://%s:%s/%s/%s", env.SERVER_HOST, env.SERVER_PORT, strings.Replace(env.MEDIA_ENDPOINT, "/", "", -1), strings.TrimPrefix(path, "/"))
}

func (b *LocalStorageBackend) ensureMediaPathExists(path string) error {
	return os.MkdirAll(filepath.Dir(path), os.ModePerm)
}

// Save saves the file to the local media directory
func (b *LocalStorageBackend) Save(ctx context.Context, file io.Reader, path string) error {
	logger := common.LoggerFromContext(ctx)

	fp := b.storageFP(path)

	if err := b.ensureMediaPathExists(fp); err != nil {
		logger.Err(err).Msg("unable to create directory path")
		return err
	}

	w, err := os.Create(fp)

	if err != nil {
		return err
	}

	defer w.Close()

	if _, err := io.Copy(w, file); err != nil {
		logger.Err(err).Msg("unable to copy file")
		return err
	}

	logger.Info().Str("path", path).Msg("file uploaded")
	return nil
}
func (b *LocalStorageBackend) Get(ctx context.Context, path string) (string, error) {
	logger := common.LoggerFromContext(ctx)

	sfp := b.storageFP(path)

	if _, err := os.Stat(sfp); err != nil {
		logger.Err(err).Str("path", sfp).Msg("file does not exist")
		return "", err
	}

	return b.mediaFP(path), nil
}

func (b *LocalStorageBackend) MustGet(ctx context.Context, path string) string {
	logger := common.LoggerFromContext(ctx)

	fp, err := b.Get(ctx, path)

	if err != nil {
		logger.Panic().Msg(err.Error())
	}

	return fp
}

// Returns the appropriate Filebackend based on the
func GetFileBackend() FileBackend {
	switch env.APP_ENV {
	case "test":
		return &MockFileBackend{}
	case "dev":
		return &LocalStorageBackend{}
	default:
		return &MockFileBackend{}
	}
}
