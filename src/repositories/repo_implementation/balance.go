package repo_implementation

import (
	"Avito-Challenge/src/database/pgdb"
	"Avito-Challenge/src/utility"
	"fmt"
)

const (
	selectAmountByUserId = `select amount from balances where user_id = $1;`
)

type BalanceRepository struct {
	db *pgdb.DBManager
}

func (br *BalanceRepository) GetByUserId(id int) (float64, error) {
	data, err := br.db.Query(selectAmountByUserId, id)
	if err != nil {
		fmt.Printf("Failed to get balance from db\n")
	}

	res := utility.BytesToFloat64(data[0][0])

	return res, err
}
