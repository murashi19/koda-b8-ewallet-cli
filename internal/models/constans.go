package models

const (
	TransactionStatusSuccess = "SUCCESS"
	TransactionStatusFailed  = "FAILED"
	TransactionStatusPending = "PENDING"
)
const (
	TransactionTypeTransfer int64 = 1
	TransactionTypeTopUp    int64 = 2
	TransactionTypeWithdraw int64 = 3
)
const DefaultCurrency = "IDR"
const DefaultBalance = 0
