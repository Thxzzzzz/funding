SELECT *
FROM
(
SELECT *
FROM
	products tp 
WHERE
	tp.product_type = 1 
	AND tp.deleted_at IS NULL
	LIMIT 1
	OFFSET 0
) p
	LEFT JOIN 
	product_packages pkg ON p.id = pkg.product_id