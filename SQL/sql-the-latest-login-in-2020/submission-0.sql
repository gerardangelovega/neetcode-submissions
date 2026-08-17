-- Write your query below
SELECT
    user_id,
    MAX(time_stamp) AS last_stamp
FROM logins
WHERE time_stamp > '2019-12-31 23:59:59' AND time_stamp < '2021'
GROUP BY user_id
ORDER BY last_stamp DESC;