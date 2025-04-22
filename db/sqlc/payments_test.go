package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"Payment/util"
)

func CreateRandomPayment(t *testing.T, userID, orderID, tranID uuid.UUID) Payment {
	data := CreatePaymentParams{
		PaymentID:     util.CreateUUID(),
		OrderID:       orderID,
		UserID:        userID,
		Amount:        util.GenerateNumeric(),
		Method:        util.GenerateMethod(),
		Status:        util.GeneratePaymentStatus(),
		TransactionID: tranID,
	}

	payment, err := testStore.CreatePayment(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, payment)
	require.Equal(t, data.PaymentID, payment.PaymentID)
	require.Equal(t, data.OrderID, payment.OrderID)
	require.Equal(t, data.UserID, payment.UserID)
	require.Equal(t, data.Amount, payment.Amount)
	require.Equal(t, data.Method, payment.Method)
	require.Equal(t, data.Status, payment.Status)
	require.Equal(t, data.TransactionID, payment.TransactionID)
	require.NotZero(t, payment.CreatedAt)

	return payment
}

func TestCreatePayment(t *testing.T) {
	id := util.CreateUUID()
	orderId := util.CreateUUID()
	tranId := CreateRandomTransaction(t, id)
	CreateRandomPayment(t, id, orderId, tranId.TransactionID)
}

func TestGetPaymentByUserID(t *testing.T) {
	id := util.CreateUUID()
	for i := 0; i < 10; i++ {
		orderId := util.CreateUUID()
		tranId := CreateRandomTransaction(t, id)
		CreateRandomPayment(t, id, orderId, tranId.TransactionID)
	}

	payments, err := testStore.GetPaymentByUserID(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, payments)
	require.GreaterOrEqual(t, len(payments), 10)
}

func TestGetPaymentsList(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := util.CreateUUID()
		orderId := util.CreateUUID()
		tranId := CreateRandomTransaction(t, id)
		CreateRandomPayment(t, id, orderId, tranId.TransactionID)
	}

	payments, err := testStore.GetPaymentsList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, payments)
	require.GreaterOrEqual(t, len(payments), 10)

}

func TestUpdatePaymentMethod(t *testing.T) {
	id := util.CreateUUID()
	orderId := util.CreateUUID()
	tranId := CreateRandomTransaction(t, id)
	payment := CreateRandomPayment(t, id, orderId, tranId.TransactionID)

	newData := UpdatePaymentMethodParams{
		PaymentID:     payment.PaymentID,
		Method:        util.GenerateMethod(),
		TransactionID: payment.TransactionID,
	}

	updatedPayment, err := testStore.UpdatePaymentMethod(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, updatedPayment)
	require.Equal(t, payment.PaymentID, updatedPayment.PaymentID)
	require.Equal(t, newData.TransactionID, updatedPayment.TransactionID)
	require.NotZero(t, updatedPayment.UpdatedAt)
}
