-- Write your query below
SELECT c.customer_id, c.customer_name
FROM customers AS C
WHERE EXISTS (
    SELECT 1
    FROM orders AS o
    WHERE c.customer_id = o.customer_id AND product_name = 'A'
)
AND EXISTS (
    SELECT 1
    FROM orders AS o
    WHERE c.customer_id = o.customer_id AND product_name = 'B'
)
AND NOT EXISTS (
    SELECT 1
    FROM orders AS o
    WHERE c.customer_id = o.customer_id AND product_name = 'C'
)
ORDER BY c.customer_name ASC;