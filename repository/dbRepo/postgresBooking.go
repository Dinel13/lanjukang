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

	stmt := `INSERT INTO bookings (  user_id, service_id, owner_id,	start_at, end_at)
				VALUES ( $1, $2, $3, $4, $5)
				RETURNING id, user_id, service_id, owner_id,	start_at, end_at`

	row := m.DB.QueryRowContext(ctx, stmt,
		booking.UserId,
		booking.ServiceId,
		booking.OwnerId,
		booking.StartAt,
		booking.EndAt,
	)

	var bookingResponse models.BookingResponse
	err := row.Scan(
		&bookingResponse.Id,
		&bookingResponse.UserId,
		&bookingResponse.ServiceId,
		&bookingResponse.OwnerId,
		&bookingResponse.StartAt,
		&bookingResponse.EndAt,
	)
	if err != nil {
		return nil, err
	}

	return &bookingResponse, nil
}
