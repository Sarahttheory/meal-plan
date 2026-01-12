-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ingredients (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredients_in_dish (
  dish_id INTEGER REFERENCES dishes(id) ON DELETE CASCADE,
  ingredient_id INTEGER REFERENCES ingredients(id) ON DELETE CASCADE,
  PRIMARY KEY (dish_id, ingredient_id)
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS ingredients_in_dish;
-- +goose StatementBegin
-- +goose StatementEnd
