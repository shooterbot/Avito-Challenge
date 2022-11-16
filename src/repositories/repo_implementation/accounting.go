package repo_implementation

import (
	"Avito-Challenge/src/database/pgdb"
	"Avito-Challenge/src/utility"
	"fmt"
)

const (
	insertReservation = `insert into accounting(service_id, amount, completion_date) values($1, $2, $3) returning id;`
	selectSum         = `select service_id, sum(amount) from accounting where $1 <= completion_date and completion_date < $2 group by service_id;`
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

func (ar *AccountingRepository) CalculateSum(startDate string, endDate string) (map[int]float64, error) {
	data, err := ar.db.Query(selectSum, startDate, endDate)
	if err != nil {
		fmt.Println("Failed to calculate accounting sum:", err)
		return nil, err
	}

	res := make(map[int]float64)
	for _, row := range data {
		id := utility.BytesToInt(row[0])
		sum := utility.BytesToFloat64(row[1])
		res[id] = sum
	}

	return res, nil
}
