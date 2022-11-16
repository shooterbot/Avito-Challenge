package usecases

import "Avito-Challenge/src/models"

type IAccountingUsecase interface {
	RecordProfit(serviceId int, amount float64) error
	GenerateReport(year int, month int) (string, error)
	LogTransaction(transaction *models.Transaction) error
}
