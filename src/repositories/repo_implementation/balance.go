package repo_implementation

import (
	"Avito-Challenge/src/database/pgdb"
	"Avito-Challenge/src/utility"
	"errors"
	"fmt"
)

const (
	selectAmountByUserId = `select amount from balances where user_id = $1;`
	AddAmountByUserId    = `update balances set amount = amount + $2 where user_id = $1;`
)

type BalanceRepository struct {
	db *pgdb.DBManager
}

func (br *BalanceRepository) GetByUserId(id int) (float64, error) {
	var res float64
	data, err := br.db.Query(selectAmountByUserId, id)
	if err != nil {
		fmt.Printf("Failed to get balance from db\n")
	} else if len(data) == 1 {
		res = utility.BytesToFloat64(data[0][0])
	} else {
		err = errors.New("Could not find a user matching the given id")
		res = -1
	}

	return res, err
}

func (br *BalanceRepository) AddByUserId(id int, amount float64) error {
	affected, err := br.db.Exec(AddAmountByUserId, id, amount)
	if err != nil && affected == 0 {
		err = errors.New("Could not find a user matching the given id")
	}
	return err
}
