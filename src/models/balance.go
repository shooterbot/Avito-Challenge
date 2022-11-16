package models

type IncomingTransaction struct {
	UserId int     `json:"userId"`
	Amount float64 `json:"amount"`
	Other  string  `json:"other"`
	Reason string  `json:"reason"`
}

type Reservation struct {
	UserId    int     `json:"userId"`
	OrderId   int     `json:"orderId"`
	ServiceId int     `json:"serviceId"`
	Amount    float64 `json:"amount"`
}

type Transfer struct {
	SourceUserId int     `json:"sourceUserId"`
	DestUserId   int     `json:"destinationUserId"`
	Amount       float64 `json:"amount"`
	Reason       string  `json:"reason"`
}
