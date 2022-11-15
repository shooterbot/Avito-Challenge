package uc_implementation

import (
	"Avito-Challenge/src/accounting"
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

func (bc *BalanceUsecases) AddReservation(userId, orderId, serviceId int, amount float64) error {
	current, err := bc.br.GetByUserId(userId)
	if err != nil {
		return err
	}
	if current < amount {
		return errors.New("User balance is too low")
	}

	err = bc.br.Withdraw(userId, amount)
	if err != nil {
		return err
	}
	err = bc.br.AddReservation(userId, orderId, serviceId, amount)
	if err != nil {
		// Must return money if reservation has failed
		// Creating new local err to return the actual error message
		for err := errors.New(""); err != nil; err = bc.br.AddByUserId(userId, amount) {
		}
	}
	return err
}

func (bc *BalanceUsecases) CommitReservation(userId, orderId, serviceId int) error {
	amount, err := bc.br.GetReserved(orderId, serviceId)
	if err != nil {
		return err
	}
	accounting.RecordProfit(serviceId, amount)
	return bc.br.CommitReservation(userId, orderId, serviceId)
}
