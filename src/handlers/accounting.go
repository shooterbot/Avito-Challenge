package handlers

import (
	"Avito-Challenge/src/usecases"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type AccountingHandlers struct {
	ac usecases.IAccountingUsecase
}

func NewAccountingHandlers(accountingCase usecases.IAccountingUsecase) *AccountingHandlers {
	return &AccountingHandlers{ac: accountingCase}
}

func (ah *AccountingHandlers) GetReport(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	// Acquiring parameters
	query := r.URL.Query()
	strYear, yearPresent := query["year"]
	strMonth, monthPresent := query["month"]
	if !(yearPresent && monthPresent) || len(strYear) != 1 || len(strMonth) != 1 {
		fmt.Println("Received a wrong query parameter for GetReport")
		http.Error(w, "Failed to get report: wrong query parameter", http.StatusBadRequest)
		return
	}

	// Validating parameters
	year, err := strconv.Atoi(strYear[0])
	if err != nil {
		fmt.Println("Received an invalid query parameter for GetReport")
		http.Error(w, "Failed to get user balance: invalid query parameter (year must be integer)", http.StatusBadRequest)
		return
	}
	month, err := strconv.Atoi(strMonth[0])
	if err != nil {
		fmt.Println("Received an invalid query parameter for GetReport")
		http.Error(w, "Failed to get user balance: invalid query parameter (month must be integer)", http.StatusBadRequest)
		return
	}

	// Getting result
	path, err := ah.ac.GenerateReport(year, month)
	if err != nil {
		fmt.Println("Failed to get balance from UC")
		http.Error(w, "Error while getting balance", http.StatusInternalServerError)
		return
	}

	path = "file:///" + path

	// Encoding result
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(path)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}
