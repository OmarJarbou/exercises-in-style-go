-- +goose Up
CREATE TABLE words(
    id UUID PRIMARY KEY,
    value TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    document_id UIID,
    CONSTRAINT fk_documents_words
    FOREIGN KEY (document_id)
    REFERENCES documents(id)
);

-- +goose Down
DROP TABLE words;