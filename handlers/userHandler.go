package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dinel13/lanjukang/dto"
	"github.com/dinel13/lanjukang/pkg/utilities"
)

// SignupHandler handles the signup request
func (m *Repository) SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user dto.UserReuqest
	err := json.NewDecoder(r.Body).Decode(&user)
	// err = errors.New("invalid request body")
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	// Validate the user data
	if user.Email == "" {
		utilities.WriteJsonError(w, errors.New("Please provide an email"), http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		http.Error(w, "Please provide a password", http.StatusBadRequest)
		return
	}

	// // Hash the password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }

	// // Store the user
	// user.Password = string(hashedPassword)
	// err = m.UserRepo.Store(user)
	// if err != nil {
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }

	// // Send a response
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)

	// Send a response
	utilities.WriteJson(w, http.StatusOK, user, "user")

}
