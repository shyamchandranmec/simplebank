-- name: CreateTransfer :one
INSERT into transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * from transfers
WHERE id = $1;

-- name: UpdateTransfer :one
UPDATE transfers 
SET amount = $2 
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE from transfers
WHERE id = $1;