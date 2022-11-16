package repositories

type IBalanceRepository interface {
	GetByUserId(id int) (float64, error)
	AddByUserId(id int, amount float64) error
	Withdraw(id int, amount float64) error
	AddReservation(userId int, orderId int, serviceId int, amount float64) error
	DeleteReservation(userId int, orderId int, serviceId int, amount float64) error
	GetReserved(orderId int, serviceId int) (float64, error)
}
