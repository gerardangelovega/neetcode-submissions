-- Write your query below
SELECT s.seller_name
FROM seller AS s
WHERE s.seller_id NOT IN (
    SELECT o.seller_id
    FROM orders AS o
    WHERE EXTRACT(YEAR FROM sale_date) = 2020
)
ORDER BY seller_name ASC;