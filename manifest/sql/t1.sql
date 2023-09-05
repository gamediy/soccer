-- --------------------------------------------------------
-- 主机:                           8.219.177.71
-- 服务器版本:                        8.0.24 - Source distribution
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  12.5.0.6677
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 star_net 的数据库结构
DROP DATABASE IF EXISTS `ab_admin`;
CREATE DATABASE IF NOT EXISTS `ab_admin` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `ab_admin`;

-- 导出  表 star_net.b_bank 结构
DROP TABLE IF EXISTS `b_bank`;
CREATE TABLE IF NOT EXISTS `b_bank` (
  `id` int NOT NULL AUTO_INCREMENT,
  `icon` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `currency` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `net` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `protocol` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` int NOT NULL,
  `class` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='银行';

-- 正在导出表  star_net.b_bank 的数据：~4 rows (大约)
DELETE FROM `b_bank`;
INSERT INTO `b_bank` (`id`, `icon`, `currency`, `net`, `protocol`, `status`, `class`) VALUES
	(1, 'fEVIrJ.png', 'PHP', 'BANK', 'BDO', 1, 'success'),
	(2, 'ZKkc8T.png', 'PHP', 'BANK', 'Union Bank', 1, 'error'),
	(3, 'MD5YT8.png', 'USDT', 'TRON', 'Trc20', 1, 'warning'),
	(4, 'DeQjNM.png', 'PHP', 'GCASH', 'GCASH', 1, 'info');

-- 导出  表 star_net.c_banner 结构
DROP TABLE IF EXISTS `c_banner`;
CREATE TABLE IF NOT EXISTS `c_banner` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `image` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `link` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sort` int DEFAULT '0',
  `status` tinyint unsigned DEFAULT '1' COMMENT '1open 2close',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=98 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  star_net.c_banner 的数据：~2 rows (大约)
DELETE FROM `c_banner`;
INSERT INTO `c_banner` (`id`, `title`, `image`, `link`, `desc`, `sort`, `status`, `created_at`) VALUES
	(96, 'sdfsdf', 'sdf', 'sdf', 'sdf', 0, 1, '2023-08-14 07:56:01'),
	(97, '测试', 'vG6zG3.png', '测试', '测试', 0, 1, '2023-08-15 10:58:49');

-- 导出  表 star_net.c_language 结构
DROP TABLE IF EXISTS `c_language`;
CREATE TABLE IF NOT EXISTS `c_language` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `en` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '英语',
  `fr` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '法语',
  `ja` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '日语',
  `hi` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '印度语',
  `zh` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '中文',
  `ko` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '汉语',
  `pt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '葡萄牙语',
  `es` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '西班牙语',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `hu` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '匈牙利语',
  `ar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '阿拉伯语',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  star_net.c_language 的数据：~0 rows (大约)
DELETE FROM `c_language`;
INSERT INTO `c_language` (`id`, `name`, `desc`, `en`, `fr`, `ja`, `hi`, `zh`, `ko`, `pt`, `es`, `created_at`, `updated_at`, `hu`, `ar`) VALUES
	(71, '期样图术十条就', 'labore minim consectetur Duis amet', 'minim adipisicing velit amet', 'deserunt eu', 'ullamco', 'Duis aliqua ipsum magna amet', 'nulla', 'do commodo ex magna Excepteur\nd\n\ndfdafaf\n\n\nfdafa', 'sunt fugiat', 'proident consequat est irure nisi', '2023-08-13 20:35:41', '2023-08-14 20:31:28', 'quis ea laboris magna', 'Duis sit id');

-- 导出  表 star_net.o_deposit 结构
DROP TABLE IF EXISTS `o_deposit`;
CREATE TABLE IF NOT EXISTS `o_deposit` (
  `order_no` bigint NOT NULL COMMENT '订单号',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `uid` bigint NOT NULL COMMENT 'UID',
  `pid` bigint NOT NULL COMMENT '上级ID',
  `status` int NOT NULL COMMENT '状态',
  `finish_at` datetime NOT NULL COMMENT '完成时间',
  `detail` varchar(500) NOT NULL COMMENT '详情',
  `status_remark` varchar(250) NOT NULL COMMENT '状态说明',
  `amount` decimal(18,2) NOT NULL COMMENT '充值金额',
  `sys_remark` varchar(250) NOT NULL COMMENT '系统说明',
  `address` varchar(100) NOT NULL COMMENT '地址或者是卡号',
  `net` varchar(100) DEFAULT NULL COMMENT '类型',
  `amount_item_id` int NOT NULL COMMENT '充值的编号',
  `currency` varchar(20) NOT NULL COMMENT '货币单位',
  `protocol` varchar(20) NOT NULL COMMENT '协议',
  `parent_path` varchar(1000) NOT NULL COMMENT '上级代理全路经',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `admin_operate` varchar(50) NOT NULL COMMENT '操作用户',
  `transfer_order_no` varchar(100) NOT NULL COMMENT '用户转账订单号',
  `transfer_img` varchar(200) NOT NULL COMMENT '用户转账图片',
  PRIMARY KEY (`order_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='充值订单';

-- 正在导出表  star_net.o_deposit 的数据：~0 rows (大约)
DELETE FROM `o_deposit`;

-- 导出  表 star_net.o_withdraw 结构
DROP TABLE IF EXISTS `o_withdraw`;
CREATE TABLE IF NOT EXISTS `o_withdraw` (
  `order_no` bigint NOT NULL COMMENT '订单号',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `uid` bigint NOT NULL COMMENT 'UID',
  `pid` bigint NOT NULL COMMENT '上级ID',
  `status` int NOT NULL COMMENT '状态 0 提现中，1提现成功，-1提现失败',
  `finish_at` datetime NOT NULL COMMENT '完成时间',
  `detail` varchar(500) NOT NULL COMMENT '详情',
  `status_remark` varchar(250) NOT NULL COMMENT '状态说明',
  `amount` decimal(18,2) NOT NULL COMMENT '提现金额',
  `sys_remark` varchar(250) NOT NULL COMMENT '系统说明',
  `address` varchar(100) NOT NULL COMMENT '地址（如果是银行卡，就是账号）',
  `amount_finally` bigint NOT NULL COMMENT '最后倒账金额',
  `fee` bigint NOT NULL COMMENT '手续费',
  `net` varchar(100) NOT NULL COMMENT '网络（数字货币和银行卡代号）',
  `amount_item_id` int NOT NULL COMMENT '编号',
  `currency` varchar(20) NOT NULL COMMENT '货币单位（USDT PHP CNY）',
  `protocol` varchar(20) NOT NULL COMMENT '协议',
  `title` varchar(250) NOT NULL DEFAULT '' COMMENT '标题',
  `note` varchar(100) NOT NULL DEFAULT '' COMMENT '说明',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `admin_operate` varchar(50) NOT NULL COMMENT '操作用户',
  `parent_path` varchar(2000) NOT NULL,
  PRIMARY KEY (`order_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='提现订单';

-- 正在导出表  star_net.o_withdraw 的数据：~0 rows (大约)
DELETE FROM `o_withdraw`;

-- 导出  表 star_net.phone_record 结构
DROP TABLE IF EXISTS `phone_record`;
CREATE TABLE IF NOT EXISTS `phone_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `phone` varchar(45) COLLATE utf8mb4_general_ci NOT NULL,
  `from_site` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int DEFAULT '0' COMMENT '0等待审核1已审核',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  star_net.phone_record 的数据：~0 rows (大约)
DELETE FROM `phone_record`;

-- 导出  表 star_net.r_report_wallet_day 结构
DROP TABLE IF EXISTS `r_report_wallet_day`;
CREATE TABLE IF NOT EXISTS `r_report_wallet_day` (
  `id` bigint NOT NULL,
  `uid` bigint NOT NULL,
  `account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `pid` bigint NOT NULL,
  `parent_path` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '上级路经',
  `balance_code` int NOT NULL COMMENT '余额编号',
  `date` date NOT NULL COMMENT '日期',
  `amount` decimal(18,2) NOT NULL COMMENT '余额',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `count` int DEFAULT NULL COMMENT '数量',
  `title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包每天报表';

-- 正在导出表  star_net.r_report_wallet_day 的数据：~0 rows (大约)
DELETE FROM `r_report_wallet_day`;

-- 导出  表 star_net.s_admin 结构
DROP TABLE IF EXISTS `s_admin`;
CREATE TABLE IF NOT EXISTS `s_admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `rid` int NOT NULL COMMENT '角色ID',
  `uname` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `pwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` int DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `key_secret` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `rid` (`rid`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  star_net.s_admin 的数据：~2 rows (大约)
DELETE FROM `s_admin`;
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `pwd`, `nickname`, `email`, `phone`, `status`, `created_at`, `updated_at`, `key_secret`) VALUES
	(42, 1, 'admin', '$2a$10$iruhy1SZdCrKZIlToWrWlOafLtHyIxzMOeuEMkGCbwlBqIYeCvqLC', 'admin', '', '', 1, '2022-07-02 11:28:52', '2023-08-18 05:01:28', ''),
	(52, 1, 'guest', '$2a$10$i4zx9ccYBUu32p9kNwWFieommIgFHdewojOSoXY2iZ7zKkK/8OUC6', 'guest', '', '', 1, '2023-04-04 05:49:59', '2023-08-02 06:57:00', NULL);

-- 导出  表 star_net.s_admin_login_log 结构
DROP TABLE IF EXISTS `s_admin_login_log`;
CREATE TABLE IF NOT EXISTS `s_admin_login_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int DEFAULT NULL,
  `account` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  CONSTRAINT `s_admin_login_log_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `s_admin` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=582 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  star_net.s_admin_login_log 的数据：~32 rows (大约)
DELETE FROM `s_admin_login_log`;
INSERT INTO `s_admin_login_log` (`id`, `uid`, `account`, `ip`, `created_at`, `updated_at`) VALUES
	(550, 42, '', '103.171.144.43', '2023-08-18 05:26:45', '2023-08-18 05:26:45'),
	(551, 42, '', '103.171.144.43', '2023-08-18 05:26:50', '2023-08-18 05:26:50'),
	(552, 42, '', '103.171.144.43', '2023-08-18 05:34:26', '2023-08-18 05:34:26'),
	(553, 42, '', '103.171.144.43', '2023-08-18 05:34:34', '2023-08-18 05:34:34'),
	(554, 42, '', '103.171.144.43', '2023-08-18 05:35:37', '2023-08-18 05:35:37'),
	(555, 42, '', '103.171.144.43', '2023-08-18 06:16:09', '2023-08-18 06:16:09'),
	(556, 42, '', '103.171.144.43', '2023-08-18 06:16:25', '2023-08-18 06:16:25'),
	(557, 42, '', '103.171.144.43', '2023-08-18 06:38:23', '2023-08-18 06:38:23'),
	(558, 42, '', '103.171.144.43', '2023-08-18 06:42:29', '2023-08-18 06:42:29'),
	(559, 42, NULL, '::1', '2023-08-17 17:13:50', '2023-08-17 17:13:50'),
	(560, 42, NULL, '::1', '2023-08-17 17:14:14', '2023-08-17 17:14:14'),
	(561, 42, NULL, '::1', '2023-08-17 17:24:26', '2023-08-17 17:24:26'),
	(562, 42, NULL, '::1', '2023-08-17 17:26:17', '2023-08-17 17:26:17'),
	(563, 42, NULL, '::1', '2023-08-17 17:29:30', '2023-08-17 17:29:30'),
	(564, 42, NULL, '::1', '2023-08-17 17:30:28', '2023-08-17 17:30:28'),
	(565, 42, NULL, '::1', '2023-08-17 17:31:08', '2023-08-17 17:31:08'),
	(566, 42, NULL, '::1', '2023-08-17 17:31:39', '2023-08-17 17:31:39'),
	(567, 42, NULL, '::1', '2023-08-17 17:42:48', '2023-08-17 17:42:48'),
	(568, 42, NULL, '103.104.101.122', '2023-08-17 09:54:47', '2023-08-17 09:54:47'),
	(569, 42, NULL, '103.104.101.122', '2023-08-17 09:54:47', '2023-08-17 09:54:47'),
	(570, 42, '', '::1', '2023-08-17 18:10:38', '2023-08-17 18:10:38'),
	(571, 42, '', '103.171.144.43', '2023-08-17 10:15:56', '2023-08-17 10:15:56'),
	(572, 42, '', '103.171.144.43', '2023-08-17 10:16:15', '2023-08-17 10:16:15'),
	(573, 42, '', '::1', '2023-08-17 18:16:36', '2023-08-17 18:16:36'),
	(574, 42, '', '103.104.101.122', '2023-08-17 10:26:49', '2023-08-17 10:26:49'),
	(575, 42, '', '::1', '2023-08-17 18:33:07', '2023-08-17 18:33:07'),
	(576, 42, '', '116.50.161.5', '2023-08-18 06:39:55', '2023-08-18 06:39:55'),
	(577, 42, '', '116.50.161.5', '2023-08-18 06:40:06', '2023-08-18 06:40:06'),
	(578, 42, '', '116.50.161.58', '2023-08-18 06:57:20', '2023-08-18 06:57:20'),
	(579, 42, '', '103.171.144.43', '2023-08-25 06:06:50', '2023-08-25 06:06:50'),
	(580, 42, '', '47.87.212.126', '2023-08-30 03:34:10', '2023-08-30 03:34:10'),
	(581, 42, '', '::1', '2023-09-03 15:40:40', '2023-09-03 15:40:40');

-- 导出  表 star_net.s_api 结构
DROP TABLE IF EXISTS `s_api`;
CREATE TABLE IF NOT EXISTS `s_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `method` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '1 get 2 post 3 put 4 delete',
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=277 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  star_net.s_api 的数据：~67 rows (大约)
DELETE FROM `s_api`;
INSERT INTO `s_api` (`id`, `group`, `url`, `method`, `desc`, `created_at`, `updated_at`) VALUES
	(210, '/sys/admin', '/api/sys/admin/list', 'GET', ' 管理员列表', '2023-08-09 19:24:04', '2023-08-10 18:25:30'),
	(211, '/sys/admin', '/api/sys/admin', 'DELETE', ' 管理员删除单个数据', '2023-08-09 19:24:04', '2023-08-10 18:25:32'),
	(212, '/sys/admin', '/api/sys/admin/user_info', 'GET', ' 用户信息', '2023-08-09 19:24:06', '2023-08-10 18:25:32'),
	(213, '/sys/admin', '/api/sys/admin/edit_pwd', 'PUT', ' 管理员修该自己密码', '2023-08-09 19:24:06', '2023-08-10 18:25:32'),
	(214, '/sys/adminLoginLog', '/api/sys/adminLoginLog/list', 'GET', ' 管理员查询登录日志', '2023-08-09 19:24:06', '2023-08-10 18:25:33'),
	(215, '/sys/api', '/api/sys/api', 'DELETE', ' API删除单个数据', '2023-08-09 19:24:07', '2023-08-10 18:25:35'),
	(216, '/sys/api', '/api/sys/api/list', 'GET', ' API查询列表', '2023-08-09 19:24:07', '2023-08-10 18:25:34'),
	(217, '/sys/dict', '/api/sys/dict', 'DELETE', ' 字典删除单个', '2023-08-09 19:24:08', '2023-08-10 18:25:37'),
	(218, '/sys/dict', '/api/sys/dict/list', 'GET', ' 字典查询列表', '2023-08-09 19:24:09', '2023-08-10 18:25:36'),
	(219, '/sys/file', '/api/sys/file', 'DELETE', ' 文件删除单个', '2023-08-09 19:24:10', '2023-08-10 18:25:38'),
	(220, '/sys/file', '/api/sys/file/list', 'GET', ' 文件查询列表', '2023-08-09 19:24:10', '2023-08-10 18:25:37'),
	(221, '/sys/menu', '/api/sys/menu', 'DELETE', ' 菜单列表删除', '2023-08-09 19:24:12', '2023-08-10 18:25:41'),
	(222, '/sys/menu', '/api/sys/menu/getMenuByPath', 'GET', ' 菜单获取根据路经', '2023-08-09 19:24:12', '2023-08-10 18:25:39'),
	(223, '/sys/menu', '/api/sys/menu/list', 'GET', ' 菜单列表', '2023-08-09 19:24:12', '2023-08-10 18:25:39'),
	(224, '/sys/menu', '/api/sys/menu/listWithChildren', 'GET', ' 列表查询菜单带子级', '2023-08-09 19:24:13', '2023-08-10 18:25:39'),
	(225, '/sys/pusher', '/api/sys/pusher/sendAdmins', 'POST', ' 推送通知管理员', '2023-08-09 19:24:14', '2023-08-10 18:25:41'),
	(226, '/sys/role', '/api/sys/role', 'PUT', ' 角色修该', '2023-08-09 19:24:14', '2023-08-10 18:25:43'),
	(227, '/sys/role', '/api/sys/role/list', 'GET', ' 角色列表', '2023-08-09 19:24:15', '2023-08-10 18:25:42'),
	(228, '/sys/ws', '/api/sys/ws/connect', 'GET', ' 长链接', '2023-08-09 19:24:16', '2023-08-10 18:25:43'),
	(229, '/sys/ws', '/api/sys/ws/send', 'POST', ' 通知单个管理员', '2023-08-09 19:24:16', '2023-08-10 18:25:44'),
	(230, '/sys/ws', '/api/sys/ws/noticeAdmins', 'POST', ' 通知所有管理员', '2023-08-09 19:24:17', '2023-08-10 18:25:44'),
	(231, '/sys/admin', '/api/sys/admin/edit_pwd_user', 'PUT', ' 管理员修该用户密码', '2023-08-10 18:19:41', '2023-08-10 18:25:33'),
	(232, '/sys/admin', '/api/sys/admin', 'POST', ' 管理员添加 (waiting)', '2023-08-10 18:25:31', '2023-08-10 18:25:31'),
	(233, '/sys/admin', '/api/sys/admin', 'GET', ' 管理员查询单个', '2023-08-10 18:25:31', '2023-08-10 18:25:31'),
	(234, '/sys/admin', '/api/sys/admin', 'PUT', ' 管理员修改单个数据', '2023-08-10 18:25:31', '2023-08-10 18:25:31'),
	(235, '/sys/api', '/api/sys/api', 'GET', ' API查询单个', '2023-08-10 18:25:33', '2023-08-10 18:25:33'),
	(236, '/sys/api', '/api/sys/api', 'POST', ' API添加', '2023-08-10 18:25:34', '2023-08-10 18:25:34'),
	(237, '/sys/api', '/api/sys/api', 'PUT', ' API修改单个数据', '2023-08-10 18:25:35', '2023-08-10 18:25:35'),
	(238, '/sys/dict', '/api/sys/dict', 'GET', ' 字典查询单个', '2023-08-10 18:25:35', '2023-08-10 18:25:35'),
	(239, '/sys/dict', '/api/sys/dict', 'POST', ' 字典添加', '2023-08-10 18:25:36', '2023-08-10 18:25:36'),
	(240, '/sys/dict', '/api/sys/dict', 'PUT', ' 字典修改', '2023-08-10 18:25:36', '2023-08-10 18:25:36'),
	(241, '/sys/file', '/api/sys/file', 'POST', ' 上传图片', '2023-08-10 18:25:37', '2023-08-10 18:25:37'),
	(242, '/sys/menu', '/api/sys/menu', 'GET', ' 菜单根据ID查询', '2023-08-10 18:25:38', '2023-08-10 18:25:38'),
	(243, '/sys/menu', '/api/sys/menu', 'POST', ' 菜单列表添加', '2023-08-10 18:25:40', '2023-08-10 18:25:40'),
	(244, '/sys/menu', '/api/sys/menu', 'PUT', ' 菜单修该 ', '2023-08-10 18:25:40', '2023-08-10 18:25:40'),
	(245, '/sys/role', '/api/sys/role', 'GET', ' 角色查询根据ID', '2023-08-10 18:25:41', '2023-08-10 18:25:41'),
	(246, '/sys/role', '/api/sys/role', 'POST', ' 角色添加', '2023-08-10 18:25:42', '2023-08-10 18:25:42'),
	(247, '/sys/role', '/api/sys/role', 'DELETE', ' 角色删除', '2023-08-10 18:25:43', '2023-08-10 18:25:43'),
	(248, '/setting/amountCategory', '/api/setting/amountCategory', 'GET', ' 查询AmountCategory', '2023-08-11 12:36:33', '2023-08-11 12:36:33'),
	(249, '/setting/amountCategory', '/api/setting/amountCategory/list', 'GET', ' 查询AmountCategory列表', '2023-08-11 12:36:33', '2023-08-11 12:36:33'),
	(250, '/setting/amountCategory', '/api/setting/amountCategory', 'POST', ' 添加AmountCategory', '2023-08-11 12:36:34', '2023-08-11 12:36:34'),
	(251, '/setting/amountCategory', '/api/setting/amountCategory', 'PUT', ' 修改AmountCategory', '2023-08-11 12:36:34', '2023-08-11 12:36:34'),
	(252, '/setting/amountCategory', '/api/setting/amountCategory', 'DELETE', ' 删除AmountCategory', '2023-08-11 12:36:34', '2023-08-11 12:36:34'),
	(253, '/sys/balanceCode', '/api/setting/balanceCode', 'GET', ' 查询账变类型', '2023-08-11 12:36:35', '2023-08-11 12:36:35'),
	(254, '/sys/balanceCode', '/api/setting/balanceCode/list', 'GET', ' 查询账变类型列表', '2023-08-11 12:36:35', '2023-08-11 12:36:35'),
	(255, '/sys/balanceCode', '/api/setting/balanceCode', 'POST', ' 添加账变类型', '2023-08-11 12:36:36', '2023-08-11 12:36:36'),
	(256, '/sys/balanceCode', '/api/setting/balanceCode', 'PUT', ' 修改账变类型', '2023-08-11 12:36:36', '2023-08-11 12:36:36'),
	(257, '/sys/balanceCode', '/api/setting/balanceCode', 'DELETE', ' 删除账变类型', '2023-08-11 12:36:36', '2023-08-11 12:36:36'),
	(258, '/user/loginLog', '/api/user/loginLog', 'GET', ' 登录日志查询', '2023-08-11 12:36:37', '2023-08-11 12:36:37'),
	(259, '/user/loginLog', '/api/user/loginLog/list', 'GET', ' 登录日志查询列表', '2023-08-11 12:36:37', '2023-08-11 12:36:37'),
	(260, '/user/loginLog', '/api/user/loginLog', 'POST', ' 登录日志添加', '2023-08-11 12:36:38', '2023-08-11 12:36:38'),
	(261, '/user/loginLog', '/api/user/loginLog', 'PUT', ' 登录日志修改', '2023-08-11 12:36:38', '2023-08-11 12:36:38'),
	(262, '/user/loginLog', '/api/user/loginLog', 'DELETE', ' 登录日志删除', '2023-08-11 12:36:39', '2023-08-11 12:36:39'),
	(263, '/user', '/api/user/', 'GET', ' 用户查询', '2023-08-11 12:36:40', '2023-08-11 12:36:40'),
	(264, '/user', '/api/user/list', 'GET', ' 用户查询列表', '2023-08-11 12:36:40', '2023-08-11 12:36:40'),
	(265, '/user', '/api/user/', 'POST', ' 用户添加', '2023-08-11 12:36:40', '2023-08-11 12:36:40'),
	(266, '/user', '/api/user/', 'PUT', ' 用户修改', '2023-08-11 12:36:41', '2023-08-11 12:36:41'),
	(267, '/user', '/api/user/', 'DELETE', ' 用户删除', '2023-08-11 12:36:41', '2023-08-11 12:36:41'),
	(268, '/user', '/api/user/pass', 'PUT', ' 用户修改用户登录密码', '2023-08-11 12:36:41', '2023-08-11 12:36:41'),
	(269, '/order', '/api/wallet/', 'GET', ' 查询钱包信息', '2023-08-11 12:36:42', '2023-08-11 12:36:42'),
	(270, '/order', '/api/wallet/update', 'PUT', ' 修改用户钱包', '2023-08-11 12:36:42', '2023-08-11 12:36:42'),
	(271, '/order', '/api/wallet/freeze', 'PUT', ' 冻结用户余额', '2023-08-11 12:36:42', '2023-08-11 12:36:42'),
	(272, '/order/walletLog', '/api/wallet/walletLog', 'GET', ' 查询账变日志', '2023-08-11 12:36:43', '2023-08-11 12:36:43'),
	(273, '/order/walletLog', '/api/wallet/walletLog/list', 'GET', ' 查询账变日志列表', '2023-08-11 12:36:43', '2023-08-11 12:36:43'),
	(274, '/order/walletLog', '/api/wallet/walletLog', 'POST', ' 添加账变日志', '2023-08-11 12:36:44', '2023-08-11 12:36:44'),
	(275, '/order/walletLog', '/api/wallet/walletLog', 'PUT', ' 修改账变日志', '2023-08-11 12:36:44', '2023-08-11 12:36:44'),
	(276, '/order/walletLog', '/api/wallet/walletLog', 'DELETE', ' 删除账变日志', '2023-08-11 12:36:45', '2023-08-18 06:16:38');

-- 导出  表 star_net.s_casbin 结构
DROP TABLE IF EXISTS `s_casbin`;
CREATE TABLE IF NOT EXISTS `s_casbin` (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='policy table';

-- 正在导出表  star_net.s_casbin 的数据：~196 rows (大约)
DELETE FROM `s_casbin`;
INSERT INTO `s_casbin` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
	('p', '财务', '/api/sys/api/list', 'GET', '', '', ''),
	('p', '财务', '/api/sys/api', 'DELETE', '', '', ''),
	('p', '财务', '/api/sys/menu/listWithChildren', 'GET', '', '', ''),
	('p', '财务', '/api/sys/menu/list', 'GET', '', '', ''),
	('p', '财务', '/api/sys/menu/getMenuByPath', 'GET', '', '', ''),
	('p', '财务', '/api/sys/menu', 'POST', '', '', ''),
	('p', '财务', '/api/sys/role/list', 'GET', '', '', ''),
	('p', '财务', '/api/sys/role', 'PUT', '', '', ''),
	('p', '财务', '/api/sys/role', 'GET', '', '', ''),
	('p', '财务', '/api/sys/role', 'POST', '', '', ''),
	('p', '财务', '/api/sys/role', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/sys/admin/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/admin', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/sys/menu', 'POST', '', '', ''),
	('p', '系统管理员', '/api/sys/role/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/role', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/role', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/role', 'POST', '', '', ''),
	('p', '系统管理员', '/api/sys/role', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/sys/menu', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/menu', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/menu', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/sys/admin/edit_pwd_user', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/admin/edit_pwd', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/admin', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/admin', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/admin', 'GET', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode', 'GET', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode', 'POST', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory', 'POST', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/menu/listWithChildren', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/menu/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/menu/getMenuByPath', 'GET', '', '', ''),
	('p', '12123', '/api/sys/admin/list', 'GET', '', '', ''),
	('p', '12123', '/api/sys/admin', 'DELETE', '', '', ''),
	('p', '12123', '/api/sys/menu', 'POST', '', '', ''),
	('p', '12123', '/api/sys/role/list', 'GET', '', '', ''),
	('p', '12123', '/api/sys/role', 'PUT', '', '', ''),
	('p', '12123', '/api/sys/role', 'GET', '', '', ''),
	('p', '12123', '/api/sys/role', 'POST', '', '', ''),
	('p', '12123', '/api/sys/role', 'DELETE', '', '', ''),
	('p', '12123', '/api/sys/menu', 'GET', '', '', ''),
	('p', '12123', '/api/sys/menu', 'PUT', '', '', ''),
	('p', '12123', '/api/sys/menu', 'DELETE', '', '', ''),
	('p', '12123', '/api/sys/admin/edit_pwd_user', 'PUT', '', '', ''),
	('p', '12123', '/api/sys/admin/edit_pwd', 'PUT', '', '', ''),
	('p', '12123', '/api/sys/admin', 'PUT', '', '', ''),
	('p', '12123', '/api/sys/admin', 'GET', '', '', ''),
	('p', '12123', '/api/sys/admin', 'GET', '', '', ''),
	('p', '12123', '/api/sys/api', 'PUT', '', '', ''),
	('p', '12123', '/api/sys/api', 'POST', '', '', ''),
	('p', '12123', '/api/sys/api', 'GET', '', '', ''),
	('p', '12123', '/api/sys/api/list', 'GET', '', '', ''),
	('p', '12123', '/api/sys/api', 'DELETE', '', '', ''),
	('p', '12123', '/api/sys/adminLoginLog/list', 'GET', '', '', ''),
	('p', '12123', '/api/setting/balanceCode', 'GET', '', '', ''),
	('p', '12123', '/api/setting/balanceCode/list', 'GET', '', '', ''),
	('p', '12123', '/api/setting/balanceCode', 'POST', '', '', ''),
	('p', '12123', '/api/setting/balanceCode', 'PUT', '', '', ''),
	('p', '12123', '/api/setting/balanceCode', 'DELETE', '', '', ''),
	('p', '12123', '/api/setting/amountCategory', 'DELETE', '', '', ''),
	('p', '12123', '/api/setting/amountCategory', 'PUT', '', '', ''),
	('p', '12123', '/api/setting/amountCategory', 'POST', '', '', ''),
	('p', '12123', '/api/setting/amountCategory/list', 'GET', '', '', ''),
	('p', '12123', '/api/setting/amountCategory', 'GET', '', '', ''),
	('p', '12123', '/api/sys/dict', 'GET', '', '', ''),
	('p', '12123', '/api/sys/dict', 'POST', '', '', ''),
	('p', '12123', '/api/sys/dict', 'PUT', '', '', ''),
	('p', '12123', '/api/sys/dict', 'GET', '', '', ''),
	('p', '12123', '/api/sys/dict', 'DELETE', '', '', ''),
	('p', '12123', '/api/sys/api', 'POST', '', '', ''),
	('p', '12123', '/api/sys/api', 'PUT', '', '', ''),
	('p', '12123', '/api/sys/api', 'GET', '', '', ''),
	('p', '12123', '/api/sys/api', 'DELETE', '', '', ''),
	('p', '12123', '/api/user/loginLog', 'DELETE', '', '', ''),
	('p', '12123', '/api/user/loginLog', 'PUT', '', '', ''),
	('p', '12123', '/api/user/loginLog', 'POST', '', '', ''),
	('p', '12123', '/api/sys/menu/listWithChildren', 'GET', '', '', ''),
	('p', '12123', '/api/sys/menu/list', 'GET', '', '', ''),
	('p', '12123', '/api/sys/menu/getMenuByPath', 'GET', '', '', ''),
	('p', '12123', '/api/sys/dict', 'DELETE', '', '', ''),
	('p', '12123', '/api/sys/dict/list', 'GET', '', '', ''),
	('p', '12123', '/api/sys/dict', 'GET', '', '', ''),
	('p', '12123', '/api/sys/dict', 'POST', '', '', ''),
	('p', '12123', '/api/sys/dict', 'PUT', '', '', ''),
	('p', '财务', '/api/sys/api/list', 'GET', '', '', ''),
	('p', '财务', '/api/sys/api', 'DELETE', '', '', ''),
	('p', '财务', '/api/sys/menu/listWithChildren', 'GET', '', '', ''),
	('p', '财务', '/api/sys/menu/list', 'GET', '', '', ''),
	('p', '财务', '/api/sys/menu/getMenuByPath', 'GET', '', '', ''),
	('p', '财务', '/api/sys/menu', 'POST', '', '', ''),
	('p', '财务', '/api/sys/role/list', 'GET', '', '', ''),
	('p', '财务', '/api/sys/role', 'PUT', '', '', ''),
	('p', '财务', '/api/sys/role', 'GET', '', '', ''),
	('p', '财务', '/api/sys/role', 'POST', '', '', ''),
	('p', '财务', '/api/sys/role', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/sys/admin/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/admin', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/sys/menu', 'POST', '', '', ''),
	('p', '系统管理员', '/api/sys/role/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/role', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/role', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/role', 'POST', '', '', ''),
	('p', '系统管理员', '/api/sys/role', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/sys/menu', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/menu', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/menu', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/sys/admin/edit_pwd_user', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/admin/edit_pwd', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/admin', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/sys/admin', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/admin', 'GET', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode', 'GET', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode', 'POST', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/setting/balanceCode', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory', 'DELETE', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory', 'PUT', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory', 'POST', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/setting/amountCategory', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/menu/listWithChildren', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/menu/list', 'GET', '', '', ''),
	('p', '系统管理员', '/api/sys/menu/getMenuByPath', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/admin/list', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/admin', 'DELETE', '', '', ''),
	('p', 'admin', '/api/sys/menu', 'POST', '', '', ''),
	('p', 'admin', '/api/sys/role/list', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/role', 'PUT', '', '', ''),
	('p', 'admin', '/api/sys/role', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/role', 'POST', '', '', ''),
	('p', 'admin', '/api/sys/role', 'DELETE', '', '', ''),
	('p', 'admin', '/api/sys/menu', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/menu', 'PUT', '', '', ''),
	('p', 'admin', '/api/sys/menu', 'DELETE', '', '', ''),
	('p', 'admin', '/api/sys/admin/edit_pwd_user', 'PUT', '', '', ''),
	('p', 'admin', '/api/sys/admin/edit_pwd', 'PUT', '', '', ''),
	('p', 'admin', '/api/sys/admin', 'PUT', '', '', ''),
	('p', 'admin', '/api/sys/admin', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/admin', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/api', 'PUT', '', '', ''),
	('p', 'admin', '/api/sys/api', 'POST', '', '', ''),
	('p', 'admin', '/api/sys/api', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/api/list', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/api', 'DELETE', '', '', ''),
	('p', 'admin', '/api/sys/adminLoginLog/list', 'GET', '', '', ''),
	('p', 'admin', '/api/setting/balanceCode', 'GET', '', '', ''),
	('p', 'admin', '/api/setting/balanceCode/list', 'GET', '', '', ''),
	('p', 'admin', '/api/setting/balanceCode', 'POST', '', '', ''),
	('p', 'admin', '/api/setting/balanceCode', 'PUT', '', '', ''),
	('p', 'admin', '/api/setting/balanceCode', 'DELETE', '', '', ''),
	('p', 'admin', '/api/setting/amountCategory', 'DELETE', '', '', ''),
	('p', 'admin', '/api/setting/amountCategory', 'PUT', '', '', ''),
	('p', 'admin', '/api/setting/amountCategory', 'POST', '', '', ''),
	('p', 'admin', '/api/setting/amountCategory/list', 'GET', '', '', ''),
	('p', 'admin', '/api/setting/amountCategory', 'GET', '', '', ''),
	('p', 'admin', '/api/wallet/update', 'PUT', '', '', ''),
	('p', 'admin', '/api/wallet/freeze', 'PUT', '', '', ''),
	('p', 'admin', '/api/user/loginLog', 'DELETE', '', '', ''),
	('p', 'admin', '/api/user/loginLog', 'PUT', '', '', ''),
	('p', 'admin', '/api/user/loginLog', 'POST', '', '', ''),
	('p', 'admin', '/api/user/loginLog/list', 'GET', '', '', ''),
	('p', 'admin', '/api/user/loginLog', 'GET', '', '', ''),
	('p', 'admin', '/api/wallet/walletLog', 'DELETE', '', '', ''),
	('p', 'admin', '/api/wallet/walletLog', 'PUT', '', '', ''),
	('p', 'admin', '/api/wallet/walletLog', 'POST', '', '', ''),
	('p', 'admin', '/api/wallet/walletLog/list', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/dict', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/dict', 'POST', '', '', ''),
	('p', 'admin', '/api/sys/dict', 'PUT', '', '', ''),
	('p', 'admin', '/api/sys/dict', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/dict', 'DELETE', '', '', ''),
	('p', 'admin', '/api/sys/api', 'POST', '', '', ''),
	('p', 'admin', '/api/sys/api', 'PUT', '', '', ''),
	('p', 'admin', '/api/sys/api', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/api', 'DELETE', '', '', ''),
	('p', 'admin', '/api/user/loginLog', 'DELETE', '', '', ''),
	('p', 'admin', '/api/user/loginLog', 'PUT', '', '', ''),
	('p', 'admin', '/api/user/loginLog', 'POST', '', '', ''),
	('p', 'admin', '/api/user/pass', 'PUT', '', '', ''),
	('p', 'admin', '/api/sys/menu/listWithChildren', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/menu/list', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/menu/getMenuByPath', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/dict', 'DELETE', '', '', ''),
	('p', 'admin', '/api/sys/dict/list', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/dict', 'GET', '', '', ''),
	('p', 'admin', '/api/sys/dict', 'POST', '', '', ''),
	('p', 'admin', '/api/sys/dict', 'PUT', '', '', '');

-- 导出  表 star_net.s_dict 结构
DROP TABLE IF EXISTS `s_dict`;
CREATE TABLE IF NOT EXISTS `s_dict` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `k` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `v` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'sys',
  `status` int DEFAULT NULL,
  `type` int NOT NULL DEFAULT '1' COMMENT '1 文本，2 img',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `k` (`k`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  star_net.s_dict 的数据：~3 rows (大约)
DELETE FROM `s_dict`;
INSERT INTO `s_dict` (`id`, `title`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES
	(42, '系统白名单', 'white_ips', 'localhost', '系统白名单', '1', 1, 3, '2022-12-28 12:25:41', '2023-08-13 03:35:34'),
	(44, '欢迎语', 'great', 'hello', '', '2', 1, 1, '2022-12-28 12:25:41', '2023-08-07 15:12:16'),
	(63, 'CloudFlare公共访问域', 'cloudflare_pub', 'https://pub-e700c1631e3a4dedaf546e73002fac2f.r2.dev/', 'CloudFlare公共访问域', '1', 1, 1, '2023-08-14 12:35:27', '2023-08-14 12:50:12');

-- 导出  表 star_net.s_file 结构
DROP TABLE IF EXISTS `s_file`;
CREATE TABLE IF NOT EXISTS `s_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `group` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=276 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  star_net.s_file 的数据：~16 rows (大约)
DELETE FROM `s_file`;
INSERT INTO `s_file` (`id`, `url`, `group`, `status`, `created_at`, `updated_at`) VALUES
	(253, 'hf5yEB.png', 2, 1, '2023-08-15 16:27:50', '2023-08-15 16:27:50'),
	(254, 'sKIquM.jpeg', 2, 1, '2023-08-15 16:36:32', '2023-08-15 16:36:32'),
	(261, 'W2HAa7.png', 2, 1, '2023-08-15 20:46:55', '2023-08-15 20:46:55'),
	(262, 'TE8xC2.jpeg', 2, 1, '2023-08-15 13:42:16', '2023-08-15 13:42:16'),
	(263, 'rJVUz4.png', 2, 1, '2023-08-15 13:42:29', '2023-08-15 13:42:29'),
	(264, '5T3wUb.jpeg', 2, 1, '2023-08-15 13:42:47', '2023-08-15 13:42:47'),
	(266, 'DvoyBR.png', 2, 1, '2023-08-15 13:43:56', '2023-08-15 13:43:56'),
	(267, 'qfhadk.jpeg', 2, 1, '2023-08-16 10:12:32', '2023-08-16 10:12:32'),
	(268, 'ZjJy7M.jpeg', 2, 1, '2023-08-16 10:13:09', '2023-08-16 10:13:09'),
	(269, 'vMj7EP.png', 2, 1, '2023-08-16 10:20:45', '2023-08-16 10:20:45'),
	(270, 'auRWSQ.png', 2, 1, '2023-08-16 13:49:24', '2023-08-16 13:49:24'),
	(271, 'eydjs6.jpeg', 2, 1, '2023-08-16 13:54:52', '2023-08-16 13:54:52'),
	(272, 'MD5YT8.png', 2, 1, '2023-08-16 11:55:14', '2023-08-16 11:55:14'),
	(273, 'ZKkc8T.png', 2, 1, '2023-08-16 11:55:27', '2023-08-16 11:55:27'),
	(274, 'fEVIrJ.png', 2, 1, '2023-08-16 11:55:39', '2023-08-16 11:55:39'),
	(275, 'XGczat.jpeg', 2, 1, '2023-08-17 10:17:06', '2023-08-17 10:17:06');

-- 导出  表 star_net.s_menu 结构
DROP TABLE IF EXISTS `s_menu`;
CREATE TABLE IF NOT EXISTS `s_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pid` int DEFAULT NULL,
  `icon` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `bg_img` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sort` decimal(7,2) NOT NULL DEFAULT '0.00',
  `type` int NOT NULL DEFAULT '1' COMMENT '1normal 2menu 3 button',
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `file_path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `permission` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '权限标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=351 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  star_net.s_menu 的数据：~52 rows (大约)
DELETE FROM `s_menu`;
INSERT INTO `s_menu` (`id`, `pid`, `icon`, `bg_img`, `name`, `path`, `sort`, `type`, `desc`, `file_path`, `status`, `created_at`, `updated_at`, `permission`) VALUES
	(1, -1, '', '', '系统', '#', 1.00, 2, '系统管理', '', 1, NULL, '2023-08-09 03:29:09', ''),
	(281, 1, 'DvoyBR.png', '', '菜单', '/admin/sys/menu', 1.10, 1, '这里是菜单管理页面', '', 1, NULL, '2023-08-15 05:44:00', '菜单管理'),
	(282, 1, 'DvoyBR.png', '', '角色', '/admin/sys/role', 1.30, 1, '这里是角色管理页面', '', 1, NULL, '2023-08-10 03:25:58', '角色管理'),
	(283, 1, 'ZjJy7M.jpeg', '', '管理员', '/admin/sys/admin', 1.50, 1, '这里是管理员管理页面', '', 1, NULL, '2023-08-09 03:30:28', '用户管理'),
	(284, 1, 'vMj7EP.png', '#', '字典', '/admin/sys/dict', 1.60, 1, '这里是字典管理页面', '', 1, NULL, '2023-08-16 02:20:50', '字典管理'),
	(285, 1, 'ZjJy7M.jpeg', '#', 'API', '/admin/sys/api', 1.20, 1, '这里是API管理页面', '', 1, NULL, '2023-08-10 20:27:35', 'API管理'),
	(286, 283, '', '', '用户删除', '#', 1.50, 3, '用户删除', '', 1, NULL, '2023-08-09 03:32:14', 'admin_del'),
	(287, 283, '', '#', '用户修改', '#', 1.50, 3, '用户修改', '', 1, NULL, '2023-08-10 20:26:07', 'admin_edit'),
	(288, 1, '', '#', '登录日志', '/admin/sys/adminLoginLog', 1.70, 1, '管理员登录日志', '', 1, NULL, '2023-08-10 20:27:58', 'API管理'),
	(289, -1, NULL, NULL, '配置', NULL, 2.00, 2, NULL, NULL, 1, '2023-08-06 08:22:03', '2023-08-06 08:22:03', '配置'),
	(290, 289, '', '', '账变类型', '/admin/setting/balance_code', 2.10, 1, 'u_balance_code', '', 1, NULL, '2023-08-10 20:46:30', '账变类型'),
	(291, 289, '', '', '充值提现类型', '/admin/setting/amount_category', 2.20, 1, 'u_amount_category', '', 1, NULL, '2023-08-10 20:47:14', '充值提现类型'),
	(292, 289, '', '', '分类项目', '/admin/setting/category', 2.30, 1, 'u_amount_item', '', 1, NULL, '2023-08-15 03:30:31', '分类项目'),
	(293, -1, NULL, NULL, '通用', NULL, 3.00, 2, NULL, NULL, 1, '2023-08-06 11:22:06', '2023-08-06 11:22:06', '通用'),
	(294, 293, 'auRWSQ.png', '', 'banner', '/admin/common/banner', 3.10, 1, 'banner', '', 1, NULL, '2023-08-16 05:49:27', 'banner'),
	(295, 293, '', '', '多语言', '/admin/common/language', 3.20, 1, '多语言', '', 1, NULL, '2023-08-13 19:17:26', ''),
	(296, -1, NULL, NULL, '统计', NULL, 4.00, 2, NULL, NULL, 1, '2023-08-06 11:26:06', '2023-08-06 11:26:06', '统计'),
	(297, 296, NULL, NULL, '平台报表', NULL, 4.10, 1, NULL, NULL, 1, '2023-08-06 11:27:04', '2023-08-06 11:27:04', '平台报表'),
	(298, 296, NULL, NULL, '平台统计', NULL, 4.20, 1, NULL, NULL, 1, '2023-08-06 11:27:05', '2023-08-06 11:27:05', '平台统计'),
	(299, 296, '', '', '账变统计', '/admin/wallet/reportWalletDay', 4.30, 1, '#', '', 1, NULL, '2023-08-15 07:07:46', '账变统计'),
	(300, -1, NULL, NULL, '用户', NULL, 5.00, 2, NULL, NULL, 1, '2023-08-06 11:27:36', '2023-08-06 11:27:36', '用户'),
	(301, 300, NULL, NULL, '用户列表', '/admin/user/user', 5.10, 1, NULL, NULL, 1, '2023-08-06 11:29:50', '2023-08-06 11:29:50', '用户列表'),
	(302, 300, '', '', '登录日志', '/admin/user/loginLog', 5.20, 1, ' 登录日志', '', 1, NULL, '2023-08-10 22:03:32', '登录日志'),
	(303, 300, '', '', '账变', '/admin/user/walletLog', 5.30, 1, '账变', '', 1, NULL, '2023-08-10 22:04:10', '账变'),
	(304, 300, '', '', '充值订单', '/admin/wallet/deposit', 5.40, 1, '充值订单页面', '', 1, NULL, '2023-08-13 03:14:32', '充值订单'),
	(305, 300, '', '', '提现订单', '/admin/wallet/withdraw', 5.50, 1, '这里是提现订单页面', '', 1, NULL, '2023-08-13 12:41:20', '提现订单'),
	(317, 314, '', '', '测试菜单', '', 1.00, 1, '测试啊啊', '测试啊', 2, '2023-08-07 23:16:15', '2023-08-07 23:16:15', ''),
	(319, 318, '', '', '测试2', '测试2', 1.00, 1, '测试2', '', 1, NULL, '2023-08-07 23:30:44', ''),
	(320, 319, '', '', '测试按钮', '测试测试测试测试', 1.00, 3, '测试123', '', 1, '2023-08-07 23:25:15', '2023-08-07 23:25:15', ''),
	(323, 281, '', '', '添加', '#', 1.00, 3, '添加', '', 1, NULL, '2023-08-10 02:51:47', 'menu_add'),
	(324, 281, '', '', '编辑', '#', 1.00, 3, '编辑', '', 1, NULL, '2023-08-10 20:23:31', 'menu_edit'),
	(325, 281, '', '', '删除', '#', 1.00, 3, '删除', '', 1, NULL, '2023-08-10 20:24:03', 'menu_del'),
	(326, 282, '', '', '添加', '#', 1.00, 3, '添加', '', 1, NULL, '2023-08-09 03:10:47', 'role_add'),
	(327, 301, '', '', '打开钱包', '#', 1.00, 3, '#', '', 1, NULL, '2023-08-10 20:51:04', 'user_打开钱包'),
	(329, 282, '', '', '编辑', '#', 1.00, 3, '角色编辑', '', 1, NULL, '2023-08-08 22:41:52', 'role_edit'),
	(330, 282, '', '', '删除', '#', 1.00, 3, '角色删除', '', 1, '2023-08-08 22:42:19', '2023-08-08 22:42:19', 'role_del'),
	(331, 283, '', '', '用户添加', '#', 1.00, 3, '用户添加', '', 1, NULL, '2023-08-10 20:26:49', 'admin_add'),
	(334, 301, '', '', '修改用户钱包', '#', 1.00, 3, '修改用户钱包', '', 1, '2023-08-10 20:52:16', '2023-08-10 20:52:16', 'user_修改用户钱包'),
	(336, 301, '', '', '冻结用户余额', '#', 1.00, 3, '冻结用户余额', '', 1, '2023-08-10 20:52:55', '2023-08-10 20:52:55', 'user_冻结用户余额'),
	(337, 301, '', '', '修改用户登录密码', '#', 1.00, 3, '修改用户登录密码', '', 1, NULL, '2023-08-15 00:02:46', 'user_用户修改用户登录密码'),
	(338, 284, '', '', '字典添加', '#', 1.00, 3, '字典添加', '', 1, NULL, '2023-08-10 23:48:37', 'dict_add'),
	(339, 284, '', '', '字典修该', '#', 1.00, 3, '字典修该', '', 1, '2023-08-10 23:49:42', '2023-08-10 23:49:42', 'dict_edit'),
	(340, 284, '', '', '字典删除', '#', 1.00, 3, '字典', '', 1, '2023-08-10 23:50:13', '2023-08-10 23:50:13', 'dict_del'),
	(341, 285, '', '', 'API添加', '#', 1.00, 3, 'API添加', '', 1, '2023-08-10 23:51:15', '2023-08-10 23:51:15', 'api_add'),
	(342, 285, '', '', 'API修该', '#', 1.00, 3, 'API修该', '', 1, '2023-08-10 23:51:54', '2023-08-10 23:51:54', 'api_edit'),
	(343, 285, '', '', 'API删除', '#', 1.00, 3, 'API删除', '', 1, '2023-08-10 23:52:33', '2023-08-10 23:52:33', 'api_del'),
	(344, 288, '', '', '登录日志删除', '#', 1.00, 3, '登录日志删除清空', '', 1, NULL, '2023-08-11 00:00:44', 'loginLog_del'),
	(345, 1, '', '', '文件', '/admin/sys/file', 1.00, 1, '这里是文件管理页面', '', 1, '2023-08-11 20:22:46', '2023-08-11 20:22:46', ''),
	(346, 345, '', '', '添加按钮', '#', 1.00, 3, '添加按钮', '', 1, '2023-08-12 20:44:03', '2023-08-12 20:44:03', 'file_add'),
	(347, 345, '', '', '编辑按钮', '#', 1.00, 3, '编辑按钮', '', 1, '2023-08-12 20:44:53', '2023-08-12 20:44:53', 'file_edit'),
	(348, 345, '', '', '删除按钮', '#', 1.00, 3, '删除按钮', '', 1, '2023-08-12 20:47:44', '2023-08-12 20:47:44', 'file_del'),
	(350, 289, '', '', '银行表', '/admin/setting/bank', 1.00, 1, '#', '', 1, NULL, '2023-08-14 22:30:31', '');

-- 导出  表 star_net.s_menu_api_rule 结构
DROP TABLE IF EXISTS `s_menu_api_rule`;
CREATE TABLE IF NOT EXISTS `s_menu_api_rule` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menu_id` bigint NOT NULL,
  `api_id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=354 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='功能菜单绑定API接口';

-- 正在导出表  star_net.s_menu_api_rule 的数据：~64 rows (大约)
DELETE FROM `s_menu_api_rule`;
INSERT INTO `s_menu_api_rule` (`id`, `menu_id`, `api_id`) VALUES
	(250, 283, 210),
	(251, 286, 211),
	(259, 323, 243),
	(265, 282, 227),
	(266, 282, 226),
	(267, 282, 245),
	(268, 282, 246),
	(269, 282, 247),
	(270, 324, 242),
	(271, 324, 244),
	(272, 325, 221),
	(273, 287, 231),
	(274, 287, 213),
	(275, 287, 234),
	(276, 287, 233),
	(277, 331, 233),
	(283, 285, 237),
	(284, 285, 236),
	(285, 285, 235),
	(286, 285, 216),
	(287, 285, 215),
	(288, 288, 214),
	(289, 290, 253),
	(290, 290, 254),
	(291, 290, 255),
	(292, 290, 256),
	(293, 290, 257),
	(294, 291, 252),
	(295, 291, 251),
	(296, 291, 250),
	(297, 291, 249),
	(298, 291, 248),
	(299, 334, 270),
	(300, 336, 271),
	(302, 302, 262),
	(303, 302, 261),
	(304, 302, 260),
	(305, 302, 259),
	(306, 302, 258),
	(307, 303, 276),
	(308, 303, 275),
	(309, 303, 274),
	(310, 303, 273),
	(313, 338, 238),
	(314, 338, 239),
	(315, 339, 240),
	(316, 339, 238),
	(317, 340, 217),
	(318, 341, 236),
	(319, 342, 237),
	(320, 342, 235),
	(321, 343, 215),
	(325, 344, 262),
	(326, 344, 261),
	(327, 344, 260),
	(329, 337, 268),
	(336, 281, 224),
	(337, 281, 223),
	(338, 281, 222),
	(349, 284, 217),
	(350, 284, 218),
	(351, 284, 238),
	(352, 284, 239),
	(353, 284, 240);

-- 导出  表 star_net.s_operation_log 结构
DROP TABLE IF EXISTS `s_operation_log`;
CREATE TABLE IF NOT EXISTS `s_operation_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `use_time` int NOT NULL,
  `created_at` datetime NOT NULL,
  `menu_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=7596 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  star_net.s_operation_log 的数据：~0 rows (大约)
DELETE FROM `s_operation_log`;

-- 导出  表 star_net.s_role 结构
DROP TABLE IF EXISTS `s_role`;
CREATE TABLE IF NOT EXISTS `s_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4834 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色';

-- 正在导出表  star_net.s_role 的数据：~4 rows (大约)
DELETE FROM `s_role`;
INSERT INTO `s_role` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
	(1, 'admin', 1, NULL, '2023-08-13 11:46:46'),
	(4827, '系统管理员', 1, NULL, '2023-08-17 15:18:37'),
	(4828, '财务', 1, NULL, '2023-08-10 20:26:42'),
	(4833, '12123', 1, NULL, '2023-08-17 07:20:35');

-- 导出  表 star_net.s_role_menu 结构
DROP TABLE IF EXISTS `s_role_menu`;
CREATE TABLE IF NOT EXISTS `s_role_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1394 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  star_net.s_role_menu 的数据：~166 rows (大约)
DELETE FROM `s_role_menu`;
INSERT INTO `s_role_menu` (`id`, `role_id`, `menu_id`) VALUES
	(985, 4828, 1),
	(986, 4828, 284),
	(987, 4828, 285),
	(988, 4828, 288),
	(989, 4828, 281),
	(990, 4828, 323),
	(991, 4828, 324),
	(992, 4828, 325),
	(993, 4828, 282),
	(994, 4828, 326),
	(995, 4828, 329),
	(996, 4828, 330),
	(1061, 1, 1),
	(1062, 1, 281),
	(1063, 1, 282),
	(1064, 1, 283),
	(1065, 1, 284),
	(1066, 1, 285),
	(1067, 1, 288),
	(1068, 1, 345),
	(1069, 1, 323),
	(1070, 1, 324),
	(1071, 1, 325),
	(1072, 1, 326),
	(1073, 1, 329),
	(1074, 1, 330),
	(1075, 1, 286),
	(1076, 1, 287),
	(1077, 1, 331),
	(1078, 1, 338),
	(1079, 1, 339),
	(1080, 1, 340),
	(1081, 1, 341),
	(1082, 1, 342),
	(1083, 1, 343),
	(1084, 1, 344),
	(1085, 1, 346),
	(1086, 1, 347),
	(1087, 1, 348),
	(1088, 1, 289),
	(1089, 1, 290),
	(1090, 1, 291),
	(1091, 1, 292),
	(1092, 1, 293),
	(1093, 1, 294),
	(1094, 1, 295),
	(1095, 1, 296),
	(1096, 1, 297),
	(1097, 1, 298),
	(1098, 1, 299),
	(1099, 1, 300),
	(1100, 1, 301),
	(1101, 1, 302),
	(1102, 1, 303),
	(1103, 1, 304),
	(1104, 1, 305),
	(1105, 1, 327),
	(1106, 1, 334),
	(1107, 1, 336),
	(1108, 1, 337),
	(1109, 1, 306),
	(1110, 1, 307),
	(1111, 1, 308),
	(1112, 1, 309),
	(1113, 1, 310),
	(1114, 1, 311),
	(1115, 1, 332),
	(1116, 1, 333),
	(1198, 4827, 1),
	(1199, 4827, 281),
	(1200, 4827, 323),
	(1201, 4827, 324),
	(1202, 4827, 325),
	(1203, 4827, 282),
	(1204, 4827, 326),
	(1205, 4827, 329),
	(1206, 4827, 330),
	(1207, 4827, 283),
	(1208, 4827, 286),
	(1209, 4827, 287),
	(1210, 4827, 331),
	(1211, 4827, 289),
	(1212, 4827, 290),
	(1213, 4827, 291),
	(1214, 4827, 292),
	(1215, 4833, 1),
	(1216, 4833, 281),
	(1217, 4833, 323),
	(1218, 4833, 324),
	(1219, 4833, 325),
	(1220, 4833, 282),
	(1221, 4833, 326),
	(1222, 4833, 329),
	(1223, 4833, 330),
	(1224, 4833, 283),
	(1225, 4833, 286),
	(1226, 4833, 287),
	(1227, 4833, 331),
	(1228, 4833, 284),
	(1229, 4833, 338),
	(1230, 4833, 339),
	(1231, 4833, 340),
	(1232, 4833, 285),
	(1233, 4833, 341),
	(1234, 4833, 342),
	(1235, 4833, 343),
	(1236, 4833, 288),
	(1237, 4833, 344),
	(1238, 4833, 345),
	(1239, 4833, 346),
	(1240, 4833, 347),
	(1241, 4833, 348),
	(1242, 4833, 289),
	(1243, 4833, 290),
	(1244, 4833, 291),
	(1245, 4833, 292),
	(1246, 4833, 350),
	(1345, 1, 1),
	(1346, 1, 281),
	(1347, 1, 323),
	(1348, 1, 324),
	(1349, 1, 325),
	(1350, 1, 282),
	(1351, 1, 326),
	(1352, 1, 329),
	(1353, 1, 330),
	(1354, 1, 283),
	(1355, 1, 286),
	(1356, 1, 287),
	(1357, 1, 331),
	(1358, 1, 284),
	(1359, 1, 338),
	(1360, 1, 339),
	(1361, 1, 340),
	(1362, 1, 285),
	(1363, 1, 341),
	(1364, 1, 342),
	(1365, 1, 343),
	(1366, 1, 288),
	(1367, 1, 344),
	(1368, 1, 345),
	(1369, 1, 346),
	(1370, 1, 347),
	(1371, 1, 348),
	(1372, 1, 289),
	(1373, 1, 290),
	(1374, 1, 291),
	(1375, 1, 292),
	(1376, 1, 293),
	(1377, 1, 294),
	(1378, 1, 295),
	(1379, 1, 296),
	(1380, 1, 297),
	(1381, 1, 298),
	(1382, 1, 299),
	(1383, 1, 300),
	(1384, 1, 302),
	(1385, 1, 303),
	(1386, 1, 304),
	(1387, 1, 305),
	(1388, 1, 301),
	(1389, 1, 327),
	(1390, 1, 334),
	(1391, 1, 336),
	(1392, 1, 337),
	(1393, 1, 350);

-- 导出  表 star_net.u_amount_category 结构
DROP TABLE IF EXISTS `u_amount_category`;
CREATE TABLE IF NOT EXISTS `u_amount_category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(250) DEFAULT NULL COMMENT '标题',
  `category` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '1:区块链，银行卡',
  `status` int DEFAULT NULL,
  `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'deposit,withdraw',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='充值或提现分类';

-- 正在导出表  star_net.u_amount_category 的数据：~2 rows (大约)
DELETE FROM `u_amount_category`;
INSERT INTO `u_amount_category` (`id`, `title`, `category`, `status`, `type`, `created_at`, `updated_at`) VALUES
	(1, '银行卡账转', '100', 1, '银行卡账转', '2023-08-11 16:49:34', '2023-08-11 16:49:34'),
	(2, 'GCASH', '200', 1, 'GCASH', '2023-08-11 16:49:57', '2023-08-11 16:49:57');

-- 导出  表 star_net.u_amount_item 结构
DROP TABLE IF EXISTS `u_amount_item`;
CREATE TABLE IF NOT EXISTS `u_amount_item` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(150) NOT NULL COMMENT '标题',
  `status` int NOT NULL COMMENT '状态',
  `detail` varchar(1000) DEFAULT NULL COMMENT '详情',
  `amount_category_id` int NOT NULL,
  `net` varchar(50) NOT NULL COMMENT 'tron eth',
  `min` bigint NOT NULL COMMENT '最小金额',
  `max` bigint NOT NULL COMMENT '最大金额',
  `fee` bigint NOT NULL COMMENT '手续费',
  `type` varchar(20) NOT NULL COMMENT '类型',
  `logo` varchar(200) NOT NULL COMMENT 'Logo',
  `sort` int NOT NULL COMMENT '排序大到小',
  `country` varchar(50) NOT NULL COMMENT '国家地区',
  `currency` varchar(20) NOT NULL COMMENT '货币单位',
  `protocol` varchar(20) NOT NULL COMMENT '协议',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `address` varchar(250) NOT NULL COMMENT '地址或卡号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='分类项目';

-- 正在导出表  star_net.u_amount_item 的数据：~5 rows (大约)
DELETE FROM `u_amount_item`;
INSERT INTO `u_amount_item` (`id`, `title`, `status`, `detail`, `amount_category_id`, `net`, `min`, `max`, `fee`, `type`, `logo`, `sort`, `country`, `currency`, `protocol`, `created_at`, `updated_at`, `address`) VALUES
	(2, 'GCASH 转账', 1, 'BDO 转账', 1, 'GCASH', 1, 100, 0, 'Deposit', 'eydjs6.jpeg', 1, 'PHP', 'PHP', 'GCASH', '2023-08-11 16:54:24', '2023-08-16 13:54:55', '123456789'),
	(3, 'USDT(TRC20)', 1, 'USDT(TRC20) 指定账号', 1, 'TRON', 1, 100, 0, 'Deposit', '5T3wUb.jpeg', 1, 'PHP', 'PHP', 'Trc20', '2023-08-11 16:54:24', '2023-08-15 13:42:49', '1234567890'),
	(4, 'USDT(TRC20)', 1, 'USDT(TRC20) 每个用户生成一个', 1, 'TRON', 1, 100, 0, 'Deposit', 'rJVUz4.png', 1, 'PHP', 'PHP', 'Trc20', '2023-08-11 16:54:24', '2023-08-15 13:42:30', ''),
	(5, 'BDO', 1, '卡2', 1, 'BANK', 1, 100, 0, 'Withdraw', 'TE8xC2.jpeg', 1, 'PHP', 'PHP', 'BDO', '2023-08-11 16:54:24', '2023-08-15 13:42:18', ''),
	(83, '非制象张正自', 73, 'quis labore culpa incididunt in', 66, 'do elit incididunt ex veniam', 91, 18, 97, 'consectetur', 'W2HAa7.png', 94, 'dolore non ipsum Excepteur nulla', 'anim ad', 'commodo non enim', '2023-08-15 19:20:30', '2023-08-15 20:46:59', '上海金门县通榆县');

-- 导出  表 star_net.u_balance_code 结构
DROP TABLE IF EXISTS `u_balance_code`;
CREATE TABLE IF NOT EXISTS `u_balance_code` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL COMMENT '标题',
  `code` int NOT NULL,
  `remark` varchar(500) DEFAULT NULL COMMENT '说明',
  `type` varchar(50) NOT NULL COMMENT '类型',
  `class` varchar(32) DEFAULT NULL COMMENT 'default,success,primary,info,warning,error,secondary,',
  `status` int NOT NULL COMMENT '状态',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='改用户余额的编码。100 -200 正加钱，负扣钱';

-- 正在导出表  star_net.u_balance_code 的数据：~10 rows (大约)
DELETE FROM `u_balance_code`;
INSERT INTO `u_balance_code` (`id`, `title`, `code`, `remark`, `type`, `class`, `status`, `created_at`, `updated_at`) VALUES
	(1, '充值', 500, '充值', '1', 'warning', 1, '2023-08-07 15:59:06', '2023-08-08 12:01:13'),
	(2, '提现', -100, '提现', '2', 'success', 1, '2023-08-07 15:59:20', '2023-08-14 19:55:45'),
	(3, '赠送', 400, '赠送', '1', 'warning', 1, '2023-08-07 15:59:38', '2023-08-09 10:12:40'),
	(4, '下注', -700, '下注', '2', 'success', 1, '2023-08-07 16:00:01', '2023-08-08 12:08:54'),
	(5, '银行卡充值', 501, '银行卡充值', '1', 'secondary', 1, '2023-08-07 16:16:21', '2023-08-08 12:02:31'),
	(6, 'GCASH充值', 502, 'GCASH充值', '1', 'error', 1, '2023-08-07 16:16:46', '2023-08-08 12:03:49'),
	(7, '人工扣除', -102, '人工扣除', '2', 'error', 1, '2023-08-08 11:05:38', '2023-08-08 11:05:53'),
	(8, '解冻', 800, '解冻', '1', 'success', 1, '2023-08-08 12:51:29', '2023-08-08 12:51:29'),
	(9, '冻结', -800, '冻结', '2', 'primary', 1, '2023-08-08 12:51:53', '2023-08-08 12:51:53'),
	(10, '提现失败', 100, '提现失败', '2', 'error', 1, '2023-08-07 15:59:20', '2023-08-14 19:55:32');

-- 导出  表 star_net.u_digital_account 结构
DROP TABLE IF EXISTS `u_digital_account`;
CREATE TABLE IF NOT EXISTS `u_digital_account` (
  `id` int NOT NULL AUTO_INCREMENT,
  `address` varchar(100) NOT NULL COMMENT '地址',
  `net` varchar(20) NOT NULL COMMENT '网络 TRON ETH',
  `status` int NOT NULL COMMENT '状态 0，1，打开',
  `count_deposit` int NOT NULL COMMENT '充值次数',
  `private_key` varchar(250) NOT NULL COMMENT '私密',
  `total_deposit` bigint NOT NULL COMMENT '总充值',
  `uid` bigint NOT NULL,
  `account` varchar(120) NOT NULL COMMENT '账户',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='数字货币账户';

-- 正在导出表  star_net.u_digital_account 的数据：~0 rows (大约)
DELETE FROM `u_digital_account`;

-- 导出  表 star_net.u_user 结构
DROP TABLE IF EXISTS `u_user`;
CREATE TABLE IF NOT EXISTS `u_user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account` varchar(100) NOT NULL DEFAULT '' COMMENT '账号',
  `email` varchar(50) NOT NULL COMMENT '邮箱',
  `password` varchar(128) NOT NULL COMMENT '密码',
  `status` int DEFAULT '1' COMMENT '1:正常，2：冻结',
  `xid` varchar(10) NOT NULL COMMENT 'short code 邀请 短码',
  `ip` varchar(50) DEFAULT NULL COMMENT 'IP',
  `client_agent` varchar(50) DEFAULT NULL COMMENT '注册clientAgen头',
  `phone` char(20) NOT NULL COMMENT '电报',
  `level_bet` int NOT NULL DEFAULT '0' COMMENT '下注的等级',
  `level_deposit` int DEFAULT NULL COMMENT '充值的等级',
  `level_agent` int NOT NULL DEFAULT '0' COMMENT '代理的等级',
  `pid` bigint DEFAULT NULL COMMENT '上级ID',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `parent_path` varchar(1000) DEFAULT NULL COMMENT '上级全路经 /1/2/3/',
  `country` varchar(50) DEFAULT NULL COMMENT '国家地区',
  `lang` varchar(10) DEFAULT NULL COMMENT '用户语言',
  `area_code` varchar(45) DEFAULT NULL,
  `city` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

-- 正在导出表  star_net.u_user 的数据：~3 rows (大约)
DELETE FROM `u_user`;
INSERT INTO `u_user` (`id`, `account`, `email`, `password`, `status`, `xid`, `ip`, `client_agent`, `phone`, `level_bet`, `level_deposit`, `level_agent`, `pid`, `created_at`, `updated_at`, `parent_path`, `country`, `lang`, `area_code`, `city`) VALUES
	(1, 'account', ' ', '$2a$10$fsDJDdTGn5pWlL4NZO4UPOtlkTf22Ih8Cj3hUDkH4bbMDzK9StzRO', 1, '1', '1', '1', '1', 1, 1, 1, 0, '2023-08-04 17:45:20', '2023-08-08 14:31:18', '1', 'Japan', 'zh', NULL, NULL),
	(112, 'user33', 'test@gmail.com', '$2a$10$AE0buvJowQXDqDBFqv6.TuB8Jsn3FpLJ4nve5XIZBOBxSCNwO2uy.', 1, 'OGNE55', '', '', '98787775676', 0, NULL, 0, 0, '2023-08-15 21:20:56', '2023-08-15 21:20:56', NULL, '', NULL, '', ''),
	(113, 'join1', '', '$2a$10$igEP2gQl.3HwL1LL/zwcyu1TpN6zpf.rHrH2G9vqk7sumwjyzQwhe', 1, '', '127.0.0.1', '', '', 0, 0, 0, 0, '2023-08-21 18:56:43', '2023-08-21 18:56:43', '', 'join1', '', '', 'gaoxiong');

-- 导出  表 star_net.u_user_login_log 结构
DROP TABLE IF EXISTS `u_user_login_log`;
CREATE TABLE IF NOT EXISTS `u_user_login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint unsigned DEFAULT NULL,
  `account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  star_net.u_user_login_log 的数据：~0 rows (大约)
DELETE FROM `u_user_login_log`;

-- 导出  表 star_net.u_wallet 结构
DROP TABLE IF EXISTS `u_wallet`;
CREATE TABLE IF NOT EXISTS `u_wallet` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint NOT NULL,
  `balance` bigint NOT NULL COMMENT '余额',
  `total_bet` decimal(18,2) NOT NULL COMMENT '总投注',
  `total_deposit` decimal(18,2) NOT NULL COMMENT '总充值',
  `total_withdraw` decimal(18,2) NOT NULL COMMENT '总提现',
  `freeze` bigint NOT NULL COMMENT '冻结',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `parent_path` varchar(2000) NOT NULL COMMENT '上级路经',
  `total_profit` decimal(18,2) NOT NULL COMMENT '总盈利',
  `total_gift` decimal(18,2) NOT NULL COMMENT '总赠送',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包';

-- 正在导出表  star_net.u_wallet 的数据：~3 rows (大约)
DELETE FROM `u_wallet`;
INSERT INTO `u_wallet` (`id`, `uid`, `balance`, `total_bet`, `total_deposit`, `total_withdraw`, `freeze`, `account`, `parent_path`, `total_profit`, `total_gift`, `created_at`, `updated_at`) VALUES
	(1, 1, 2074000, 0.00, 4010.00, 382.00, 21000, 'account', '1', 0.00, 0.00, '2023-08-08 15:20:38', '2023-08-16 19:34:32'),
	(5, 112, 927000, 0.00, 0.00, 0.00, 2000, 'user33', '', 0.00, 929.00, '2023-08-15 21:20:56', '2023-08-16 15:56:15'),
	(6, 113, 0, 0.00, 0.00, 0.00, 0, 'join1', '', 0.00, 0.00, '2023-08-21 18:56:43', '2023-08-21 18:56:43');

-- 导出  表 star_net.u_wallet_log 结构
DROP TABLE IF EXISTS `u_wallet_log`;
CREATE TABLE IF NOT EXISTS `u_wallet_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_no` bigint NOT NULL COMMENT '订单号，（有可能是充值提现的订单号）',
  `uid` bigint NOT NULL COMMENT 'UID',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `pid` bigint NOT NULL COMMENT '上级ID',
  `balance_code` int NOT NULL COMMENT '余额编号',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `balance_before` decimal(18,2) NOT NULL COMMENT '之前余额',
  `balance_after` decimal(18,2) NOT NULL COMMENT '之后余额',
  `note` varchar(500) NOT NULL COMMENT '说明',
  `order_no_relation` bigint NOT NULL COMMENT '关联订单号，（有可能是充值提现的订单号）',
  `parent_path` varchar(500) NOT NULL COMMENT '上级全路经',
  `amount` decimal(18,2) NOT NULL COMMENT '金额',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=171 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包账变';

-- 正在导出表  star_net.u_wallet_log 的数据：~154 rows (大约)
DELETE FROM `u_wallet_log`;
INSERT INTO `u_wallet_log` (`id`, `order_no`, `uid`, `account`, `pid`, `balance_code`, `title`, `balance_before`, `balance_after`, `note`, `order_no_relation`, `parent_path`, `amount`, `created_at`, `update_at`) VALUES
	(1, 1, 1, '1', 1, 501, '1', 1.00, 1.00, '1', 1, '1', 1.00, '2023-08-08 12:35:47', '2023-08-08 12:35:49'),
	(2, 2, 1, 'test2', 1, 502, '1', 1.00, 1.00, '13', 1, '1', 1.00, '2023-08-08 12:35:47', '2023-08-08 12:42:23'),
	(3, 1688862053034364928, 1, 'account', 0, 400, '', 10.00, 11.00, '', 0, '1', 1.00, '2023-08-08 18:37:46', '2023-08-08 18:37:46'),
	(4, 1688862363345752064, 1, 'account', 0, 400, '', 11.00, 12.00, '', 0, '1', 1.00, '2023-08-08 18:39:00', '2023-08-08 18:39:00'),
	(5, 1688862582074511360, 1, 'account', 0, 400, '', 12.00, 15.00, '', 0, '1', 3.00, '2023-08-08 18:39:52', '2023-08-08 18:39:52'),
	(6, 1688862650450055168, 1, 'account', 0, 400, '', 15.00, 16.00, '', 0, '1', 1.00, '2023-08-08 18:40:08', '2023-08-08 18:40:08'),
	(7, 1688862665843150848, 1, 'account', 0, 400, '', 16.00, 18.00, '', 0, '1', 2.00, '2023-08-08 18:40:12', '2023-08-08 18:40:12'),
	(8, 1688862805337313280, 1, 'account', 0, 400, '', 18.00, 19.00, '', 0, '1', 1.00, '2023-08-08 18:40:45', '2023-08-08 18:40:45'),
	(9, 1688862835548884992, 1, 'account', 0, 400, '', 19.00, 22.00, '', 0, '1', 3.00, '2023-08-08 18:40:52', '2023-08-08 18:40:52'),
	(10, 1688862867274600448, 1, 'account', 0, 400, '', 22.00, 23.00, '', 0, '1', 1.00, '2023-08-08 18:41:00', '2023-08-08 18:41:00'),
	(11, 1688862897674915840, 1, 'account', 0, 400, '', 23.00, 25.00, '', 0, '1', 2.00, '2023-08-08 18:41:07', '2023-08-08 18:41:07'),
	(12, 1688863683368718336, 1, 'account', 0, 400, '', 25.00, 27.00, '', 0, '1', 2.00, '2023-08-08 10:44:14', '2023-08-08 10:44:14'),
	(13, 1688869481452933120, 1, 'account', 0, -102, '', 27.00, 7.00, '', 0, '1', -20.00, '2023-08-08 11:07:17', '2023-08-08 11:07:17'),
	(14, 1688869610356477952, 1, 'account', 0, -102, '', 7.00, 6.00, '', 0, '1', -1.00, '2023-08-08 11:07:47', '2023-08-08 11:07:47'),
	(15, 1688869612906614784, 1, 'account', 0, -102, '', 6.00, 5.00, '', 0, '1', -1.00, '2023-08-08 11:07:48', '2023-08-08 11:07:48'),
	(16, 1688869613825167360, 1, 'account', 0, -102, '', 5.00, 4.00, '', 0, '1', -1.00, '2023-08-08 11:07:48', '2023-08-08 11:07:48'),
	(17, 1688869614647250944, 1, 'account', 0, -102, '', 4.00, 3.00, '', 0, '1', -1.00, '2023-08-08 11:07:48', '2023-08-08 11:07:48'),
	(18, 1688869616266252288, 1, 'account', 0, -102, '', 3.00, 2.00, '', 0, '1', -1.00, '2023-08-08 11:07:49', '2023-08-08 11:07:49'),
	(19, 1688869813771833344, 1, 'account', 0, -102, '', 2.00, 1.00, '', 0, '1', -1.00, '2023-08-08 11:08:36', '2023-08-08 11:08:36'),
	(20, 1688869814317092864, 1, 'account', 0, -102, '', 1.00, 0.00, '', 0, '1', -1.00, '2023-08-08 11:08:36', '2023-08-08 11:08:36'),
	(21, 1688869856583094272, 1, 'account', 0, 400, '', 0.00, 20.00, '', 0, '1', 20.00, '2023-08-08 11:08:46', '2023-08-08 11:08:46'),
	(22, 1688869886341681152, 1, 'account', 0, 400, '', 20.00, 21.00, '', 0, '1', 1.00, '2023-08-08 11:08:53', '2023-08-08 11:08:53'),
	(23, 1688870079111892992, 1, 'account', 0, 400, '', 21.00, 22.00, '', 0, '1', 1.00, '2023-08-08 11:09:39', '2023-08-08 11:09:39'),
	(24, 1688870124808835072, 1, 'account', 0, 400, '', 22.00, 23.00, '', 0, '1', 1.00, '2023-08-08 11:09:50', '2023-08-08 11:09:50'),
	(25, 1688870280761446400, 1, 'account', 0, -102, '', 23.00, 22.00, '', 0, '1', -1.00, '2023-08-08 11:10:27', '2023-08-08 11:10:27'),
	(26, 1688870281344454656, 1, 'account', 0, -102, '', 22.00, 21.00, '', 0, '1', -1.00, '2023-08-08 11:10:27', '2023-08-08 11:10:27'),
	(27, 1688870282112012288, 1, 'account', 0, -102, '', 21.00, 20.00, '', 0, '1', -1.00, '2023-08-08 11:10:28', '2023-08-08 11:10:28'),
	(28, 1688870283714236416, 1, 'account', 0, -102, '', 20.00, 19.00, '', 0, '1', -1.00, '2023-08-08 11:10:28', '2023-08-08 11:10:28'),
	(29, 1688870284897030144, 1, 'account', 0, -102, '', 19.00, 18.00, '', 0, '1', -1.00, '2023-08-08 11:10:28', '2023-08-08 11:10:28'),
	(30, 1688870341255892992, 1, 'account', 0, 400, '', 18.00, 19.00, '', 0, '1', 1.00, '2023-08-08 11:10:42', '2023-08-08 11:10:42'),
	(31, 1688870362235801600, 1, 'account', 0, -102, '', 19.00, 18.00, '', 0, '1', -1.00, '2023-08-08 11:10:47', '2023-08-08 11:10:47'),
	(32, 1688870362634260480, 1, 'account', 0, -102, '', 18.00, 17.00, '', 0, '1', -1.00, '2023-08-08 11:10:47', '2023-08-08 11:10:47'),
	(33, 1688870363397623808, 1, 'account', 0, -102, '', 17.00, 16.00, '', 0, '1', -1.00, '2023-08-08 11:10:47', '2023-08-08 11:10:47'),
	(34, 1688870364098072576, 1, 'account', 0, -102, '', 16.00, 15.00, '', 0, '1', -1.00, '2023-08-08 11:10:47', '2023-08-08 11:10:47'),
	(35, 1688870364727218176, 1, 'account', 0, -102, '', 15.00, 14.00, '', 0, '1', -1.00, '2023-08-08 11:10:47', '2023-08-08 11:10:47'),
	(36, 1688870563931492352, 1, 'account', 0, -102, '', 14.00, 13.00, '', 0, '1', -1.00, '2023-08-08 11:11:35', '2023-08-08 11:11:35'),
	(37, 1688870581824393216, 1, 'account', 0, -102, '', 13.00, 12.00, '', 0, '1', -1.00, '2023-08-08 11:11:39', '2023-08-08 11:11:39'),
	(38, 1688870597334929408, 1, 'account', 0, 400, '', 12.00, 13.00, '', 0, '1', 1.00, '2023-08-08 11:11:43', '2023-08-08 11:11:43'),
	(39, 1689091945613234176, 1, 'account', 0, -800, '冻结金额', 13.00, 3.00, '1', 0, '1', -10.00, '2023-08-09 09:51:16', '2023-08-09 09:51:16'),
	(40, 1689092582560239616, 1, 'account', 0, -800, '冻结金额', 3.00, 2.00, '1', 0, '1', -1.00, '2023-08-09 09:53:48', '2023-08-09 09:53:48'),
	(41, 1689092630622769152, 1, 'account', 0, 800, '解冻金额', 2.00, 22.00, '1', 0, '1', 20.00, '2023-08-09 09:54:00', '2023-08-09 09:54:00'),
	(42, 1689096329592049664, 1, 'account', 0, 400, '', 22.00, 23.00, '', 0, '1', 1.00, '2023-08-09 10:08:42', '2023-08-09 10:08:42'),
	(43, 1689096364278943744, 1, 'account', 0, -102, '', 23.00, 20.00, '', 0, '1', -3.00, '2023-08-09 10:08:50', '2023-08-09 10:08:50'),
	(44, 1689519315549687808, 1, 'account', 0, -800, '', 20.00, 10.00, '1', 0, '1', -10.00, '2023-08-10 06:09:29', '2023-08-10 06:09:29'),
	(45, 1689519350727315456, 1, 'account', 0, 800, '', 10.00, 20.00, '1', 0, '1', 10.00, '2023-08-10 06:09:38', '2023-08-10 06:09:38'),
	(46, 1689519385921720320, 1, 'account', 0, 400, '', 20.00, 22.00, '1', 0, '1', 2.00, '2023-08-10 06:09:46', '2023-08-10 06:09:46'),
	(47, 1689551424737775616, 1, 'account', 0, 400, '', 22.00, 23.00, '', 0, '1', 1.00, '2023-08-10 08:17:05', '2023-08-10 08:17:05'),
	(48, 1689614977943474176, 1, 'account', 0, 400, '', 23.00, 56.00, '', 0, '1', 33.00, '2023-08-10 12:29:37', '2023-08-10 12:29:37'),
	(50, 1689946212238626816, 1, 'account', 0, 501, 'BDO 转账 BDO 转账', 10000000.00, 10001000.00, '', 1689946210212777984, '1', 1000.00, '2023-08-11 18:25:56', '2023-08-11 18:25:56'),
	(51, 1689946937463148544, 1, 'account', 0, 501, 'BDO 转账 BDO 转账', 10001000.00, 10002000.00, '', 1689946935361802240, '1', 1000.00, '2023-08-11 18:28:42', '2023-08-11 18:28:42'),
	(62, 1690614278123425792, 1, 'account', 0, -100, 'Withdraw', 10002.00, 10001.00, 'BDO_address', 1690614275933999104, '1', -1.00, '2023-08-13 14:40:29', '2023-08-13 14:40:29'),
	(63, 1690614700103962624, 1, 'account', 0, -100, 'Withdraw', 100.00, 99.00, 'BDO_address', 1690614696714964992, '1', -1.00, '2023-08-13 14:42:09', '2023-08-13 14:42:09'),
	(64, 1690614772040470528, 1, 'account', 0, -100, 'Withdraw', 99.00, 98.00, 'BDO_address', 1690614769821683712, '1', -1.00, '2023-08-13 14:42:26', '2023-08-13 14:42:26'),
	(65, 1690619285187072000, 1, 'account', 0, 501, 'BANK', 98.00, 1098.00, 'BDO_1234567890', 1690584953198219264, '1', 1000.00, '2023-08-13 15:00:28', '2023-08-13 15:00:28'),
	(66, 1690622535101583360, 1, 'account', 0, 501, 'BANK', 1098.00, 1099.00, 'BDO_1234567890', 1690584953198219264, '1', 1.00, '2023-08-13 15:13:20', '2023-08-13 15:13:20'),
	(67, 1690623173097164800, 1, 'account', 0, 501, 'BANK', 1099.00, 2099.00, 'BDO_1234567890', 1689946935361802240, '1', 1000.00, '2023-08-13 15:15:51', '2023-08-13 15:15:51'),
	(68, 1690623660752113664, 1, 'account', 0, 501, 'BANK', 2099.00, 2100.00, 'BDO_1234567890', 1689983101553348608, '1', 1.00, '2023-08-13 15:17:46', '2023-08-13 15:17:46'),
	(69, 1690641115499204608, 1, 'account', 0, 501, 'BANK', 2100.00, 2101.00, 'BDO_1234567890', 1690636285145780224, '1', 1.00, '2023-08-13 16:27:07', '2023-08-13 16:27:07'),
	(75, 1690644654187352064, 1, 'account', 0, 501, 'BANK', 2101.00, 2104.00, 'BDO_1234567890', 1690644326264082432, '1', 3.00, '2023-08-13 16:41:11', '2023-08-13 16:41:11'),
	(76, 1690703281401106432, 1, 'account', 0, 100, 'Fail 提现失败', 2104.00, 2105.00, '提现失败', 1690614275933999104, '1', 1.00, '2023-08-13 20:34:09', '2023-08-13 20:34:09'),
	(77, 1690703641725374464, 1, 'account', 0, 100, 'Fail 提现失败', 2105.00, 2106.00, '提现失败', 1690614275933999104, '1', 1.00, '2023-08-13 20:35:35', '2023-08-13 20:35:35'),
	(78, 1690704106643001344, 1, 'account', 0, 100, 'Fail 提现失败', 2106.00, 2107.00, '提现失败', 1690614696714964992, '1', 1.00, '2023-08-13 20:37:25', '2023-08-13 20:37:25'),
	(79, 1690898090623504384, 1, 'account', 0, 501, 'BANK', 2107.00, 2108.00, 'BDO_1234567890', 1690687788648763392, '1', 1.00, '2023-08-14 01:28:15', '2023-08-14 01:28:15'),
	(80, 1691055331217510400, 1, 'account', 0, -100, 'Withdraw', 2108.00, 2107.00, 'BDO_address', 1691055327694295040, '1', -1.00, '2023-08-14 19:53:04', '2023-08-14 19:53:04'),
	(81, 1691055416877780992, 1, 'account', 0, 100, 'Fail ', 2107.00, 2108.00, '', 1691055327694295040, '1', 1.00, '2023-08-14 19:53:24', '2023-08-14 19:53:24'),
	(82, 1691055473681240064, 1, 'account', 0, -100, 'Withdraw', 2108.00, 2107.00, 'BDO_address', 1691055465884028928, '1', -1.00, '2023-08-14 19:53:38', '2023-08-14 19:53:38'),
	(83, 1691056736758140928, 1, 'account', 0, -100, 'Withdraw', 2107.00, 2094.00, 'BDO_address', 1691056733302034432, '1', -13.00, '2023-08-14 19:58:39', '2023-08-14 19:58:39'),
	(84, 1691063134355197952, 1, 'account', 0, -100, 'Withdraw', 2094.00, 2081.00, 'BDO_address', 1691063131368853504, '1', -13.00, '2023-08-14 20:24:04', '2023-08-14 20:24:04'),
	(85, 1691063369399799808, 1, 'account', 0, 100, 'Fail ', 2081.00, 2094.00, '', 1691063131368853504, '1', 13.00, '2023-08-14 20:25:00', '2023-08-14 20:25:00'),
	(86, 1691345825839452160, 1, 'account', 0, 500, 'title', 2094.00, 2095.00, '自己写', 1691345820013563904, '1', 1.00, '2023-08-15 15:07:23', '2023-08-15 15:07:23'),
	(87, 1691346008551723008, 1, 'account', 0, 500, 'title', 2095.00, 2096.00, '自己写', 1691346005288554496, '1', 1.00, '2023-08-15 15:08:07', '2023-08-15 15:08:07'),
	(88, 1691346759223087104, 1, 'account', 0, -800, 'title', 2096.00, 2095.00, '自己写', 1691346753615302656, '1', -1.00, '2023-08-15 15:11:06', '2023-08-15 15:11:06'),
	(89, 1691348006655234048, 1, 'account', 0, -800, 'title', 2095.00, 2094.00, '自己写', 1691347964536033280, '1', -1.00, '2023-08-15 15:17:38', '2023-08-15 15:17:38'),
	(90, 1691349314795737088, 1, 'account', 0, -800, 'title', 2094.00, 2093.00, '自己写', 1691349301046808576, '1', -1.00, '2023-08-15 15:21:15', '2023-08-15 15:21:15'),
	(91, 1691349370508677120, 1, 'account', 0, -800, 'title', 2093.00, 2092.00, '自己写', 1691349364124946432, '1', -1.00, '2023-08-15 15:21:28', '2023-08-15 15:21:28'),
	(92, 1691416633584652288, 109, 'join233', 0, 400, '', 0.00, 30.00, '', 0, '', 30.00, '2023-08-15 19:48:45', '2023-08-15 19:48:45'),
	(93, 1691417114037981184, 109, 'join233', 0, -800, '', 30.00, 20.00, '1', 0, '', -10.00, '2023-08-15 19:50:40', '2023-08-15 19:50:40'),
	(94, 1691417186133872640, 109, 'join233', 0, 800, '', 20.00, 25.00, '1', 0, '', 5.00, '2023-08-15 19:50:57', '2023-08-15 19:50:57'),
	(95, 1691424935689326592, 109, 'join233', 0, 400, '', 25.00, 35.00, '', 0, '', 10.00, '2023-08-15 20:21:44', '2023-08-15 20:21:44'),
	(96, 1691425015779561472, 109, 'join233', 0, -800, '', 35.00, 25.00, '1', 0, '', -10.00, '2023-08-15 20:22:04', '2023-08-15 20:22:04'),
	(97, 1691425100307369984, 109, 'join233', 0, 400, '', 25.00, 35.00, '1', 0, '', 10.00, '2023-08-15 20:22:24', '2023-08-15 20:22:24'),
	(98, 1691425134008602624, 109, 'join233', 0, -102, '', 35.00, 25.00, '1', 0, '', -10.00, '2023-08-15 20:22:32', '2023-08-15 20:22:32'),
	(99, 1691425794661814272, 1, 'account', 0, 501, '充值成功', 2092.00, 2093.00, 'BDO_1234567890', 1690687829874577408, '1', 1.00, '2023-08-15 12:25:09', '2023-08-15 12:25:09'),
	(100, 1691426638539001856, 109, 'join233', 0, -800, '', 25.00, 24.00, '1', 0, '', -1.00, '2023-08-15 20:28:30', '2023-08-15 20:28:30'),
	(101, 1691427229721956352, 109, 'join233', 0, -800, '', 24.00, 14.00, '1', 0, '', -10.00, '2023-08-15 20:30:51', '2023-08-15 20:30:51'),
	(102, 1691427308314824704, 109, 'join233', 0, 800, '', 14.00, 24.00, '1', 0, '', 10.00, '2023-08-15 20:31:10', '2023-08-15 20:31:10'),
	(103, 1691428065990676480, 109, 'join233', 0, -800, '', 24.00, 14.00, '1', 0, '', -10.00, '2023-08-15 20:34:11', '2023-08-15 20:34:11'),
	(104, 1691428096147722240, 109, 'join233', 0, -800, '', 14.00, 4.00, '1', 0, '', -10.00, '2023-08-15 20:34:18', '2023-08-15 20:34:18'),
	(105, 1691440476181237760, 112, 'user33', 0, 400, '赠送', 0.00, 333.00, 'd', 0, '', 333.00, '2023-08-15 21:23:30', '2023-08-15 21:23:30'),
	(106, 1691441409158025216, 1, 'account', 0, -800, '冻结', 2093.00, 2083.00, 'test', 0, '1', -10.00, '2023-08-15 21:27:13', '2023-08-15 21:27:13'),
	(107, 1691441806073401344, 112, 'user33', 0, 400, '赠送', 333.00, 343.00, 'test', 0, '', 10.00, '2023-08-15 21:28:47', '2023-08-15 21:28:47'),
	(108, 1691441816823402496, 112, 'user33', 0, 400, '赠送', 343.00, 353.00, 'test', 0, '', 10.00, '2023-08-15 21:28:50', '2023-08-15 21:28:50'),
	(109, 1691442158852116480, 1, 'account', 0, 800, '解冻', 2083.00, 2084.00, 'd', 0, '1', 1.00, '2023-08-15 21:30:11', '2023-08-15 21:30:11'),
	(110, 1691442399089266688, 112, 'user33', 0, 400, '赠送', 353.00, 586.00, 'd', 0, '', 233.00, '2023-08-15 21:31:09', '2023-08-15 21:31:09'),
	(111, 1691442422342488064, 112, 'user33', 0, 400, '赠送', 586.00, 819.00, 'd', 0, '', 233.00, '2023-08-15 21:31:15', '2023-08-15 21:31:15'),
	(112, 1691700862017081344, 112, 'user33', 0, -800, '冻结', 819.00, 818.00, 'd', 0, '', -1.00, '2023-08-16 14:38:11', '2023-08-16 14:38:11'),
	(113, 1691700957672378368, 112, 'user33', 0, 400, '赠送', 818.00, 828.00, 'd', 0, '', 10.00, '2023-08-16 14:38:34', '2023-08-16 14:38:34'),
	(114, 1691701870755581952, 112, 'user33', 0, 400, '赠送', 828.00, 928.00, 'd', 0, '', 100.00, '2023-08-16 14:42:12', '2023-08-16 14:42:12'),
	(115, 1691718421491748864, 1, 'account', 0, -800, '冻结', 2084.00, 2074.00, 'd', 0, '1', -10.00, '2023-08-16 15:47:59', '2023-08-16 15:47:59'),
	(116, 1691720510229975040, 112, 'user33', 0, -800, '冻结', 928.00, 927.00, 'd', 0, '', -1.00, '2023-08-16 15:56:17', '2023-08-16 15:56:17'),
	(117, 1691757453357617152, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691757449192673280, '1', -13.00, '2023-08-16 18:23:04', '2023-08-16 18:23:04'),
	(118, 1691757746027761664, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691757449192673280, '1', 13.00, '2023-08-16 18:24:14', '2023-08-16 18:24:14'),
	(119, 1691757791837949952, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691757788620918784, '1', -13.00, '2023-08-16 18:24:24', '2023-08-16 18:24:24'),
	(120, 1691757839833370624, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691757788620918784, '1', 13.00, '2023-08-16 18:24:35', '2023-08-16 18:24:35'),
	(121, 1691758146420215808, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691758143085744128, '1', -13.00, '2023-08-16 18:25:49', '2023-08-16 18:25:49'),
	(122, 1691758462846898176, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691758143085744128, '1', 13.00, '2023-08-16 18:27:04', '2023-08-16 18:27:04'),
	(123, 1691758587493224448, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691758582728495104, '1', -13.00, '2023-08-16 18:27:34', '2023-08-16 18:27:34'),
	(124, 1691758865923706880, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691758582728495104, '1', 13.00, '2023-08-16 18:28:40', '2023-08-16 18:28:40'),
	(125, 1691758908718190592, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691758904460972032, '1', -13.00, '2023-08-16 18:28:51', '2023-08-16 18:28:51'),
	(126, 1691758985591394304, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691758904460972032, '1', 13.00, '2023-08-16 18:29:09', '2023-08-16 18:29:09'),
	(127, 1691759249274703872, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691759238717640704, '1', -13.00, '2023-08-16 18:30:14', '2023-08-16 18:30:14'),
	(128, 1691759529210941440, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691759238717640704, '1', 13.00, '2023-08-16 18:31:18', '2023-08-16 18:31:18'),
	(129, 1691759611742261248, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691759600153399296, '1', -13.00, '2023-08-16 18:31:38', '2023-08-16 18:31:38'),
	(130, 1691760493569511424, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691759600153399296, '1', 13.00, '2023-08-16 18:35:08', '2023-08-16 18:35:08'),
	(131, 1691760561554984960, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691760557796888576, '1', -13.00, '2023-08-16 18:35:25', '2023-08-16 18:35:25'),
	(132, 1691760635194380288, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691760557796888576, '1', 13.00, '2023-08-16 18:35:42', '2023-08-16 18:35:42'),
	(133, 1691760732196048896, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691760728744136704, '1', -13.00, '2023-08-16 18:36:05', '2023-08-16 18:36:05'),
	(134, 1691760955769229312, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691760728744136704, '1', 13.00, '2023-08-16 18:36:58', '2023-08-16 18:36:58'),
	(135, 1691760990791667712, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691760986454757376, '1', -13.00, '2023-08-16 18:37:07', '2023-08-16 18:37:07'),
	(136, 1691761159532711936, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691760986454757376, '1', 13.00, '2023-08-16 18:37:47', '2023-08-16 18:37:47'),
	(137, 1691761192856457216, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691761189215801344, '1', -13.00, '2023-08-16 18:37:55', '2023-08-16 18:37:55'),
	(138, 1691761553025536000, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691761189215801344, '1', 13.00, '2023-08-16 18:39:21', '2023-08-16 18:39:21'),
	(139, 1691761586445750272, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691761582469550080, '1', -13.00, '2023-08-16 18:39:29', '2023-08-16 18:39:29'),
	(140, 1691761861919248384, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691761582469550080, '1', 13.00, '2023-08-16 18:40:34', '2023-08-16 18:40:34'),
	(141, 1691761899613458432, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691761891803664384, '1', -13.00, '2023-08-16 18:40:44', '2023-08-16 18:40:44'),
	(142, 1691762160289452032, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691761891803664384, '1', 13.00, '2023-08-16 18:41:46', '2023-08-16 18:41:46'),
	(143, 1691762298659540992, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691762294502985728, '1', -13.00, '2023-08-16 18:42:19', '2023-08-16 18:42:19'),
	(144, 1691762407384289280, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691762294502985728, '1', 13.00, '2023-08-16 18:42:44', '2023-08-16 18:42:44'),
	(145, 1691762541337776128, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691762537453850624, '1', -13.00, '2023-08-16 18:43:17', '2023-08-16 18:43:17'),
	(146, 1691762751044587520, 1, 'account', 0, -100, 'Withdraw', 2061.00, 2048.00, 'BDO_address', 1691762747764641792, '1', -13.00, '2023-08-16 18:44:07', '2023-08-16 18:44:07'),
	(147, 1691762813300641792, 1, 'account', 0, 100, 'Fail ', 2048.00, 2061.00, '', 1691762537453850624, '1', 13.00, '2023-08-16 18:44:21', '2023-08-16 18:44:21'),
	(148, 1691762832602828800, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691762747764641792, '1', 13.00, '2023-08-16 18:44:26', '2023-08-16 18:44:26'),
	(149, 1691762969777541120, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691762966568898560, '1', -13.00, '2023-08-16 18:44:59', '2023-08-16 18:44:59'),
	(150, 1691763089462005760, 1, 'account', 0, -100, 'Withdraw', 2061.00, 2048.00, 'BDO_address', 1691763086232391680, '1', -13.00, '2023-08-16 18:45:28', '2023-08-16 18:45:28'),
	(151, 1691763190133690368, 1, 'account', 0, 100, 'Fail ', 2048.00, 2061.00, '', 1691762966568898560, '1', 13.00, '2023-08-16 18:45:51', '2023-08-16 18:45:51'),
	(152, 1691763207028346880, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691763086232391680, '1', 13.00, '2023-08-16 18:45:55', '2023-08-16 18:45:55'),
	(153, 1691763347487199232, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691763344173699072, '1', -13.00, '2023-08-16 18:46:29', '2023-08-16 18:46:29'),
	(154, 1691763525199859712, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691763344173699072, '1', 13.00, '2023-08-16 18:47:11', '2023-08-16 18:47:11'),
	(155, 1691765443431239680, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691765440054824960, '1', -13.00, '2023-08-16 18:54:49', '2023-08-16 18:54:49'),
	(156, 1691765551916912640, 1, 'account', 0, -100, 'Withdraw', 2061.00, 2048.00, 'BDO_address', 1691765548263673856, '1', -13.00, '2023-08-16 18:55:14', '2023-08-16 18:55:14'),
	(157, 1691765623853420544, 1, 'account', 0, -100, 'Withdraw', 2048.00, 2035.00, 'BDO_address', 1691765620552503296, '1', -13.00, '2023-08-16 18:55:32', '2023-08-16 18:55:32'),
	(158, 1691765761078464512, 1, 'account', 0, -100, 'Withdraw', 2035.00, 2022.00, 'BDO_address', 1691765747342118912, '1', -13.00, '2023-08-16 18:56:04', '2023-08-16 18:56:04'),
	(159, 1691765851360858112, 1, 'account', 0, 100, 'Fail ', 2022.00, 2035.00, '', 1691765747342118912, '1', 13.00, '2023-08-16 18:56:26', '2023-08-16 18:56:26'),
	(160, 1691765873603252224, 1, 'account', 0, 100, 'Fail ', 2035.00, 2048.00, '', 1691765620552503296, '1', 13.00, '2023-08-16 18:56:31', '2023-08-16 18:56:31'),
	(161, 1691765895711428608, 1, 'account', 0, 100, 'Fail ', 2048.00, 2061.00, '', 1691765548263673856, '1', 13.00, '2023-08-16 18:56:37', '2023-08-16 18:56:37'),
	(162, 1691765970386817024, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691765440054824960, '1', 13.00, '2023-08-16 18:56:54', '2023-08-16 18:56:54'),
	(163, 1691766388269518848, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691766384742109184, '1', -13.00, '2023-08-16 18:58:34', '2023-08-16 18:58:34'),
	(164, 1691767685437722624, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691766384742109184, '1', 13.00, '2023-08-16 19:03:43', '2023-08-16 19:03:43'),
	(165, 1691771386793562112, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691771352828088320, '1', -13.00, '2023-08-16 19:18:31', '2023-08-16 19:18:31'),
	(166, 1691772322807025664, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691771352828088320, '1', 13.00, '2023-08-16 11:22:08', '2023-08-16 11:22:08'),
	(167, 1691774404393963520, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691774399994138624, '1', -13.00, '2023-08-16 19:30:25', '2023-08-16 19:30:25'),
	(168, 1691774497239076864, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691774399994138624, '1', 13.00, '2023-08-16 19:30:47', '2023-08-16 19:30:47'),
	(169, 1691774530009174016, 1, 'account', 0, -100, 'Withdraw', 2074.00, 2061.00, 'BDO_address', 1691774524850180096, '1', -13.00, '2023-08-16 19:30:55', '2023-08-16 19:30:55'),
	(170, 1691775443612471296, 1, 'account', 0, 100, 'Fail ', 2061.00, 2074.00, '', 1691774524850180096, '1', 13.00, '2023-08-16 19:34:33', '2023-08-16 19:34:33');

-- 导出  表 star_net.u_withdraw_account 结构
DROP TABLE IF EXISTS `u_withdraw_account`;
CREATE TABLE IF NOT EXISTS `u_withdraw_account` (
  `id` bigint NOT NULL,
  `net` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `protocol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `uid` int NOT NULL,
  `account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `currency` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'currency',
  `status` int NOT NULL,
  `default` int NOT NULL COMMENT '是否默认的',
  `title` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='提现账户';

-- 正在导出表  star_net.u_withdraw_account 的数据：~0 rows (大约)
DELETE FROM `u_withdraw_account`;
INSERT INTO `u_withdraw_account` (`id`, `net`, `protocol`, `uid`, `account`, `address`, `currency`, `status`, `default`, `title`) VALUES
	(1, 'BANK', 'BDO', 1, 'account', 'address', 'PHP', 1, 1, '');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
