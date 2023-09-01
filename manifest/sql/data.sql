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


-- 导出 soccer 的数据库结构
DROP DATABASE IF EXISTS `soccer`;
CREATE DATABASE IF NOT EXISTS `soccer` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `soccer`;

-- 导出  表 soccer.b_bank 结构
DROP TABLE IF EXISTS `b_bank`;
CREATE TABLE IF NOT EXISTS `b_bank` (
  `id` int NOT NULL AUTO_INCREMENT,
  `icon` varchar(500) COLLATE utf8mb4_general_ci NOT NULL COMMENT '图标',
  `currency` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `net` varchar(10) COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型',
  `protocol` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `summary` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '简介',
  `url` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '官网',
  `status` int NOT NULL COMMENT '状态 1开启 2关闭',
  `class` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='银行';

-- 正在导出表  soccer.b_bank 的数据：~42 rows (大约)
DELETE FROM `b_bank`;
INSERT INTO `b_bank` (`id`, `icon`, `currency`, `net`, `protocol`, `summary`, `url`, `status`, `class`) VALUES
	(1, 'fEVIrJ.png', 'PHP', 'BANK', 'BDO', '', '', 2, 'success'),
	(2, 'ZKkc8T.png', 'PHP', 'BANK', 'Union Bank', '', '', 2, 'error'),
	(3, 'MD5YT8.png', 'USDT', 'TRON', 'Trc20', '', '', 2, 'warning'),
	(4, 'DeQjNM.png', 'PHP', 'GCASH', 'GCASH', '', '', 2, 'info'),
	(7, 'aYv1L1.png', 'VND', 'BANK', 'Vietin Bank', 'Ngân hàng Thương mại Cổ phần Công thương Việt Nam', 'https://www.vietinbank.vn/web/home/vn/index.html', 1, 'default'),
	(8, 'SxyHWq.png', 'VND', 'BANK', 'BIDV', 'Ngân hàng Đầu tư và Phát triển Việt Nam', 'https://www.bidv.com.vn/en/ca-nhan', 1, 'default'),
	(9, 'E5Cu96.png', 'VND', 'BANK', 'VCB', 'Ngân hàng thương mại cổ phần Ngoại Thương Việt Nam', 'https://portal.vietcombank.com.vn/Pages/Home.aspx?devicechannel=default', 1, 'default'),
	(10, 'WaUkXn.png', 'VND', 'BANK', 'AGRIBANK', 'Ngân hàng Nông nghiệp và Phát triển Nông thôn Việt Nam', 'https://www.agribank.com.vn', 1, 'default'),
	(11, 'qPIiIc.png', 'VND', 'BANK', 'ACB', 'Ngân hàng thương mại cổ phần Á Châu', 'https://acb.com.vn', 1, 'default'),
	(12, 'Uvteyd.png', 'VND', 'BANK', 'TCB', 'Ngân hàng thương mại cổ phần Kỹ Thương Việt Nam', 'https://www.techcombank.com.vn/home', 1, 'default'),
	(13, 'dcOatu.png', 'VND', 'BANK', 'MBBANK', 'Ngân hàng Thương mại Cổ phần Quân đội ', 'https://www.mbbank.com.vn/', 1, 'default'),
	(14, 'P8jfCv.png', 'VND', 'BANK', 'SACOMBANK', 'Ngân hàng thương mại cổ phần Sài Gòn Thương Tín', 'https://www.sacombank.com.vn/', 1, 'default'),
	(15, '3UV6VN.png', 'VND', 'BANK', 'SCB', 'Ngân hàng Thương mại Cổ phần Sài gòn', 'https://www.scb.com.vn/', 1, 'default'),
	(16, 'KoPeOP.png', 'VND', 'BANK', 'TPBank', 'Ngân hàng Thương mại Cổ phần Tiên Phong', 'https://tpb.vn/', 1, 'default'),
	(17, 'fZvP6c.png', 'VND', 'BANK', 'SHB', 'Ngân hàng Sài Gòn-Hà Nội', 'https://www.shb.com.vn/', 1, 'default'),
	(18, '2a59uj.png', 'VND', 'BANK', 'VPB', 'Ngân hàng Việt Nam Thịnh Vượng', 'https://www.vpbank.com.vn/ca-nhan', 1, 'default'),
	(19, 'yQdtae.png', 'VND', 'BANK', 'OCB', 'Ngân hàng Phương Đông', 'https://www.ocb.com.vn/', 1, 'default'),
	(20, 'oySMOg.png', 'VND', 'BANK', 'SBV', 'Ngân hàng Nhà nước Việt Nam', 'https://www.sbv.gov.vn', 1, 'default'),
	(21, 'GJanwE.png', 'VND', 'BANK', 'VIB', 'Ngân hàng Thương mại Cổ phần Quốc tế Việt Nam', 'https://www.vib.com.vn/wps/portal/vn/ca-nhan', 1, 'default'),
	(22, 'MozBdi.png', 'VND', 'BANK', 'PVcomBank', 'Ngân hàng TMCP Đại Chúng Việt Nam', '', 1, 'default'),
	(23, 'ydsvML.png', 'VND', 'BANK', 'Exim Bank', 'Ngân hàng thương mại cổ phần Xuất Nhập Khẩu Việt Nam', '', 1, 'default'),
	(24, '6zlMYW.png', 'VND', 'BANK', 'DongA Bank', 'Ngân Hàng TMCP Đông Á', '', 1, 'default'),
	(25, 'JefmgS.png', 'VND', 'BANK', 'IVB', 'Ngân Hàng TNHH INDOVINA', '', 1, 'default'),
	(26, '4sdaeE.png', 'VND', 'BANK', 'ABBANK', 'Ngân hàng TMCP An Bình', '', 1, 'default'),
	(27, '4hCvPR.png', 'VND', 'BANK', 'LPB', 'Ngân hàng Bưu điện Liên Việt', 'https://www.lienvietpostbank.com.vn/', 1, 'default'),
	(28, 'g1dyH9.png', 'VND', 'BANK', 'MOMO', 'MOMO', '', 1, 'default'),
	(29, 'KHya0K.png', 'VND', 'BANK', 'Shinhan BANK', 'Ngân hàng Shinhan', '', 1, 'default'),
	(30, 'VZ0Kho.png', 'VND', 'BANK', 'BAN VIET', 'Ngân hàng bản việt', '', 1, 'default'),
	(31, 'wdRbmh.png', 'VND', 'BANK', 'NAM A BANK', 'Ngân hàng Thương mại Cổ phần Nam Á', '', 1, 'default'),
	(32, 'TDhgNg.png', 'VND', 'BANK', 'OCEAN BANK', 'Ngân hàng Đại Dương', '', 1, 'default'),
	(33, 'j2CYVd.png', 'VND', 'BANK', 'NCB BANK', 'Ngân hàng TMCP Quốc Dân', '', 1, 'default'),
	(34, 'bfuxnW.png', 'VND', 'BANK', 'PG Bank', 'Ngân Hàng TMCP Xăng Dầu Petrolimex', '', 1, 'default'),
	(35, '4fRPx8.png', 'VND', 'BANK', 'SeA Bank', 'Ngân Hàng TMCP Đông Nam Á', '', 1, 'default'),
	(36, 'rshp1F.png', 'VND', 'BANK', 'VIETA BANK', 'Ngân Hàng Tmcp Việt Á', '', 1, 'default'),
	(37, '58Du2E.png', 'VND', 'BANK', 'BAOVIET BANK', 'Ngân hàng TMCP Bảo Việt', '', 1, 'default'),
	(38, '6lfPGk.png', 'VND', 'BANK', 'CAKE VPBANK', 'Ngân hàng số Cake by VPBank', '', 1, 'default'),
	(39, 'bJs0pg.png', 'VND', 'BANK', 'CB BANK', 'Ngân Hàng Xây Dựng CB', '', 1, 'default'),
	(40, 'WpodMg.png', 'VND', 'BANK', 'CIMB BANK', 'Ngân hàng CIMB Việt Nam', '', 1, 'default'),
	(41, 'BhL7G5.png', 'VND', 'BANK', 'GP BANK', 'Ngân Hàng Tmcp Dầu Khí Toàn Cầu', '', 1, 'default'),
	(42, 'lCn583.png', 'VND', 'BANK', 'KIENLONG BANK', 'Ngân Hàng TMCP Kiên Long', '', 1, 'default'),
	(43, 'gOaTLb.png', 'VND', 'BANK', 'VIET BANK', 'Ngân hàng Thương mại Cổ phần Việt Nam Thương Tín', '', 1, 'default'),
	(44, '06aWgR.png', 'VND', 'BANK', 'UBANK', 'UBANK', '', 1, 'default'),
	(45, 'Wm3e0J.png', 'VND', 'BANK', 'HD', 'Ngân hàng thương mại cổ phần Phát triển Thành phố Hồ Chí Minh', 'https://www.hdbank.com.vn/', 1, 'default'),
	(46, 'hH9s8w.png', 'VND', 'BANK', 'MSB', 'Ngân hàng Hàng Hải Việt Nam', 'https://www.msb.com.vn/', 1, 'default'),
	(47, 'VqPVah.png', 'VND', 'BANK', 'BAB', 'Ngân hàng Bắc Á', 'https://www.baca-bank.vn/SitePages/trang-chu.aspx', 1, 'default'),
	(48, '7T06cV.png', 'VND', 'BANK', 'PBVN', 'Ngân hàng TNHaH MTV Public Việt Nam', 'https://publicbank.com.vn/index.aspx', 1, 'default');

-- 导出  表 soccer.c_banner 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- 正在导出表  soccer.c_banner 的数据：~0 rows (大约)
DELETE FROM `c_banner`;
INSERT INTO `c_banner` (`id`, `title`, `image`, `link`, `desc`, `sort`, `status`, `created_at`) VALUES
	(101, 'banner1', '2/2023/8/DIygkD.png', '', '', 0, 1, '2023-08-31 20:33:20'),
	(102, 'banner2', '2/2023/8/wNmDTD.png', '', '', 0, 1, '2023-08-31 20:33:47'),
	(103, 'banner3', '2/2023/8/MQbmP7.png', '', '', 0, 1, '2023-08-31 20:34:10');

-- 导出  表 soccer.c_language 结构
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

-- 正在导出表  soccer.c_language 的数据：~0 rows (大约)
DELETE FROM `c_language`;

-- 导出  表 soccer.o_deposit 结构
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

-- 正在导出表  soccer.o_deposit 的数据：~6 rows (大约)
DELETE FROM `o_deposit`;
INSERT INTO `o_deposit` (`order_no`, `account`, `uid`, `pid`, `status`, `finish_at`, `detail`, `status_remark`, `amount`, `sys_remark`, `address`, `net`, `amount_item_id`, `currency`, `protocol`, `parent_path`, `created_at`, `updated_at`, `admin_operate`, `transfer_order_no`, `transfer_img`) VALUES
	(1694318277255237632, 'join', 0, 0, -1, '2023-08-23 20:00:03', 'GCASH 转账 BDO 转账', '', 100.00, '', '123456789', 'GCASH', 2, 'PHP', 'GCASH', '', '2023-08-23 19:58:51', '2023-08-23 20:00:03', 'admin_42_103.171.144.43', '1231ff', 'Uvteyd.png'),
	(1694318683372916736, 'join', 0, 0, -1, '2023-08-23 20:00:43', 'GCASH 转账 BDO 转账', '', 100.00, '', '123456789', 'GCASH', 2, 'PHP', 'GCASH', '', '2023-08-23 20:00:28', '2023-08-23 20:00:43', 'admin_42_103.171.144.43', '1231ff', 'Uvteyd.png'),
	(1694319123795808256, 'join', 0, 0, -1, '2023-08-23 20:02:41', 'TCB Name:John', '', 100.00, '', '4123813103412323', 'BANK', 85, 'VND', 'TCB', '', '2023-08-23 20:02:13', '2023-08-23 20:02:41', 'admin_42_103.171.144.43', '1231ff', 'Uvteyd.png'),
	(1694319386690588672, 'join', 121, 0, 1, '2023-08-23 20:03:50', 'TCB Name:John', '充值成功', 100.00, '', '4123813103412323', 'BANK', 85, 'VND', 'TCB', '', '2023-08-23 20:03:15', '2023-08-23 20:03:51', 'admin_42_103.171.144.43', '1231ff', 'Uvteyd.png'),
	(1694319620439150592, 'join', 121, 0, 1, '2023-08-23 20:04:16', 'TCB Name:John', '充值成功', 200.00, '', '4123813103412323', 'BANK', 85, 'VND', 'TCB', '', '2023-08-23 20:04:11', '2023-08-23 20:04:17', 'admin_42_103.171.144.43', '1231ff', 'Uvteyd.png'),
	(1694320009880276992, 'join', 121, 0, 1, '2023-08-23 20:06:02', 'TCB Name:John', '充值成功', 100.00, '', '4123813103412323', 'BANK', 85, 'VND', 'TCB', '', '2023-08-23 20:05:44', '2023-08-23 20:06:04', 'admin_42_103.171.144.43', 'dfasfewfawf', 'Uvteyd.png');

-- 导出  表 soccer.o_soccer_order 结构
DROP TABLE IF EXISTS `o_soccer_order`;
CREATE TABLE IF NOT EXISTS `o_soccer_order` (
  `order_no` bigint NOT NULL COMMENT '编号',
  `uid` bigint NOT NULL DEFAULT '0' COMMENT '用户编号',
  `account` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '时间',
  `events_id` bigint NOT NULL DEFAULT '0' COMMENT '赛事编号',
  `events_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赛事名称',
  `odds_id` bigint NOT NULL DEFAULT '0' COMMENT '赔率编号',
  `odds_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赔率标题',
  `amount` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '投注金额',
  `profit` decimal(18,6) NOT NULL DEFAULT '0.000000' COMMENT '赢利',
  `calc_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '结算时间',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态 0 撤单，1:投注成功，2：中奖，3：未中奖',
  `odds_calc_rule` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '结算规则',
  `odds` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '赔率',
  `league_id` bigint NOT NULL DEFAULT '0' COMMENT '联盟编号',
  `league_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联盟',
  `events_start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '赛事开始时间',
  `fee` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '手续费',
  `events_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赛事开奖结果',
  `play_code` int NOT NULL DEFAULT '0' COMMENT 'PlayId',
  `bout_status` int NOT NULL COMMENT '哪个场次',
  `pid` bigint NOT NULL,
  `parent_path` varchar(2000) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`order_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.o_soccer_order 的数据：~10 rows (大约)
DELETE FROM `o_soccer_order`;
INSERT INTO `o_soccer_order` (`order_no`, `uid`, `account`, `created_at`, `events_id`, `events_title`, `odds_id`, `odds_title`, `amount`, `profit`, `calc_at`, `status`, `odds_calc_rule`, `odds`, `league_id`, `league_title`, `events_start_time`, `fee`, `events_open_result`, `play_code`, `bout_status`, `pid`, `parent_path`) VALUES
	(1694582258989535232, 121, 'join', '2023-08-24 18:28:35', 1, '1 vs 1', 1, '小', 1.00, -1.000000, '2023-08-24 18:28:31', 3, '4', 0.00, 1, '1', '2023-08-24 12:50:43', 0.00, '', 1000, 2, 0, ''),
	(1694584144580841472, 121, 'join', '2023-08-24 18:28:35', 1, '1 vs 1', 1, '小', 1.00, -1.000000, '2023-08-24 18:28:33', 3, '4', 0.00, 1, '1', '2023-08-24 12:50:43', 0.00, '', 1000, 2, 0, ''),
	(1694584530737827840, 121, 'join', '2023-08-24 18:49:25', 1, '1 vs 1', 2, '大', 1.00, 0.800000, '2023-08-24 18:49:11', 2, '4', 1.80, 1, '1', '2023-08-24 12:50:43', 0.00, '', 1001, 2, 0, ''),
	(1694916423035392000, 121, 'join', '2023-08-25 11:39:11', 1, '1 vs 1', 2, '大', 1.00, 0.500000, '2023-08-25 11:38:03', 2, '4', 1.50, 1, '1', '2023-08-24 12:50:43', 0.00, '1-4', 1001, 2, 0, ''),
	(1694916502785888256, 121, 'join', '2023-08-25 11:38:35', 1, '1 vs 1', 1, '小', 1.00, -1.000000, '2023-08-25 11:38:21', 3, '4', 1.90, 1, '1', '2023-08-24 12:50:43', 0.00, '1-4', 1000, 2, 0, ''),
	(1694916532317982720, 121, 'join', '2023-08-25 11:38:35', 1, '1 vs 1', 1, '小', 1.00, -1.000000, '2023-08-25 11:38:27', 3, '4', 1.90, 1, '1', '2023-08-24 12:50:43', 0.00, '1-4', 1000, 2, 0, ''),
	(1696780628218875904, 121, 'join', '2023-08-30 15:22:33', 1, '曼彻斯特联 vs 阿森纳', 21005, 'Over/Under', 1.00, 44.000000, '2023-08-30 15:21:14', 2, '大 2.5/3', 45.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '1-4', 1005, 2, 0, ''),
	(1696780631435907072, 121, 'join', '2023-08-30 15:21:57', 1, '曼彻斯特联 vs 阿森纳', 22000, '1x2', 1.00, -1.000000, '2023-08-30 15:21:50', 3, '主', 1.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '1-4', 2000, 2, 0, ''),
	(1696780634644549632, 121, 'join', '2023-08-30 15:22:43', 1, '曼彻斯特联 vs 阿森纳', 22001, '1x2', 1.00, 1.000000, '2023-08-30 15:21:57', 2, '客', 2.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '1-4', 2001, 2, 0, ''),
	(1696780637869969408, 121, 'join', '2023-08-30 15:21:57', 1, '曼彻斯特联 vs 阿森纳', 22002, '1x2', 1.00, -1.000000, '2023-08-30 15:21:57', 3, '和', 3.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '1-4', 2002, 2, 0, '');

-- 导出  表 soccer.o_soccer_order_settle 结构
DROP TABLE IF EXISTS `o_soccer_order_settle`;
CREATE TABLE IF NOT EXISTS `o_soccer_order_settle` (
  `order_no` bigint NOT NULL COMMENT '编号',
  `uid` bigint NOT NULL DEFAULT '0' COMMENT '用户编号',
  `account` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '时间',
  `events_id` bigint NOT NULL DEFAULT '0' COMMENT '赛事编号',
  `events_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赛事名称',
  `odds_id` bigint NOT NULL DEFAULT '0' COMMENT '赔率编号',
  `odds_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赔率标题',
  `amount` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '投注金额',
  `profit` decimal(18,6) NOT NULL DEFAULT '0.000000' COMMENT '赢利',
  `calc_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '结算时间',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态 0 撤单，1:投注成功，2：中奖，3：未中奖',
  `odds_calc_rule` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '结算规则',
  `odds` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '赔率',
  `league_id` bigint NOT NULL DEFAULT '0' COMMENT '联盟编号',
  `league_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联盟',
  `events_start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '赛事开始时间',
  `fee` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '手续费',
  `events_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '赛事开奖结果',
  `play_code` int NOT NULL DEFAULT '0' COMMENT 'PlayId',
  `bout_status` int NOT NULL COMMENT '哪个场次',
  `pid` bigint NOT NULL,
  `parent_path` varchar(2000) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`order_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.o_soccer_order_settle 的数据：~7 rows (大约)
DELETE FROM `o_soccer_order_settle`;
INSERT INTO `o_soccer_order_settle` (`order_no`, `uid`, `account`, `created_at`, `events_id`, `events_title`, `odds_id`, `odds_title`, `amount`, `profit`, `calc_at`, `status`, `odds_calc_rule`, `odds`, `league_id`, `league_title`, `events_start_time`, `fee`, `events_open_result`, `play_code`, `bout_status`, `pid`, `parent_path`) VALUES
	(1696107296448319488, 121, 'join', '2023-08-28 18:27:47', 1, '曼彻斯特联 vs 阿森纳', 3, '小', 1.00, 0.000000, '1970-01-01 08:00:00', 1, '大 2', 1.90, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '', 1000, 1, 0, ''),
	(1696780616298663936, 121, 'join', '2023-08-30 15:03:19', 1, '曼彻斯特联 vs 阿森纳', 12000, '1x2', 1.00, 0.000000, '1970-01-01 08:00:00', 1, '主', 1.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '', 2000, 1, 0, ''),
	(1696780621126307840, 121, 'join', '2023-08-30 15:03:20', 1, '曼彻斯特联 vs 阿森纳', 12001, '1x2', 1.00, 0.000000, '1970-01-01 08:00:00', 1, '客', 2.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '', 2001, 1, 0, ''),
	(1696780625006039040, 121, 'join', '2023-08-30 15:03:21', 1, '曼彻斯特联 vs 阿森纳', 12002, '1x2', 1.00, 0.000000, '1970-01-01 08:00:00', 1, '和', 3.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '', 2002, 1, 0, ''),
	(1696780641649037312, 121, 'join', '2023-08-30 15:03:25', 1, '曼彻斯特联 vs 阿森纳', 32000, '1x2', 1.00, 0.000000, '1970-01-01 08:00:00', 1, '主', 1.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '', 2000, 3, 0, ''),
	(1696780644866068480, 121, 'join', '2023-08-30 15:03:25', 1, '曼彻斯特联 vs 阿森纳', 32001, '1x2', 1.00, 0.000000, '1970-01-01 08:00:00', 1, '客', 2.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '', 2001, 3, 0, ''),
	(1696780648083099648, 121, 'join', '2023-08-30 15:03:26', 1, '曼彻斯特联 vs 阿森纳', 32002, '1x2', 1.00, 0.000000, '1970-01-01 08:00:00', 1, '和', 3.00, 1, '英格兰足球超级联赛', '2023-08-28 02:02:00', 0.00, '', 2002, 3, 0, '');

-- 导出  表 soccer.o_withdraw 结构
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

-- 正在导出表  soccer.o_withdraw 的数据：~2 rows (大约)
DELETE FROM `o_withdraw`;
INSERT INTO `o_withdraw` (`order_no`, `account`, `uid`, `pid`, `status`, `finish_at`, `detail`, `status_remark`, `amount`, `sys_remark`, `address`, `amount_finally`, `fee`, `net`, `amount_item_id`, `currency`, `protocol`, `title`, `note`, `created_at`, `updated_at`, `admin_operate`, `parent_path`) VALUES
	(1694650127437795328, 'join', 121, 0, -1, '2023-08-24 17:57:43', 'TCB Name:John', '', 100.00, '', '3123441333', 0, 0, 'BANK', 85, 'PHP', 'BDO', '', '2', '2023-08-24 17:57:32', '2023-08-25 06:45:41', 'admin_42_103.171.144.43', ''),
	(1694651645213806592, 'join', 121, 0, 1, '2023-08-24 18:03:44', 'TCB Name:John', '提现成功', 100.00, '', '3333331233333', 0, 0, 'BANK', 85, 'VND', 'PBVN', '', '1', '2023-08-24 18:03:34', '2023-08-25 06:45:34', 'admin_42_103.171.144.43', '');

-- 导出  表 soccer.p_events 结构
DROP TABLE IF EXISTS `p_events`;
CREATE TABLE IF NOT EXISTS `p_events` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `home_team_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '主队名称',
  `away_team_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '客队名称',
  `home_team_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '主队Id',
  `away_team_id` int NOT NULL DEFAULT '0' COMMENT '客队Id',
  `rest_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '进行时间',
  `start_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '开始时间',
  `end_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '结束时间',
  `league_id` bigint NOT NULL DEFAULT '0' COMMENT '联盟编号',
  `league_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联盟',
  `first_status` int NOT NULL DEFAULT '0' COMMENT '上半场状态 0未开始，1已开始，2已结束',
  `start_date` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '开始日期',
  `first_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '上半场开奖结果',
  `first_open_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上半场开奖时间',
  `second_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '下半场开奖结果',
  `second_open_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '下半场开奖结果',
  `second_status` int NOT NULL DEFAULT '0' COMMENT '0未开始，1已开始，2已结束',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `is_hot` int NOT NULL DEFAULT '0' COMMENT '热门',
  `status` int NOT NULL DEFAULT '1' COMMENT '0：未开始 1接受下注，2：停止下注，3 已结算',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '注释',
  `bet_money` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '交易量',
  `all_open_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '全场开奖结果',
  `all_open_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '全场开奖时间',
  `handicap` int NOT NULL DEFAULT '0' COMMENT '让球',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='赛事';

-- 正在导出表  soccer.p_events 的数据：~0 rows (大约)
DELETE FROM `p_events`;
INSERT INTO `p_events` (`id`, `home_team_name`, `away_team_name`, `home_team_id`, `away_team_id`, `rest_time`, `start_time`, `end_time`, `league_id`, `league_name`, `first_status`, `start_date`, `first_open_result`, `first_open_time`, `second_open_result`, `second_open_time`, `second_status`, `add_time`, `is_hot`, `status`, `remark`, `bet_money`, `all_open_result`, `all_open_time`, `handicap`) VALUES
	(1, '曼彻斯特联', '阿森纳', '1', 2, '30:33', '2023-08-28 02:02:00', '2023-08-24 12:50:46', 1, '英格兰足球超级联赛', 1, '2023-08-01 12:51:12', '1:1', '2023-08-01 12:51:19', '0:2', '2023-08-01 12:51:24', 1, '1970-01-01 00:00:03', 1, 1, '', 1.00, '1:3', '2023-08-24 12:51:54', 1);

-- 导出  表 soccer.p_events_odds 结构
DROP TABLE IF EXISTS `p_events_odds`;
CREATE TABLE IF NOT EXISTS `p_events_odds` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `events_id` bigint NOT NULL COMMENT '赛事编号',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `calc_rule` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '结算规则',
  `bout_status` int NOT NULL DEFAULT '0' COMMENT '类型 1:上半场 2：下半场，3：全场',
  `total_amount` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '投注金额',
  `total_profit` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '赢利',
  `header` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '主球队 1：主队，2：客队，draw 平局',
  `odds` decimal(18,2) NOT NULL DEFAULT '0.00' COMMENT '赔率',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `play_code` int NOT NULL,
  `status` int NOT NULL,
  `play_type_code` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.p_events_odds 的数据：~10 rows (大约)
DELETE FROM `p_events_odds`;
INSERT INTO `p_events_odds` (`id`, `events_id`, `title`, `calc_rule`, `bout_status`, `total_amount`, `total_profit`, `header`, `odds`, `created_at`, `updated_at`, `play_code`, `status`, `play_type_code`) VALUES
	(11001, 1, '大小球-大 1/1.5', '大 1/1.5', 1, 0.00, 0.00, '', 2.00, '2023-08-31 18:41:42', '2023-08-31 18:41:42', 1001, 1, 100),
	(12000, 1, '1x2', '主', 1, 1.00, 0.00, '', 1.00, '2023-08-30 03:36:15', '2023-08-30 15:03:19', 2000, 1, 200),
	(12001, 1, '1x2', '客', 1, 1.00, 0.00, '', 2.00, '2023-08-30 03:36:15', '2023-08-30 15:03:20', 2001, 1, 200),
	(12002, 1, '1x2', '和', 1, 1.00, 0.00, '', 3.00, '2023-08-30 03:36:15', '2023-08-30 15:03:21', 2002, 1, 200),
	(14000, 1, '让球-0', '0', 1, 0.00, 0.00, '', 2.00, '2023-08-31 20:16:26', '2023-08-31 20:16:26', 4000, 1, 400),
	(14001, 1, '让球-0/0.5', '0/0.5', 1, 0.00, 0.00, '', 1.00, '2023-08-31 20:13:44', '2023-08-31 20:13:44', 4001, 1, 400),
	(21005, 1, 'Over/Under', '大 2.5/3', 2, 1.00, -44.00, '', 45.00, '2023-08-30 03:36:48', '2023-08-30 15:22:34', 1005, 1, 100),
	(22000, 1, '1x2', '主', 2, 1.00, 0.00, '', 1.00, '2023-08-30 03:36:48', '2023-08-30 15:03:22', 2000, 1, 200),
	(22001, 1, '1x2', '客', 2, 1.00, -1.00, '', 2.00, '2023-08-30 03:36:48', '2023-08-30 15:22:43', 2001, 1, 200),
	(22002, 1, '1x2', '和', 2, 1.00, 0.00, '', 3.00, '2023-08-30 03:36:48', '2023-08-30 15:03:24', 2002, 1, 200),
	(32000, 1, '独赢-主', '主', 3, 0.00, 0.00, '', 1.00, '2023-08-31 16:01:23', '2023-08-31 16:01:23', 2000, 1, 200),
	(32001, 1, '独赢-客', '客', 3, 0.00, 0.00, '', 1.00, '2023-08-31 16:01:24', '2023-08-31 16:01:24', 2001, 1, 200),
	(32002, 1, '独赢-和', '和', 3, 0.00, 0.00, '', 1.00, '2023-08-31 16:01:24', '2023-08-31 16:01:24', 2002, 1, 200);

-- 导出  表 soccer.p_league 结构
DROP TABLE IF EXISTS `p_league`;
CREATE TABLE IF NOT EXISTS `p_league` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `country` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '国家',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态',
  `zh_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '中文名称',
  `en_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '英文名称',
  `icon` int DEFAULT NULL,
  `vd_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联盟';

-- 正在导出表  soccer.p_league 的数据：~0 rows (大约)
DELETE FROM `p_league`;
INSERT INTO `p_league` (`id`, `country`, `status`, `zh_name`, `en_name`, `icon`, `vd_name`) VALUES
	(1, 'PHL', 1, '英格兰足球超级联赛', 'lg_en_name', 1212, 'lg_vd_name');

-- 导出  表 soccer.p_play 结构
DROP TABLE IF EXISTS `p_play`;
CREATE TABLE IF NOT EXISTS `p_play` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `en_name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int NOT NULL,
  `type_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` int NOT NULL,
  `sort` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='玩法';

-- 正在导出表  soccer.p_play 的数据：~55 rows (大约)
DELETE FROM `p_play`;
INSERT INTO `p_play` (`id`, `type_code`, `en_name`, `name`, `status`, `type_name`, `code`, `sort`) VALUES
	(4, '200', 'Home', '主', 1, '独赢', 2000, 0),
	(5, '200', 'Away', '客', 1, '独赢', 2001, 0),
	(6, '200', 'Draw', '和', 1, '独赢', 2002, 0),
	(10, '300', '1-0', '1-0', 1, '波胆', 3000, 0),
	(11, '300', '1-1', '1-1', 1, '波胆', 3001, 0),
	(12, '300', '2-0', '2-0', 1, '波胆', 3002, 0),
	(13, '300', '2-1', '2-1', 1, '波胆', 3003, 0),
	(14, '300', '2-2', '2-2', 1, '波胆', 3004, 0),
	(15, '300', '3-0', '3-0', 1, '波胆', 3005, 0),
	(16, '300', '3-1', '3-1', 1, '波胆', 3006, 0),
	(17, '300', '3-2', '3-2', 1, '波胆', 3007, 0),
	(18, '300', '4-0', '4-0', 1, '波胆', 3008, 0),
	(19, '300', '4-1', '4-1', 1, '波胆', 3009, 0),
	(20, '300', '4-2', '4-2', 1, '波胆', 3010, 0),
	(21, '300', '4-3', '4-3', 1, '波胆', 3011, 0),
	(22, '300', '4-4', '4-4', 1, '波胆', 3012, 0),
	(23, '400', '0', '0', 1, '让球', 4000, 0),
	(24, '400', '+0/0.5', '0/0.5', 1, '让球', 4001, 0),
	(25, '400', '-0/0.5', '-0/0.5', 1, '让球', 4002, 0),
	(26, '400', '-0.5', '-0.5', 1, '让球', 4003, 0),
	(27, '400', '+0.5', '+0.5', 1, '让球', 4004, 0),
	(28, '400', '-0.5/1', '-0.5/1', 1, '让球', 4005, 0),
	(29, '400', '+0.5/1', '+0.5/1', 1, '让球', 4006, 0),
	(30, '400', '-1', '-1', 1, '让球', 4007, 0),
	(31, '400', '+1', '+1', 1, '让球', 4008, 0),
	(32, '400', '-1/1.5', '-1/1.5', 1, '让球', 4009, 0),
	(33, '400', '+1/1.5', '+1/1.5', 1, '让球', 4010, 0),
	(34, '400', '-1.5', '-1.5', 1, '让球', 4011, 0),
	(35, '400', '+1.5', '+1.5', 1, '让球', 4012, 0),
	(36, '400', '-1.5/2', '-1.5/2', 1, '让球', 4013, 0),
	(37, '400', '+1.5/2', '+1.5/2', 1, '让球', 4014, 0),
	(38, '400', '-2', '-2', 1, '让球', 4015, 0),
	(39, '400', '+2', '+2', 1, '让球', 4016, 0),
	(40, '400', '-2/2.5', '-2/2.5', 1, '让球', 4017, 0),
	(41, '400', '+2/2.5', '+2/2.5', 1, '让球', 4018, 0),
	(43, '100', 'Over 1/1.5', '大 1/1.5', 1, '大小球', 1001, 0),
	(44, '100', 'Over 2', '大 2', 1, '大小球', 1002, 0),
	(45, '100', 'Over 2/2.5', '大 2/2.5', 1, '大小球', 1003, 0),
	(46, '100', 'Over 2.5', '大 2.5', 1, '大小球', 1004, 0),
	(47, '100', 'Over 2.5/3', '大 2.5/3', 1, '大小球', 1005, 0),
	(48, '100', 'Over 3', '大 3', 1, '大小球', 1006, 0),
	(49, '100', 'Over 3/3.5', '大 3/3.5', 1, '大小球', 1007, 0),
	(50, '100', 'Over 3.5', '大 3.5', 1, '大小球', 1008, 0),
	(51, '100', 'Over 3.5/4', '大 3.5/4', 1, '大小球', 1009, 0),
	(52, '100', 'Under 1', '小 1', 1, '大小球', 1100, 0),
	(53, '100', 'Under 1/1.5', '小 1/1.5', 1, '大小球', 1101, 0),
	(54, '100', 'Under 2', '小 2', 1, '大小球', 1102, 0),
	(55, '100', 'Under 2/2.5', '小 2/2.5', 1, '大小球', 1103, 0),
	(56, '100', 'Under 2.5', '小 2.5', 1, '大小球', 1104, 0),
	(57, '100', 'Under 2.5/3', '小 2.5/3', 1, '大小球', 1105, 0),
	(58, '100', 'Under 3', '小 3', 1, '大小球', 1106, 0),
	(59, '100', 'Under 3/3.5', '小 3/3.5', 1, '大小球', 1107, 0),
	(60, '100', 'Under 3.5', '小 3.5', 1, '大小球', 1108, 0),
	(61, '100', 'Under 3.5/4', '小 3.5/4', 1, '大小球', 1109, 0),
	(62, '100', 'Over 1', '大 1', 1, '大小球', 1000, 99);

-- 导出  表 soccer.p_play_type 结构
DROP TABLE IF EXISTS `p_play_type`;
CREATE TABLE IF NOT EXISTS `p_play_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` int NOT NULL,
  `en_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` int NOT NULL,
  `class` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='玩法类型';

-- 正在导出表  soccer.p_play_type 的数据：~4 rows (大约)
DELETE FROM `p_play_type`;
INSERT INTO `p_play_type` (`id`, `code`, `en_name`, `name`, `status`, `class`) VALUES
	(1, 100, 'Over/Under', '大小球', 1, 'secondary'),
	(2, 200, '1x2', '独赢', 1, 'warning'),
	(3, 300, 'Correct Score', '波胆', 1, 'success'),
	(4, 400, 'Over/Under', '让球', 1, 'default');

-- 导出  表 soccer.p_team 结构
DROP TABLE IF EXISTS `p_team`;
CREATE TABLE IF NOT EXISTS `p_team` (
  `id` int NOT NULL AUTO_INCREMENT,
  `zh_name` varchar(250) COLLATE utf8mb4_general_ci NOT NULL,
  `en_name` varchar(250) COLLATE utf8mb4_general_ci NOT NULL,
  `vd_name` varchar(250) COLLATE utf8mb4_general_ci NOT NULL,
  `status` int NOT NULL,
  `icon` varchar(500) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='队伍';

-- 正在导出表  soccer.p_team 的数据：~2 rows (大约)
DELETE FROM `p_team`;
INSERT INTO `p_team` (`id`, `zh_name`, `en_name`, `vd_name`, `status`, `icon`) VALUES
	(1, '队1zh', '队1en', '队1vd', 1, '2/2023/8/Yahwcy.jpg'),
	(2, '队2zh', '队2en', '队2vd', 1, '2/2023/8/cq4SPx.png');

-- 导出  表 soccer.r_report_wallet_day 结构
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

-- 正在导出表  soccer.r_report_wallet_day 的数据：~6 rows (大约)
DELETE FROM `r_report_wallet_day`;
INSERT INTO `r_report_wallet_day` (`id`, `uid`, `account`, `pid`, `parent_path`, `balance_code`, `date`, `amount`, `created_at`, `updated_at`, `count`, `title`) VALUES
	(1694285301180010496, 121, 'join', 0, '', 400, '2023-08-23', 110.00, '2023-08-23 17:47:49', '2023-08-23 17:48:43', 2, '赠送'),
	(1694319539174510592, 121, 'join', 0, '', 501, '2023-08-23', 400.00, '2023-08-23 20:03:52', '2023-08-23 20:06:04', 3, '银行卡充值'),
	(1694650139647414272, 121, 'join', 0, '', -100, '2023-08-24', 200.00, '2023-08-24 17:57:33', '2023-08-24 18:03:35', 2, '提现'),
	(1694650189664489472, 121, 'join', 0, '', 100, '2023-08-24', 100.00, '2023-08-24 17:57:45', '2023-08-24 17:57:45', 1, '提现失败'),
	(1696780620941758464, 121, 'join', 0, '', -700, '2023-08-30', 9.00, '2023-08-30 15:03:19', '2023-08-30 15:03:26', 9, '下注'),
	(1696785479669649408, 121, 'join', 0, '', 0, '2023-08-30', 45.00, '2023-08-30 15:22:38', '2023-08-30 15:22:38', 1, '');

-- 导出  表 soccer.s_admin 结构
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

-- 正在导出表  soccer.s_admin 的数据：~2 rows (大约)
DELETE FROM `s_admin`;
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `pwd`, `nickname`, `email`, `phone`, `status`, `created_at`, `updated_at`, `key_secret`) VALUES
	(42, 1, 'admin', '$2a$10$i4zx9ccYBUu32p9kNwWFieommIgFHdewojOSoXY2iZ7zKkK/8OUC6', 'admin', '', '', 1, '2022-07-02 11:28:52', '2023-08-17 10:34:28', ''),
	(52, 1, 'guest', '$2a$10$i4zx9ccYBUu32p9kNwWFieommIgFHdewojOSoXY2iZ7zKkK/8OUC6', 'guest', '', '', 1, '2023-04-04 05:49:59', '2023-08-02 06:57:00', NULL);

-- 导出  表 soccer.s_admin_login_log 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=618 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  soccer.s_admin_login_log 的数据：~11 rows (大约)
DELETE FROM `s_admin_login_log`;
INSERT INTO `s_admin_login_log` (`id`, `uid`, `account`, `ip`, `created_at`, `updated_at`) VALUES
	(603, 42, 'admin', '103.171.144.43', '2023-08-25 06:59:45', '2023-08-25 06:59:45'),
	(604, 42, 'admin', '47.87.212.126', '2023-08-25 07:03:10', '2023-08-25 07:03:10'),
	(605, 42, 'admin', '::1', '2023-08-25 19:12:20', '2023-08-25 19:12:20'),
	(606, 42, 'admin', '::1', '2023-08-25 20:27:58', '2023-08-25 20:27:58'),
	(607, 42, 'admin', '::1', '2023-08-25 20:28:22', '2023-08-25 20:28:22'),
	(608, 42, 'admin', '103.171.144.43', '2023-08-25 12:41:09', '2023-08-25 12:41:09'),
	(609, 42, 'admin', '::1', '2023-08-28 14:07:04', '2023-08-28 14:07:04'),
	(610, 42, 'admin', '103.171.144.43', '2023-08-28 07:41:48', '2023-08-28 07:41:48'),
	(611, 42, 'admin', '::1', '2023-08-28 19:48:51', '2023-08-28 19:48:51'),
	(612, 42, 'admin', '::1', '2023-08-28 20:24:56', '2023-08-28 20:24:56'),
	(613, 42, 'admin', '103.171.144.43', '2023-08-29 14:04:33', '2023-08-29 14:04:33'),
	(614, 42, 'admin', '103.171.144.43', '2023-08-31 11:08:09', '2023-08-31 11:08:09'),
	(615, 42, 'admin', '47.87.212.126', '2023-08-31 12:40:25', '2023-08-31 12:40:25'),
	(616, 42, 'admin', '154.198.5.249', '2023-09-01 03:24:34', '2023-09-01 03:24:34'),
	(617, 42, 'admin', '154.198.5.249', '2023-09-01 03:24:43', '2023-09-01 03:24:43');

-- 导出  表 soccer.s_api 结构
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

-- 正在导出表  soccer.s_api 的数据：~67 rows (大约)
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
	(276, '/order/walletLog', '/api/wallet/walletLog', 'DELETE', ' 删除账变日志', '2023-08-11 12:36:45', '2023-08-11 12:36:45');

-- 导出  表 soccer.s_casbin 结构
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

-- 正在导出表  soccer.s_casbin 的数据：~104 rows (大约)
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

-- 导出  表 soccer.s_dict 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  soccer.s_dict 的数据：~3 rows (大约)
DELETE FROM `s_dict`;
INSERT INTO `s_dict` (`id`, `title`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES
	(42, '系统白名单', 'white_ips', 'localhost', '系统白名单', '1', 1, 3, '2022-12-28 12:25:41', '2023-08-13 03:35:34'),
	(44, '欢迎语', 'great', 'hello', 'great', '2', 1, 1, '2022-12-28 12:25:41', '2023-08-07 15:12:16'),
	(63, 'CloudFlare公共访问域', 'cloudflare_pub', 'https://pub-e700c1631e3a4dedaf546e73002fac2f.r2.dev/', 'CloudFlare公共访问域', '1', 1, 1, '2023-08-14 12:35:27', '2023-08-14 12:50:12');

-- 导出  表 soccer.s_file 结构
DROP TABLE IF EXISTS `s_file`;
CREATE TABLE IF NOT EXISTS `s_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `group` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=332 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  soccer.s_file 的数据：~59 rows (大约)
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
	(275, 'XGczat.jpeg', 2, 1, '2023-08-17 10:17:06', '2023-08-17 10:17:06'),
	(276, 'aYv1L1.png', 2, 1, '2023-08-22 20:23:56', '2023-08-22 20:23:56'),
	(277, 'SxyHWq.png', 2, 1, '2023-08-22 20:25:50', '2023-08-22 20:25:50'),
	(278, 'E5Cu96.png', 2, 1, '2023-08-22 20:29:19', '2023-08-22 20:29:19'),
	(279, 'WaUkXn.png', 2, 1, '2023-08-22 20:31:00', '2023-08-22 20:31:00'),
	(280, 'qPIiIc.png', 2, 1, '2023-08-22 20:31:58', '2023-08-22 20:31:58'),
	(281, 'Uvteyd.png', 2, 1, '2023-08-22 20:32:42', '2023-08-22 20:32:42'),
	(282, 'dcOatu.png', 2, 1, '2023-08-22 20:33:38', '2023-08-22 20:33:38'),
	(283, 'P8jfCv.png', 2, 1, '2023-08-22 20:34:16', '2023-08-22 20:34:16'),
	(284, '3UV6VN.png', 2, 1, '2023-08-22 20:34:53', '2023-08-22 20:34:53'),
	(285, 'KoPeOP.png', 2, 1, '2023-08-22 20:35:31', '2023-08-22 20:35:31'),
	(286, 'fZvP6c.png', 2, 1, '2023-08-22 20:36:05', '2023-08-22 20:36:05'),
	(287, '2a59uj.png', 2, 1, '2023-08-22 20:36:42', '2023-08-22 20:36:42'),
	(288, 'yQdtae.png', 2, 1, '2023-08-22 20:37:12', '2023-08-22 20:37:12'),
	(289, 'oySMOg.png', 2, 1, '2023-08-22 20:37:43', '2023-08-22 20:37:43'),
	(290, 'GJanwE.png', 2, 1, '2023-08-22 20:38:15', '2023-08-22 20:38:15'),
	(291, 'MozBdi.png', 2, 1, '2023-08-22 20:38:49', '2023-08-22 20:38:49'),
	(292, 'ydsvML.png', 2, 1, '2023-08-22 20:39:16', '2023-08-22 20:39:16'),
	(293, '6zlMYW.png', 2, 1, '2023-08-22 20:39:42', '2023-08-22 20:39:42'),
	(294, 'JefmgS.png', 2, 1, '2023-08-22 20:40:25', '2023-08-22 20:40:25'),
	(295, 'ikQ2e3.png', 2, 1, '2023-08-22 20:41:08', '2023-08-22 20:41:08'),
	(296, '4hCvPR.png', 2, 1, '2023-08-22 20:43:17', '2023-08-22 20:43:17'),
	(297, '4sdaeE.png', 2, 1, '2023-08-22 20:44:04', '2023-08-22 20:44:04'),
	(298, 'g1dyH9.png', 2, 1, '2023-08-22 20:44:35', '2023-08-22 20:44:35'),
	(299, 'KHya0K.png', 2, 1, '2023-08-22 20:45:12', '2023-08-22 20:45:12'),
	(300, 'VZ0Kho.png', 2, 1, '2023-08-22 20:45:45', '2023-08-22 20:45:45'),
	(301, 'wdRbmh.png', 2, 1, '2023-08-22 20:46:12', '2023-08-22 20:46:12'),
	(302, 'TDhgNg.png', 2, 1, '2023-08-22 20:46:39', '2023-08-22 20:46:39'),
	(303, 'j2CYVd.png', 2, 1, '2023-08-22 20:47:09', '2023-08-22 20:47:09'),
	(304, 'bfuxnW.png', 2, 1, '2023-08-22 20:47:49', '2023-08-22 20:47:49'),
	(305, '4fRPx8.png', 2, 1, '2023-08-22 20:48:14', '2023-08-22 20:48:14'),
	(306, 'rshp1F.png', 2, 1, '2023-08-22 20:48:35', '2023-08-22 20:48:35'),
	(307, '58Du2E.png', 2, 1, '2023-08-22 20:48:58', '2023-08-22 20:48:58'),
	(308, '6lfPGk.png', 2, 1, '2023-08-22 20:49:30', '2023-08-22 20:49:30'),
	(309, 'bJs0pg.png', 2, 1, '2023-08-22 20:50:12', '2023-08-22 20:50:12'),
	(310, 'WpodMg.png', 2, 1, '2023-08-22 20:50:41', '2023-08-22 20:50:41'),
	(311, 'BhL7G5.png', 2, 1, '2023-08-22 20:51:05', '2023-08-22 20:51:05'),
	(312, 'lCn583.png', 2, 1, '2023-08-22 20:51:29', '2023-08-22 20:51:29'),
	(313, 'gOaTLb.png', 2, 1, '2023-08-22 20:51:54', '2023-08-22 20:51:54'),
	(314, '06aWgR.png', 2, 1, '2023-08-22 20:52:19', '2023-08-22 20:52:19'),
	(315, 'Wm3e0J.png', 2, 1, '2023-08-22 20:52:53', '2023-08-22 20:52:53'),
	(316, 'hH9s8w.png', 2, 1, '2023-08-22 20:53:26', '2023-08-22 20:53:26'),
	(317, 'VqPVah.png', 2, 1, '2023-08-22 20:53:51', '2023-08-22 20:53:51'),
	(318, '7T06cV.png', 2, 1, '2023-08-22 20:54:24', '2023-08-22 20:54:24'),
	(325, '2/2023/8/cq4SPx.png', 2, 1, '2023-08-31 18:50:48', '2023-08-31 18:50:48'),
	(326, '2/2023/8/JcM5Zg.png', 2, 1, '2023-08-31 18:52:28', '2023-08-31 18:52:28'),
	(328, '2/2023/8/Yahwcy.jpg', 2, 1, '2023-08-31 11:09:08', '2023-08-31 11:09:08'),
	(329, '2/2023/8/DIygkD.png', 2, 1, '2023-08-31 20:33:13', '2023-08-31 20:33:13'),
	(330, '2/2023/8/wNmDTD.png', 2, 1, '2023-08-31 20:33:44', '2023-08-31 20:33:44'),
	(331, '2/2023/8/MQbmP7.png', 2, 1, '2023-08-31 20:34:04', '2023-08-31 20:34:04');

-- 导出  表 soccer.s_menu 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=359 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  soccer.s_menu 的数据：~59 rows (大约)
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
	(292, 289, '', '', '平台充值方式', '/admin/setting/category', 2.30, 1, 'u_amount_item', '', 1, NULL, '2023-08-23 18:30:21', '分类项目'),
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
	(345, 1, '', '', '文件', '/admin/sys/file', 1.20, 1, '这里是文件管理页面', '', 1, NULL, '2023-08-28 13:38:38', ''),
	(346, 345, '', '', '添加按钮', '#', 1.00, 3, '添加按钮', '', 1, '2023-08-12 20:44:03', '2023-08-12 20:44:03', 'file_add'),
	(347, 345, '', '', '编辑按钮', '#', 1.00, 3, '编辑按钮', '', 1, '2023-08-12 20:44:53', '2023-08-12 20:44:53', 'file_edit'),
	(348, 345, '', '', '删除按钮', '#', 1.00, 3, '删除按钮', '', 1, '2023-08-12 20:47:44', '2023-08-12 20:47:44', 'file_del'),
	(350, 289, '', '', '银行表', '/admin/setting/bank', 1.00, 1, '#', '', 1, NULL, '2023-08-14 22:30:31', ''),
	(351, -1, '', '', '足球', '#', 6.00, 2, '足球', '', 1, '2023-08-25 07:04:30', '2023-08-25 07:04:30', ''),
	(352, 351, '', '', '赛事', '/admin/soccer/event', 6.30, 1, '赛事', '', 1, NULL, '2023-08-25 16:13:16', ''),
	(353, 351, '', '', '已结订单', '#', 6.60, 1, '已结订单', '', 1, NULL, '2023-08-31 11:10:27', ''),
	(354, 351, '', '', '未结订单', '#', 6.50, 1, '未结订单', '', 1, NULL, '2023-08-31 11:10:36', ''),
	(355, 351, '', '', '赛事赔率', '/admin/soccer/evensOdds', 6.40, 1, '赔率', '', 1, NULL, '2023-08-28 15:58:55', ''),
	(356, 351, '', '', '玩法类型', '/admin/soccer/playType', 6.10, 1, '#', '', 1, NULL, '2023-08-28 15:59:16', ''),
	(357, 351, '', '', '玩法', '/admin/soccer/play', 6.20, 1, '#', '', 1, NULL, '2023-08-28 15:59:28', ''),
	(358, 351, '', '', '队伍', '/admin/soccer/team', 1.00, 1, '#', '', 1, '2023-08-31 18:46:03', '2023-08-31 18:46:03', '');

-- 导出  表 soccer.s_menu_api_rule 结构
DROP TABLE IF EXISTS `s_menu_api_rule`;
CREATE TABLE IF NOT EXISTS `s_menu_api_rule` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menu_id` bigint NOT NULL,
  `api_id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=354 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='功能菜单绑定API接口';

-- 正在导出表  soccer.s_menu_api_rule 的数据：~64 rows (大约)
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

-- 导出  表 soccer.s_operation_log 结构
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

-- 正在导出表  soccer.s_operation_log 的数据：~0 rows (大约)
DELETE FROM `s_operation_log`;

-- 导出  表 soccer.s_role 结构
DROP TABLE IF EXISTS `s_role`;
CREATE TABLE IF NOT EXISTS `s_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4831 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色';

-- 正在导出表  soccer.s_role 的数据：~3 rows (大约)
DELETE FROM `s_role`;
INSERT INTO `s_role` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
	(1, 'admin', 1, NULL, '2023-08-17 07:32:20'),
	(4827, '系统管理员', 1, NULL, '2023-08-17 15:18:37'),
	(4828, '财务', 1, NULL, '2023-08-10 20:26:42');

-- 导出  表 soccer.s_role_menu 结构
DROP TABLE IF EXISTS `s_role_menu`;
CREATE TABLE IF NOT EXISTS `s_role_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1394 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.s_role_menu 的数据：~78 rows (大约)
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

-- 导出  表 soccer.u_amount_category 结构
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

-- 正在导出表  soccer.u_amount_category 的数据：~2 rows (大约)
DELETE FROM `u_amount_category`;
INSERT INTO `u_amount_category` (`id`, `title`, `category`, `status`, `type`, `created_at`, `updated_at`) VALUES
	(1, '银行卡账转', '100', 1, '银行卡账转', '2023-08-11 16:49:34', '2023-08-11 16:49:34'),
	(2, 'GCASH', '200', 1, 'GCASH', '2023-08-11 16:49:57', '2023-08-11 16:49:57');

-- 导出  表 soccer.u_amount_item 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=86 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='分类项目';

-- 正在导出表  soccer.u_amount_item 的数据：~4 rows (大约)
DELETE FROM `u_amount_item`;
INSERT INTO `u_amount_item` (`id`, `title`, `status`, `detail`, `amount_category_id`, `net`, `min`, `max`, `fee`, `type`, `logo`, `sort`, `country`, `currency`, `protocol`, `created_at`, `updated_at`, `address`) VALUES
	(2, 'GCASH 转账', -1, 'BDO 转账', 1, 'GCASH', 1, 100, 0, 'Deposit', 'eydjs6.jpeg', 1, 'PHP', 'PHP', 'GCASH', '2023-08-11 16:54:24', '2023-08-23 18:35:44', '123456789'),
	(3, 'USDT(TRC20)', -1, 'USDT(TRC20) 指定账号', 1, 'TRON', 1, 100, 0, 'Deposit', '5T3wUb.jpeg', 1, 'PHP', 'PHP', 'Trc20', '2023-08-11 16:54:24', '2023-08-23 18:35:49', '1234567890'),
	(4, 'USDT(TRC20)', -1, 'USDT(TRC20) 每个用户生成一个', 1, 'TRON', 1, 100, 0, 'Deposit', 'rJVUz4.png', 1, 'PHP', 'PHP', 'Trc20', '2023-08-11 16:54:24', '2023-08-23 18:35:53', ''),
	(85, 'TCB', 1, 'Name:John', 0, 'BANK', 1, 1000, 0, 'Deposit', 'Uvteyd.png', 0, 'Vietnam', 'VND', 'TCB', '2023-08-23 18:38:27', '2023-08-23 18:39:24', '4123813103412323');

-- 导出  表 soccer.u_balance_code 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='改用户余额的编码。100 -200 正加钱，负扣钱';

-- 正在导出表  soccer.u_balance_code 的数据：~11 rows (大约)
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
	(10, '提现失败', 100, '提现失败', '2', 'error', 1, '2023-08-07 15:59:20', '2023-08-14 19:55:32'),
	(11, '利润', 400, '利润', 'profit', 'error', 1, '2023-08-21 20:03:42', '2023-08-21 20:03:45');

-- 导出  表 soccer.u_digital_account 结构
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

-- 正在导出表  soccer.u_digital_account 的数据：~0 rows (大约)
DELETE FROM `u_digital_account`;

-- 导出  表 soccer.u_user 结构
DROP TABLE IF EXISTS `u_user`;
CREATE TABLE IF NOT EXISTS `u_user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account` varchar(100) NOT NULL DEFAULT '' COMMENT '账号',
  `email` varchar(50) DEFAULT '' COMMENT '邮箱',
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
  `real_name` varchar(32) DEFAULT NULL COMMENT 'real name',
  `pay_pass` varchar(64) DEFAULT NULL COMMENT '交易密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=124 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

-- 正在导出表  soccer.u_user 的数据：~3 rows (大约)
DELETE FROM `u_user`;
INSERT INTO `u_user` (`id`, `account`, `email`, `password`, `status`, `xid`, `ip`, `client_agent`, `phone`, `level_bet`, `level_deposit`, `level_agent`, `pid`, `created_at`, `updated_at`, `parent_path`, `country`, `lang`, `area_code`, `city`, `real_name`, `pay_pass`) VALUES
	(121, 'join', 'test@gmail.com', '$2a$10$lxDAD0qEZ0/wb6GdYQJ66.mEWYWztJZxd70w0Yj7DZg8owT9k5gZK', 1, '', 'localhost', '', '38112093123', 0, 0, 0, 0, '2023-08-21 20:57:52', '2023-08-24 17:37:40', '', 'Japan', '', '', 'Tokyo', 'Tom', '$2a$10$e4HAoJZLse6SKXXsc7/KWuE2.y8gNDudM/86/WRaEL1CzhojA5MqO'),
	(122, 'tom2', '', '$2a$10$EnxMolGPLvMeukCGj3YOQ.rFJTVovoDPcEuJymOKN1pAQV83XDzyO', 1, '', '127.0.0.1', '', '', 0, 0, 0, 0, '2023-08-21 20:58:50', '2023-08-21 20:58:50', '', 'tom2', '', '', 'gaoxiong', 'tom2', NULL),
	(123, 'test001', '', '$2a$10$JajEdj4EA7TmVzl3phF5su53grsuuvsYRfIuDE6jSLkAFT678X0fS', 1, '', '154.198.5.249', 'PostmanRuntime/7.32.3', '', 0, 0, 0, 0, '2023-08-29 05:54:39', '2023-08-29 05:54:39', '', 'test001', '', '', 'string', 'string', '');

-- 导出  表 soccer.u_user_login_log 结构
DROP TABLE IF EXISTS `u_user_login_log`;
CREATE TABLE IF NOT EXISTS `u_user_login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint unsigned DEFAULT NULL,
  `account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  soccer.u_user_login_log 的数据：~29 rows (大约)
DELETE FROM `u_user_login_log`;
INSERT INTO `u_user_login_log` (`id`, `uid`, `account`, `ip`, `created_at`) VALUES
	(1, 121, 'join', '127.0.0.1', '2023-08-22 15:23:11'),
	(2, 121, 'join', '127.0.0.1', '2023-08-22 15:54:55'),
	(3, 121, 'join', '127.0.0.1', '2023-08-22 15:55:05'),
	(4, 121, 'join', '127.0.0.1', '2023-08-22 15:57:36'),
	(5, 121, 'join', '127.0.0.1', '2023-08-22 18:45:38'),
	(6, 121, 'join', '127.0.0.1', '2023-08-22 18:46:02'),
	(7, 121, 'join', '127.0.0.1', '2023-08-22 19:11:08'),
	(8, 121, 'join', '127.0.0.1', '2023-08-23 15:40:51'),
	(9, 121, 'join', '127.0.0.1', '2023-08-23 17:44:45'),
	(10, 121, 'join', '127.0.0.1', '2023-08-23 17:44:59'),
	(11, 121, 'join', '103.171.144.43', '2023-08-25 05:30:28'),
	(12, 121, 'join', '103.171.144.43', '2023-08-25 10:39:05'),
	(13, 123, 'test001', '154.198.5.249', '2023-08-29 05:54:54'),
	(14, 123, 'test001', '154.198.5.249', '2023-08-29 05:56:10'),
	(15, 123, 'test001', '154.198.5.249', '2023-08-29 05:57:03'),
	(16, 123, 'test001', '154.198.5.249', '2023-08-29 06:03:22'),
	(17, 123, 'test001', '154.198.5.249', '2023-08-29 06:17:38'),
	(18, 123, 'test001', '154.198.5.249', '2023-08-29 06:21:28'),
	(19, 123, 'test001', '154.198.5.249', '2023-08-29 06:46:03'),
	(20, 123, 'test001', '154.198.5.249', '2023-08-29 06:47:45'),
	(21, 123, 'test001', '154.198.5.249', '2023-08-29 06:48:20'),
	(22, 123, 'test001', '154.198.5.249', '2023-08-29 06:53:48'),
	(23, 123, 'test001', '154.198.5.249', '2023-08-29 06:56:00'),
	(24, 123, 'test001', '154.198.5.249', '2023-08-29 06:56:34'),
	(25, 123, 'test001', '154.198.5.249', '2023-08-29 07:00:16'),
	(26, 123, 'test001', '154.198.5.249', '2023-08-29 06:58:59'),
	(27, 123, 'test001', '154.198.5.249', '2023-08-29 07:01:28'),
	(28, 123, 'test001', '154.198.5.249', '2023-08-29 07:28:08'),
	(29, 123, 'test001', '154.198.5.249', '2023-08-29 07:29:37');

-- 导出  表 soccer.u_wallet 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包';

-- 正在导出表  soccer.u_wallet 的数据：~3 rows (大约)
DELETE FROM `u_wallet`;
INSERT INTO `u_wallet` (`id`, `uid`, `balance`, `total_bet`, `total_deposit`, `total_withdraw`, `freeze`, `account`, `parent_path`, `total_profit`, `total_gift`, `created_at`, `updated_at`) VALUES
	(14, 121, 440000, 20.00, 400.00, 200.00, 0, 'Tom', '', 32.10, 110.00, '2023-08-21 20:57:52', '2023-08-30 15:22:38'),
	(15, 122, 0, 0.00, 0.00, 0.00, 0, 'tom2', '', 0.00, 0.00, '2023-08-21 20:58:51', '2023-08-21 20:58:51'),
	(16, 123, 0, 0.00, 0.00, 0.00, 0, 'test001', '', 0.00, 0.00, '2023-08-29 05:54:39', '2023-08-29 05:54:39');

-- 导出  表 soccer.u_wallet_log 结构
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
) ENGINE=InnoDB AUTO_INCREMENT=208 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='钱包账变';

-- 正在导出表  soccer.u_wallet_log 的数据：~32 rows (大约)
DELETE FROM `u_wallet_log`;
INSERT INTO `u_wallet_log` (`id`, `order_no`, `uid`, `account`, `pid`, `balance_code`, `title`, `balance_before`, `balance_after`, `note`, `order_no_relation`, `parent_path`, `amount`, `created_at`, `update_at`) VALUES
	(171, 1694285295882604544, 121, 'join', 0, 400, '赠送', 0.00, 10.00, '', 0, '', 10.00, '2023-08-23 17:47:48', '2023-08-23 17:47:48'),
	(172, 1694285522697981952, 121, 'join', 0, 400, '赠送', 10.00, 110.00, '', 0, '', 100.00, '2023-08-23 17:48:42', '2023-08-23 17:48:42'),
	(173, 1694319533977767936, 121, 'join', 0, 501, '充值成功', 110.00, 210.00, 'TCB_4123813103412323', 1694319386690588672, '', 100.00, '2023-08-23 20:03:51', '2023-08-23 20:03:51'),
	(174, 1694319643700760576, 121, 'join', 0, 501, '充值成功', 210.00, 410.00, 'TCB_4123813103412323', 1694319620439150592, '', 200.00, '2023-08-23 20:04:17', '2023-08-23 20:04:17'),
	(175, 1694320089752408064, 121, 'join', 0, 501, '充值成功', 410.00, 510.00, 'TCB_4123813103412323', 1694320009880276992, '', 100.00, '2023-08-23 20:06:03', '2023-08-23 20:06:03'),
	(180, 1694582261493534720, 121, 'join', 0, -700, '小', 510.00, 509.00, '', 1694582258989535232, '', -1.00, '2023-08-24 13:27:50', '2023-08-24 13:27:50'),
	(181, 1694584145742663680, 121, 'join', 0, -700, '小', 509.00, 508.00, '', 1694584144580841472, '', -1.00, '2023-08-24 13:35:19', '2023-08-24 13:35:19'),
	(182, 1694584531895455744, 121, 'join', 0, -700, '大', 508.00, 507.00, '', 1694584530737827840, '', -1.00, '2023-08-24 13:36:51', '2023-08-24 13:36:51'),
	(183, 1694650130692575232, 121, 'join', 0, -100, 'Withdraw', 507.00, 407.00, 'BDO_3123441333', 1694650127437795328, '', -100.00, '2023-08-24 17:57:31', '2023-08-24 17:57:31'),
	(184, 1694650185172389888, 121, 'join', 0, 100, 'Fail ', 407.00, 507.00, '', 1694650127437795328, '', 100.00, '2023-08-24 17:57:44', '2023-08-24 17:57:44'),
	(185, 1694651649261309952, 121, 'join', 0, -100, 'Withdraw', 507.00, 407.00, 'PBVN_3333331233333', 1694651645213806592, '', -100.00, '2023-08-24 18:03:33', '2023-08-24 18:03:33'),
	(187, 1694663171429634048, 121, 'join', 0, 201, '1 vs 1', 408.00, 409.80, '大', 0, '', 1.80, '2023-08-24 18:49:22', '2023-08-24 18:49:22'),
	(188, 1694916424226574336, 121, 'join', 0, -700, '大', 409.00, 408.00, '', 1694916423035392000, '', -1.00, '2023-08-25 11:35:40', '2023-08-25 11:35:40'),
	(189, 1694916503977070592, 121, 'join', 0, -700, '小', 408.00, 407.00, '', 1694916502785888256, '', -1.00, '2023-08-25 11:35:59', '2023-08-25 11:35:59'),
	(190, 1694916533509165056, 121, 'join', 0, -700, '小', 407.00, 406.00, '', 1694916532317982720, '', -1.00, '2023-08-25 11:36:06', '2023-08-25 11:36:06'),
	(191, 1694917274521047040, 121, 'join', 0, 201, '1 vs 1', 406.00, 407.50, '大', 0, '', 1.50, '2023-08-25 11:39:07', '2023-08-25 11:39:07'),
	(192, 1696044296936886272, 121, 'join', 0, -700, '小', 407.00, 406.00, '', 1696044269132845056, '', -1.00, '2023-08-28 14:17:27', '2023-08-28 14:17:27'),
	(193, 1696044625573187584, 121, 'join', 0, -700, '大', 406.00, 405.00, '', 1696044624402976768, '', -1.00, '2023-08-28 14:18:45', '2023-08-28 14:18:45'),
	(194, 1696044688642936832, 121, 'join', 0, -700, '小', 405.00, 404.00, '', 1696044687468531712, '', -1.00, '2023-08-28 14:19:00', '2023-08-28 14:19:00'),
	(195, 1696107297631113216, 121, 'join', 0, -700, '小', 404.00, 403.00, '', 1696107296448319488, '', -1.00, '2023-08-28 18:27:47', '2023-08-28 18:27:47'),
	(196, 1696780617569538048, 121, 'join', 0, -700, '1x2', 403.00, 402.00, '', 1696780616298663936, '', -1.00, '2023-08-30 15:03:19', '2023-08-30 15:03:19'),
	(197, 1696780622162300928, 121, 'join', 0, -700, '1x2', 402.00, 401.00, '', 1696780621126307840, '', -1.00, '2023-08-30 15:03:20', '2023-08-30 15:03:20'),
	(198, 1696780625849094144, 121, 'join', 0, -700, '1x2', 401.00, 400.00, '', 1696780625006039040, '', -1.00, '2023-08-30 15:03:21', '2023-08-30 15:03:21'),
	(199, 1696780629061931008, 121, 'join', 0, -700, 'Over/Under', 400.00, 399.00, '', 1696780628218875904, '', -1.00, '2023-08-30 15:03:21', '2023-08-30 15:03:21'),
	(200, 1696780632278962176, 121, 'join', 0, -700, '1x2', 399.00, 398.00, '', 1696780631435907072, '', -1.00, '2023-08-30 15:03:22', '2023-08-30 15:03:22'),
	(201, 1696780635487604736, 121, 'join', 0, -700, '1x2', 398.00, 397.00, '', 1696780634644549632, '', -1.00, '2023-08-30 15:03:23', '2023-08-30 15:03:23'),
	(202, 1696780639279255552, 121, 'join', 0, -700, '1x2', 397.00, 396.00, '', 1696780637869969408, '', -1.00, '2023-08-30 15:03:24', '2023-08-30 15:03:24'),
	(203, 1696780642492092416, 121, 'join', 0, -700, '1x2', 396.00, 395.00, '', 1696780641649037312, '', -1.00, '2023-08-30 15:03:25', '2023-08-30 15:03:25'),
	(204, 1696780645717512192, 121, 'join', 0, -700, '1x2', 395.00, 394.00, '', 1696780644866068480, '', -1.00, '2023-08-30 15:03:25', '2023-08-30 15:03:25'),
	(205, 1696780648930349056, 121, 'join', 0, -700, '1x2', 394.00, 393.00, '', 1696780648083099648, '', -1.00, '2023-08-30 15:03:26', '2023-08-30 15:03:26'),
	(206, 1696785433263869952, 121, 'join', 0, 201, '曼彻斯特联 vs 阿森纳', 393.00, 438.00, 'Over/Under', 0, '', 45.00, '2023-08-30 15:22:31', '2023-08-30 15:22:31'),
	(207, 1696785482752462848, 121, 'join', 0, 201, '曼彻斯特联 vs 阿森纳', 438.00, 440.00, '1x2', 0, '', 2.00, '2023-08-30 15:22:40', '2023-08-30 15:22:40');

-- 导出  表 soccer.u_withdraw_account 结构
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

-- 正在导出表  soccer.u_withdraw_account 的数据：~2 rows (大约)
DELETE FROM `u_withdraw_account`;
INSERT INTO `u_withdraw_account` (`id`, `net`, `protocol`, `uid`, `account`, `address`, `currency`, `status`, `default`, `title`) VALUES
	(1694634395836616704, 'BANK', 'BDO', 121, 'john', '3123441333', 'PHP', 1, -1, 'John'),
	(1694635985192620032, 'BANK', 'PBVN', 121, 'join', '3333331233333', 'VND', 1, 1, 'John');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
