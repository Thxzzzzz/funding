CREATE TABLE `orders`(
	`id` BIGINT(24) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Id',
	`user_id` BIGINT(24) UNSIGNED NOT NULL COMMENT '用户 ID',
	`seller_id` BIGINT(24) UNSIGNED NOT NULL COMMENT '卖家 ID (可简化字段)',
	`product_id` BIGINT(24) UNSIGNED NOT NULL COMMENT '产品 ID(可简化字段)',
	`product_package_id` BIGINT(24) UNSIGNED NOT NULL COMMENT '产品套餐 ID',
	`nums` INT(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '购买数量',
	`unit_price` DOUBLE NOT NULL DEFAULT 0 COMMENT '单价(可简化字段)',
	`total_price` DOUBLE NOT NULL DEFAULT 0 COMMENT '总价(或许也是可简化)',
	`status` INT(5) NOT NULL DEFAULT 0 COMMENT '订单装填，0:下单,1:付款,2:配货,3:出库,200:交易成功,400:交易关闭',
	`checking_number` VARCHAR(24) NOT NULL DEFAULT '' COMMENT '物流单号',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT	'删除时间(软删除) NULL 为未删除',
	PRIMARY KEY	(`id`) USING BTREE,
	INDEX `idx_orders_deleted_at` (`deleted_at`),
	 INDEX `idx_orders_user_id` (`user_id`),
	 INDEX `idx_orders_seller_id` (`seller_id`),
	 INDEX `idx_product_id` (`product_id`),
	 INDEX `idx_orders_package_id` (`product_package_id`)
)


-- CREATE TABLE `orders` (
-- `id` BIGINT(24) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Id',
-- `user_id` BIGINT(24) UNSIGNED NOT NULL COMMENT '用户ID',
-- `product_package_id` BIGINT(24) UNSIGNED NOT NULL COMMENT '套餐ID',
-- `nums` INT(8) UNSIGNED NOT NULL COMMENT '数量',
-- `checked` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否勾选',
-- `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
-- `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
-- `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间(软删除) NULL为未删除',
-- PRIMARY KEY (`id`) USING BTREE,
-- INDEX `idx_carts_deleted_at` (`deleted_at`),
-- INDEX `idx_carts_user_id` (`user_id`) USING BTREE,
-- INDEX `idx_product_package_id` (`product_package_id`)
-- )
-- COMMENT='购物车'
-- COLLATE='utf8mb4_general_ci'
-- ENGINE=InnoDB
-- ROW_FORMAT=COMPACT
-- AUTO_INCREMENT=14
-- ;
-- 