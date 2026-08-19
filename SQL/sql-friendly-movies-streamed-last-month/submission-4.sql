-- Write your query below
SELECT DISTINCT c.title
FROM content AS c
JOIN tv_program AS t ON c.content_id = t.content_id
WHERE c.kids_content = 'Y'
  AND c.content_type = 'Movies'
  AND t.program_date LIKE '2020-06-%'
;