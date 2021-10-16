package dbrepo

import (
	"context"
	"time"

	"github.com/dinel13/lanjukang/models"
)

// CReateBooking creates a new booking in the database
func (m *postgresDbRepo) CreateBooking(booking models.BookingRequest) (*models.BookingResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `INSERT INTO bookings (  user_id, service_id, owner_id, amount,	start_at, end_at)
				VALUES ( $1, $2, $3, $4, $5, $6)
				RETURNING id, user_id, service_id, owner_id, amount,	start_at, end_at`

	row := m.DB.QueryRowContext(ctx, stmt,
		booking.UserId,
		booking.ServiceId,
		booking.OwnerId,
		booking.Amount,
		booking.StartAt,
		booking.EndAt,
	)

	var bookingResponse models.BookingResponse
	err := row.Scan(
		&bookingResponse.Id,
		&bookingResponse.UserId,
		&bookingResponse.ServiceId,
		&bookingResponse.OwnerId,
		&bookingResponse.Amount,
		&bookingResponse.StartAt,
		&bookingResponse.EndAt,
	)
	if err != nil {
		return nil, err
	}

	return &bookingResponse, nil
}

// UpdateBooking updates a booking in the database
func (m *postgresDbRepo) UpdateBooking(booking models.BookingRequestUpdate) (*models.BookingResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `UPDATE bookings SET amount = $1, start_at = $2, end_at = $3 
				WHERE id = $4 AND user_id = $5
				RETURNING id, user_id, service_id, owner_id, amount, start_at, end_at`

	row := m.DB.QueryRowContext(ctx, stmt,
		booking.Amount,
		booking.StartAt,
		booking.EndAt,
		booking.Id,
		booking.UserId,
	)

	var bookingResponse models.BookingResponse
	err := row.Scan(
		&bookingResponse.Id,
		&bookingResponse.UserId,
		&bookingResponse.ServiceId,
		&bookingResponse.OwnerId,
		&bookingResponse.Amount,
		&bookingResponse.StartAt,
		&bookingResponse.EndAt,
	)
	if err != nil {
		return nil, err
	}

	return &bookingResponse, nil
}

// GetAllBookingSbYUserId returns all bookings of a user
func (m *postgresDbRepo) GetAllBookingSbYUserId(userId int) ([]models.BookingResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `SELECT id, user_id, service_id, owner_id, start_at, end_at
				FROM bookings
				WHERE user_id = $1`

	rows, err := m.DB.QueryContext(ctx, stmt, userId)
	if err != nil {
		return nil, err
	}

	var bookings []models.BookingResponse
	for rows.Next() {
		var booking models.BookingResponse
		err := rows.Scan(
			&booking.Id,
			&booking.UserId,
			&booking.ServiceId,
			&booking.OwnerId,
			&booking.StartAt,
			&booking.EndAt,
		)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// DEleteBooking deletes a booking from the database
func (m *postgresDbRepo) DeleteBooking(bookingId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `DELETE FROM bookings WHERE id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, bookingId)
	if err != nil {
		return err
	}

	return nil
}
