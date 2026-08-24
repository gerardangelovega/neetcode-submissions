-- Write your query below
SELECT c.name AS customer_name,
       c.customer_id,
       o.order_id,
       o.order_date
FROM customers AS c
JOIN orders AS o on c.customer_id = o.customer_id
WHERE o.order_date IN (
    SELECT o2.order_date
    FROM orders AS o2
    WHERE o2.customer_id = o.customer_id
    ORDER BY o2.order_date DESC
    LIMIT 3
)
ORDER BY c.name, c.customer_id ASC, o.order_date DESC
;