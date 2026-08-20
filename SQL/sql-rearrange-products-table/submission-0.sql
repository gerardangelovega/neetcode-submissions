-- Write your query below
WITH product_cte AS (
    SELECT product_id, 'store1' AS store, store1 AS price
    FROM products
    UNION ALL
    SELECT product_id, 'store2' AS store, store2 AS price
    FROM products
    UNION ALL
    SELECT product_id, 'store3' AS store, store3 AS price
    FROM products
)
SELECT *
FROM product_cte
WHERE price IS NOT NULL
;