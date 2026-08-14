-- Write your query below
SELECT c.customer_id, c.customer_name
FROM customers AS c
WHERE c.customer_id IN (
    SELECT o.customer_id FROM orders AS o WHERE product_name = 'A'
)
AND c.customer_id IN (
    SELECT o.customer_id FROM orders AS o WHERE product_name = 'B'
)
AND c.customer_id NOT IN (
    SELECT o.customer_id FROM orders AS o WHERE product_name = 'C'
)
ORDER BY c.customer_name ASC;
