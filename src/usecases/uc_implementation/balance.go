package uc_implementation

type BalanceUsecases struct {
	br repositories.IBalanceRepository
}

func NewBalanceUsecases(br repositories.IBalanceRepository) *BalanceUsecases {
	return &BalanceUsecases{br: br}
}

func (bc *BalanceUsecases) GetByUserId(id int) (float64, error) {
	return bc.br.GetByUserId(id)
}
