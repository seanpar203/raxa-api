package api

import (
	"errors"
	"time"

	"github.com/ogen-go/ogen/middleware"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/seanpar203/go-api/internal/common"
)

var Middlewares = []middleware.Middleware{
	RequestLoggerMiddleware(),
	AtomicRequestsMiddleware(),
}

// RequestLoggerMiddleware is a middleware function that logs information about incoming requests.
//
// # Creates a new RequestLogger
//
// Tries to get the user from the context and set it to the logger.
// Logs the request start and end with an elapsed.
func RequestLoggerMiddleware() middleware.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		start := time.Now()
		logger := common.GetRequestLogger(req.Raw)

		if user, err := common.UserFromContext(req.Context); err == nil {
			logger = common.AddUserToLogger(logger, user)
		}

		logger.Info().Msg("request started")

		defer func() {
			logger.Info().Dur("elapsed_ms", time.Since(start)).Msg("request completed")
		}()

		req.Context = logger.WithContext(req.Context)

		return next(req)
	}
}

// AtomicRequestsMiddleware is a middleware function that wraps a handler function in a transaction.
//
// It takes in a request object of type middleware.Request and a next function of type middleware.Next.
// The next function is responsible for calling the next middleware in the chain and returning the response and error.
//
// It returns a response object of type middleware.Response and an error object.
func AtomicRequestsMiddleware() middleware.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		logger := common.LoggerFromContext(req.Context)

		tx, _ := boil.BeginTx(req.Context, nil)

		res, err := next(req)

		if err := tx.Commit(); err != nil {
			logger.Err(err).Msg("unable to commit transaction")
			return res, errors.New("unable to process request")
		}

		return res, err
	}
}
