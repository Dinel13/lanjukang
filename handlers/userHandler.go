package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dinel13/lanjukang/middleware"
	"github.com/dinel13/lanjukang/models"
	"github.com/dinel13/lanjukang/pkg/utilities"
	"github.com/julienschmidt/httprouter"
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
		Id:    newUser.Id,
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
		Id:    existUser.Id,
		Token: token,
		Name:  existUser.NickName,
	}
	utilities.WriteJson(w, http.StatusOK, userResponse, "user")

}

// ForgetPasswordHandler handles the forget password request
func (m *Repository) ForgetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserByEmail
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

	// generate reset password token
	token, err := utilities.CreateResePasswordToken(existUser.Id, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// send link to reset password via email
	to := []string{user.Email}
	subject := "Reset Password"
	body := "Click the link to reset your password: " + m.App.Frontend + "/reset-password/" + token

	// make chanel as receiver for sending email
	mailEror := make(chan error)
	go func() {
		mailEror <- utilities.SendMail(to, subject, body)
	}()
	err = <-mailEror
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, "send link reset to your email", "user")

}

// ResetPasswordHandler handles the reset password request
func (m *Repository) ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserResetPasswordRequest

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// Validate the user data
	if user.Password == "" {
		utilities.WriteJsonError(w, errors.New("please provide password"), http.StatusBadRequest)
		return
	}
	if user.PasswordConfirm == "" {
		utilities.WriteJsonError(w, errors.New("please provide confirm password"), http.StatusBadRequest)
		return
	}
	if user.Password != user.PasswordConfirm {
		utilities.WriteJsonError(w, errors.New("password and confirm password not match"), http.StatusBadRequest)
		return
	}

	// cek if request have valid token
	id, err := middleware.CheckResetPasswordToken(r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// check if the user already exists
	// return
	existUser, err := m.DB.GetUserForResetPassword(id)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if existUser == nil {
		utilities.WriteJsonError(w, errors.New("user bot found"), http.StatusBadRequest)
		return
	}

	// check if old password not match with new password
	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password))
	if err == nil {
		utilities.WriteJsonError(w, errors.New("old password stil match with new password"), http.StatusBadRequest)
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// update user password
	existUser.Password = string(hashedPassword)
	updatedUser, err := m.DB.UpdateUserPasword(existUser.Id, existUser.Password)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// generate new token
	token, err := utilities.CreateToken(updatedUser.Id, updatedUser.Role, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	userResponse := models.UserResponse{
		Token: token,
		Name:  updatedUser.NickName,
	}

	utilities.WriteJson(w, http.StatusOK, userResponse, "user")
}

// UpdateUserImageHandler handles the update user info request
func (m *Repository) UpdateUserImageHandler(w http.ResponseWriter, r *http.Request) {
	// get user id from token
	id, _, err := middleware.ChecToken(r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// get user info from database
	user, err := m.DB.GetUserForUpdateImage(id)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// get data from request
	if err := r.ParseMultipartForm(1024); err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	uploadedImage, header, err := r.FormFile("image")
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}
	defer uploadedImage.Close()

	if uploadedImage == nil {
		utilities.WriteJsonError(w, errors.New("please provide image"), http.StatusBadRequest)
		return
	}

	filename, err := utilities.UploadedImage(uploadedImage, header, "user")
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}

	// update user info
	updatedUser, err := m.DB.UpdateUserImage(id, filename)

	if err != nil {
		utilities.DeleteImage(filename, "user")
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// delete old image if exist
	if user.Image != nil {
		oldImage := *user.Image // convert to string
		_ = utilities.DeleteImage(oldImage, "user")
	}

	utilities.WriteJson(w, http.StatusOK, updatedUser, "user")

}

// BecomeAdminHandler handles the become admin request
func (m *Repository) BecomeAdminHandler(w http.ResponseWriter, r *http.Request) {
	// cek if request have valid token
	id, role, err := middleware.ChecToken(r, m.App.JwtSecret)
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

// GetUserHandler handles the get user request
func (m *Repository) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// get the user id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// get the user
	user, err := m.DB.GetUserForOtherUser(id)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, user, "user")
}

// UpdateUserHandler handles the update user request
func (m *Repository) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	// Cek if request have valid token
	userId, _, err := middleware.ChecToken(r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if userId == 0 {
		utilities.WriteJsonError(w, errors.New("not alowed"), http.StatusBadRequest)
		return
	}

	// decode the request body
	var update models.UserUpdateRequset
	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// Update the user
	user, err := m.DB.UpdateUserProfile(userId, update)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, user, "user")
}
