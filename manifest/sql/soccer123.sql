-- --------------------------------------------------------
-- 主机:                           104.233.171.253
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


-- 导出 soccer 的数据库结构
CREATE DATABASE IF NOT EXISTS `soccer` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `soccer`;

-- 导出  表 soccer.b_bank 结构
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

-- 正在导出表  soccer.b_bank 的数据：~0 rows (大约)
DELETE FROM `b_bank`;

-- 导出  表 soccer.c_banner 结构
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

-- 正在导出表  soccer.c_banner 的数据：~0 rows (大约)
DELETE FROM `c_banner`;

-- 导出  表 soccer.c_language 结构
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

-- 正在导出表  soccer.c_language 的数据：~0 rows (大约)
DELETE FROM `c_language`;

-- 导出  表 soccer.o_deposit 结构
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

-- 正在导出表  soccer.o_deposit 的数据：~0 rows (大约)
DELETE FROM `o_deposit`;

-- 导出  表 soccer.o_soccer_order 结构
CREATE TABLE IF NOT EXISTS `o_soccer_order` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_no` bigint NOT NULL COMMENT '编号',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户编号',
  `user_account` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `add_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '时间',
  `events_id` bigint NOT NULL DEFAULT '0' COMMENT '赛事编号',
  `events_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赛事名称',
  `odds_id` bigint NOT NULL DEFAULT '0' COMMENT '赔率编号',
  `odds_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赔率标题',
  `bet_money` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '投注金额',
  `profit` decimal(18,6) NOT NULL DEFAULT '0.000000' COMMENT '赢利',
  `calc_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '结算时间',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态 0 撤单，1:投注成功，2：中奖，3：未中奖',
  `company_code` int NOT NULL DEFAULT '0' COMMENT '公司编号',
  `odds_calc_rule` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '结算规则',
  `odds_profit_rate` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '利率',
  `league_id` bigint NOT NULL DEFAULT '0' COMMENT '联盟编号',
  `league_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联盟',
  `odds_type` int NOT NULL DEFAULT '0' COMMENT '1:半场，2：全场',
  `events_start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '赛事开始时间',
  `fee` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '手续费',
  `events_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赛事开奖结果',
  `play_id` bigint NOT NULL DEFAULT '0' COMMENT 'PlayId',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.o_soccer_order 的数据：~0 rows (大约)
DELETE FROM `o_soccer_order`;

-- 导出  表 soccer.o_soccer_order_settle 结构
CREATE TABLE IF NOT EXISTS `o_soccer_order_settle` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_no` bigint NOT NULL COMMENT '编号',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户编号',
  `user_account` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `add_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '时间',
  `events_id` bigint NOT NULL DEFAULT '0' COMMENT '赛事编号',
  `events_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赛事名称',
  `odds_id` bigint NOT NULL DEFAULT '0' COMMENT '赔率编号',
  `odds_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赔率标题',
  `bet_money` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '投注金额',
  `profit` decimal(18,6) NOT NULL DEFAULT '0.000000' COMMENT '赢利',
  `calc_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '结算时间',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态 0 撤单，1:投注成功，2：中奖，3：未中奖',
  `company_code` int NOT NULL DEFAULT '0' COMMENT '公司编号',
  `odds_calc_rule` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '结算规则',
  `odds_profit_rate` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '利率',
  `league_id` bigint NOT NULL DEFAULT '0' COMMENT '联盟编号',
  `league_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联盟',
  `odds_type` int NOT NULL DEFAULT '0' COMMENT '1:半场，2：全场',
  `events_start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '赛事开始时间',
  `fee` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '手续费',
  `events_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赛事开奖结果',
  `play_id` bigint NOT NULL DEFAULT '0' COMMENT 'PlayId',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.o_soccer_order_settle 的数据：~0 rows (大约)
DELETE FROM `o_soccer_order_settle`;

-- 导出  表 soccer.o_withdraw 结构
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

-- 正在导出表  soccer.o_withdraw 的数据：~0 rows (大约)
DELETE FROM `o_withdraw`;

