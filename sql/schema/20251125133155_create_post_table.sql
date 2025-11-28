-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post
(
    id       SERIAL PRIMARY KEY,
    title    text NOT NULL,
    body     text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post;
-- +goose StatementEnd
