-- Write your query below
WITH cte AS (
    SELECT p.product_id, MAX(o.order_date) AS recent
    FROM orders AS o
    JOIN products AS p ON o.product_id = p.product_id
    GROUP BY p.product_id
)
SELECT p.product_name, p.product_id, o.order_id, o.order_date
FROM orders AS o
JOIN products AS p ON o.product_id = p.product_id
JOIN cte AS c ON p.product_id = c.product_id
WHERE o.order_date = c.recent
ORDER BY p.product_name, p.product_id, o.order_id ASC
;