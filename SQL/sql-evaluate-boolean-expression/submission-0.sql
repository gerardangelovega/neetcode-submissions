-- Write your query below
SELECT e.left_operand,
       e.operator,
       e.right_operand,
       CASE operator
            WHEN '>' THEN v1.value > v2.value
            WHEN '<' THEN v1.value < v2.value 
            WHEN '=' THEN v1.value = v2.value
       END AS value
FROM expressions AS e
JOIN variables AS v1 on e.left_operand = v1.name
JOIN variables AS v2 on e.right_operand = v2.name;