package uc_implementation

import (
	"Avito-Challenge/src/models"
	"Avito-Challenge/src/repositories"
	"Avito-Challenge/src/usecases"
	"errors"
	"time"
)

type BalanceUsecases struct {
	br repositories.IBalanceRepository
	ac usecases.IAccountingUsecase
}

func NewBalanceUsecases(br repositories.IBalanceRepository, ac usecases.IAccountingUsecase) *BalanceUsecases {
	return &BalanceUsecases{br: br, ac: ac}
}

func (bc *BalanceUsecases) GetByUserId(id int) (float64, error) {
	return bc.br.GetByUserId(id)
}

func (bc *BalanceUsecases) AddByUserId(income *models.IncomingTransaction) error {
	if income.Amount < 0 {
		return errors.New("Wrong parameter: amount must be positive")
	}
	err := bc.ac.LogTransaction(&models.Transaction{
		UserId: income.UserId,
		Other:  income.Other,
		Reason: income.Reason,
		Date:   time.Now().Format("2006-01-02"),
		Amount: income.Amount,
	})
	if err != nil {
		return err
	}
	return bc.br.AddByUserId(income.UserId, income.Amount)
}

func (bc *BalanceUsecases) AddReservation(reservation *models.Reservation) error {
	current, err := bc.br.GetByUserId(reservation.UserId)
	if err != nil {
		return err
	}
	if current < reservation.Amount {
		return errors.New("User balance is too low")
	}
	if reservation.Amount < 0 {
		return errors.New("Wrong parameter: amount must be positive")
	}

	err = bc.ac.LogTransaction(&models.Transaction{
		UserId: reservation.UserId,
		Other:  "User reservation bill",
		// Можно сделать запрос, получающий название услуги по ID или
		// полностью заменить использование ID названием услуги
		// и вставить название в поле Reason
		Reason: "Reserved for a service",
		Date:   time.Now().Format("2006-01-02"),
		Amount: -reservation.Amount,
	})
	if err != nil {
		return err
	}

	err = bc.br.Withdraw(reservation.UserId, reservation.Amount)
	if err != nil {
		return err
	}
	err = bc.br.AddReservation(reservation.UserId, reservation.OrderId, reservation.ServiceId, reservation.Amount)
	if err != nil {
		// Must return money if reservation has failed
		// Creating new local err to return the actual error message
		for err := errors.New(""); err != nil; err = bc.br.AddByUserId(reservation.UserId, reservation.Amount) {
		}
	}
	return err
}

func (bc *BalanceUsecases) CommitReservation(reservation *models.Reservation) error {
	// DeleteReservation: This method verifies every parameter including 'amount'
	err := bc.br.DeleteReservation(reservation.UserId, reservation.OrderId, reservation.ServiceId, reservation.Amount)
	if err != nil {
		return err
	}
	return bc.ac.RecordProfit(reservation.ServiceId, reservation.Amount)
}

func (bc *BalanceUsecases) AbortReservation(reservation *models.Reservation) error {
	err := bc.br.DeleteReservation(reservation.UserId, reservation.OrderId, reservation.ServiceId, reservation.Amount)
	if err != nil {
		return err
	}
	// Must return money if reservation has been canceled
	for err = errors.New(""); err != nil; err = bc.br.AddByUserId(reservation.UserId, reservation.Amount) {
	}

	return bc.ac.LogTransaction(&models.Transaction{
		UserId: reservation.UserId,
		Other:  "User reservation bill",
		Reason: "Reservation has been canceled",
		Date:   time.Now().Format("2006-01-02"),
		Amount: reservation.Amount,
	})
}

func (bc *BalanceUsecases) Transfer(transfer *models.Transfer) error {

	if transfer.Amount < 0 {
		return errors.New("Wrong parameter: amount must be positive")
	}

	err := bc.br.Withdraw(transfer.SourceUserId, transfer.Amount)
	if err != nil {
		return err
	}
	err = bc.ac.LogTransaction(&models.Transaction{
		UserId: transfer.SourceUserId,
		Other:  "Another user",
		Reason: transfer.Reason,
		Date:   time.Now().Format("2006-01-02"),
		Amount: -transfer.Amount,
	})
	if err != nil {
		// Must return money if transaction has failed
		// Creating new local err to return the actual error message
		for err := errors.New(""); err != nil; err = bc.br.AddByUserId(transfer.SourceUserId, transfer.Amount) {
		}
		return err
	}
	err = bc.br.AddByUserId(transfer.DestUserId, transfer.Amount)
	if err != nil {
		// Must return money if transaction has failed
		// Creating new local err to return the actual error message
		for err := errors.New(""); err != nil; err = bc.br.AddByUserId(transfer.SourceUserId, transfer.Amount) {
		}
		return err
	}
	err = bc.ac.LogTransaction(&models.Transaction{
		UserId: transfer.DestUserId,
		Other:  "Another user",
		Reason: transfer.Reason,
		Date:   time.Now().Format("2006-01-02"),
		Amount: transfer.Amount,
	})
	return err
}
