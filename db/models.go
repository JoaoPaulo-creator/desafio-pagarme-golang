package db

import (
	"time"

	"github.com/google/uuid"
)

type TransactionModel struct {
	ID                 uuid.UUID
	value              float64
	productDescription string
	cardNumber         int16
	nameInCard         string
	cardExpirationDate time.Time
	cvv                int8
}

type Payable struct {
	ID           uuid.UUID
	status       string
	payementDate time.Time
	fee          float64
	transactions []TransactionModel
}
