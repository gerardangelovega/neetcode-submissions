-- Write your query below
SELECT DISTINCT customer_id
FROM customers
WHERE revenue > 0 AND year = 2020;