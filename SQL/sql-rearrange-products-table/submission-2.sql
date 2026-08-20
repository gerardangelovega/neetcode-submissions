-- Write your query below
SELECT p.product_id, v.store, v.price
FROM products AS p
CROSS JOIN LATERAL (
    VALUES ('store1', p.store1),
           ('store2', p.store2),
           ('store3', p.store3)
) AS v(store, price)
WHERE v.price IS NOT NULL
;