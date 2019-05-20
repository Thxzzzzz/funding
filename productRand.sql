SELECT *
FROM `products`
WHERE (verify_status = 1)
AND end_time < CURRENT_TIMESTAMP()
ORDER BY RAND() LIMIT 5 