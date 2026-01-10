CREATE TABLE days_of_week (
  id INT PRIMARY KEY,
  name VARCHAR(20) NOT NULL
);

CREATE TABLE meal_types (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  sort_order INT
);

CREATE TABLE meal_plan (
  id SERIAL PRIMARY KEY,
  day_id INT REFERENCES days_of_week(id),
  meal_type_id INT REFERENCES meal_types(id),
  dish_id INT REFERENCES dishes(id) ON DELETE CASCADE,

  CONSTRAINT unique_meal_entry UNIQUE(day_id, meal_type_id, dish_id)
);
