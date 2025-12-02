-- name: GetAccountByName :one
SELECT *
FROM account
WHERE name = $1 LIMIT 1;

