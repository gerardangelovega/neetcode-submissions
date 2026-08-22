-- Write your query below
WITH cte AS (
    SELECT d.id, MAX(e.salary) AS max_salary
    FROM employee AS e
    JOIN department AS d ON e.department_id = d.id
    GROUP BY d.id
)
SELECT d.name AS department,
       e.name AS employee,
       e.salary
FROM employee AS e
JOIN department AS d ON e.department_id = d.id
JOIN cte AS c ON d.id = c.id
WHERE e.department_id = c.id
  AND e.salary = c.max_salary
ORDER BY e.salary DESC
;