package models

type IncomingTransaction struct {
	UserId int     `json:"userId"`
	Amount float64 `json:"amount"`
	Origin string  `json:"origin"`
}

type Reservation struct {
	UserId    int     `json:"userId"`
	OrderId   int     `json:"orderId"`
	ServiceId int     `json:"serviceId"`
	Amount    float64 `json:"amount"`
}
