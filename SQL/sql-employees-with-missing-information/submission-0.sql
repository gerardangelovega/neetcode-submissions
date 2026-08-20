-- Write your query below
SELECT COALESCE(e.employee_id, s.employee_id) AS employee_id
FROM employees AS e
FULL JOIN salaries AS s ON e.employee_id = s.employee_id
WHERE name IS NULL OR salary IS NULL
;