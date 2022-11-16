package repo_implementation

import (
	"Avito-Challenge/src/database/pgdb"
	"Avito-Challenge/src/models"
	"Avito-Challenge/src/utility"
	"fmt"
)

const (
	insertReservation  = `insert into accounting(service_id, amount, completion_date) values($1, $2, $3) returning id;`
	selectSum          = `select service_id, sum(amount) from accounting where $1 <= completion_date and completion_date < $2 group by service_id;`
	insertTransaction  = `insert into transactions(user_id, other, reason, date, amount) values($1, $2, $3, $4, $5);`
	selectTransactions = `select s.* from (select row_number() over (order by %s) rn, * from transactions where user_id = $1) s where $2 <= rn and rn < $3;`
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

func (ar *AccountingRepository) LogTransaction(transaction *models.Transaction) error {
	_, err := ar.db.Exec(insertTransaction, transaction.UserId, transaction.Other, transaction.Reason, transaction.Date, transaction.Amount)
	return err
}

func (ar *AccountingRepository) GetTransactions(id int, size int, page int, sortBy string) (*[]models.Transaction, error) {
	start := (page - 1) * size
	end := start + size
	var orderBy string
	if sortBy == "date" {
		orderBy = "completion_date"
	} else {
		orderBy = "amount"
	}
	data, err := ar.db.Query(fmt.Sprintf(selectTransactions, orderBy), id, start, end)
	if err != nil {
		fmt.Println("Failed to get transactions:", err)
		return nil, err
	}
	res := make([]models.Transaction, 0)
	for _, row := range data {
		t := models.Transaction{
			UserId: utility.BytesToInt(row[0]),
			Other:  utility.BytesToString(row[1]),
			Reason: utility.BytesToString(row[2]),
			Date:   utility.BytesToString(row[3]),
			Amount: utility.BytesToFloat64(row[4]),
		}
		res = append(res, t)
	}
	return &res, nil
}
