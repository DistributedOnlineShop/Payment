package util

import (
	"math/big"
	"math/rand/v2"

	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateNumeric() pgtype.Numeric {
	intPart := rand.IntN(100000)
	fracPart := rand.IntN(100)
	value := int64(intPart) + int64(fracPart)/100.0

	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   -2,
		Valid: true,
	}
}

func GenerateMethod() string {
	paymentMethods := []string{
		"CREDIT_CARD",
		"DEBIT_CARD",
		"PAYPAL",
		"STRIPE",
		"APPLE_PAY",
		"GOOGLE_PAY",
		"BANK_TRANSFER",
		"CASH_ON_DELIVERY",
		"CRYPTO",
		"GIFT_CARD",
		"STORE_CREDIT",
	}
	return paymentMethods[rand.IntN(len(paymentMethods))]
}

func GeneratePaymentStatus() string {
	paymentStatuses := []string{
		"PENDING",
		"PROCESSING",
		"COMPLETED",
		"FAILED",
		"CANCELLED",
		"REFUNDED",
		"PARTIALLY_REFUNDED",
		"CHARGEBACK",
	}
	return paymentStatuses[rand.IntN(len(paymentStatuses))]
}

func GenerateTransactionStatus() string {
	Statuses := []string{
		"PENDING",
		"COMPLETED",
		"FAILED",
		"CANCELLED",
		"REFUNDED",
		"IN PROGRESS",
		"UNDER REVIEW",
	}

	return Statuses[rand.IntN(len(Statuses))]
}

func GenerateTransactionTypes() string {
	Types := []string{
		"ORDER",
		"PAYMENT",
		"REFUND",
		"CANCELLATION",
		"DISCOUNT",
		"SHIPMENT",
		"POINTS",
		"ADJUSTMENT",
		"TOP-UP",
	}
	return Types[rand.IntN(len(Types))]
}
