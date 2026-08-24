-- Write your query below
SELECT t1.transaction_id
FROM transactions AS t1
WHERE t1.amount = (
    SELECT MAX(t2.amount)
    FROM transactions AS t2
    WHERE t2.day::date = t1.day::date
)
ORDER BY transaction_id ASC
;