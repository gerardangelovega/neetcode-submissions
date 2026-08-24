-- Write your query below
SELECT t1.transaction_id
FROM transactions AS t1
WHERE (t1.day::date, t1.amount) IN (
    SELECT t2.day::date, MAX(t2.amount)
    FROM transactions AS t2
    GROUP BY t2.day::date
)
ORDER BY transaction_id ASC
;