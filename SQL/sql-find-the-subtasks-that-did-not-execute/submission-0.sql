-- Write your query below
WITH RECURSIVE subtasks AS (
    SELECT task_id,
           1 AS subtask_id,
           subtasks_count
    FROM tasks
    UNION ALL
    SELECT task_id,
           subtask_id + 1 AS subtask_id,
           subtasks_count
    FROM subtasks
    WHERE subtask_id < subtasks_count
)
SELECT s.task_id, s.subtask_id
FROM subtasks AS s
WHERE (s.task_id, s.subtask_id) NOT IN (
    SELECT *
    FROM executed
)
;