-- +goose Up
CREATE TABLE characters(
    id UUID PRIMARY KEY,
    value CHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    word_id UIID,
    CONSTRAINT fk_words_characters
    FOREIGN KEY (word_id)
    REFERENCES words(id)
);

-- +goose Down
DROP TABLE characters;