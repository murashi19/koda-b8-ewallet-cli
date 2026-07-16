package models

import "time"

type TransactionType struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}
