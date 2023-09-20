package common

import (
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	log     *zerolog.Logger
	logOnce sync.Once
)

// getLogger returns a logger instance with customized configuration.
//
// It retrieves the value of the `GO_ENV` environment variable and based on its value, it initializes the output writer. If `GO_ENV` is not equal to "dev", the output writer is set to `os.Stderr`, otherwise it is set to a `zerolog.ConsoleWriter` with a specified time format.
//
// The logger is then created with the configured output writer and additional customizations such as log level, timestamp, caller information, and process ID. The logger is returned as a pointer to a `zerolog.Logger`.
func getLogger() *zerolog.Logger {

	logOnce.Do(func() {

		var output io.Writer = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}

		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		if GetEnv("APP_ENV", "dev") != "dev" {
			output = os.Stderr
		}

		level := GetEnvAsInt("LOG_LEVEL", 1)

		l := zerolog.New(output).
			Level(zerolog.Level(level)).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()

		log = &l
	})

	return log
}

// GetRequestLogger returns a logger for logging HTTP request information.
//
// It takes a pointer to an http.Request as a parameter.
// It returns a pointer to a zerolog.Logger.
func GetRequestLogger(r *http.Request) *zerolog.Logger {

	l := getLogger().With().
		Str("service", "api").
		Str("url", r.URL.RequestURI()).
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Str("user_agent", r.UserAgent()).
		Str("referer", r.Referer()).
		Str("proto", r.Proto).
		Str("ip", r.RemoteAddr).
		Logger()

	return &l
}

// GetAuthLogger returns the logger for the auth service.
//
// It does not take any parameters.
// It returns a pointer to a zerolog.Logger.
func GetAuthLogger() *zerolog.Logger {
	l := getLogger().With().Str("service", "auth").Logger()

	return &l
}

func init() {
	zerolog.DefaultContextLogger = getLogger()
}
