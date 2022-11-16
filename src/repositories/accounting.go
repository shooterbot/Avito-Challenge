package repositories

import "Avito-Challenge/src/models"

type IAccountingRepository interface {
	RecordProfit(serviceId int, amount float64, date string) error
	CalculateSum(startDate string, endDate string) (map[int]float64, error)
	LogTransaction(transaction *models.Transaction) error
	GetTransactions(id int, size int, page int, sortBy string) (*[]models.Transaction, error)
}
