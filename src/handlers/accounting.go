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

func (ah *AccountingHandlers) GetTransactions(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body")
		}
	}(r.Body)

	query := r.URL.Query()
	strId, idPresent := query["userId"]
	strSize, sizePresent := query["size"]
	strPage, pagePresent := query["page"]
	qSortBy, sortPresent := query["sortBy"]
	if !(idPresent && sizePresent && pagePresent) || len(strId) != 1 || len(strSize) != 1 || len(strPage) != 1 {
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
	size, err := strconv.Atoi(strSize[0])
	if err != nil {
		fmt.Println("Received an invalid query parameter for GetByUserId")
		http.Error(w, "Failed to get user balance: invalid query parameter (id must be integer)", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(strPage[0])
	if err != nil {
		fmt.Println("Received an invalid query parameter for GetByUserId")
		http.Error(w, "Failed to get user balance: invalid query parameter (id must be integer)", http.StatusBadRequest)
		return
	}

	var sortBy string
	if !sortPresent || (qSortBy[0] != "sum") {
		sortBy = "date"
	} else {
		sortBy = "sum"
	}

	transactions, err := ah.ac.GetTransactions(id, size, page, sortBy)
	if err != nil {
		fmt.Println("Failed to get transactions")
		http.Error(w, "Error while getting transactions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		fmt.Println("Encoding json error: ", err)
		http.Error(w, "Failed to encode data to json", http.StatusInternalServerError)
		return
	}
}
