package usecases

type IAccountingUsecase interface {
	RecordProfit(serviceId int, amount float64) error
}
