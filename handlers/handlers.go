package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dinel13/lanjukang/config"
	"github.com/dinel13/lanjukang/db/driver"
	"github.com/dinel13/lanjukang/models"
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
	user := models.UserSignUp{
		FullName: "Dinel",
		Email:    "",
		Password: "",
	}

	// get token from header
	authorizationHeader := r.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	fmt.Println(tokenString)
	id, err := utilities.ParseToken(tokenString, m.App.JwtSecret)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	userId, err := m.DB.CreateUser(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userId)
	w.Write([]byte("Hello World!"))
}
