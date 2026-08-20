-- Write your query below
SELECT ABS(s.x - e.x) AS shortest
FROM point AS s
CROSS JOIN point AS e
WHERE s.x != e.x
ORDER BY shortest ASC
LIMIT 1
;