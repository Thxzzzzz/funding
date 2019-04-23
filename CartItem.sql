SELECT
c.*,pkg.product_id,p.name,pkg.image_url,pkg.description
FROM
	carts c
JOIN
	product_packages pkg ON c.product_package_id = pkg.id
JOIN
	products p ON pkg.product_id = p.id
WHERE
	c.user_id = 20003 AND
	c.product_package_id = 111111114