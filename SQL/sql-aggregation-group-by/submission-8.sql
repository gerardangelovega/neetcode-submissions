CREATE TABLE olympic_medals (
    country TEXT,
    athlete_name TEXT,
    gold INTEGER,
    silver INTEGER,
    bronze INTEGER
);


INSERT INTO olympic_medals (country, athlete_name, gold, silver, bronze) VALUES
('USA', 'John', 0, 1, 0),
('CAN', 'Jesse', 1, 0, 0),
('MEX', 'Juan', 0, 0, 1),
('BRA', 'Rafael', 0, 1, 0),
('USA', 'Jill', 0, 0, 1),
('USA', 'Jack', 1, 1, 1),
('UK', 'James', 0, 0, 2),
('CAN', 'Joseph', 1, 0, 1);
-- Do not modify above this line. --

SELECT country, SUM(gold) as gold_medals, SUM(silver) as silver_medals, SUM(bronze) as bronze_medals FROM olympic_medals GROUP BY country;