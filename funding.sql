-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        8.0.15 - MySQL Community Server - GPL
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  10.1.0.5464
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 funding 的数据库结构
CREATE DATABASE IF NOT EXISTS `funding` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */;
USE `funding`;

-- 导出  表 funding.addresses 结构
CREATE TABLE IF NOT EXISTS `addresses` (
  `id` bigint(24) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `user_id` bigint(24) unsigned NOT NULL COMMENT '对应的用户id',
  `name` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '收件人姓名',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '收件人地址',
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '收件人联系电话',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间(软删除) NULL为未删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_addresses_deleted_at` (`deleted_at`),
  KEY `idx_addresses_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='用户收货地址表';

-- 正在导出表  funding.addresses 的数据：~6 rows (大约)
/*!40000 ALTER TABLE `addresses` DISABLE KEYS */;
REPLACE INTO `addresses` (`id`, `user_id`, `name`, `address`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 20003, '李小明5', '广西壮族自治区桂林市七星区桂林电子科技大学', '18500000012', '2019-04-15 14:24:56', '2019-04-23 21:24:44', NULL),
	(2, 20003, '测试大佬2', '广西壮族自治区桂林市七星区南方小清华', '18512345432', '2019-04-16 00:31:47', '2019-04-16 00:33:11', '2019-04-16 00:33:11'),
	(3, 20003, '测试大佬2', '广西壮族自治区桂林市七星区南方小清华', '18512345432', '2019-04-16 01:54:39', '2019-04-30 00:29:33', '2019-04-30 00:29:34'),
	(4, 20003, '测试大佬2', '广西壮族自治区桂林市七星区南方小清华', '18512345432', '2019-04-16 01:56:50', '2019-04-16 01:59:21', '2019-04-16 01:59:21'),
	(5, 20003, '测试大佬update', '广西壮族自治区桂林市四电之一', '18512345521', '2019-04-16 02:00:30', '2019-04-20 15:58:26', NULL),
	(6, 20003, '什么鬼', '桂林电子科技大学', '18511123456', '2019-04-20 14:51:14', '2019-04-20 15:30:24', '2019-04-20 15:30:24'),
	(7, 20003, '测试大佬123', '桂林电子科技大学', '185123123123', '2019-04-30 00:31:19', '2019-04-30 00:31:19', NULL);
/*!40000 ALTER TABLE `addresses` ENABLE KEYS */;

-- 导出  表 funding.base_table 结构
CREATE TABLE IF NOT EXISTS `base_table` (
  `id` bigint(24) unsigned NOT NULL COMMENT 'Id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间(软删除) NULL为未删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_base_table_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='基本表，用来方便创建新表';

-- 正在导出表  funding.base_table 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `base_table` DISABLE KEYS */;
/*!40000 ALTER TABLE `base_table` ENABLE KEYS */;

-- 导出  表 funding.carts 结构
CREATE TABLE IF NOT EXISTS `carts` (
  `id` bigint(24) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `user_id` bigint(24) unsigned NOT NULL COMMENT '用户ID',
  `product_package_id` bigint(24) unsigned NOT NULL COMMENT '套餐ID',
  `nums` int(8) unsigned NOT NULL COMMENT '数量',
  `checked` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否勾选',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间(软删除) NULL为未删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_carts_deleted_at` (`deleted_at`),
  KEY `idx_carts_user_id` (`user_id`) USING BTREE,
  KEY `idx_product_package_id` (`product_package_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='购物车';

-- 正在导出表  funding.carts 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `carts` DISABLE KEYS */;
REPLACE INTO `carts` (`id`, `user_id`, `product_package_id`, `nums`, `checked`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 20003, 111111113, 4, 1, '2019-04-22 20:54:49', '2019-04-29 00:17:13', NULL),
	(2, 20003, 111111114, 10, 0, '2019-04-22 21:08:45', '2019-04-25 21:11:22', '2019-04-25 21:11:23'),
	(3, 20003, 111111112, 5, 0, '2019-04-25 17:02:18', '2019-04-25 21:14:07', NULL),
	(4, 20003, 111111112, 1, 0, '2019-04-23 01:07:43', '2019-04-23 01:07:43', NULL),
	(13, 20003, 111111114, 1, 0, '2019-04-25 21:18:24', '2019-04-29 00:12:12', NULL),
	(14, 20003, 0, 0, 0, '2019-04-29 00:12:06', '2019-04-30 00:25:55', NULL),
	(15, 20003, 0, 0, 0, '2019-04-29 00:12:06', '2019-04-29 00:12:06', NULL),
	(16, 20003, 0, 0, 0, '2019-04-29 00:12:06', '2019-04-29 00:12:06', NULL);
/*!40000 ALTER TABLE `carts` ENABLE KEYS */;

-- 导出  表 funding.licenses 结构
CREATE TABLE IF NOT EXISTS `licenses` (
  `id` bigint(24) unsigned NOT NULL COMMENT 'Id',
  `compony_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公司名称',
  `user_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对应账号ID',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '联系地址',
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '官方电话',
  `license_image_url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '执照图片地址',
  `verify` tinyint(1) NOT NULL DEFAULT '2' COMMENT '审核情况 0->审核不通过 1->审核通过 2->未审核',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间(软删除) NULL为未删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_licenses_user_id` (`user_id`) USING BTREE,
  KEY `idx_licenses_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='执照信息 项目发起人资质信息';

-- 正在导出表  funding.licenses 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `licenses` DISABLE KEYS */;
REPLACE INTO `licenses` (`id`, `compony_name`, `user_id`, `address`, `phone`, `license_image_url`, `verify`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(11111111, '北京驰讯通科技有限公司', 'chixuntech', '北京海淀区三环到四环之间  中关村东路18号财智国际大厦702', '13570828180', 'https://img30.360buyimg.com/cf/jfs/t1/4374/9/6975/444185/5ba469f2E849b4a18/9ad35bddbd55accf.jpg', 1, '2019-03-09 18:40:38', '2019-03-09 18:43:35', NULL);
/*!40000 ALTER TABLE `licenses` ENABLE KEYS */;

-- 导出  表 funding.orders 结构
CREATE TABLE IF NOT EXISTS `orders` (
  `id` bigint(24) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `user_id` bigint(24) unsigned NOT NULL COMMENT '用户 ID',
  `seller_id` bigint(24) unsigned NOT NULL COMMENT '卖家 ID (可简化字段)',
  `product_id` bigint(24) unsigned NOT NULL COMMENT '产品 ID(可简化字段)',
  `product_package_id` bigint(24) unsigned NOT NULL COMMENT '产品套餐 ID',
  `nums` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '购买数量',
  `unit_price` double NOT NULL DEFAULT '0' COMMENT '单价(可简化字段)',
  `total_price` double NOT NULL DEFAULT '0' COMMENT '总价(或许也是可简化)',
  `status` int(5) NOT NULL DEFAULT '0' COMMENT '订单装填，0:下单,1:付款,2:配货,3:出库,200:交易成功,400:交易关闭',
  `checking_number` varchar(24) NOT NULL DEFAULT '' COMMENT '物流单号',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间(软删除) NULL 为未删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_orders_deleted_at` (`deleted_at`),
  KEY `idx_orders_user_id` (`user_id`),
  KEY `idx_orders_seller_id` (`seller_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_orders_package_id` (`product_package_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单信息';

-- 正在导出表  funding.orders 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;

-- 导出  表 funding.permissions 结构
CREATE TABLE IF NOT EXISTS `permissions` (
  `id` int(3) NOT NULL COMMENT '权限类型的 ID',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='权限表，用来存储各种权限的名称';

-- 正在导出表  funding.permissions 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
REPLACE INTO `permissions` (`id`, `name`) VALUES
	(1, 'All');
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;

-- 导出  表 funding.products 结构
CREATE TABLE IF NOT EXISTS `products` (
  `id` bigint(24) unsigned NOT NULL COMMENT '商品ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品名',
  `big_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '大图（用于轮播图以及详情页顶部）',
  `small_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '小图（用于商品列表）',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '发起者ID',
  `product_type` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '产品类型',
  `current_price` double unsigned NOT NULL DEFAULT '0' COMMENT '已经筹集的金额',
  `target_price` double unsigned NOT NULL DEFAULT '0' COMMENT '目标金额',
  `VerifyStatus` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '0:未通过 1：已通过 2：待审核',
  `backers` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '支持者人数',
  `end_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '众筹截止时间',
  `detail_html` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '详情介绍Html',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_products_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_products_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='产品表';

-- 正在导出表  funding.products 的数据：~8 rows (大约)
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
REPLACE INTO `products` (`id`, `name`, `big_img`, `small_img`, `user_id`, `product_type`, `current_price`, `target_price`, `VerifyStatus`, `backers`, `end_time`, `detail_html`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(11111, '鑫乐迪运动手环蓝牙耳机二合一', 'https://img30.360buyimg.com/cf/jfs/t1/21081/18/6891/106614/5c6639a2E4d110821/c5802dca70419338.jpg', 'https://img30.360buyimg.com/cf/jfs/t1/28544/16/6906/44121/5c66399bE78db06bd/0b919fb33eaccc26.jpg', 20002, 1, 1098900, 100000, 1, 2897, '2019-04-15 21:00:47', '<div class="tab-div tab-current">\r\n                <!--无缝滚动公告-->\r\n                <div class="tab-public-mess clearfix" style="display:none" id="officeTopic">\r\n                    <span class="mess-public-title"><i class="laba"></i>众筹官方：</span>\r\n                    <div class="mess-box-w" id="MBW">\r\n                        <div class="scroll-box clearfix">\r\n                            <ul class="mess-box" id="messBox1">\r\n                                <li class="mess-list"></li>\r\n                            </ul>\r\n                            <ul class="mess-box mess-box2" id="messBox2">\r\n                                <li class="mess-list"></li>\r\n                            </ul>\r\n                        </div>\r\n                    </div>\r\n                    <div class="close-btn-area">\r\n                        <span>×</span>\r\n                    </div>\r\n                </div>\r\n                <!--图片部分-->\r\n                <div class=" tab-img-group-old" style="width:733px;margin:0 auto;padding:0;text-align:center;">\r\n					                        <iframe src="http://newbuz.360buyimg.com/video/4.5/jdvideo.html?autoplay=false&amp;fuScrnEnabled=true&amp;playbackRateEnabled=true&amp;fileid=150308713247805441&amp;appid=1251412368&amp;sw=1280&amp;sh=720" width="100%" height="422px" frameborder="0"></iframe>\r\n					                </div>\r\n                <div class="                            new-story-container\r\n                            ">\r\n                        \r\n \r\n \r\n  <p class="title">众筹故事</p>\r\n  <p><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/12114/22/10496/190054/5c85f8c0Ead5ae439/a39092bf70d60ea7.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/12114/22/10496/190054/5c85f8c0Ead5ae439/a39092bf70d60ea7.jpg" alt="鑫乐迪运动手环蓝牙耳机二合一" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/20186/11/7010/369749/5c6639cfEf0deed9e/f21e34a7f63fa92a.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/20186/11/7010/369749/5c6639cfEf0deed9e/f21e34a7f63fa92a.jpg" alt="京东众筹" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/30810/29/2136/263567/5c6639d6E49863816/e55c0299b1fa7391.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/30810/29/2136/263567/5c6639d6E49863816/e55c0299b1fa7391.jpg" alt="众筹网" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/19762/14/6984/242480/5c6639dfEebbd8f69/cbad99b8d73553c2.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/19762/14/6984/242480/5c6639dfEebbd8f69/cbad99b8d73553c2.jpg" alt="科技众筹" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/32345/10/2129/315696/5c6639e7E43010088/04f21a674e084ed2.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/32345/10/2129/315696/5c6639e7E43010088/04f21a674e084ed2.jpg" alt="京东产品众筹平台" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/23583/17/6920/201057/5c6639eeEf05a4001/8c26bd69db316547.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/23583/17/6920/201057/5c6639eeEf05a4001/8c26bd69db316547.jpg" alt="众筹网站" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/21145/9/6967/281372/5c6639f5Eb796cd19/121156915b61031f.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/21145/9/6967/281372/5c6639f5Eb796cd19/121156915b61031f.jpg" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/11422/39/7833/284106/5c663a02E0e3ef54d/c27db2a8e67fb24c.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/11422/39/7833/284106/5c663a02E0e3ef54d/c27db2a8e67fb24c.jpg" alt="鑫乐迪运动手环蓝牙耳机二合一" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/22161/19/7014/173831/5c663a09E41fbe85a/d8dec990605e93dd.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/22161/19/7014/173831/5c663a09E41fbe85a/d8dec990605e93dd.jpg" alt="京东众筹" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/19222/7/7813/477969/5c6fac10E6fb5857f/76d21f15cee97561.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/19222/7/7813/477969/5c6fac10E6fb5857f/76d21f15cee97561.jpg" alt="众筹网" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/8956/33/14423/184511/5c663a17E05c4388b/b74dc7364b17e567.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/8956/33/14423/184511/5c663a17E05c4388b/b74dc7364b17e567.jpg" alt="科技众筹" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/28885/10/6847/154142/5c663a26Ec10ed136/327f2b6d17fabc7f.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/28885/10/6847/154142/5c663a26Ec10ed136/327f2b6d17fabc7f.jpg" alt="京东产品众筹平台" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/10699/9/10551/124702/5c663a2eE4a7af395/87521de309725121.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/10699/9/10551/124702/5c663a2eE4a7af395/87521de309725121.jpg" alt="众筹网站" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/21964/31/7067/137559/5c663a35Eb5769f56/c7d3a2c5afc320fc.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/21964/31/7067/137559/5c663a35Eb5769f56/c7d3a2c5afc320fc.jpg" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/14476/8/6983/133632/5c663a3dE354a4d92/a6f0372963795a1c.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/14476/8/6983/133632/5c663a3dE354a4d92/a6f0372963795a1c.jpg" alt="鑫乐迪运动手环蓝牙耳机二合一" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/10899/39/10466/81372/5c663a44Ef7603a92/cf184a14801ce22e.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/10899/39/10466/81372/5c663a44Ef7603a92/cf184a14801ce22e.jpg" alt="京东众筹" style="display: inline;"></p>\r\n  <p class="title">为什么众筹</p>\r\n  <p><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/19412/15/7293/93187/5c6cefc0E687a6a61/00ae401a398e61e5.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/19412/15/7293/93187/5c6cefc0E687a6a61/00ae401a398e61e5.jpg" alt="众筹网" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/9106/5/14712/288452/5c663a53Edfb55609/261f2d28b3cc49ce.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/9106/5/14712/288452/5c663a53Edfb55609/261f2d28b3cc49ce.jpg" alt="科技众筹" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/32790/25/2106/211297/5c663a5fE70e434ea/758bfd3c1890112c.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/32790/25/2106/211297/5c663a5fE70e434ea/758bfd3c1890112c.jpg" alt="京东产品众筹平台" style="display: inline;"><img class="lazyout-detail" src="http://img30.360buyimg.com/cf/jfs/t1/23457/21/7012/41323/5c663a65E23ac333c/9d0b44eeb672e339.jpg" data-original="http://img30.360buyimg.com/cf/jfs/t1/23457/21/7012/41323/5c663a65E23ac333c/9d0b44eeb672e339.jpg" alt="众筹网站" style="display: inline;"></p>\r\n  <p class="zc-qrcode"><img class="lazyout-detail" src="http://storage.jd.com/zc-ued-fe/zc_oa_qrcode.jpg" data-original="http://storage.jd.com/zc-ued-fe/zc_oa_qrcode.jpg" style="display: inline;"></p>\r\n  <p class="para al-center">推荐关注「京东众筹」公众号，我们会为您提供咨询服务，及时同步最新项目进展和优惠活动。</p>\r\n \r\n                </div>\r\n\r\n                <!--图片部分end-->\r\n\r\n            </div>', '2019-03-06 00:16:54', '2019-04-15 21:00:47', NULL),
	(11112, '鑫乐迪运动手环蓝牙耳机二合一2', 'https://img30.360buyimg.com/cf/jfs/t1/21081/18/6891/106614/5c6639a2E4d110821/c5802dca70419338.jpg', 'https://img30.360buyimg.com/cf/jfs/t1/28544/16/6906/44121/5c66399bE78db06bd/0b919fb33eaccc26.jpg', 20002, 1, 1098, 100000, 1, 10, '2019-04-15 21:00:49', '', '2019-03-06 00:16:54', '2019-04-15 21:00:49', NULL),
	(11113, '鑫乐迪运动手环蓝牙耳机二合一3', 'https://img30.360buyimg.com/cf/jfs/t1/21081/18/6891/106614/5c6639a2E4d110821/c5802dca70419338.jpg', 'https://img30.360buyimg.com/cf/jfs/t1/28544/16/6906/44121/5c66399bE78db06bd/0b919fb33eaccc26.jpg', 20002, 1, 100, 100000, 1, 1, '2019-04-15 21:00:51', '', '2019-03-06 00:16:54', '2019-04-15 21:00:51', NULL),
	(11114, '鑫乐迪运动手环蓝牙耳机二合一4', 'https://img30.360buyimg.com/cf/jfs/t1/21081/18/6891/106614/5c6639a2E4d110821/c5802dca70419338.jpg', 'https://img30.360buyimg.com/cf/jfs/t1/28544/16/6906/44121/5c66399bE78db06bd/0b919fb33eaccc26.jpg', 20002, 1, 78900, 100000, 1, 789, '2019-04-15 21:00:53', '', '2019-03-06 00:16:54', '2019-04-15 21:00:53', NULL),
	(11115, '鑫乐迪运动手环蓝牙耳机二合一5', 'https://img30.360buyimg.com/cf/jfs/t1/21081/18/6891/106614/5c6639a2E4d110821/c5802dca70419338.jpg', 'https://img30.360buyimg.com/cf/jfs/t1/28544/16/6906/44121/5c66399bE78db06bd/0b919fb33eaccc26.jpg', 20002, 1, 48900, 100000, 1, 21, '2019-04-15 21:00:55', '', '2019-03-06 00:16:54', '2019-04-15 21:00:55', NULL),
	(11116, '鑫乐迪运动手环蓝牙耳机二合一6', 'https://img30.360buyimg.com/cf/jfs/t1/21081/18/6891/106614/5c6639a2E4d110821/c5802dca70419338.jpg', 'https://img30.360buyimg.com/cf/jfs/t1/28544/16/6906/44121/5c66399bE78db06bd/0b919fb33eaccc26.jpg', 20002, 1, 98900, 100000, 1, 2894, '2019-04-15 21:00:56', '', '2019-03-06 00:16:54', '2019-04-15 21:00:56', NULL),
	(11117, '鑫乐迪运动手环蓝牙耳机二合一7', 'https://img30.360buyimg.com/cf/jfs/t1/21081/18/6891/106614/5c6639a2E4d110821/c5802dca70419338.jpg', 'https://img30.360buyimg.com/cf/jfs/t1/28544/16/6906/44121/5c66399bE78db06bd/0b919fb33eaccc26.jpg', 20002, 1, 1098, 100000, 1, 2897, '2019-04-15 21:00:58', '', '2019-03-06 00:16:54', '2019-04-15 21:00:58', NULL),
	(11118, '鑫乐迪运动手环蓝牙耳机二合一7', 'https://img30.360buyimg.com/cf/jfs/t1/21081/18/6891/106614/5c6639a2E4d110821/c5802dca70419338.jpg', 'https://img30.360buyimg.com/cf/jfs/t1/28544/16/6906/44121/5c66399bE78db06bd/0b919fb33eaccc26.jpg', 20002, 2, 1098900, 100000, 1, 2897, '2019-04-15 21:01:00', '', '2019-03-06 00:16:54', '2019-04-15 21:01:00', NULL);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;

-- 导出  表 funding.product_packages 结构
CREATE TABLE IF NOT EXISTS `product_packages` (
  `id` bigint(24) unsigned NOT NULL COMMENT 'Id',
  `product_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '商品 ID',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '套餐描述',
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片 Url',
  `price` double NOT NULL DEFAULT '0' COMMENT '商品价格',
  `stock` int(10) unsigned NOT NULL COMMENT '库存 剩余数量',
  `total` int(10) unsigned NOT NULL COMMENT '限额数 总备货数',
  `backer` int(10) NOT NULL DEFAULT '0' COMMENT '支持人数',
  `freight` double NOT NULL DEFAULT '0' COMMENT '运费',
  `delivery_day` int(10) unsigned NOT NULL COMMENT '项目众筹成功后多少天内发送回报（发货）',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间(软删除) NULL为未删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_product_packages_product_id` (`product_id`) USING BTREE COMMENT '商品 ID 索引',
  KEY `idx_product_packages_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='产品套餐表，对应不同的众筹套餐';

-- 正在导出表  funding.product_packages 的数据：~6 rows (大约)
/*!40000 ALTER TABLE `product_packages` DISABLE KEYS */;
REPLACE INTO `product_packages` (`id`, `product_id`, `description`, `image_url`, `price`, `stock`, `total`, `backer`, `freight`, `delivery_day`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(111111111, '11111', '感谢您的支持，您将以众筹专属价格269元获得智能手环运动蓝牙耳机二合一（运动版）1副', 'https://img30.360buyimg.com/cf/jfs/t1/8310/15/14533/53034/5c663b99E2c623c38/d9786ddce73c06b2.jpg', 269, 0, 700, 700, 0, 30, '2019-03-09 21:13:25', '2019-03-21 20:22:33', NULL),
	(111111112, '11111', '感谢您的支持，您将以众筹专属价格299元获得智能手环运动蓝牙耳机二合一（运动版）1副', 'https://img30.360buyimg.com/cf/jfs/t1/14505/24/6901/53034/5c663becE003e7921/060318eb1da6e9d1.jpg', 299, 2320, 3000, 679, 0, 30, '2019-03-09 21:20:34', '2019-03-21 20:22:48', NULL),
	(111111113, '11111', '感谢您的支持，您将以众筹专属价格339元获得智能手环运动蓝牙耳机二合一（商务版）1副', 'https://img30.360buyimg.com/cf/jfs/t1/29636/4/6858/52719/5c663c31E7812cb0d/77ef257f10a2ac2f.jpg', 339, 1036, 2000, 963, 0, 30, '2019-03-09 21:23:53', '2019-03-21 20:23:07', NULL),
	(111111114, '11111', '感谢您的支持，您将以众筹专属价格599元获得智能手环运动蓝牙耳机二合一（运动版+商务版）总2副', 'https://img30.360buyimg.com/cf/jfs/t1/16497/2/6865/69028/5c663c68E5502510c/f803fe92aefba4f8.jpg', 599, 455, 500, 43, 0, 30, '2019-03-09 21:24:33', '2019-03-21 20:23:25', NULL),
	(111111115, '11111', '感谢您的支持，您将以众筹专属价格599元获得智能手环运动蓝牙耳机二合一（运动版+商务版）总2副', 'https://img30.360buyimg.com/cf/jfs/t1/16497/2/6865/69028/5c663c68E5502510c/f803fe92aefba4f8.jpg', 2800, 180, 200, 20, 0, 30, '2019-03-09 21:24:33', '2019-03-21 20:23:25', NULL),
	(111111116, '11111', '感谢您的支持，您将以众筹专属价格599元获得智能手环运动蓝牙耳机二合一（运动版+商务版）总2副', 'https://img30.360buyimg.com/cf/jfs/t1/16497/2/6865/69028/5c663c68E5502510c/f803fe92aefba4f8.jpg', 27000, 89, 100, 11, 0, 30, '2019-03-09 21:24:33', '2019-03-21 20:23:25', NULL);
/*!40000 ALTER TABLE `product_packages` ENABLE KEYS */;

-- 导出  表 funding.product_types 结构
CREATE TABLE IF NOT EXISTS `product_types` (
  `id` int(3) unsigned NOT NULL,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='产品类型';

-- 正在导出表  funding.product_types 的数据：~3 rows (大约)
/*!40000 ALTER TABLE `product_types` DISABLE KEYS */;
REPLACE INTO `product_types` (`id`, `name`) VALUES
	(1, '科技'),
	(2, '生活'),
	(99, '其他');
/*!40000 ALTER TABLE `product_types` ENABLE KEYS */;

-- 导出  表 funding.roles 结构
CREATE TABLE IF NOT EXISTS `roles` (
  `id` int(2) NOT NULL COMMENT 'ID',
  `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色类型名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='角色表，记录角色类型';

-- 正在导出表  funding.roles 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
REPLACE INTO `roles` (`id`, `name`) VALUES
	(0, '普通用户'),
	(1, '审核员'),
	(2, '项目发起者'),
	(999, '超级管理员');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;

-- 导出  表 funding.role_permissions 结构
CREATE TABLE IF NOT EXISTS `role_permissions` (
  `id` int(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色-权限 ID',
  `role_id` int(2) NOT NULL COMMENT '角色 ID',
  `permission_id` int(3) NOT NULL COMMENT '角色所有的权限 ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_role_permissions_role_id` (`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='角色权限表，记录每个角色对应的权限';

-- 正在导出表  funding.role_permissions 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `role_permissions` DISABLE KEYS */;
REPLACE INTO `role_permissions` (`id`, `role_id`, `permission_id`) VALUES
	(1, 0, 1);
/*!40000 ALTER TABLE `role_permissions` ENABLE KEYS */;

-- 导出  表 funding.users 结构
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(24) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '手机号',
  `role_id` int(2) NOT NULL DEFAULT '0' COMMENT '角色类型 可见 roles 表',
  `person_id` int(16) NOT NULL DEFAULT '0' COMMENT '身份证号',
  `icon_url` varchar(0) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `default_address_id` bigint(24) unsigned NOT NULL DEFAULT '0' COMMENT '默认地址',
  `license_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '执照信息ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间(软删除) NULL为未删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `username_idx` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20008 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='用户';

-- 正在导出表  funding.users 的数据：~7 rows (大约)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
REPLACE INTO `users` (`id`, `username`, `password`, `nickname`, `email`, `phone`, `role_id`, `person_id`, `icon_url`, `default_address_id`, `license_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(10000, 'admin', '123456', '超级管理员', '', '0', 999, 0, '', 0, '', '2019-03-09 17:02:45', '2019-04-15 20:49:48', NULL),
	(20002, 'chixuntech', '123456', '驰讯通科技', '', '0', 2, 0, '', 0, '11111111', '2019-03-09 17:03:37', '2019-04-08 18:47:28', NULL),
	(20003, 'test1', '123456', '一号测试用户', '123456@123.com', '18512345678', 0, 0, '', 3, '', '2019-04-15 14:22:26', '2019-04-30 00:27:32', NULL),
	(20004, 'test2', '123456', '测试2', '123456@123.com', '18512345678', 0, 0, '', 0, '', '2019-04-15 20:47:28', '2019-04-15 20:50:00', NULL),
	(20005, 'test3', '123456', '测试3', '123456@123.com', '18512345678', 0, 0, '', 0, '', '2019-04-15 21:11:16', '2019-04-15 21:12:23', NULL),
	(20006, 'test4', '123456', '测试4', '123456@123.com', '18512345678', 0, 0, '', 0, '', '2019-04-15 21:14:30', '2019-04-15 21:14:30', NULL),
	(20007, 'test1415', '123456', 'test1415111', '3121@123.com', '18512345678', 0, 0, '', 0, '', '2019-04-17 20:35:13', '2019-04-17 20:35:13', NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
