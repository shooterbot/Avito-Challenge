package repositories

type IBalanceRepository interface {
	GetByUserId(id int) (float64, error)
	AddByUserId(id int, amount float64) error
}
