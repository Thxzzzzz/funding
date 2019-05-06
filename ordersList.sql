SELECT
	o.id,o.user_id,p.user_id AS seller_id,su.nickname AS seller_nickname,
	o.product_package_id,o.nums,o.unit_price,pkg.product_id,
	p.name AS product_name,pkg.price,pkg.stock,pkg.image_url,pkg.description,
	o.created_at,o.status AS order_status,o.total_price
FROM
	orders o
JOIN
	product_packages pkg ON o.product_package_id = pkg.id
JOIN
	products p ON p.id = o.product_id
JOIN 
	users su ON su.id = p.user_id	
WHERE
	o.deleted_at IS NULL  AND
	p.deleted_at IS NULL  AND
	pkg.deleted_at IS NULL AND
	o.user_id = 20003
ORDER BY
	o.created_at DESC
LIMIT 5 OFFSET 0