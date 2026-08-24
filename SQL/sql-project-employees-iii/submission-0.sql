-- Write your query below
SELECT p1.project_id, p1.employee_id
FROM project AS p1
JOIN employee AS e1 ON p1.employee_id = e1.employee_id
WHERE e1.experience_years = (
    SELECT MAX(experience_years)
    FROM project AS p2
    JOIN employee AS e2 ON p2.employee_id = e2.employee_id
    WHERE p2.project_id = p1.project_id
)
;