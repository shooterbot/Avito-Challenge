package uc_implementation

import (
	"Avito-Challenge/src/repositories"
	"errors"
)

type BalanceUsecases struct {
	br repositories.IBalanceRepository
}

func NewBalanceUsecases(br repositories.IBalanceRepository) *BalanceUsecases {
	return &BalanceUsecases{br: br}
}

func (bc *BalanceUsecases) GetByUserId(id int) (float64, error) {
	return bc.br.GetByUserId(id)
}

func (bc *BalanceUsecases) AddByUserId(id int, amount float64) error {
	if amount < 0 {
		return errors.New("Wrong parameter: amount must be positive")
	}
	return bc.br.AddByUserId(id, amount)
}
