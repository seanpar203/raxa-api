package api

import (
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/seanpar203/go-api/internal/api/oas"
	"github.com/seanpar203/go-api/internal/common"
	"github.com/seanpar203/go-api/internal/db"
)

type API struct {
	Port string
}

func New() (*oas.Server, error) {

	db, err := db.Postgres()

	if err != nil {
		log.Panic().Msg(err.Error())
	}

	boil.SetDB(db)

	return oas.NewServer(&API{
		Port: common.GetEnv("API_PORT", "8080"),
	})
}
