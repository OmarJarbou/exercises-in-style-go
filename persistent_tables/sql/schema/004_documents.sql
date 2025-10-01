-- +goose Up
CREATE UNIQUE INDEX idx_documents_name ON documents(name);

-- +goose Down
DROP INDEX IF EXISTS idx_documents_name;
