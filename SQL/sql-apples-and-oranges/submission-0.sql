-- Write your query below
WITH sales_details AS (
    SELECT
        sale_date,
        SUM(CASE
                WHEN fruit = 'apples' THEN sold_num
                ELSE 0
            END
        ) AS apples_sales,
        SUM(CASE
                WHEN fruit = 'oranges' THEN sold_num
                ELSE 0
            END
        ) AS oranges_sales
    FROM sales
    GROUP BY sale_date
)
SELECT sale_date, (apples_sales - oranges_sales) AS diff
FROM sales_details;