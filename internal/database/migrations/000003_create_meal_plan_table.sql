-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS days_of_week (
  id INT PRIMARY KEY,
  name VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS meal_types (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  sort_order INT
);

CREATE TABLE IF NOT EXISTS meal_plan (
  id SERIAL PRIMARY KEY,
  day_id INT REFERENCES days_of_week(id),
  meal_type_id INT REFERENCES meal_types(id),
  dish_id INT REFERENCES dishes(id) ON DELETE CASCADE,

  CONSTRAINT unique_meal_entry UNIQUE(day_id, meal_type_id)
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS days_of_week;
DROP TABLE IF EXISTS meal_types;
DROP TABLE IF EXISTS meal_plan;
-- +goose StatementBegin
-- +goose StatementEnd
