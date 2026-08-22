-- Write your query below
WITH teams AS (
    SELECT team_id, COUNT(*) AS team_size
    FROM employee
    GROUP BY team_id
)
SELECT e.employee_id, t.team_size 
FROM employee AS e
JOIN teams AS t ON e.team_id = t.team_id
;