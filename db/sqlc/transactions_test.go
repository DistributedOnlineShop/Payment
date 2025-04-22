package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"Payment/util"
)

func CreateRandomTransaction(t *testing.T, userId uuid.UUID) Transaction {
	data := CreateTransactionParams{
		TransactionID: util.CreateUUID(),
		UserID:        userId,
		VendorID:      util.CreateUUID(),
		OrderID:       util.CreateUUID(),
		Type:          util.GenerateTransactionTypes(),
		Amount:        util.GenerateNumeric(),
		Status:        util.GenerateTransactionStatus(),
	}

	tran, err := testStore.CreateTransaction(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, tran)
	require.Equal(t, tran.TransactionID, data.TransactionID)
	require.Equal(t, tran.UserID, data.UserID)
	require.Equal(t, tran.VendorID, data.VendorID)
	require.Equal(t, tran.OrderID, data.OrderID)
	require.Equal(t, tran.Type, data.Type)
	require.Equal(t, tran.Amount, data.Amount)
	require.Equal(t, tran.Status, data.Status)
	require.NotZero(t, tran.CreatedAt)

	return tran
}

func TestCreateTransaction(t *testing.T) {
	id := util.CreateUUID()
	CreateRandomTransaction(t, id)
}

func TestGetTransactionsByUserId(t *testing.T) {
	id := util.CreateUUID()
	for i := 0; i < 10; i++ {
		CreateRandomTransaction(t, id)
	}

	transactions, err := testStore.GetTransactionsByUserId(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
	require.GreaterOrEqual(t, len(transactions), 10)
}

func TestGetTransactionsList(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := util.CreateUUID()
		CreateRandomTransaction(t, id)
	}

	transactions, err := testStore.GetTransactionsList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
	require.GreaterOrEqual(t, len(transactions), 10)
}

func TestUpdateTransactionStatus(t *testing.T) {
	id := util.CreateUUID()
	tran := CreateRandomTransaction(t, id)

	newData := UpdateTransactionStatusParams{
		TransactionID: tran.TransactionID,
		Status:        util.GenerateTransactionStatus(),
	}

	transactions, err := testStore.UpdateTransactionStatus(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
	require.Equal(t, transactions.TransactionID, tran.TransactionID)
	require.NotZero(t, transactions.UpdatedAt)
}
