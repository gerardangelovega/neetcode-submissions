-- Write your query below
SELECT main.seat_id
FROM cinema AS main
WHERE EXISTS (
    SELECT sub.seat_id 
    FROM cinema AS sub 
    WHERE (sub.seat_id = main.seat_id + 1 OR sub.seat_id = main.seat_id - 1)
      AND sub.free = 1
) AND main.free = 1
ORDER BY main.seat_id ASC
;