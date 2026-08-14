-- Write your query below
SELECT u.name,
       COALESCE(SUM(r.distance), 0) AS travelled_distance
FROM users AS u
LEFT JOIN rides AS r ON u.id = r.user_id
GROUP BY u.name
ORDER BY travelled_distance DESC, u.name ASC;