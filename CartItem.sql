SELECT
	c.id,c.user_id,c.product_package_id,c.nums,c.checked,pkg.product_id,
	p.name AS product_name,pkg.price,pkg.stock,pkg.image_url,pkg.description
FROM
	carts c
JOIN
	product_packages pkg ON c.product_package_id = pkg.id
JOIN
	products p ON pkg.product_id = p.id
WHERE
	c.deleted_at IS NULL  AND
	p.deleted_at IS NULL  AND
	pkg.deleted_at IS NULL AND
	c.user_id = 20003 AND
	c.product_package_id = 111111114