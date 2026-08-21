-- Write your query below
SELECT a1.player_id,
       a1.event_date,
       (
            SELECT SUM(a2.games_played)
            FROM activity AS a2
            WHERE a2.event_date <= a1.event_date
              AND a2.player_id = a1.player_id
       ) AS games_played_so_far
FROM activity AS a1
GROUP BY player_id, event_date
;