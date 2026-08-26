SELECT DISTINCT s1.student_id, s1.student_name
FROM student AS s1
JOIN exam AS e1 ON s1.student_id = e1.student_id
WHERE s1.student_id NOT IN (
    SELECT e2.student_id
    FROM exam AS e2
    WHERE (e2.exam_id, e2.score) IN (
        SELECT e3.exam_id, MAX(e3.score) FROM exam AS e3 GROUP BY e3.exam_id
    ) OR (e2.exam_id, e2.score) IN (
        SELECT e3.exam_id, MIN(e3.score) FROM exam AS e3 GROUP BY e3.exam_id
    )
);