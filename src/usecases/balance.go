package usecases

type IBalanceUsecase interface {
	GetByUserId(id int) (float64, error)
	AddByUserId(id int, amount float64) error
	AddReservation(userId int, orderId int, serviceId int, amount float64) error
	CommitReservation(userId int, orderId int, serviceId int, amount float64) error
}
