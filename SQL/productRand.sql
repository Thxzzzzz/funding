SELECT * # 获取所有字段
FROM `products` # 从 products 表查询
WHERE (verify_status = 1) # 只查询通过验证的产品
AND product_type = 1 # 产品类型为 1 
AND end_time > CURRENT_TIMESTAMP() # 只查询截止时间大于当前时间的产品（众筹未结束）
ORDER BY RAND()  # 按随机数排序（即可随机查询）
LIMIT 4 # 限制返回前四个结果