-- SELECT noc,medal,
-- CASE
-- WHEN medal = "Bronze" THEN 'A'
-- WHEN medal = "Gold" THEN 'B'
-- WHEN medal = "Silver" THEN 'C'
-- END as test
-- FROM athlete_events;

SELECT noc,
    CASE
    WHEN medal = "Gold" THEN "A"
    END as total
FROM athlete_events
WHERE noc = "USA";