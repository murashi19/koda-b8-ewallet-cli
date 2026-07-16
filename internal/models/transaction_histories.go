package models

import "time"

type TransactionDetail struct {
	Type             string
	SenderWalletID   *int64
	ReceiverWalletID *int64
	Amount           int64
	Status           string
	CreatedAt        time.Time
}

type TransactionHistory struct {
	Type      string
	Amount    int64
	Status    string
	CreatedAt time.Time
}
