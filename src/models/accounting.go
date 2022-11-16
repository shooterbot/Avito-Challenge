package models

type Transaction struct {
	UserId int
	Other  string
	Reason string
	Date   string
	Amount float64
}
