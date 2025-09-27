-- name: AddWord :one
INSERT INTO words (id, value, created_at, updated_at, document_id)
VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)
RETURNING *;

-- name: CountWordsFrequencies :many
SELECT value AS word, COUNT(*) AS freq
FROM words
WHERE document_id = ?
GROUP BY value
ORDER BY freq DESC
LIMIT 25; 