package services

import (
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/seanpar203/go-api/internal/db"
)

var (
	userSvc         *user
	accessTokenSvc  *accessToken
	refreshTokenSvc *refreshToken
)

type Services struct {
	User         *user
	AccessToken  *accessToken
	RefreshToken *refreshToken
}

// New initializes a new Services struct.
//
// It takes a pointer to a sql.DB as a parameter and returns a pointer to Services and an error.
func New(_db *sql.DB) (*Services, error) {
	var err error

	if _db == nil {
		_db, err = db.Postgres()

		if err != nil {
			return &Services{}, fmt.Errorf("failed to connect to database: %w", err)
		}
	}

	boil.SetDB(_db)

	return &Services{
		User:         userSvc,
		AccessToken:  accessTokenSvc,
		RefreshToken: refreshTokenSvc,
	}, err
}

func init() {
	userSvc = &user{}
	accessTokenSvc = &accessToken{}
	refreshTokenSvc = &refreshToken{}
}
