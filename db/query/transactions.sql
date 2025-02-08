-- name: CreateTransaction :one
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
    STATUS = $2
WHERE 
    transaction_id = $1 RETURNING *;
