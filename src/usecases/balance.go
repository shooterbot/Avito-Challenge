package usecases

type IBalanceUsecase interface {
	GetByUserId(id int) (float64, error)
}
