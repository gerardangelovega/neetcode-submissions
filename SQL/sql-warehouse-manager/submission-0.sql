-- Write your query below
SELECT w.name as warehouse_name, SUM(p.width * p.height * p.length * w.units) AS volume
FROM warehouse AS w
LEFT JOIN products AS p ON w.product_id = p.product_id
GROUP BY w.name;