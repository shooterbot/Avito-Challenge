package handlers

import (
	"Avito-Challenge/src/usecases"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type BalanceHandlers struct {
	bc usecases.IBalanceUsecase
}

func NewBalanceHandlers(balanceCase usecases.IBalanceUsecase) *BalanceHandlers {
	return &BalanceHandlers{bc: balanceCase}
}

func (bh *BalanceHandlers) GetByUserId(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	// Acquiring parameters
	query := r.URL.Query()
	strId, present := query["userId"]
	if !present || len(strId) != 1 {
		// Received none or several values, error
		fmt.Println("Received a wrong query parameter for GetByUserId")
		http.Error(w, "Failed to get user balance: wrong query parameter", http.StatusBadRequest)
		return
	}

	// Validating parameters
	id, err := strconv.Atoi(strId[0])
	if err != nil {
		fmt.Println("Received an invalid query parameter for GetByUserId")
		http.Error(w, "Failed to get user balance: invalid query parameter (id must be integer)", http.StatusBadRequest)
		return
	}

	// Getting result
	balance, err := bh.bc.GetByUserId(id)
	if err != nil {
		fmt.Println("Failed to get balance from UC")
		http.Error(w, "Error while getting balance", http.StatusInternalServerError)
		return
	}

	// Encoding result
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}

func (bh *BalanceHandlers) AddByUserId(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	// Acquiring parameters
	query := r.URL.Query()
	strId, idPresent := query["userId"]
	strAmount, amountPresent := query["amount"]
	if !(idPresent && amountPresent) || len(strId) != 1 || len(strAmount) != 1 {
		fmt.Println("Received a wrong query parameter for AddByUserId")
		http.Error(w, "Failed to update user balance: wrong query parameter", http.StatusBadRequest)
		return
	}

	// Validating parameters
	id, err := strconv.Atoi(strId[0])
	if err != nil {
		fmt.Println("Received an invalid query parameter for AddByUserId")
		http.Error(w, "Failed to get user balance: invalid query parameter (id must be integer)", http.StatusBadRequest)
		return
	}
	amount, err := strconv.ParseFloat(strAmount[0], 64)
	if err != nil {
		fmt.Println("Received an invalid query parameter for AddByUserId")
		http.Error(w, "Failed to get user balance: invalid query parameter (amount must be a number)", http.StatusBadRequest)
		return
	}

	// Processing request
	err = bh.bc.AddByUserId(id, amount)
	if err != nil {
		fmt.Println("Failed to update user balance")
		http.Error(w, "Error while updating balance", http.StatusInternalServerError)
		return
	}

	// Completing request
	w.WriteHeader(http.StatusOK)
}

func (bh *BalanceHandlers) Reserve(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	// Acquiring parameters
	query := r.URL.Query()
	strUserId, userIdPresent := query["userId"]
	strServiceId, serviceIdPresent := query["serviceId"]
	strAmount, amountPresent := query["amount"]
	if !(userIdPresent && serviceIdPresent && amountPresent) ||
		len(strUserId) != 1 || len(strAmount) != 1 || len(strServiceId) != 1 {
		fmt.Println("Received a wrong query parameter for AddByUserId")
		http.Error(w, "Failed to update user balance: wrong query parameter", http.StatusBadRequest)
		return
	}

	// Validating parameters
	userId, err := strconv.Atoi(strUserId[0])
	if err != nil {
		fmt.Println("Received an invalid query parameter for AddReservation")
		http.Error(w, "Failed to get user balance: invalid query parameter (id must be integer)", http.StatusBadRequest)
		return
	}
	serviceId, err := strconv.Atoi(strUserId[0])
	if err != nil {
		fmt.Println("Received an invalid query parameter for AddReservation")
		http.Error(w, "Failed to get user balance: invalid query parameter (id must be integer)", http.StatusBadRequest)
		return
	}
	amount, err := strconv.ParseFloat(strAmount[0], 64)
	if err != nil {
		fmt.Println("Received an invalid query parameter for AddReservation")
		http.Error(w, "Failed to get user balance: invalid query parameter (amount must be a number)", http.StatusBadRequest)
		return
	}

	err = bh.bc.AddReservation(userId, serviceId, amount)
	if err != nil {
		fmt.Println("Failed to update user balance")
		http.Error(w, "Error while updating balance", http.StatusInternalServerError)
		return
	}

	// Completing request
	w.WriteHeader(http.StatusOK)

}
