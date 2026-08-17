-- Write your query below
WITH delivery_status AS (
    SELECT 
        CASE
            WHEN customer_pref_delivery_date = order_date THEN 1
            ELSE 0
        END AS status
    FROM delivery
)
SELECT ROUND(AVG(status) * 100, 2) AS immediate_percentage
FROM delivery_status;