-- name: CreateTargetAccount :execresult
INSERT INTO target_account (
    name,
    atm_bank_code,
    bank_detail,
    account_type,
    account_number,
    bank,
    bank_branch,
    description,
    target_type,
    status,
    fourth_digit,
    customer_id,
    currency,
    account_type_code,
    amount,
    is_favorite 
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetTargetAccount :one
SELECT * FROM target_account
WHERE id = ? LIMIT 1;

-- name: GetTargetAccountForUpdate :one
SELECT * FROM target_account
WHERE id = ? LIMIT 1
FOR UPDATE;

-- name: ListTargetAccount :many
SELECT * FROM target_account
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateTargetAccount :execresult
UPDATE target_account
SET description = ?
WHERE id = ?;

-- name: DeleteTargetAccount :execresult
UPDATE target_account
SET status = ?
WHERE id = ?;