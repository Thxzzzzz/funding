-- 开启事务
-- BEGIN;

-- 自动检查并退款
UPDATE
orders  o
LEFT JOIN products p
ON o.product_id = p.id
SET
o.status = 7,
o.last_status = 2,
o.refund_reason = "众筹失败，自动退款"
WHERE
o.status = 2
AND p.end_time < DATE(NOW())
AND p.current_price < p.target_price
AND p.deleted_at IS NULL
AND o.deleted_at IS NULL;

-- 查询结果
SELECT *
FROM orders;

-- 回滚
-- ROLLBACK;
 
-- SELECT o.id
-- FROM orders AS o
-- LEFT JOIN products p
-- ON o.product_id = p.id
-- WHERE
-- o.`status` = 2
-- AND p.end_time < DATE(NOW())
-- AND p.current_price < p.target_price
-- AND p.deleted_at IS NULL
-- 	AND o.deleted_at IS NULL
 
-- SELECT *
-- FROM products
-- WHERE
-- end_time < DATE(NOW())
-- AND current_price < target_price
-- AND deleted_at IS NULL