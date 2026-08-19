-- Write your query below
SELECT c.customer_id, c.name
FROM customers AS c
JOIN orders AS o ON c.customer_id = o.customer_id
JOIN product AS p ON o.product_id = p.product_id
GROUP BY c.customer_id
HAVING SUM(CASE
        WHEN o.order_date NOT BETWEEN '2020-06-01' AND '2020-06-30' THEN 0
        ELSE o.quantity * p.price
    END) >= 100
AND SUM(CASE
        WHEN o.order_date NOT BETWEEN '2020-07-01' AND '2020-07-31' THEN 0
        ELSE o.quantity * p.price
    END) >= 100
;