-- +goose Up
-- +goose StatementBegin
CREATE TABLE widgets (
    id SERIAL PRIMARY KEY,
    color TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE widgets;
-- +goose StatementEnd
