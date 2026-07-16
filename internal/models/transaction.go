package models

import "time"

type Transaction struct {
	ID                int64
	SenderWalletID    int64
	ReceiverWalletID  int64
	TransactionTypeID int64
	Amount            int64
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
