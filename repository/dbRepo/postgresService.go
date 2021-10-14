package dbrepo

import (
	"context"
	"time"

	"github.com/dinel13/lanjukang/models"
)

func (m *postgresDbRepo) CreateService(service models.ServiceRequest) (*models.ServicePostCreate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO services (name, price, image, type_id, capacity,
		 						location, description)
				VALUES ($1, $2, $3, $4, $5, $6, $7)
				RETURNING id, name, price, image, type_id, capacity,
							location, description`

	row := m.DB.QueryRowContext(ctx, stmt,
		service.Name,
		service.Price,
		service.Image,
		service.TypeId,
		service.Capacity,
		service.LocationId,
		service.Description,
	)

	var newServices models.ServicePostCreate

	err := row.Scan(
		&newServices.Id,
		&newServices.Name,
		&newServices.Price,
		&newServices.Image,
		&newServices.TypeId,
		&newServices.Capacity,
		&newServices.LocationId,
		&newServices.Description,
	)

	if err != nil {
		return nil, err
	}

	return &newServices, nil
}
