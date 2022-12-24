package handlers

import (
	dto "dewetour/dto/result"
	tripdto "dewetour/dto/trip"
	"dewetour/models"

	// "dewetour/pkg/jwt"
	"dewetour/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
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

	for i, p := range trips {
		trips[i].Image = path_file + p.Image
	}

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

	trips.Image = path_file + trips.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trips}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userInfo := r.Context().Value("userLogin").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	// qty, _ := strconv.Atoi(r.FormValue("qty"))
	country_id, _ := strconv.Atoi(r.FormValue("country_id"))
	request := tripdto.TripRequest{
		Title:       r.FormValue("name"),
		Description: r.FormValue("desc"),
		Price:       price,
		// Qty:        qty,
		CountryID: country_id,
		// Image:     filename,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trip := models.Trip{
		Title:          request.Title,
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Image:          filename,
		// Day:            request.Day,
		// Night:          request.Night,
		// DateTrip:       request.DateTrip,
		Price:       request.Price,
		Quota:       request.Quota,
		Description: request.Description,
		CountryID:   request.CountryID,
		// CountryID:   request.C,
		// CountryId:      request.CountryID,
		// UserId:         userId,
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

	// price, _ := strconv.Atoi(r.FormValue("price"))
	// qty, _ := strconv.Atoi(r.FormValue("quota"))
	// countryId, _ := strconv.Atoi(r.FormValue("countryId"))
	// request := new(tripdto.UpdateTripRequest)

	// request.Title = r.FormValue("title")
	// request.CountryID = countryId
	// request.Price = price
	// request.Quota = qty

	// 	Title:       r.FormValue("title"),
	// 	Description: r.FormValue("desc"),
	// 	Price:       price,
	// 	// Qty:        qty,
	// 	CountryID: country_id,
	// 	// Image:     filename,
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

	price, _ := strconv.Atoi(r.FormValue("price"))
	if r.FormValue("price") != "" {
		trip.Price = price
	}

	data, err := h.TripRepository.UpdateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// tripInserted, err := h.TripRepository.GetTrip(data.ID)
	// // tripInserted.Image = os.Getenv("PATH_FILE") + tripInserted.Image
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
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
		Image:          u.Image,
		Day:            u.Day,
		Night:          u.Night,
		DateTrip:       u.DateTrip,
		Price:          u.Price,
		Quota:          u.Quota,
		Description:    u.Description,
		// Country:        u.Country,
	}
}
