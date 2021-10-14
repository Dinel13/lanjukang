package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dinel13/lanjukang/models"
	"github.com/dinel13/lanjukang/pkg/utilities"
	"golang.org/x/crypto/bcrypt"
)

// SignupHandler handles the signup request
func (m *Repository) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserSignUp
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// Validate the user data
	if user.Email == "" {
		utilities.WriteJsonError(w, errors.New("please provide an email"), http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		utilities.WriteJsonError(w, errors.New("please provide password"), http.StatusBadRequest)
		return
	}
	if user.FullName == "" {
		utilities.WriteJsonError(w, errors.New("please provide fullname"), http.StatusBadRequest)
		return
	}

	// check if the user already exists
	existUser, err := m.DB.GetUserByEmail(user.Email)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if existUser != nil {
		utilities.WriteJsonError(w, errors.New("user already exits, please login"), http.StatusBadRequest)
		return
	}

	// // Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// change the password to hash
	user.Password = string(hashedPassword)

	NewUserId, err := m.DB.CreateUser(user)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	defaultRole := 0

	token, err := utilities.CreateToken(NewUserId, defaultRole, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	userResponse := models.UserResponse{
		Token: token,
		Name:  user.FullName,
	}
	utilities.WriteJson(w, http.StatusOK, userResponse, "user")
}
