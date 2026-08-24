-- Write your query below
WITH cte AS (
    SELECT CASE
            WHEN (log_id-1) NOT IN (SELECT log_id FROM logs) THEN log_id
            ELSE NULL
        END AS start_id,
        CASE
            WHEN (log_id+1) NOT IN (SELECT log_id FROM logs) THEN log_id
            ELSE NULL
       END AS end_id
    FROM logs
)
SELECT c1.start_id, MIN(c2.end_id) AS end_id
FROM (SELECT start_id FROM cte WHERE start_id IS NOT NULL) AS c1
JOIN (SELECT end_id FROM cte WHERE end_id IS NOT NULL) AS c2 ON c1.start_id <= c2.end_id
GROUP BY c1.start_id
ORDER BY c1.start_id
;
