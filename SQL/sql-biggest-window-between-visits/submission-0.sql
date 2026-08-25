-- Write your query below
WITH cte AS (
    SELECT user_id,
       visit_date,
       LEAD(visit_date, 1, '2021-1-1') OVER (
                PARTITION BY user_id 
                ORDER BY visit_date
       ) AS next_visit
    FROM user_visits
)
SELECT user_id,
       MAX(next_visit - visit_date) AS biggest_window
FROM cte
GROUP BY user_id
;