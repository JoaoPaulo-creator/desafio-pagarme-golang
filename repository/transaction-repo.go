package repository

import "github.com/desafio-pagarme-golang/db"

type TransactionRepository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) *TransactionRepository {
	return &TransactionRepository{
		q: q,
	}
}

func (t *TransactionRepository) Create() {

}
