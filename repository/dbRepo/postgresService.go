package dbrepo

import (
	"context"
	"time"

	"github.com/dinel13/lanjukang/models"
)

// CReatteService create service to db
func (m *postgresDbRepo) CreateService(service models.ServiceRequest) (*models.ServicePostCreate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO services (name, price, image, owner_id, type_id, capacity,
		 						location, description)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				RETURNING id, name, price, image, type_id, capacity,
							location, description`

	row := m.DB.QueryRowContext(ctx, stmt,
		service.Name,
		service.Price,
		service.Image,
		service.OwnerId,
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

// GetServiceByID get service by id
func (m *postgresDbRepo) GetDetailServiceByID(id int) (*models.ServiceDetailResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT s.id, s.name, s.price, s.image, s.capacity, s.description, u.nick_name, u.id, t.name,  l.name
				FROM services s
				JOIN users u ON s.owner_id = u.id
				JOIN type_services t ON s.type_id = t.id
				JOIN locations l ON s.location = l.id
				WHERE s.id = $1`

	row := m.DB.QueryRowContext(ctx, stmt, id)

	var service models.ServiceDetailResponse

	err := row.Scan(
		&service.Id,
		&service.Name,
		&service.Price,
		&service.Image,
		&service.Capacity,
		&service.Description,
		// &service.Comments,
		&service.Owner,
		&service.OwnerId,
		&service.Type,
		&service.Location,
	)

	if err != nil {
		return nil, err
	}

	return &service, nil
}

// ListAllServices list all services
func (m *postgresDbRepo) ListAllServices(limit int) ([]models.ServiceResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT s.id, s.name, s.price, s.image, s.capacity, u.nick_name, u.id, t.name,  l.name
				FROM services s
				LEFT JOIN locations l ON s.location = l.id
				LEFT JOIN type_services t ON s.type_id = t.id
				LEFT JOIN users u ON s.owner_id = u.id
				LIMIT $1`

	rows, err := m.DB.QueryContext(ctx, stmt, limit)

	if err != nil {
		return nil, err
	}

	var services []models.ServiceResponse

	for rows.Next() {
		var service models.ServiceResponse

		err = rows.Scan(
			&service.Id,
			&service.Name,
			&service.Price,
			&service.Image,
			&service.Capacity,
			&service.Owner,
			&service.OwnerId,
			&service.Type,
			&service.Location,
		)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

// ListAllServicesByType list all services by type
func (m *postgresDbRepo) ListAllServicesByType(typeId int, limit int) ([]models.ServiceResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT s.id, s.name, s.price, s.image, s.capacity, u.nick_name, u.id, t.name,  l.name
				FROM services s
				LEFT JOIN locations l ON s.location = l.id
				LEFT JOIN type_services t ON s.type_id = t.id
				LEFT JOIN users u ON s.owner_id = u.id
				WHERE s.type_id = $1
				LIMIT $2`

	rows, err := m.DB.QueryContext(ctx, stmt, typeId, limit)

	if err != nil {
		return nil, err
	}

	var services []models.ServiceResponse

	for rows.Next() {
		var service models.ServiceResponse

		err = rows.Scan(
			&service.Id,
			&service.Name,
			&service.Price,
			&service.Image,
			&service.Capacity,
			&service.Owner,
			&service.OwnerId,
			&service.Type,
			&service.Location,
		)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

// ListAllServicesByLocation list all services by location
func (m *postgresDbRepo) ListAllServicesByLocation(locationId int, limit int) ([]models.ServiceResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT s.id, s.name, s.price, s.image, s.capacity, u.nick_name, u.id, t.name,  l.name
				FROM services s
				LEFT JOIN locations l ON s.location = l.id
				LEFT JOIN type_services t ON s.type_id = t.id
				LEFT JOIN users u ON s.owner_id = u.id
				WHERE s.location = $1
				LIMIT $2`

	rows, err := m.DB.QueryContext(ctx, stmt, locationId, limit)

	if err != nil {
		return nil, err
	}

	var services []models.ServiceResponse

	for rows.Next() {
		var service models.ServiceResponse

		err = rows.Scan(
			&service.Id,
			&service.Name,
			&service.Price,
			&service.Image,
			&service.Capacity,
			&service.Owner,
			&service.OwnerId,
			&service.Type,
			&service.Location,
		)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}
