package models

import "time"

type Wallet struct {
	ID        int64
	UserID    int64
	Balance   int64
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WalletDetail struct {
	UserName string
	Balance  int64
	Currency string
}
