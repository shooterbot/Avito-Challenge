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

	query := r.URL.Query()
	strId, present := query["userId"]
	if !present || len(strId) != 1 {
		fmt.Println("Received a wrong query parameter for GetByUserId")
		http.Error(w, "Failed to get user balance: wrong query parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(strId[0])
	if err != nil {
		fmt.Println("Received an invalid query parameter for GetByUserId")
		http.Error(w, "Failed to get user balance: invalid query parameter (id must be integer)", http.StatusBadRequest)
		return
	}

	balance, err := bh.bc.GetByUserId(id)
	if err != nil {
		fmt.Println("Failed to get balance from UC")
		http.Error(w, "Error while getting balance", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}
