-- Write your query below
WITH logs AS (
    SELECT 'failed' AS state,
            fail_date AS log_date
    FROM failed
    WHERE fail_date >= '2019-01-01' AND fail_date  <= '2019-12-31'
    UNION ALL
    SELECT 'succeeded' AS state,
            success_date AS log_date
    FROM succeeded
    WHERE success_date >= '2019-01-01' AND success_date <= '2019-12-31'
),
ranked AS (
    SELECT state,
       log_date,
       ROW_NUMBER() OVER (ORDER BY log_date) AS date_rank,
       ROW_NUMBER() OVER (PARTITION BY state ORDER BY log_date) AS state_rank
    FROM logs
),
grouped AS (
    SELECT state,
       log_date,
       date_rank - state_rank AS date_group
    FROM ranked
)
SELECT state AS period_state,
       MIN(log_date) AS start_date,
       MAX(log_date) AS end_date
FROM grouped
GROUP BY state, date_group
ORDER BY start_date
;