-- name: CreateTransaction :one
INSERT INTO transactions (
    TRANSACTION_ID,
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
    $6,
    $7
) RETURNING *;

-- name: GetTransactionsList :many
SELECT 
    * 
FROM 
    transactions;

-- name: GetTransactionsByUserId :many
SELECT 
    * 
FROM 
    transactions 
WHERE 
    user_id = $1;

-- name: UpdateTransactionStatus :one
UPDATE 
    transactions
SET 
    STATUS = $2,
    UPDATED_AT = NOW()
WHERE 
    transaction_id = $1 RETURNING *;
