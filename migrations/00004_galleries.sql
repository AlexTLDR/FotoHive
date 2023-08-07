-- +goose Up
-- +goose StatementBegin
CREATE TABLE galleries (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id),
  title TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE galleries;
-- +goose StatementEnd
