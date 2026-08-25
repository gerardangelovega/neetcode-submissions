-- Write your query below
WITH cte AS (
    SELECT e1.employee_id
    FROM employees AS e1
    JOIN employees AS e2 ON e1.manager_id = e2.employee_id
    WHERE e2.manager_id = 1 AND e1.employee_id != 1
)
SELECT employee_id
FROM employees
WHERE employee_id != 1 
  AND (manager_id = 1 OR manager_id IN (SELECT c.employee_id FROM cte AS c))
;