CREATE TABLE books (
  id INTEGER,
  title TEXT,
  author TEXT
);
-- Do not modify above this line --

ALTER TABLE IF EXISTS books ADD COLUMN published_year INTEGER;

ALTER TABLE IF EXISTS books RENAME COLUMN id to isbn;

ALTER TABLE IF EXISTS books DROP COLUMN author;

-- Do not modify below this line --
SELECT column_name, data_type, column_default
FROM information_schema.columns
WHERE table_name = 'books'
ORDER BY column_name;
