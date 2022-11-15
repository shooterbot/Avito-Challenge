package repositories

type IAccountingRepository interface {
	RecordProfit(serviceId int, amount float64, date string) error
}
