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
		 						location, description, start, destiny, date, time, distance, duration )
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13 , $14)
				RETURNING id, name, price, image, type_id, capacity, location, description,
								start, destiny, date, time, distance, duration`

	row := m.DB.QueryRowContext(ctx, stmt,
		service.Name,
		service.Price,
		service.Image,
		service.OwnerId,
		service.TypeId,
		service.Capacity,
		service.Location,
		service.Description,
		service.Start,
		service.Destiny,
		service.Date,
		service.Time,
		service.Distance,
		service.Duration,
	)

	var newServices models.ServicePostCreate

	err := row.Scan(
		&newServices.Id,
		&newServices.Name,
		&newServices.Price,
		&newServices.Image,
		&newServices.TypeId,
		&newServices.Capacity,
		&newServices.Location,
		&newServices.Description,
		&newServices.Start,
		&newServices.Destiny,
		&newServices.Date,
		&newServices.Time,
		&newServices.Distance,
		&newServices.Duration,
	)

	if err != nil {
		return nil, err
	}

	typeServiceQuery := `SELECT name FROM type_services WHERE id = $1`
	rowType := m.DB.QueryRowContext(ctx, typeServiceQuery, service.TypeId)
	err = rowType.Scan(&newServices.Type)
	if err != nil {
		return nil, err
	}

	return &newServices, nil
}

// GetServiceByID get service by id
func (m *postgresDbRepo) GetDetailServiceByID(id int) (*models.ServiceDetailResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT s.id, s.name, s.price, s.image, s.capacity, s.location, s.description, u.nick_name, u.id, t.name
				FROM services s
				JOIN users u ON s.owner_id = u.id
				JOIN type_services t ON s.type_id = t.id
				WHERE s.id = $1`

	row := m.DB.QueryRowContext(ctx, stmt, id)

	var service models.ServiceDetailResponse

	err := row.Scan(
		&service.Id,
		&service.Name,
		&service.Price,
		&service.Image,
		&service.Capacity,
		&service.Location,
		&service.Description,
		&service.Owner,
		&service.OwnerId,
		&service.Type,
	)

	if err != nil {
		return nil, err
	}

	return &service, nil
}

// GetSortDetailServiceByID get service by id
func (m *postgresDbRepo) GetSortDetailServiceByID(id int) (*models.ServiceSortDetailResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT owner_id, image FROM services WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, stmt, id)

	var service models.ServiceSortDetailResponse
	err := row.Scan(&service.OwnerId, &service.Image)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

// UpdateService update service
func (m *postgresDbRepo) UpdateService(id int, service models.ServiceUpdateRequest) (*models.ServicePostCreate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE services
				SET name = $1, price = $2, image = $3, type_id = $4, capacity = $5,
					location = $6, description = $7
				WHERE id = $8
				RETURNING id, name, price, image, type_id, capacity,	
							location, description`

	row := m.DB.QueryRowContext(ctx, stmt,
		service.Name,
		service.Price,
		service.Image,
		service.TypeId,
		service.Capacity,
		service.Location,
		service.Description,
		id,
	)

	var newServices models.ServicePostCreate

	err := row.Scan(
		&newServices.Id,
		&newServices.Name,
		&newServices.Price,
		&newServices.Image,
		&newServices.TypeId,
		&newServices.Capacity,
		&newServices.Location,
		&newServices.Description,
	)

	if err != nil {
		return nil, err
	}

	return &newServices, nil
}

// ListAllServices list all services
func (m *postgresDbRepo) ListAllServices(limit int) ([]models.ServiceResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT s.id, s.name, s.price, s.image, s.capacity, s.location, u.nick_name, u.id, t.name
				FROM services s
				LEFT JOIN type_services t ON s.type_id = t.id
				LEFT JOIN users u ON s.owner_id = u.id
				ORDER BY s.rating DESC
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
			&service.Location,
			&service.Owner,
			&service.OwnerId,
			&service.Type,
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

// DeleteService delete service
func (m *postgresDbRepo) DeleteService(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `DELETE FROM services WHERE id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	return nil
}

// ListPopularServices list popular services
func (m *postgresDbRepo) ListPopularServices(limit int) ([]models.ServiceResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT s.id, s.name, s.price, s.image, s.capacity, s.location, u.nick_name, u.id, t.name
				FROM services s
				LEFT JOIN type_services t ON s.type_id = t.id
				LEFT JOIN users u ON s.owner_id = u.id
				ORDER BY s.rating DESC
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
			&service.Location,
			&service.Owner,
			&service.OwnerId,
			&service.Type,
		)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}
