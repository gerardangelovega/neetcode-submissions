-- Write your query below
WITH reports AS (
    SELECT u.name,
       SUM(t.amount) AS balance
    FROM users AS u
    LEFT JOIN transactions AS t ON u.account = t.account
    GROUP BY u.name
)
SELECT *
FROM reports
WHERE balance > 10_000;