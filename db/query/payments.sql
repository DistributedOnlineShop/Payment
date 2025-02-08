-- name: CreatePayment :one
INSERT INTO payments (
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
    $6
) RETURNING *;

-- name: GetPaymentByUserID :one
SELECT 
    * 
FROM 
    payments 
WHERE 
    user_id = $1;

-- name: GetPaymentsList :many
SELECT 
    * 
FROM 
    payments;

-- name: UpdatePaymentMethod :one
UPDATE 
    payments
SET 
    method = COALESCE($2,method),
    transaction_id = COALESCE($3,transaction_id),
    updated_at = NOW()
WHERE
    payment_id = $1 RETURNING *;

-- name: UpdatePaymentStatus :one
UPDATE 
    payments
SET 
    STATUS = $2
WHERE
    payment_id = $1 RETURNING *;