-- Write your query below
SELECT team_id,
       team_name,
       COALESCE(SUM(
            CASE
                WHEN t.team_id = m.host_team AND m.host_goals > m.guest_goals THEN 3
                WHEN t.team_id = m.guest_team AND m.host_goals < m.guest_goals THEN 3
                WHEN m.host_goals = m.guest_goals THEN 1
                ELSE 0
            END
       ), 0) AS num_points
FROM teams AS t
LEFT JOIN matches AS m ON t.team_id = m.host_team OR t.team_id = m.guest_team
GROUP BY team_id, team_name
ORDER BY num_points DESC, team_id ASC;