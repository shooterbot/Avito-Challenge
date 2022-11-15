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
	return &BalancesHandlers{bc: balanceCase}
}

func (bh *BalanceHandlers) GetByUserId(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	query := r.URL.Query()
	id, present := query["userId"]
	if !present || len(id) != 1 {
		fmt.Println("Received a wrong query parameter for GetByUserId")
		http.Error(w, "Failed to get user balance: wrong query parameter", http.StatusBadRequest)
		return
	}

	balance, err := bh.bc.GetByUserId(strconv.Atoi(id[0]))
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
