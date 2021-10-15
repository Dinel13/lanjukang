package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dinel13/lanjukang/middleware"
	"github.com/dinel13/lanjukang/models"
	"github.com/dinel13/lanjukang/pkg/utilities"
)

func (m *Repository) CreateService(w http.ResponseWriter, r *http.Request) {

	// cek if request have valid token
	_, role, err := middleware.ChecToken(w, r, m.App.JwtSecret)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if role != 1 {
		utilities.WriteJsonError(w, errors.New("not allowed, become admin first"), http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	price := r.FormValue("price")
	typeId := r.FormValue("type")
	capacity := r.FormValue("capacity")
	location := r.FormValue("location")
	description := r.FormValue("description")

	// validate

	// end validate

	uploadedImage, header, err := r.FormFile("image")
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}
	defer uploadedImage.Close()

	filename, err := utilities.UploadedImage(uploadedImage, header)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}

	priceInt, err := strconv.Atoi(price)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}

	typeInt, err := strconv.Atoi(typeId)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}

	capacityInt, err := strconv.Atoi(capacity)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}

	locationInt, err := strconv.Atoi(location)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}

	service := models.ServiceRequest{
		Name:        name,
		Price:       priceInt,
		Image:       filename,
		TypeId:      typeInt,
		Capacity:    capacityInt,
		LocationId:  locationInt,
		Description: description,
	}

	newService, err := m.DB.CreateService(service)
	if err != nil {
		utilities.WriteJsonError(w, err, http.StatusInternalServerError)
	}

	utilities.WriteJson(w, http.StatusOK, newService, "service")
}
