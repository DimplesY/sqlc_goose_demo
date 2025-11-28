-- name: GetAccountById :one
SELECT *
FROM account
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO account (name, email, password)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :exec
UPDATE account
SET name = $1, email = $2, password = $3
WHERE id = $4;

-- name: DeleteUser :exec
DELETE FROM account
WHERE id = $1;
