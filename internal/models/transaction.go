package models

import "time"

type Transaction struct {
	ID                int64
	SenderWalletID    *int64
	ReceiverWalletID  *int64
	TransactionTypeID int64
	Amount            int64
	Status            string
	CreatedAt         time.Time
}

type TopUpRequest struct {
	UserID int64
	Amount int64
}

type TransferRequest struct {
	SenderUserID   int64
	ReceiverUserID int64
	Amount         int64
}

type WithdrawRequest struct {
	UserID int64
	Amount int64
}
