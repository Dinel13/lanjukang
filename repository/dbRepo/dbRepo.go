package dbrepo

import (
	"database/sql"

	"github.com/dinel13/lanjukang/config"
	"github.com/dinel13/lanjukang/repository"
)

type postgresDbRepo struct {
	app *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, app *config.AppConfig) repository.DatabaseRepo {
	return &postgresDbRepo{
		app: app,
		DB:  conn,
	}
}
