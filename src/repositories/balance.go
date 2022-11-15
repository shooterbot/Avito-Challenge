package repositories

type IBalanceRepository interface {
	GetByUserId(id int) (float64, error)
}
