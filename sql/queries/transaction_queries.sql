
-- name: GetTransactionById :one
SELECT * FROM transactions WHERE id = $1;

-- name: Gettransactions :many
SELECT * FROM transactions;

-- name: UpdatetransactionsName :exec
UPDATE transactions SET card_number = $1 WHERE id = $2;

-- name: CreateTransactions :exec
INSERT INTO transactions (id, transaction_value, product_description, card_number, name_in_card, card_expiration_date, cvv) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: DeleteTransactions :exec
DELETE FROM transactions WHERE id = $1;
