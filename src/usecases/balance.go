package usecases

import "Avito-Challenge/src/models"

type IBalanceUsecase interface {
	GetByUserId(id int) (float64, error)
	AddByUserId(income *models.IncomingTransaction) error
	AddReservation(reservation *models.Reservation) error
	CommitReservation(reservation *models.Reservation) error
	AbortReservation(reservation *models.Reservation) error
	Transfer(transfer *models.Transfer) error
}
