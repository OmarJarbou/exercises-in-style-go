-- name: AddDocument :one
INSERT INTO documents (id, name, created_at, updated_at)
VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: FindDocumentBYName :one
SELECT * FROM documents
WHERE name=?;