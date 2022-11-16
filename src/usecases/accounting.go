package usecases

type IAccountingUsecase interface {
	RecordProfit(serviceId int, amount float64) error
	GenerateReport(year int, month int) (string, error)
}
