package uc_implementation

import (
	"Avito-Challenge/src/models"
	"Avito-Challenge/src/repositories"
	"fmt"
	"os"
	"time"
)

type AccountingUsecase struct {
	ar repositories.IAccountingRepository
}

func NewAccountingUsecases(repo repositories.IAccountingRepository) *AccountingUsecase {
	return &AccountingUsecase{ar: repo}
}

func (ac *AccountingUsecase) RecordProfit(serviceId int, amount float64) error {
	date := time.Now().Format("2006-01-02")
	return ac.ar.RecordProfit(serviceId, amount, date)
}

func (ac *AccountingUsecase) GenerateReport(year int, month int) (string, error) {
	start := fmt.Sprintf("01.%d.%d", month, year)
	end := fmt.Sprintf("01.%d.%d", month+1, year)
	report, err := ac.ar.CalculateSum(start, end)

	if err != nil {
		return "", err
	}

	file, err := os.Create("report.csv")
	if err != nil {
		fmt.Println("Could not create report file: ", err)
		return "", err
	}
	defer file.Close()

	for id, amount := range report {
		file.WriteString(fmt.Sprintf("%d;%f\n", id, amount))
	}

	dir, err := os.Getwd()

	return dir + "\\" + file.Name(), err
}

func (ac *AccountingUsecase) LogTransaction(transaction *models.Transaction) error {
	return ac.ar.LogTransaction(transaction)
}