-- 导出  表 soccer.p_events 结构
CREATE TABLE IF NOT EXISTS `p_events` (
  `id` bigint NOT NULL COMMENT '编号',
  `home_team` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '主队名称',
  `away_team` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '客队名称',
  `en_home_team` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '主队英文',
  `en_away_team` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '客队英文',
  `rest_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '进行时间',
  `start_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '开始时间',
  `end_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '结束时间',
  `league_id` bigint NOT NULL DEFAULT '0' COMMENT '联盟编号',
  `league_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联盟',
  `en_league_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联盟英文',
  `half_status` int NOT NULL DEFAULT '0' COMMENT '半场状态 0未开始，1已开始，2已结束',
  `first_status` int NOT NULL DEFAULT '0' COMMENT '上半场状态 0未开始，1已开始，2已结束',
  `start_date` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '开始日期',
  `first_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '上半场开奖结果',
  `first_open_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上半场开奖时间',
  `second_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '下半场开奖结果',
  `second_open_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '下半场开奖结果',
  `second_status` int NOT NULL DEFAULT '0' COMMENT '全场状态',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `is_hot` int NOT NULL DEFAULT '0' COMMENT '热门',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '注释',
  `bet_money` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '交易量',
  `open_ressult` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `half_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '上半场开奖结果',
  `half_open_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上半场开奖时间',
  `all_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '全场开奖结果',
  `all_open_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '全场开奖时间',
  `all_status` int NOT NULL DEFAULT '0' COMMENT '全场状态',
  `fi_id` bigint NOT NULL DEFAULT '0' COMMENT 'FI编号',
  `handicap` double NOT NULL COMMENT '让球',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='赛事';

-- 正在导出表  soccer.p_events 的数据：~0 rows (大约)
DELETE FROM `p_events`;

-- 导出  表 soccer.p_events_odds 结构
CREATE TABLE IF NOT EXISTS `p_events_odds` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `events_id` bigint NOT NULL COMMENT '赛事编号',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `calc_rule` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '结算规则',
  `type` int NOT NULL DEFAULT '0' COMMENT '类型 1:上半场 2：全场',
  `bet_money` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '投注金额',
  `profit_rate` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '赢利',
  `header` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '主球队 1：主队，2：客队，draw 平局',
  `odds` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '赔率',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `play_code` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.p_events_odds 的数据：~0 rows (大约)
DELETE FROM `p_events_odds`;

-- 导出  表 soccer.p_league 结构
CREATE TABLE IF NOT EXISTS `p_league` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `country` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '国家',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '中文名称',
  `en_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '英文名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联盟';

-- 正在导出表  soccer.p_league 的数据：~0 rows (大约)
DELETE FROM `p_league`;

-- 导出  表 soccer.p_play 结构
CREATE TABLE IF NOT EXISTS `p_play` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type_code` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `en_name` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int NOT NULL,
  `type_name` varchar(200) COLLATE utf8mb4_general_ci NOT NULL,
  `code` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='玩法';

-- 正在导出表  soccer.p_play 的数据：~3 rows (大约)
DELETE FROM `p_play`;
INSERT INTO `p_play` (`id`, `type_code`, `name`, `en_name`, `status`, `type_name`, `code`) VALUES
	(1, '1000', 'Small', '小', 1, '让球&大小', 10000),
	(2, '1000', 'Big', '大', 1, '让球&大小', 10001),
	(3, '1000', 'WinAlone', '独赢', 1, '让球&大小', 10002);

-- 导出  表 soccer.p_play_type 结构
CREATE TABLE IF NOT EXISTS `p_play_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `en_name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `status` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='玩法类型';

-- 正在导出表  soccer.p_play_type 的数据：~2 rows (大约)
DELETE FROM `p_play_type`;
INSERT INTO `p_play_type` (`id`, `code`, `en_name`, `name`, `status`) VALUES
	(1, '1000', '大小', '大小', 1),
	(2, '2000', '独赢&双重机会', '独赢&双重机会', 1);

-- 导出  表 soccer.r_report_wallet_day 结构
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

-- 正在导出表  soccer.r_report_wallet_day 的数据：~0 rows (大约)
DELETE FROM `r_report_wallet_day`;

-- 导出  表 soccer.s_admin 结构
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

-- 正在导出表  soccer.s_admin 的数据：~0 rows (大约)
DELETE FROM `s_admin`;

-- 导出  表 soccer.s_admin_login_log 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=546 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  soccer.s_admin_login_log 的数据：~0 rows (大约)
DELETE FROM `s_admin_login_log`;

-- 导出  表 soccer.s_api 结构
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

-- 正在导出表  soccer.s_api 的数据：~0 rows (大约)
DELETE FROM `s_api`;

-- 导出  表 soccer.s_casbin 结构
CREATE TABLE IF NOT EXISTS `s_casbin` (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='policy table';

-- 正在导出表  soccer.s_casbin 的数据：~0 rows (大约)
DELETE FROM `s_casbin`;

-- 导出  表 soccer.s_dict 结构
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

-- 正在导出表  soccer.s_dict 的数据：~0 rows (大约)
DELETE FROM `s_dict`;

-- 导出  表 soccer.s_file 结构
CREATE TABLE IF NOT EXISTS `s_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `group` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=275 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  soccer.s_file 的数据：~0 rows (大约)
DELETE FROM `s_file`;

-- 导出  表 soccer.s_menu 结构
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

-- 正在导出表  soccer.s_menu 的数据：~0 rows (大约)
DELETE FROM `s_menu`;

-- 导出  表 soccer.s_menu_api_rule 结构
CREATE TABLE IF NOT EXISTS `s_menu_api_rule` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menu_id` bigint NOT NULL,
  `api_id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=354 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='功能菜单绑定API接口';

-- 正在导出表  soccer.s_menu_api_rule 的数据：~0 rows (大约)
DELETE FROM `s_menu_api_rule`;

-- 导出  表 soccer.s_operation_log 结构
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

-- 正在导出表  soccer.s_operation_log 的数据：~0 rows (大约)
DELETE FROM `s_operation_log`;

-- 导出  表 soccer.s_role 结构
CREATE TABLE IF NOT EXISTS `s_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4831 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色';

-- 正在导出表  soccer.s_role 的数据：~0 rows (大约)
DELETE FROM `s_role`;

-- 导出  表 soccer.s_role_menu 结构
CREATE TABLE IF NOT EXISTS `s_role_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1117 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.s_role_menu 的数据：~0 rows (大约)
DELETE FROM `s_role_menu`;

-- 导出  表 soccer.u_amount_category 结构
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

-- 正在导出表  soccer.u_amount_category 的数据：~0 rows (大约)
DELETE FROM `u_amount_category`;

-- 导出  表 soccer.u_amount_item 结构
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

-- 正在导出表  soccer.u_amount_item 的数据：~0 rows (大约)
DELETE FROM `u_amount_item`;

-- 导出  表 soccer.u_balance_code 结构
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

-- 正在导出表  soccer.u_balance_code 的数据：~0 rows (大约)
DELETE FROM `u_balance_code`;

-- 导出  表 soccer.u_digital_account 结构
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

-- 正在导出表  soccer.u_digital_account 的数据：~0 rows (大约)
DELETE FROM `u_digital_account`;

-- 导出  表 soccer.u_user 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=113 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

-- 正在导出表  soccer.u_user 的数据：~0 rows (大约)
DELETE FROM `u_user`;

-- 导出  表 soccer.u_user_login_log 结构
CREATE TABLE IF NOT EXISTS `u_user_login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint unsigned DEFAULT NULL,
  `account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.u_user_login_log 的数据：~0 rows (大约)
DELETE FROM `u_user_login_log`;

-- 导出  表 soccer.u_wallet 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包';

-- 正在导出表  soccer.u_wallet 的数据：~0 rows (大约)
DELETE FROM `u_wallet`;

-- 导出  表 soccer.u_wallet_log 结构
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

-- 正在导出表  soccer.u_wallet_log 的数据：~0 rows (大约)
DELETE FROM `u_wallet_log`;

-- 导出  表 soccer.u_wallet_top_up_application 结构
CREATE TABLE IF NOT EXISTS `u_wallet_top_up_application` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `trans_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `uid` bigint unsigned NOT NULL,
  `change_type` int unsigned NOT NULL,
  `money` decimal(12,4) unsigned DEFAULT NULL,
  `receipt_amount` decimal(12,4) unsigned DEFAULT '0.0000',
  `fee` decimal(12,4) unsigned DEFAULT '0.0000',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `aid` bigint unsigned DEFAULT '0',
  `status` tinyint unsigned DEFAULT '1' COMMENT '1 wait 2 success 3 fail',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `pay_address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `receive_address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `hash` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'tx_id',
  `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `ip` (`ip`),
  KEY `description` (`description`),
  KEY `trans_id` (`trans_id`),
  KEY `change_type` (`change_type`),
  KEY `pay_address` (`pay_address`),
  KEY `aid` (`aid`),
  KEY `receive_address` (`receive_address`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=3245 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  soccer.u_wallet_top_up_application 的数据：~0 rows (大约)
DELETE FROM `u_wallet_top_up_application`;

-- 导出  表 soccer.u_wallet_withdraw_application 结构
CREATE TABLE IF NOT EXISTS `u_wallet_withdraw_application` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `trans_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `uid` bigint unsigned NOT NULL,
  `change_type` int unsigned NOT NULL,
  `money` decimal(12,4) unsigned NOT NULL,
  `receipt_amount` decimal(12,4) unsigned NOT NULL,
  `fee` decimal(12,4) unsigned NOT NULL,
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `aid` bigint unsigned DEFAULT NULL,
  `status` int DEFAULT '1',
  `pay_address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `receive_address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `desc_two` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `ip` (`ip`),
  KEY `pay_address` (`pay_address`),
  KEY `receive_address` (`receive_address`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=7430 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  soccer.u_wallet_withdraw_application 的数据：~0 rows (大约)
DELETE FROM `u_wallet_withdraw_application`;

-- 导出  表 soccer.u_withdraw_account 结构
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

-- 正在导出表  soccer.u_withdraw_account 的数据：~0 rows (大约)
DELETE FROM `u_withdraw_account`;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
