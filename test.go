// store data
data, err := h.TripRepository.CreateTrip(trip)
if err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	json.NewEncoder(w).Encode(response)
	return
}

// get data
tripInserted, err := h.TripRepository.GetTrip(data.ID)
// tripInserted.Image = os.Getenv("PATH_FILE") + tripInserted.Image
if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	json.NewEncoder(w).Encode(response)
	return
}
// success
w.WriteHeader(http.StatusOK)
response := dto.SuccessResult{Code: http.StatusOK, Data: convertTripResponse(tripInserted)}
json.NewEncoder(w).Encode(response)
}

// Update data
func (h *handlerTrip) UpdateTrip(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

// params
id, _ := strconv.Atoi(mux.Vars(r)["id"])
// get data
trip, err := h.TripRepository.GetTrip(int(id))
if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	json.NewEncoder(w).Encode(response)
	return
}
// validation
if r.FormValue("title") != "" {
	trip.Title = r.FormValue("title")
}
country_idInput, _ := strconv.Atoi(r.FormValue("country_id"))
if country_idInput != 0 {
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
dayInput, _ := strconv.Atoi(r.FormValue("day"))
if dayInput != 0 {
	trip.Day = dayInput
}
nightInput, _ := strconv.Atoi(r.FormValue("night"))
if nightInput != 0 {
	trip.Night = nightInput
}
date_tripInput, _ := time.Parse("2006-01-02", r.FormValue("date_trip"))
if date_tripInput.IsZero() {
	date_trip := trip.DateTrip
	trip.DateTrip = date_trip
}
priceInput, _ := strconv.Atoi(r.FormValue("price"))
if priceInput != 0 {
	trip.Price = priceInput
}
quotaInput, _ := strconv.Atoi(r.FormValue("price"))
if quotaInput != 0 {
	trip.Quota = quotaInput
}
if r.FormValue("description") != "" {
	trip.Description = r.FormValue("description")
}

// image
dataContexErr := r.Context().Value("Error")
// fmt.Println(dataContex)

var result []models.ImageResponse
if dataContexErr != true {
	// image
	dataContex := r.Context().Value("dataFileName")
	filename := dataContex.([]string)
	result = make([]models.ImageResponse, len(filename))

	fmt.Println(dataContex)

	// Isi slice baru dengan struktur yang diinginkan
	for i, v := range filename {
		result[i] = models.ImageResponse{Name: v}
	}
}
trip.Image = result

// fmt.Println(trip.ID)
// fmt.Println(trip.CountryID)
// update data
data, err := h.TripRepository.UpdateTrip(trip)
if err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	json.NewEncoder(w).Encode(response)
	return
}

// get data
tripInserted, err := h.TripRepository.GetTrip(data.ID)
// tripInserted.Image = os.Getenv("PATH_FILE") + tripInserted.Image
if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	json.NewEncoder(w).Encode(response)
	return
}
// success
w.WriteHeader(http.StatusOK)
response := dto.SuccessResult{Code: http.StatusOK, Data: convertTripResponse(tripInserted)}
json.NewEncoder(w).Encode(response)
}

// Delete data
func (h *handlerTrip) DeleteTrip(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
