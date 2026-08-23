-- Write your query below
SELECT p.product_name, p.product_id, o.order_id, o.order_date
FROM orders AS o
JOIN products AS p ON o.product_id = p.product_id
WHERE o.order_date = (
    SELECT MAX(o2.order_date)
    FROM orders AS o2
    WHERE o2.product_id = o.product_id
)
ORDER BY p.product_name, p.product_id, o.order_id ASC
;