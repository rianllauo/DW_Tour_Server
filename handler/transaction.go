package handlers

import (
	dto "dewetour/dto/result"
	tripdto "dewetour/dto/trip"
	"dewetour/models"
	"fmt"
	"log"
	"os"
	"time"

	transactiondto "dewetour/dto/transaction"

	// "dewetour/pkg/jwt"
	"dewetour/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"

	"gopkg.in/gomail.v2"
)

var c = coreapi.Client{
	ServerKey: "SB-Mid-server-BIcg8EbkRJTWRv68iWkcHQnq",
	ClientKey: os.Getenv("CLIENT_KEY"),
}

// var path_file = "http://localhost:5000/uploads/"

type handleTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handleTransaction {
	return &handleTransaction{TransactionRepository}
}

func (h *handleTransaction) FindTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range transaction {
		transaction[i].Attachment = path_file + p.Attachment
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handleTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// var trips models.Trip
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handleTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// dataContex := r.Context().Value("dataFile") // add this code
	// filename := dataContex.(string)

	// counter_qyt, _ := strconv.Atoi(r.FormValue("counter_qyt"))
	// total, _ := strconv.Atoi(r.FormValue("total"))

	// trip, _ := strconv.Atoi(r.FormValue("trip_id"))
	// request := transactiondto.TransactionRequest{
	// 	CounterQty: counter_qyt,
	// 	Total:      total,
	// 	Status:     r.FormValue("status"),
	// 	TripID:     trip,
	// }

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	var request transactiondto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Unique Transaction Id
	var transIdIsMatch = false
	var transactionId int
	for !transIdIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(transactionId)
		if transactionData.ID == 0 {
			transIdIsMatch = true
		}
	}

	transaction := models.Transaction{
		ID:         transactionId,
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		TripID:     request.TripID,
		// Attachment: filename,
		UserId: userId,
	}

	newTransaction, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataTransaction, err := h.TransactionRepository.GetTransaction(newTransaction.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	fmt.Println(newTransaction)

	var s = snap.Client{}
	s.New("SB-Mid-server-BIcg8EbkRJTWRv68iWkcHQnq", midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(newTransaction.ID),
			GrossAmt: int64(newTransaction.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataTransaction.User.FullName,
			Email: dataTransaction.User.Email,
		},
	}

	fmt.Println(req)
	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp, DataTrip: req}
	json.NewEncoder(w).Encode(response)
}

func (h *handleTransaction) ChangeTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// dataContex := r.Context().Value("dataFile") // add this code
	// filename := dataContex.(string)

	// id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// transaction, err := h.TransactionRepository.GetTransaction(int(id))
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// counterQyt, _ := strconv.Atoi(r.FormValue("counter_qyt"))
	// if counterQyt != 0 {
	// 	transaction.CounterQty = counterQyt
	// }

	// total, _ := strconv.Atoi(r.FormValue("total"))
	// if total != 0 {
	// 	transaction.Total = total
	// }

	// if r.FormValue("status") != "" {
	// 	transaction.Status = r.FormValue("status")
	// }

	// trip_id, _ := strconv.Atoi(r.FormValue("trip_id"))
	// if trip_id != 0 {
	// 	transaction.TripID = trip_id
	// }

	// if r.FormValue("image") != "" {
	// 	transaction.Attachment = filename
	// }

	var request transactiondto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := models.Transaction{
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		TripID:     request.TripID,
		// Attachment: filename,
		// UserId: userId,
	}

	data, err := h.TransactionRepository.ChangeTransaction(transaction)
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
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handleTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	transaction, _ := h.TransactionRepository.GetOneTransaction(orderId)
	fmt.Println(transactionStatus, fraudStatus, orderId, transaction)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			h.TransactionRepository.UpdateTransaction("pending", transaction)
		} else if fraudStatus == "accept" {
			SendMail("success", transaction)
			h.TransactionRepository.UpdateTransaction("success", transaction)
		}
	} else if transactionStatus == "settlement" {
		SendMail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", transaction)
	} else if transactionStatus == "deny" {
		SendMail("failed", transaction)
		h.TransactionRepository.UpdateTransaction("failed", transaction)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		SendMail("failed", transaction)
		h.TransactionRepository.UpdateTransaction("failed", transaction)
	} else if transactionStatus == "pending" {

		h.TransactionRepository.UpdateTransaction("pending", transaction)
	}

	w.WriteHeader(http.StatusOK)
}
func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DumbMerch <demo.dumbways@gmail.com>"
		var CONFIG_AUTH_EMAIL = "rianlauo11@gmail.com"
		var CONFIG_AUTH_PASSWORD = "avhklxgxkeqkynck"

		var tripName = transaction.Trip.Title
		var price = strconv.Itoa(transaction.Trip.Price)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
		<style>
		  h1 {
		  color: brown;
		  }
		</style>
		</head>
		<body>
		<h2>Product payment :</h2>
		<ul style="list-style-type:none;">
		  <li>Name : %s</li>
		  <li>Total payment: Rp.%s</li>
		  <li>Status : <b>%s</b></li>
		</ul>
		</body>
	  </html>`, tripName, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
}
func convertResponseTransaction(u models.Trip) tripdto.TripResponse {
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
