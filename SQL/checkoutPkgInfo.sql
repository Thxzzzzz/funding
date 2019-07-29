SELECT
	p.id AS product_id,pkg.id AS product_package_id,
	p.name AS product_name, pkg.price AS price,
	p.user_id AS seller_id ,u.nickname AS seller_nickname,
	pkg.image_url,pkg.description,
	pkg.stock,p.end_time
FROM
	products p LEFT JOIN
	product_packages pkg ON p.id = pkg.product_id
	JOIN users u ON p.user_id = u.id
WHERE
	p.deleted_at IS NULL 
	AND pkg.deleted_at IS NULL
 