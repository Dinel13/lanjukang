package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dinel13/lanjukang/middleware"
	"github.com/dinel13/lanjukang/models"
	"github.com/dinel13/lanjukang/pkg/utilities"
	"github.com/julienschmidt/httprouter"
)

// CreateService handler for get service detail
func (m *Repository) CreateService(w http.ResponseWriter, r *http.Request) {

	// cek if request have valid token
	id, role, err := middleware.ChecToken(r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if role != 1 || id == 0 {
		utilities.WriteJsonError(w, errors.New("not allowed, become admin first"), http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	price := r.FormValue("price")
	typeId := r.FormValue("type_id")
	capacity := r.FormValue("capacity")
	location := r.FormValue("location")
	description := r.FormValue("description")
	start := r.FormValue("start")
	destiny := r.FormValue("destiny")
	time := r.FormValue("time")
	date := r.FormValue("date")
	distance := r.FormValue("distance")
	duration := r.FormValue("duration")

	// validate

	// end validate

	// convert to int
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	typeInt, err := strconv.Atoi(typeId)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	capacityInt, err := strconv.Atoi(capacity)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	distanceInt, err := strconv.Atoi(distance)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// CONVERT TO DATE
	startTime, err := utilities.ConvertToTimeHour(time)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	startDate, err := utilities.ConvertToDate(date)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// upload image
	uploadedImage, header, err := r.FormFile("image")
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	defer uploadedImage.Close()

	filename, err := utilities.UploadedImage(uploadedImage, header, "service")
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	service := models.ServiceRequest{
		Name:        name,
		Price:       priceInt,
		Image:       filename,
		OwnerId:     id,
		TypeId:      typeInt,
		Capacity:    capacityInt,
		Location:    location,
		Description: description,
		Start:       start,
		Destiny:     destiny,
		Time:        startTime,
		Date:        startDate,
		Distance:    distanceInt,
		Duration:    duration,
	}

	newService, err := m.DB.CreateService(service)
	if err != nil {
		utilities.DeleteImage(filename, "service")
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, newService, "service")
}

// GetServiceDetail handler for get service detail
func (m *Repository) GetServiceDetail(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	service, err := m.DB.GetDetailServiceByID(id)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, service, "service")
}

// ListAllService handler for list all service
func (m *Repository) ListAllService(w http.ResponseWriter, r *http.Request) {
	services, err := m.DB.ListAllServices(10)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, services, "services")
}

// UpdateService handler for update service
func (m *Repository) UpdateService(w http.ResponseWriter, r *http.Request) {

	// get id from url
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// cek if request have valid token
	idUser, role, err := middleware.ChecToken(r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if role != 1 || idUser == 0 {
		utilities.WriteJsonError(w, errors.New("not allowed, become admin first"), http.StatusBadRequest)
		return
	}

	// get data from request
	if err := r.ParseMultipartForm(1024); err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	ownerId := r.FormValue("owner_id")
	price := r.FormValue("price")
	oldImage := r.FormValue("old_image")
	typeId := r.FormValue("type_id")
	capacity := r.FormValue("capacity")
	location := r.FormValue("location")
	description := r.FormValue("description")

	// validate

	// end validate

	// convert data to int
	ownerInt, err := strconv.Atoi(ownerId)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	priceInt, err := strconv.Atoi(price)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}

	typeInt, err := strconv.Atoi(typeId)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	capacityInt, err := strconv.Atoi(capacity)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// validate owner of user
	if idUser != ownerInt {
		utilities.WriteJsonError(w, errors.New("not allowed"), http.StatusBadRequest)
		return
	}

	uploadedImage, header, err := r.FormFile("image")
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	defer uploadedImage.Close()

	filename, err := utilities.UploadedImage(uploadedImage, header, "service")
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	service := models.ServiceUpdateRequest{
		Name:        name,
		Price:       priceInt,
		Image:       filename,
		TypeId:      typeInt,
		Capacity:    capacityInt,
		Location:    location,
		Description: description,
	}

	updatedService, err := m.DB.UpdateService(id, service)
	if err != nil {
		_ = utilities.DeleteImage(filename, "service")
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// delete old image
	_ = utilities.DeleteImage(oldImage, "service")

	utilities.WriteJson(w, http.StatusOK, updatedService, "service")
}

// DeleteService handler for delete service
func (m *Repository) DeleteService(w http.ResponseWriter, r *http.Request) {
	// get id from url
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// cek if request have valid token
	idUser, role, err := middleware.ChecToken(r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if role != 1 || idUser == 0 {
		utilities.WriteJsonError(w, errors.New("not allowed, become admin first"), http.StatusBadRequest)
		return
	}

	// get service data from database
	service, err := m.DB.GetSortDetailServiceByID(id)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	if service == nil {
		utilities.WriteJsonError(w, errors.New("service not found"), http.StatusBadRequest)
		return
	}

	// validate owner of service
	if idUser != service.OwnerId {
		utilities.WriteJsonError(w, errors.New("not allowed"), http.StatusBadRequest)
		return
	}

	// delete service
	err = m.DB.DeleteService(id)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	_ = utilities.DeleteImage(service.Image, "service")

	utilities.WriteJson(w, http.StatusOK, "ok", "service")
}

// ListPopService handler for list service
func (m *Repository) ListPopService(w http.ResponseWriter, r *http.Request) {
	services, err := m.DB.ListPopularServices(5)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, services, "services")
}

// SeachService handler list service by query
func (m *Repository) SeachService(w http.ResponseWriter, r *http.Request) {
	// get query from url
	queryValues := r.URL.Query()
	query := queryValues.Get("q")

	// get data from database
	services, err := m.DB.ListAllServicesBySearch(query)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, services, "services")
}

// ServiceByType handler for list service by type
func (m *Repository) ServiceByType(w http.ResponseWriter, r *http.Request) {
	// get id from url
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// get data from database
	services, err := m.DB.ListAllServicesByType(id, 10)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	utilities.WriteJson(w, http.StatusOK, services, "services")
}
