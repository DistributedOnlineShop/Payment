// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: transactions.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
    USER_ID,
    VENDOR_ID,
    ORDER_ID,
    TYPE,
    AMOUNT,
    STATUS
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING transaction_id, user_id, vendor_id, order_id, type, amount, status, created_at, updated_at
`

type CreateTransactionParams struct {
	UserID   uuid.UUID      `json:"user_id"`
	VendorID uuid.UUID      `json:"vendor_id"`
	OrderID  uuid.UUID      `json:"order_id"`
	Type     string         `json:"type"`
	Amount   pgtype.Numeric `json:"amount"`
	Status   string         `json:"status"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.UserID,
		arg.VendorID,
		arg.OrderID,
		arg.Type,
		arg.Amount,
		arg.Status,
	)
	var i Transaction
	err := row.Scan(
		&i.TransactionID,
		&i.UserID,
		&i.VendorID,
		&i.OrderID,
		&i.Type,
		&i.Amount,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTransactionsByUserId = `-- name: GetTransactionsByUserId :many
SELECT 
    transaction_id, user_id, vendor_id, order_id, type, amount, status, created_at, updated_at 
FROM 
    transactions 
WHERE 
    user_id = $1
`

func (q *Queries) GetTransactionsByUserId(ctx context.Context, userID uuid.UUID) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TransactionID,
			&i.UserID,
			&i.VendorID,
			&i.OrderID,
			&i.Type,
			&i.Amount,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransactionsList = `-- name: GetTransactionsList :many
SELECT 
    transaction_id, user_id, vendor_id, order_id, type, amount, status, created_at, updated_at 
FROM 
    transactions
`

func (q *Queries) GetTransactionsList(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionsList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TransactionID,
			&i.UserID,
			&i.VendorID,
			&i.OrderID,
			&i.Type,
			&i.Amount,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransactionStatus = `-- name: UpdateTransactionStatus :one
UPDATE 
    transactions
SET 
    STATUS = $2
WHERE 
    transaction_id = $1 RETURNING transaction_id, user_id, vendor_id, order_id, type, amount, status, created_at, updated_at
`

type UpdateTransactionStatusParams struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	Status        string    `json:"status"`
}

func (q *Queries) UpdateTransactionStatus(ctx context.Context, arg UpdateTransactionStatusParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, updateTransactionStatus, arg.TransactionID, arg.Status)
	var i Transaction
	err := row.Scan(
		&i.TransactionID,
		&i.UserID,
		&i.VendorID,
		&i.OrderID,
		&i.Type,
		&i.Amount,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
