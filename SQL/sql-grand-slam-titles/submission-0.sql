-- Write your query below
WITH grand_slams AS (
    SELECT wimbledon AS winner
    FROM championships
    UNION ALL
    SELECT fr_open AS winner
    FROM championships
    UNION ALL
    SELECT us_open AS winner
    FROM championships
    UNION ALL
    SELECT au_open AS winner
    FROM championships
)
SELECT p.player_id,
       p.player_name,
       COUNT(*) AS grand_slams_count
FROM players AS p
JOIN grand_slams AS gs ON p.player_id = gs.winner
GROUP BY p.player_id, p.player_name
;