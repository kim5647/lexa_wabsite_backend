-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByEmail :one
SELECT id, name, hash_password, phone, email
FROM users
WHERE email = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    name,
    hash_password,
    phone,
    email
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, name, hash_password, phone, email;