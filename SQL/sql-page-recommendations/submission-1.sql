-- Write your query below
WITH user_friends AS (
    SELECT 
        CASE
            WHEN user1_id = 1 THEN user2_id
            WHEN user2_id = 1 THEN user1_id
        END AS user_id
    FROM friendship
    WHERE user1_id = 1 OR user2_id = 1
)
SELECT DISTINCT l.page_id AS recommended_page
FROM likes AS l
WHERE l.user_id IN (SELECT user_id FROM user_friends)
  AND l.page_id NOT IN (SELECT page_id FROM likes WHERE user_id = 1)
;