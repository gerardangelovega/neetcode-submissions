-- Write your query below
SELECT c.title
FROM content AS c
JOIN tv_program AS t ON c.content_id = t.content_id
WHERE c.kids_content = 'Y'
  AND c.content_type  = 'Movies'
  AND t.program_date BETWEEN '2020-05-31 23:59:59' AND '2020-06-30 23:59:59'
GROUP BY c.title
;