package handlers

import (
	"Avito-Challenge/src/models"
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

	// Acquiring and validating
	income := &models.IncomingTransaction{}
	err := json.NewDecoder(r.Body).Decode(income)
	if err != nil {
		fmt.Println("Decoding json error: ", err)
		http.Error(w, "Failed to decode json data", http.StatusBadRequest)
		return
	}

	// Processing request
	err = bh.bc.AddByUserId(income)
	if err != nil {
		fmt.Println("Failed to update user balance")
		http.Error(w, "Error while updating balance", http.StatusInternalServerError)
		return
	}

	// Completing request
	w.WriteHeader(http.StatusOK)
}

func (bh *BalanceHandlers) AddReservation(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	reservation := &models.Reservation{}
	err := json.NewDecoder(r.Body).Decode(reservation)
	if err != nil {
		fmt.Println("Decoding json error: ", err)
		http.Error(w, "Failed to decode json data", http.StatusBadRequest)
		return
	}

	err = bh.bc.AddReservation(reservation)
	if err != nil {
		fmt.Println("Failed to add a reservation")
		http.Error(w, "Failed to add a reservation", http.StatusInternalServerError)
		return
	}

	// Completing request
	w.WriteHeader(http.StatusOK)

}

func (bh *BalanceHandlers) CommitReservation(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	reservation := &models.Reservation{}
	err := json.NewDecoder(r.Body).Decode(reservation)
	if err != nil {
		fmt.Println("Decoding json error: ", err)
		http.Error(w, "Failed to decode json data", http.StatusBadRequest)
		return
	}

	err = bh.bc.CommitReservation(reservation)
	if err != nil {
		fmt.Println("Failed to close reservation")
		http.Error(w, "Failed to close reservation", http.StatusInternalServerError)
		return
	}

	// Completing request
	w.WriteHeader(http.StatusOK)
}

func (bh *BalanceHandlers) AbortReservation(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	reservation := &models.Reservation{}
	err := json.NewDecoder(r.Body).Decode(reservation)
	if err != nil {
		fmt.Println("Decoding json error: ", err)
		http.Error(w, "Failed to decode json data", http.StatusBadRequest)
		return
	}

	err = bh.bc.AbortReservation(reservation)
	if err != nil {
		fmt.Println("Failed to abort reservation")
		http.Error(w, "Failed to abort reservation", http.StatusInternalServerError)
		return
	}

	// Completing request
	w.WriteHeader(http.StatusOK)
}

func (bh *BalanceHandlers) TransferBetweenUsers(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	transfer := &models.Transfer{}
	err := json.NewDecoder(r.Body).Decode(transfer)
	if err != nil {
		fmt.Println("Decoding json error: ", err)
		http.Error(w, "Failed to decode json data", http.StatusBadRequest)
		return
	}

	err = bh.bc.Transfer(transfer)
	if err != nil {
		fmt.Println("Failed to update user balance")
		http.Error(w, "Error while updating balance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
