-- Write your query below
WITH cte AS (
    SELECT c.customer_id,
           p.product_id,
           COUNT(p.product_id),
           p.product_name
    FROM customers AS c
    JOIN orders AS o ON c.customer_id = o.customer_id
    JOIN products AS p ON o.product_id = p.product_id
    GROUP BY c.customer_id, p.product_id, p.product_name
)
SELECT customer_id,
       product_id,
       product_name
FROM cte AS c1
WHERE count = (
    SELECT MAX(count)
    FROM cte AS c2
    WHERE c1.customer_id = c2.customer_id
)
;