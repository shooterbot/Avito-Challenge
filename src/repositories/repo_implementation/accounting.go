package repo_implementation

import "Avito-Challenge/src/database/pgdb"

const (
	insertReservation = `insert into accounting(service_id, amount, date) values($1, $2, $3) returning id;`
)

type AccountingRepository struct {
	db *pgdb.DBManager
}

func NewAccountingRepository(manager *pgdb.DBManager) *AccountingRepository {
	return &AccountingRepository{db: manager}
}

func (ar *AccountingRepository) RecordProfit(serviceId int, amount float64, date string) error {
	_, err := ar.db.Exec(insertReservation, serviceId, amount, date)
	return err
}
