-- --------------------------------------------------------
-- 主机:                           114.132.48.179
-- 服务器版本:                        8.0.24 - Source distribution
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 db_goskeleton 的数据库结构
CREATE DATABASE IF NOT EXISTS `db_goskeleton` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `db_goskeleton`;

-- 导出  表 db_goskeleton.tb_auth_casbin_rule 结构
CREATE TABLE IF NOT EXISTS `tb_auth_casbin_rule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT 'p',
  `v0` varchar(100) DEFAULT '',
  `v1` varchar(100) DEFAULT '',
  `v2` varchar(100) DEFAULT '*',
  `v3` varchar(100) DEFAULT '',
  `v4` varchar(100) DEFAULT '',
  `v5` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_index` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- 正在导出表  db_goskeleton.tb_auth_casbin_rule 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `tb_auth_casbin_rule` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_auth_casbin_rule` ENABLE KEYS */;

-- 导出  表 db_goskeleton.tb_oauth_access_tokens 结构
CREATE TABLE IF NOT EXISTS `tb_oauth_access_tokens` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `fr_user_id` int DEFAULT '0' COMMENT '外键:tb_users表id',
  `client_id` int unsigned DEFAULT '1' COMMENT '普通用户的授权，默认为1',
  `token` varchar(500) DEFAULT NULL,
  `action_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '' COMMENT 'login|refresh|reset表示token生成动作',
  `scopes` varchar(128) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '[*]' COMMENT '暂时预留,未启用',
  `revoked` tinyint(1) DEFAULT '0' COMMENT '是否撤销',
  `client_ip` varchar(128) DEFAULT NULL COMMENT 'ipv6最长为128位',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `expires_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `oauth_access_tokens_user_id_index` (`fr_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb3;

