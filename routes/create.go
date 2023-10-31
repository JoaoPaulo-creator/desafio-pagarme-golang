package routes

import (
	"database/sql"
	"github.com/desafio-pagarme-golang/db"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"

	"encoding/json"
)

type (
	Transaction struct {
		ID                 uuid.UUID      `json:"id"`
		TransactionValue   sql.NullInt64  `json:"transaction_value"`
		ProductDescription sql.NullString `json:"product_description"`
		CardNumber         sql.NullInt64  `json:"card_number"`
		NameInCard         sql.NullString `json:"name_in_card"`
		CardExpirationDate sql.NullTime   `json:"card_expiration_date"`
		Cvv                sql.NullInt32  `json:"cvv"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) ValidateBody(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func CreateTransactionHandler(c echo.Context) (err error) {
	var transaction db.Transaction

	parsedBody, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, parsedBody)
}
