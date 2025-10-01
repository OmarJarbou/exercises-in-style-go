-- name: AddCharacter :one
INSERT INTO characters (id, value, created_at, updated_at, word_id)
VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)
RETURNING *;