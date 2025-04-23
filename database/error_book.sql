-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        8.0.12 - MySQL Community Server - GPL
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 error_book 的数据库结构
CREATE DATABASE IF NOT EXISTS `error_book` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `error_book`;

-- 导出  表 error_book.file_tags 结构
CREATE TABLE IF NOT EXISTS `file_tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `desc` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  error_book.file_tags 的数据：0 rows
DELETE FROM `file_tags`;
/*!40000 ALTER TABLE `file_tags` DISABLE KEYS */;
/*!40000 ALTER TABLE `file_tags` ENABLE KEYS */;

-- 导出  表 error_book.file_uploads 结构
CREATE TABLE IF NOT EXISTS `file_uploads` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `file_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `file_url` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `file_size` bigint(20) NOT NULL,
  `file_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `tag_id` int(11) DEFAULT NULL,
  `status` tinyint(4) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 正在导出表  error_book.file_uploads 的数据：0 rows
DELETE FROM `file_uploads`;
/*!40000 ALTER TABLE `file_uploads` DISABLE KEYS */;
/*!40000 ALTER TABLE `file_uploads` ENABLE KEYS */;

-- 导出  表 error_book.questions 结构
CREATE TABLE IF NOT EXISTS `questions` (
  `uuid` int(11) NOT NULL,
  `id` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `question` varchar(1500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `option1` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `option2` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `option3` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `option4` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `answer` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `image` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `difficulty` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `wrong_count` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `correct_count` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='驾驶员之家抓题表';

-- 正在导出表  error_book.questions 的数据：~2 rows (大约)
DELETE FROM `questions`;
/*!40000 ALTER TABLE `questions` DISABLE KEYS */;
INSERT INTO `questions` (`uuid`, `id`, `type`, `question`, `option1`, `option2`, `option3`, `option4`, `answer`, `image`, `difficulty`, `wrong_count`, `correct_count`) VALUES
	(2088, '2250', NULL, '在高速公路上行车时，错过当前要驶出的出口时，应当在应急停车道上停车，等待车辆较少的时候再伺机倒车。', '正确', '错误', '错误', '错误', '', './pic3', NULL, NULL, NULL),
	(2089, '2251', NULL, '在高速公路上行车时，错过当前要驶出的出口时，应当在应急停车道上停车，等待车辆较少的时候再伺机倒车。', '正确', '错误', '错误', '错误', '', './pic4', NULL, NULL, NULL);
/*!40000 ALTER TABLE `questions` ENABLE KEYS */;

-- 导出  表 error_book.users 结构
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `password_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 正在导出表  error_book.users 的数据：1 rows
DELETE FROM `users`;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `username`, `password_hash`, `created_at`, `updated_at`, `status`) VALUES
	(1, 'vzjtu609', 'yujun1006', '2025-04-01 17:21:44', '2025-04-23 10:20:58', 1);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
