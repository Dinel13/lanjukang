package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dinel13/lanjukang/middleware"
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

	newUser, err := m.DB.CreateUser(user)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token, err := utilities.CreateToken(newUser.Id, newUser.Role, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	userResponse := models.UserResponse{
		Token: token,
		Name:  newUser.NickName,
	}
	utilities.WriteJson(w, http.StatusOK, userResponse, "user")
}

// LoginHandler handles the login request
func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin

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

	// check if the user already exists
	existUser, err := m.DB.GetUserByEmail(user.Email)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if existUser == nil {
		utilities.WriteJsonError(w, errors.New("user bot found"), http.StatusBadRequest)
		return
	}

	// check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password))
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token, err := utilities.CreateToken(existUser.Id, existUser.Role, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	userResponse := models.UserResponse{
		Token: token,
		Name:  existUser.NickName,
	}
	utilities.WriteJson(w, http.StatusOK, userResponse, "user")

}

// BecomeAdminHandler handles the become admin request
func (m *Repository) BecomeAdminHandler(w http.ResponseWriter, r *http.Request) {
	// cek if request have valid token
	id, role, err := middleware.ChecToken(w, r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// check if the user already become admin
	if role == 1 {
		utilities.WriteJsonError(w, errors.New("not allowed, you already admin"), http.StatusBadRequest)
		return
	}

	// update the user to become an admin
	err = m.DB.UpdateUserRole(id)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token, err := utilities.CreateToken(id, 1, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, token, "newToken")

}
