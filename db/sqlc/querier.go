// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error)
	GetPaymentByUserID(ctx context.Context, userID uuid.UUID) ([]Payment, error)
	GetPaymentsList(ctx context.Context) ([]Payment, error)
	GetTransactionsByUserId(ctx context.Context, userID uuid.UUID) ([]Transaction, error)
	GetTransactionsList(ctx context.Context) ([]Transaction, error)
	UpdatePaymentMethod(ctx context.Context, arg UpdatePaymentMethodParams) (Payment, error)
	UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (Payment, error)
	UpdateTransactionStatus(ctx context.Context, arg UpdateTransactionStatusParams) (Transaction, error)
}

var _ Querier = (*Queries)(nil)
