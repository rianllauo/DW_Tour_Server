package handlers

import (
	dto "dewetour/dto/result"
	tripdto "dewetour/dto/trip"
	"dewetour/models"
	"fmt"
	"time"

	// "dewetour/pkg/jwt"
	"dewetour/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

var path_file = "http://localhost:5000/uploads/"

type handlerTrip struct {
	TripRepository repositories.TripRepository
}

func HandlerTrip(TripRepository repositories.TripRepository) *handlerTrip {
	return &handlerTrip{TripRepository}
}

func (h *handlerTrip) FindTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trips, err := h.TripRepository.FindTrip()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// for i, p := range trips {
	// 	trips[i].Image = path_file + p.Image
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trips}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) GetTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// var trips models.Trip
	trips, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// trips.Image = path_file + trips.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trips}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// dataContex := r.Context().Value("dataFile")
	// filename := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	night, _ := strconv.Atoi(r.FormValue("night"))
	quota, _ := strconv.Atoi(r.FormValue("quota"))
	country_id, _ := strconv.Atoi(r.FormValue("country_id"))

	// images := []string{r.FormValue("image")}

	date_trip, _ := time.Parse("01 January 2002", r.FormValue("date_trip"))
	fmt.Println(date_trip)

	request := tripdto.TripRequest{
		Title:          r.FormValue("title"),
		Accomodation:   r.FormValue("accomodation"),
		Transportation: r.FormValue("transportation"),
		Eat:            r.FormValue("eat"),
		Day:            day,
		Night:          night,
		Description:    r.FormValue("description"),
		Price:          price,
		Quota:          quota,
		// Image:          []string{r.FormValue("image")},
		CountryID: country_id,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// imageUrl := request.Image  path_file + trips.Image

	trip := models.Trip{
		Title:          request.Title,
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		// Image:          string{r.FormValue("image")},
		Day:         request.Day,
		Night:       request.Night,
		DateTrip:    date_trip,
		Price:       request.Price,
		Quota:       request.Quota,
		Description: request.Description,
		CountryID:   request.CountryID,
		UserId:      userId,
	}

	trip, err = h.TripRepository.CreateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trip, _ = h.TripRepository.GetTrip(trip.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) UpdateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// dataContex := r.Context().Value("dataFile") // add this code
	// filename := dataContex.(string)

	// request := tripdto.UpdateTripRequest{
	// 	Image: filename,
	// }

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	if r.FormValue("title") != "" {
		trip.Title = r.FormValue("title")
	}

	country_idInput, _ := strconv.Atoi(r.FormValue("country_id"))
	if r.FormValue("country_id") != "" {
		trip.CountryID = country_idInput
	}

	if r.FormValue("accomodation") != "" {
		trip.Accomodation = r.FormValue("accomodation")
	}

	if r.FormValue("transportation") != "" {
		trip.Transportation = r.FormValue("transportation")
	}

	if r.FormValue("eat") != "" {
		trip.Eat = r.FormValue("eat")
	}

	day, _ := strconv.Atoi(r.FormValue("day"))
	if r.FormValue("day") != "" {
		trip.Day = day
	}

	night, _ := strconv.Atoi(r.FormValue("night"))
	if r.FormValue("night") != "" {
		trip.Night = night
	}

	// if request.Image != "" {
	// 	trip.Image = request.Image
	// }

	price, _ := strconv.Atoi(r.FormValue("price"))
	if r.FormValue("price") != "" {
		trip.Price = price
	}

	quota, _ := strconv.Atoi(r.FormValue("quota"))
	if r.FormValue("quota") != "" {
		trip.Quota = quota
	}

	data, err := h.TripRepository.UpdateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	countryIns, err := h.TripRepository.GetTrip(data.ID)
	// countryIns.Image = path_file + countryIns.Image
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: countryIns}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) DeleteTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trip, err := h.TripRepository.GetTrip(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TripRepository.DeleteTrip(trip, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data.ID}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTrip(u models.Trip) tripdto.TripResponse {
	return tripdto.TripResponse{
		// ID:             u.ID,
		Title:          u.Title,
		CountryID:      u.CountryID,
		Accomodation:   u.Accomodation,
		Transportation: u.Transportation,
		Eat:            u.Eat,
		// Image:          u.Image,
		Day:         u.Day,
		Night:       u.Night,
		DateTrip:    u.DateTrip,
		Price:       u.Price,
		Quota:       u.Quota,
		Description: u.Description,
		// Country:        u.Country,
	}
}
