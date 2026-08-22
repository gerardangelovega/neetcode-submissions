-- Write your query below
WITH cte AS (
    SELECT player_id, MIN(event_date) as event_date
    FROM activity
    GROUP BY player_id
)
SELECT a.player_id, a.device_id
FROM activity AS a
JOIN cte AS c ON a.player_id = c.player_id AND a.event_date = c.event_date
;
