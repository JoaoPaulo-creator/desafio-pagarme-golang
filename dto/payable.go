package dto

import (
	"time"

	"github.com/google/uuid"
)

type Payable struct {
	ID           uuid.UUID     `json:"id"`
	status       string        `json:"status"`
	payementDate time.Time     `json:"paymentDate"`
	fee          float64       `json:"fee"`
	transactions []Transaction `json:"transactions"`
}