-- 正在导出表  db_goskeleton.tb_oauth_access_tokens 的数据：~26 rows (大约)
/*!40000 ALTER TABLE `tb_oauth_access_tokens` DISABLE KEYS */;
REPLACE INTO `tb_oauth_access_tokens` (`id`, `fr_user_id`, `client_id`, `token`, `action_name`, `scopes`, `revoked`, `client_ip`, `created_at`, `updated_at`, `expires_at`) VALUES
	(1, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg4MzQzNzkxLCJuYmYiOjE2ODgzMTQ5ODF9.8t3wGqv7mTjweogkJMYWVs5s7ZFNYt4bsXs0P_uM5S8', 'login', '[*]', 0, '::1', '2023-07-03 00:23:14', '2023-07-03 00:23:14', '2023-07-03 08:23:11'),
	(2, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg4MzQ0Mjk2LCJuYmYiOjE2ODgzMTU0ODZ9.b7925Xzxu6ambrO3ovi44Muvu30_d3gunyQdKY3z0XI', 'login', '[*]', 0, '::1', '2023-07-03 00:31:39', '2023-07-03 00:31:39', '2023-07-03 08:31:36'),
	(3, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg4MzUwMTYxLCJuYmYiOjE2ODgzMjEzNTF9.Xn-qy-wOa-sw3wtwgbxqrPbgD3PWoqEQDDTXa0nCMEg', 'login', '[*]', 0, '27.151.69.194', '2023-07-03 02:09:21', '2023-07-03 02:09:21', '2023-07-03 10:09:21'),
	(4, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MDgxNzA1LCJuYmYiOjE2ODkwNTI4OTV9._5b1WXoFA1nl0iH00Tm7gaduaps3dYEJipG5pbH-pYE', 'login', '[*]', 0, '183.250.189.232', '2023-07-11 13:21:45', '2023-07-11 13:21:45', '2023-07-11 21:21:45'),
	(5, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MDg0MDQzLCJuYmYiOjE2ODkwNTUyMzN9._0ceEo1g37ilgOyF7aYOoo9SsdKNy-kBWomEDH4-dmw', 'login', '[*]', 0, '183.250.189.232', '2023-07-11 14:00:43', '2023-07-11 14:00:43', '2023-07-11 22:00:43'),
	(6, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MDg0MDcyLCJuYmYiOjE2ODkwNTUyNjJ9.aQ4VgZQDwdkMsdD8ZrDDbxNLRR-9hzZ1XWf7avqfhXM', 'login', '[*]', 0, '183.250.189.232', '2023-07-11 14:01:12', '2023-07-11 14:01:12', '2023-07-11 22:01:12'),
	(7, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MDg0MDkxLCJuYmYiOjE2ODkwNTUyODF9.eu1PGp7crU6564rGQBiXANJwiSF3HHGT8iyYzmhQhW4', 'login', '[*]', 0, '183.250.189.232', '2023-07-11 14:01:31', '2023-07-11 14:01:31', '2023-07-11 22:01:31'),
	(8, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MDg0MTA0LCJuYmYiOjE2ODkwNTUyOTR9.iz6biesj9QzIdpoXygrbHGkLXx3s3xBEoZnAK9QP-p0', 'login', '[*]', 0, '183.250.189.232', '2023-07-11 14:01:44', '2023-07-11 14:01:44', '2023-07-11 22:01:44'),
	(9, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MDg0MzU1LCJuYmYiOjE2ODkwNTU1NDV9._83emejcejbvzgTNdG5J620heNhv46QtBuXK78KPHbk', 'login', '[*]', 0, '183.250.189.232', '2023-07-11 14:05:55', '2023-07-11 14:05:55', '2023-07-11 22:05:55'),
	(10, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MDg0NTQ3LCJuYmYiOjE2ODkwNTU3Mzd9.BAuQEaGvardStRM8gSGyeo0Xrs73J4S9ZX_ez76zXW0', 'login', '[*]', 0, '183.250.189.232', '2023-07-11 14:09:07', '2023-07-11 14:09:07', '2023-07-11 22:09:07'),
	(11, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MjY5MjM2LCJuYmYiOjE2ODkyNDA0MjZ9.U7MjKAYZ2bP9F9Jk7f8pZWmbCZ37007hrb7cWmFTv20', 'login', '[*]', 0, '183.250.189.232', '2023-07-13 17:27:16', '2023-07-13 17:27:16', '2023-07-14 01:27:16'),
	(12, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MjY5NzQ2LCJuYmYiOjE2ODkyNDA5MzZ9.oMu7O1PoFm48PUKbpyTAMu3ZJuTTOdR1Rbu-zTZIdRo', 'login', '[*]', 0, '183.250.189.232', '2023-07-13 17:35:46', '2023-07-13 17:35:46', '2023-07-14 01:35:46'),
	(13, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MjcwNDY3LCJuYmYiOjE2ODkyNDE2NTd9.DjhlJP9LWPjSjOkxWvbuchqchDf4zklnd26AvLNY3Lo', 'login', '[*]', 0, '183.250.189.232', '2023-07-13 17:47:47', '2023-07-13 17:47:47', '2023-07-14 01:47:47'),
	(14, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MjcwNTAwLCJuYmYiOjE2ODkyNDE2OTB9.G1yFBquxola98Hd6qvgtCO9gzMmyMDTmseeX2i2A4_g', 'login', '[*]', 0, '183.250.189.232', '2023-07-13 17:48:20', '2023-07-13 17:48:20', '2023-07-14 01:48:20'),
	(15, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MzI3MDM5LCJuYmYiOjE2ODkyOTgyMjl9.Bs92wb8MJwVKGdJT9owvIHm--ZVdgYZzx2mvhi2QAXE', 'login', '[*]', 0, '183.250.189.232', '2023-07-14 09:30:39', '2023-07-14 09:30:39', '2023-07-14 17:30:39'),
	(16, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5MzQ0ODUyLCJuYmYiOjE2ODkzMTYwNDJ9.ftwD7gaUgIZTvkHLkgDkkuywFtswlGcWOGLdfyGTNP4', 'login', '[*]', 0, '183.250.189.232', '2023-07-14 14:27:32', '2023-07-14 14:27:32', '2023-07-14 22:27:32'),
	(17, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5NTk3Nzc1LCJuYmYiOjE2ODk1Njg5NjV9.FNvUia92v_pIBITnUN08EItk8baQ52Pkv7DL3pMczTc', 'login', '[*]', 0, '112.111.13.23', '2023-07-17 12:42:55', '2023-07-17 12:42:55', '2023-07-17 20:42:55'),
	(18, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5NTk3ODkxLCJuYmYiOjE2ODk1NjkwODF9.hgmP1C1weBixD73WSkD0xUO-xWvY3E1pNlIWdoYQEPE', 'login', '[*]', 0, '112.111.13.23', '2023-07-17 12:44:51', '2023-07-17 12:44:51', '2023-07-17 20:44:51'),
	(19, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5NTk4MjI2LCJuYmYiOjE2ODk1Njk0MTZ9.69xPoRulZYWalyJs1Woe_a-xswrjZxo0yzKKHOsHrpQ', 'login', '[*]', 0, '112.111.13.23', '2023-07-17 12:50:26', '2023-07-17 12:50:26', '2023-07-17 20:50:26'),
	(20, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5NTk4OTgyLCJuYmYiOjE2ODk1NzAxNzJ9.ufWuOzPV_vYHU9MRyv2Rsh1ZG28RLlw_FxtAsJsV0Hs', 'login', '[*]', 0, '112.111.13.23', '2023-07-17 13:03:02', '2023-07-17 13:03:02', '2023-07-17 21:03:02'),
	(21, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5NjAzNDA4LCJuYmYiOjE2ODk1NzQ1OTh9.yTQbrhMAura6dbZpizASmKJ4iCQuPIBqbRbTtBJ_NlU', 'login', '[*]', 0, '112.111.13.23', '2023-07-17 14:16:48', '2023-07-17 14:16:48', '2023-07-17 22:16:48'),
	(22, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5NjA2MDgwLCJuYmYiOjE2ODk1NzcyNzB9.cgs2mP890kNzklX-D4AhokJsWf89u9DM_kQi1H_ADs0', 'login', '[*]', 0, '112.111.13.23', '2023-07-17 15:01:20', '2023-07-17 15:01:20', '2023-07-17 23:01:20'),
	(23, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5Njc5NTkwLCJuYmYiOjE2ODk2NTA3ODB9.rVjuLKQlP14XCppbBpBKHAbzyc9FaNa3DbZUdDD3RM0', 'login', '[*]', 0, '112.111.13.23', '2023-07-18 11:26:30', '2023-07-18 11:26:30', '2023-07-18 19:26:30'),
	(24, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5NzcwMzkxLCJuYmYiOjE2ODk3NDE1ODF9.ofb_QAx0cwqL6gHda-TiGk8UbaPHQRGr-Sw_vHwS8ms', 'login', '[*]', 0, '112.111.13.23', '2023-07-19 12:39:51', '2023-07-19 12:39:51', '2023-07-19 20:39:51'),
	(25, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5NzcwMzk3LCJuYmYiOjE2ODk3NDE1ODd9.5P28jYI1490MNyXCJCGEwEuNLRpvdl5Rk-Rmd1__gYE', 'login', '[*]', 0, '112.111.13.23', '2023-07-19 12:39:57', '2023-07-19 12:39:57', '2023-07-19 20:39:57'),
	(26, 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJoYyIsInBob25lIjoiIiwiZXhwIjoxNjg5NzcwOTE5LCJuYmYiOjE2ODk3NDIxMDl9.x43GNqVi7k0d4IHP7J9KkgPIcO-Qkw7ihGLoabDskkk', 'login', '[*]', 0, '112.111.13.23', '2023-07-19 12:48:39', '2023-07-19 12:48:39', '2023-07-19 20:48:39');
/*!40000 ALTER TABLE `tb_oauth_access_tokens` ENABLE KEYS */;

-- 导出  表 db_goskeleton.tb_users 结构
CREATE TABLE IF NOT EXISTS `tb_users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(30) DEFAULT '' COMMENT '账号',
  `pass` varchar(128) DEFAULT '' COMMENT '密码',
  `real_name` varchar(30) DEFAULT '' COMMENT '姓名',
  `phone` char(11) DEFAULT '' COMMENT '手机',
  `status` tinyint DEFAULT '1' COMMENT '状态',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `last_login_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `last_login_ip` char(30) DEFAULT '' COMMENT '最近一次登录ip',
  `login_times` int DEFAULT '0' COMMENT '累计登录次数',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- 正在导出表  db_goskeleton.tb_users 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `tb_users` DISABLE KEYS */;
REPLACE INTO `tb_users` (`id`, `user_name`, `pass`, `real_name`, `phone`, `status`, `remark`, `last_login_time`, `last_login_ip`, `login_times`, `created_at`, `updated_at`) VALUES
	(1, 'hc', '87d9bb400c0634691f0e3baaf1e2fd0d', '', '', 1, '', '2023-07-19 12:48:39', '112.111.13.23', 26, '2023-07-03 00:21:29', '2023-07-03 00:21:29');
/*!40000 ALTER TABLE `tb_users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
