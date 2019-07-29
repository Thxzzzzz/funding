CREATE TABLE `comments_replys` (
  `id` BIGINT(24) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论主键id',
  `comment_id`  BIGINT(24) UNSIGNED  NOT NULL COMMENT '评论主表id',
  `user_id`  BIGINT(24) UNSIGNED  NOT NULL COMMENT '评论者id',
  `content` varchar(512) DEFAULT NULL COMMENT '评论内容',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  	INDEX `idx_comments_replys_deleted_at` (`deleted_at`) USING BTREE,
	INDEX `idx_comments_replys_user_id` (`user_id`) USING BTREE,
	INDEX `idx_comments_replys_comment_id` (`comment_id`) USING BTREE
) 
COMMENT='评论回复表'
COLLATE='utf8mb4_general_ci'
ENGINE=INNODB
;