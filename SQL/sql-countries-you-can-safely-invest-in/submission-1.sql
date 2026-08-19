-- Write your query below
SELECT co.name AS country
FROM country AS co
JOIN person AS p ON p.phone_number LIKE CONCAT(co.country_code, '%')
JOIN calls AS ca ON p.id = ca.caller_id OR p.id = ca.callee_id
GROUP BY co.name
HAVING AVG(ca.duration) > (
    SELECT AVG(ca.duration)
    FROM person AS p
    JOIN calls AS ca ON p.id = ca.caller_id OR p.id = ca.callee_id
)
;