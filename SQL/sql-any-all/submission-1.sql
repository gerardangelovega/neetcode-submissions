CREATE TABLE customers (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT
);

CREATE TABLE orders (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    customer_id INTEGER,
    price INTEGER
);

INSERT INTO customers (name) VALUES
  ('Alice'),
  ('Bob'),
  ('Charlie'),
  ('David'),
  ('Eve'),
  ('Frank'),
  ('Grace'),
  ('Hank');

INSERT INTO orders (customer_id, price) VALUES
  (1, 50),
  (2, 100),
  (3, 150),
  (4, 200),
  (5, 250),
  (6, 300),
  (7, 70),
  (8, 400);
-- Do not modify above this line. --

SELECT c.name
FROM customers AS c
WHERE c.id = ANY (
    SELECT o.customer_id
    FROM orders AS o
    WHERE o.price < 100
)
ORDER BY c.name ASC;