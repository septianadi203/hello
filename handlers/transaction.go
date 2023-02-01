package handlers

import (
	"encoding/json"
	dto "housy/dto/result"
	"housy/models"
	"strconv"

	transactiondto "housy/dto/transaction"
	"housy/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transactions}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// var profile models.Profile
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

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get dataFile from midleware and store to filename variable here ...

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	house_id, _ := strconv.Atoi(r.FormValue("house_id"))
	user_id, _ := strconv.Atoi(r.FormValue("user_id"))
	total, _ := strconv.Atoi(r.FormValue("total"))

	request := transactiondto.RequestTransaction{
		CheckIn:       r.FormValue("check_in"),
		CheckOut:      r.FormValue("check_out"),
		StatusPayment: r.FormValue("status_payment"),
		Attachment:    r.FormValue("attachment"),
		HouseId:       house_id,
		UserId:        user_id,
		Total:         total,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := models.Transaction{
		CheckIn:       request.CheckIn,
		CheckOut:      request.CheckOut,
		HouseId:       request.HouseId,
		UserId:        request.UserId,
		Total:         request.Total,
		StatusPayment: request.StatusPayment,
		Attachment:    filename,
	}
	// err := mysql.DB.Create(&product).Error
	transaction, err = h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction, _ = h.TransactionRepository.GetTransaction(transaction.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}
