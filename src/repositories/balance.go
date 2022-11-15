package repositories

type IBalanceRepository interface {
	GetByUserId(id int) (float64, error)
	AddByUserId(id int, amount float64) error
	Withdraw(id int, amount float64) error
	AddReservation(userId int, serviceId int, amount float64) error
}
