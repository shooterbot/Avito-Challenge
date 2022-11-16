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
	CreateBalance        = `insert into balances(user_id, amount) values($1, $2);`
	ReduceAmountByUserId = `update balances set amount = amount - $2 where user_id = $1;`
	InsertReservation    = `insert into reservations(user_id, service_id, order_id, amount) values($1, $2, $3, $4);`
	selectReserved       = `select amount from reservations where order_id = $1 and service_id = $2;`

	// DeleteReservation Вообще, для удаления записи о резерве достаточно только ID заказа, он должен быть уникальным
	// (и, например, ссылаться на заказ, в ктором уже хранится ID услуги), иначе в нем нет смысла. Но, раз по ТЗ
	// остальные параметры тоже передаются в запросе - можно сделать проверку и по ним
	DeleteReservation = `delete from reservations where user_id = $1 and service_id = $2 and order_id = $3 and amount = $4;`
)

type BalanceRepository struct {
	db *pgdb.DBManager
}

func NewBalanceRepository(manager *pgdb.DBManager) *BalanceRepository {
	return &BalanceRepository{db: manager}
}

func (br *BalanceRepository) GetByUserId(id int) (float64, error) {
	var res float64
	data, err := br.db.Query(selectAmountByUserId, id)
	if err != nil {
		fmt.Printf("Failed to get balance from db\n")
	} else if len(data) == 1 {
		res = utility.BytesToFloat64(data[0][0])
	} else {
		res = 0
	}

	return res, err
}

func (br *BalanceRepository) AddByUserId(id int, amount float64) error {
	data, err := br.db.Query(selectAmountByUserId, id)
	if err != nil {
		return errors.New("Error while getting user amount")
	}
	if len(data) == 0 {
		_, err = br.db.Exec(CreateBalance, id, amount)
	} else {
		_, err = br.db.Exec(AddAmountByUserId, id, amount)
	}
	if err != nil {
		err = errors.New("Error while updating user amount")
	}
	return err
}

func (br *BalanceRepository) Withdraw(id int, amount float64) error {
	affected, err := br.db.Exec(ReduceAmountByUserId, id, amount)
	if err != nil && affected == 0 {
		err = errors.New("Error while updating user amount")
	}
	return err
}

func (br *BalanceRepository) AddReservation(userId, orderId, serviceId int, amount float64) error {
	affected, err := br.db.Exec(InsertReservation, userId, serviceId, orderId, amount)
	if err != nil && affected == 0 {
		err = errors.New("Error while updating user reservation")
	}
	return err
}

func (br *BalanceRepository) GetReserved(orderId int, serviceId int) (float64, error) {
	var res float64
	data, err := br.db.Query(selectReserved, orderId, serviceId)
	if err != nil {
		fmt.Printf("Failed to get reserved amount from db\n")
	} else if len(data) == 1 {
		res = utility.BytesToFloat64(data[0][0])
	} else {
		err = errors.New("Could not find a user matching the given id")
		res = -1
	}

	return res, err
}

func (br *BalanceRepository) DeleteReservation(userId, orderId, serviceId int, amount float64) error {
	affected, err := br.db.Exec(DeleteReservation, userId, serviceId, orderId, amount)
	if err != nil || affected == 0 {
		err = errors.New("Error while updating user reservation")
	}
	return err
}
