package handlers

import (
	"net/http"

	"github.com/dinel13/lanjukang/config"
	"github.com/dinel13/lanjukang/db/driver"
	"github.com/dinel13/lanjukang/pkg/utilities"
	"github.com/dinel13/lanjukang/repository"
	dbrepo "github.com/dinel13/lanjukang/repository/dbRepo"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo returns a new Repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

//NewHandler set the repository for thde handler
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	type Home struct {
		Name    string
		Version string
		Status  string
	}

	response := Home{
		Name:    m.App.AppName,
		Version: m.App.AppVersion,
		Status:  "ok",
	}

	utilities.WriteJson(w, http.StatusOK, response, "respon")
}
