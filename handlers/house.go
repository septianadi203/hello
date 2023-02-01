package handlers

import (
	"encoding/json"
	housesdto "housy/dto/house"
	dto "housy/dto/result"
	"housy/models"
	"housy/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var path_file = "http://localhost:5000/uploads/"

type handlerHouse struct {
	HouseRepository repositories.HouseRepository
}

func HandlerHouse(HouseRepository repositories.HouseRepository) *handlerHouse {
	return &handlerHouse{HouseRepository}
}

func (h *handlerHouse) FindHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	houses, err := h.HouseRepository.FindHouses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range houses {
		houses[i].Image = path_file + p.Image
	  }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: houses}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse) GetHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// var profile models.Profile
	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	house.Image = path_file + house.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseHouse(house)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse) CreateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get dataFile from midleware and store to filename variable here ...

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	Bedroom, _ := strconv.Atoi(r.FormValue("Bedroom"))
	Bathroom, _ := strconv.Atoi(r.FormValue("Bathroom"))

	request := housesdto.HouseRequest{
		Name:      r.FormValue("name"),
		CityName:  r.FormValue("cityname"),
		Address:   r.FormValue("address"),
		TypeRent:  r.FormValue("type_rent"),
		Amenities: r.FormValue("amenities"),
		Image:     r.FormValue("image"),
		Bedroom:   Bedroom,
		Price:     price,
		Bathroom:  Bathroom,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	house := models.House{
		Name:      request.Name,
		CityName:  request.CityName,
		Address:   request.Address,
		Price:     request.Price,
		TypeRent:  request.TypeRent,
		Amenities: request.Amenities,
		Bedroom:   request.Bedroom,
		Bathroom:  request.Bathroom,
		Image:     filename,
	}

	// err := mysql.DB.Create(&product).Error
	house, err = h.HouseRepository.CreateHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	house, _ = h.HouseRepository.GetHouse(house.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: house}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse) DeleteHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.HouseRepository.DeleteHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse) UpdateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	bedroom, _ := strconv.Atoi(r.FormValue("Bedroom"))
	bathroom, _ := strconv.Atoi(r.FormValue("Bathroom"))
	request := housesdto.HouseRequest{
		Name:      r.FormValue("name"),
		CityName:  r.FormValue("cityname"),
		Address:   r.FormValue("address"),
		TypeRent:  r.FormValue("type_rent"),
		Amenities: r.FormValue("amenities"),
		Price:     price,
		Bedroom:   bedroom,
		Bathroom:  bathroom,
		Image:     filename,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	house, err := h.HouseRepository.GetHouse(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		house.Name = request.Name
	}

	if request.CityName != "" {
		house.CityName = request.CityName
	}

	if request.Address != "" {
		house.Address = request.Address
	}

	if request.Price != 0 {
		house.Price = request.Price
	}

	if request.TypeRent != "" {
		house.TypeRent = request.TypeRent
	}

	if request.Amenities != "" {
		house.Amenities = request.Amenities
	}

	if request.Bedroom != 0 {
		house.Bedroom = request.Bedroom
	}

	if request.Bathroom != 0 {
		house.Bathroom = request.Bathroom
	}

	if request.Image != "" {
		house.Image = request.Image
	}

	data, err := h.HouseRepository.UpdateHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func convertResponseHouse(u models.House) housesdto.ResponseHouse {
	return housesdto.ResponseHouse{
		ID:        u.ID,
		Name:      u.Name,
		CityName:  u.CityName,
		Address:   u.Address,
		Price:     u.Price,
		TypeRent:  u.TypeRent,
		Amenities: u.Amenities,
		Bedroom:   u.Bedroom,
		Bathroom:  u.Bathroom,
		Image:     u.Image,
	}
}
