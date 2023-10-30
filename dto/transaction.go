package dto

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                 uuid.UUID `json:"id"`
	value              float64   `json:"value"`
	productDescription string    `json:"productDescription"`
	cardNumber         int16     `json:"cardNumber"`
	nameInCard         string    `json:"nameInCard"`
	cardExpirationDate time.Time `json:"cardExpirationDate"`
	cvv                int8      `json:"cvv"`
}
