// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: payments.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (
    PAYMENT_ID,
    ORDER_ID,
    USER_ID,
    AMOUNT,
    METHOD,
    STATUS,
    TRANSACTION_ID
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) RETURNING payment_id, order_id, user_id, amount, method, status, transaction_id, created_at, updated_at
`

type CreatePaymentParams struct {
	PaymentID     uuid.UUID      `json:"payment_id"`
	OrderID       uuid.UUID      `json:"order_id"`
	UserID        uuid.UUID      `json:"user_id"`
	Amount        pgtype.Numeric `json:"amount"`
	Method        string         `json:"method"`
	Status        string         `json:"status"`
	TransactionID uuid.UUID      `json:"transaction_id"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, createPayment,
		arg.PaymentID,
		arg.OrderID,
		arg.UserID,
		arg.Amount,
		arg.Method,
		arg.Status,
		arg.TransactionID,
	)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.OrderID,
		&i.UserID,
		&i.Amount,
		&i.Method,
		&i.Status,
		&i.TransactionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPaymentByUserID = `-- name: GetPaymentByUserID :many
SELECT 
    payment_id, order_id, user_id, amount, method, status, transaction_id, created_at, updated_at 
FROM 
    payments 
WHERE 
    user_id = $1
`

func (q *Queries) GetPaymentByUserID(ctx context.Context, userID uuid.UUID) ([]Payment, error) {
	rows, err := q.db.Query(ctx, getPaymentByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Payment{}
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.PaymentID,
			&i.OrderID,
			&i.UserID,
			&i.Amount,
			&i.Method,
			&i.Status,
			&i.TransactionID,
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

const getPaymentsList = `-- name: GetPaymentsList :many
SELECT 
    payment_id, order_id, user_id, amount, method, status, transaction_id, created_at, updated_at 
FROM 
    payments
`

func (q *Queries) GetPaymentsList(ctx context.Context) ([]Payment, error) {
	rows, err := q.db.Query(ctx, getPaymentsList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Payment{}
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.PaymentID,
			&i.OrderID,
			&i.UserID,
			&i.Amount,
			&i.Method,
			&i.Status,
			&i.TransactionID,
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

const updatePaymentMethod = `-- name: UpdatePaymentMethod :one
UPDATE 
    payments
SET 
    method = COALESCE($2,method),
    transaction_id = COALESCE($3,transaction_id),
    updated_at = NOW()
WHERE
    payment_id = $1 RETURNING payment_id, order_id, user_id, amount, method, status, transaction_id, created_at, updated_at
`

type UpdatePaymentMethodParams struct {
	PaymentID     uuid.UUID `json:"payment_id"`
	Method        string    `json:"method"`
	TransactionID uuid.UUID `json:"transaction_id"`
}

func (q *Queries) UpdatePaymentMethod(ctx context.Context, arg UpdatePaymentMethodParams) (Payment, error) {
	row := q.db.QueryRow(ctx, updatePaymentMethod, arg.PaymentID, arg.Method, arg.TransactionID)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.OrderID,
		&i.UserID,
		&i.Amount,
		&i.Method,
		&i.Status,
		&i.TransactionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePaymentStatus = `-- name: UpdatePaymentStatus :one
UPDATE 
    payments
SET 
    STATUS = $2
WHERE
    payment_id = $1 RETURNING payment_id, order_id, user_id, amount, method, status, transaction_id, created_at, updated_at
`

type UpdatePaymentStatusParams struct {
	PaymentID uuid.UUID `json:"payment_id"`
	Status    string    `json:"status"`
}

func (q *Queries) UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (Payment, error) {
	row := q.db.QueryRow(ctx, updatePaymentStatus, arg.PaymentID, arg.Status)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.OrderID,
		&i.UserID,
		&i.Amount,
		&i.Method,
		&i.Status,
		&i.TransactionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
