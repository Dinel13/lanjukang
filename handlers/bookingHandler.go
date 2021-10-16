package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dinel13/lanjukang/middleware"
	"github.com/dinel13/lanjukang/models"
	"github.com/dinel13/lanjukang/pkg/utilities"
)

// CreateBookingHandler is a handler for creating a booking
func (m *Repository) CreateBookingHandler(w http.ResponseWriter, r *http.Request) {

	// cek token
	userId, _, err := middleware.ChecToken(r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusUnauthorized)
		return
	}

	// get data from request
	var BookingRequest models.BookingRequestFrontend
	err = json.NewDecoder(r.Body).Decode(&BookingRequest)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	layoutFormat := "2006-01-02 15:04:05"
	startTime, err := time.Parse(layoutFormat, BookingRequest.StartAt)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	endTime, err := time.Parse(layoutFormat, BookingRequest.EndAt)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	booking := models.BookingRequest{
		UserId:    userId,
		ServiceId: BookingRequest.ServiceId,
		OwnerId:   BookingRequest.OwnerId,
		Amount:    BookingRequest.Amount,
		StartAt:   startTime,
		EndAt:     endTime,
	}

	createdBooking, err := m.DB.CreateBooking(booking)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusCreated, createdBooking, "booking")
}

// UpdateBookingHandler is a handler for updating a booking
func (m *Repository) UpdateBookingHandler(w http.ResponseWriter, r *http.Request) {

	// cek token
	userId, _, err := middleware.ChecToken(r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusUnauthorized)
		return
	}

	// get data from request
	var BookingUpdate models.BookingRequestUpdateFrontend
	err = json.NewDecoder(r.Body).Decode(&BookingUpdate)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// convert string to time
	layoutFormat := "2006-01-02 15:04:05"
	startTime, err := time.Parse(layoutFormat, BookingUpdate.StartAt)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	endTime, err := time.Parse(layoutFormat, BookingUpdate.EndAt)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	booking := models.BookingRequestUpdate{
		Id:      BookingUpdate.Id,
		UserId:  userId,
		Amount:  BookingUpdate.Amount,
		StartAt: startTime,
		EndAt:   endTime,
	}

	updatedBooking, err := m.DB.UpdateBooking(booking)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, updatedBooking, "booking")

}

// GetBookingByUserHandler is a handler for getting a booking
func (m *Repository) GetBookingByUserHandler(w http.ResponseWriter, r *http.Request) {

	// cek token
	userId, _, err := middleware.ChecToken(r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusUnauthorized)
		return
	}

	bookings, err := m.DB.GetAllBookingByUserId(userId)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, bookings, "bookings")
}
