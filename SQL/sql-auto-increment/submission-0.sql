CREATE SEQUENCE gov_id START WITH 1000 INCREMENT BY 3;

CREATE TABLE IF NOT EXISTS gov_employee (
  id INTEGER GENERATED ALWAYS AS IDENTITY,
  gov_id INTEGER DEFAULT nextval('gov_id'),
  name TEXT
);

-- Do not modify below this line --
INSERT INTO gov_employee (name) 
  VALUES
      ('John Doe'),
      ('Jane Doe'),
      ('Jim Beam');

SELECT * FROM gov_employee;
