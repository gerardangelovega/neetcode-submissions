-- Write your query below
SELECT s.id, s.name
FROM students AS s
LEFT JOIN departments AS d ON s.department_id = d.id
WHERE d.id IS NULL
;