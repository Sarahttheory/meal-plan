-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS dishes (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  calories INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS dishes;
-- +goose StatementBegin
-- +goose StatementEnd
