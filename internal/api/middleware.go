package api

import (
	"time"

	"github.com/ogen-go/ogen/middleware"

	"github.com/seanpar203/go-api/internal/common"
)

// Attaches a logger to our request context.
func LoggerMiddleware() middleware.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {

		start := time.Now()

		logger := common.GetRequestLogger(req.Raw)

		logger.Info().Msg("request started")

		req.Context = logger.WithContext(req.Context)

		defer func() {
			logger.Info().
				Dur("elapsed_ms", time.Since(start)).
				Msg("request completed")
		}()

		return next(req)

	}
}
