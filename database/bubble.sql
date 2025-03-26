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


-- 导出 bubble 的数据库结构
CREATE DATABASE IF NOT EXISTS `bubble` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `bubble`;

-- 导出  表 bubble.classmates 结构
CREATE TABLE IF NOT EXISTS `classmates` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `work` varchar(255) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `education` varchar(255) DEFAULT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 正在导出表  bubble.classmates 的数据：1 rows
DELETE FROM `classmates`;
/*!40000 ALTER TABLE `classmates` DISABLE KEYS */;
INSERT INTO `classmates` (`id`, `name`, `work`, `age`, `education`, `created_at`, `updated_at`) VALUES
	(1, '叶兄', '前端', 23, '本科', 0, 0);
/*!40000 ALTER TABLE `classmates` ENABLE KEYS */;

-- 导出  表 bubble.course 结构
CREATE TABLE IF NOT EXISTS `course` (
  `cid` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `cname` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `tid` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 正在导出表  bubble.course 的数据：9 rows
DELETE FROM `course`;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;
INSERT INTO `course` (`cid`, `cname`, `tid`) VALUES
	('01', '语文', '02'),
	('02', '数学', '01'),
	('03', '英语', '03'),
	('01', '语文', '02'),
	('02', '数学', '01'),
	('03', '英语', '03'),
	('01', '语文', '02'),
	('02', '数学', '01'),
	('03', '英语', '03');
/*!40000 ALTER TABLE `course` ENABLE KEYS */;

-- 导出  表 bubble.nongs 结构
CREATE TABLE IF NOT EXISTS `nongs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `insert_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 正在导出表  bubble.nongs 的数据：17 rows
DELETE FROM `nongs`;
/*!40000 ALTER TABLE `nongs` DISABLE KEYS */;
INSERT INTO `nongs` (`id`, `name`, `content`, `insert_time`, `update_time`) VALUES
	(1, '最近更新', '哈哈', '2021-04-30 01:11:36', '2021-04-30 01:32:18'),
	(2, 'nihao', 'hhhhh', '2021-04-30 01:11:41', '2021-04-30 01:11:36'),
	(3, 'nihao', 'hhhhh', '2021-04-30 01:11:37', '2021-04-30 01:11:37'),
	(4, 'nihao', 'hhhhh', '2021-04-30 01:11:38', '2021-04-30 01:11:38'),
	(5, 'io', 'oo', '2021-04-30 01:11:39', '2021-04-30 01:11:38'),
	(6, '我', '赶紧辞职了', '2021-04-30 01:11:39', '2021-04-30 01:11:39'),
	(7, '你好啊', '最近如何', '2021-04-30 01:11:40', '2021-04-30 01:11:40'),
	(8, '你好的啊', '最近如何', '2021-04-30 01:11:41', '2021-04-30 01:11:41'),
	(9, '你好部的啊', '最近如何', '2021-04-30 01:14:11', NULL),
	(10, '你好胡部的啊', '最近如何', '2021-04-30 01:15:05', '2021-04-30 01:15:05'),
	(11, '你好胡部的啊', '最近如何', NULL, '2021-04-30 01:30:26'),
	(12, '你好胡部的啊', '最近如何', '2021-04-30 01:30:55', '2021-04-30 01:30:55'),
	(13, 'nihao', 'hhhhh', '2021-05-01 12:55:18', '2021-05-01 12:55:18'),
	(15, '我不会是哈哈', 'halou', '2021-06-08 00:56:47', '2021-06-08 01:05:23'),
	(16, '哈哈', 'xiix', '2021-06-08 00:58:17', '2021-06-08 00:58:17'),
	(18, '', '', '2021-06-08 00:59:29', '2021-06-08 00:59:29'),
	(19, '哈哈', 'halou', '2021-06-08 01:01:17', '2021-06-08 01:01:17');
/*!40000 ALTER TABLE `nongs` ENABLE KEYS */;

-- 导出  表 bubble.sc 结构
CREATE TABLE IF NOT EXISTS `sc` (
  `sid` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `cid` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `score` decimal(18,1) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 正在导出表  bubble.sc 的数据：54 rows
DELETE FROM `sc`;
/*!40000 ALTER TABLE `sc` DISABLE KEYS */;
INSERT INTO `sc` (`sid`, `cid`, `score`) VALUES
	('01', '01', 80.0),
	('01', '02', 90.0),
	('01', '03', 99.0),
	('02', '01', 70.0),
	('02', '02', 60.0),
	('02', '03', 80.0),
	('03', '01', 80.0),
	('03', '02', 80.0),
	('03', '03', 80.0),
	('04', '01', 50.0),
	('04', '02', 30.0),
	('04', '03', 20.0),
	('05', '01', 76.0),
	('05', '02', 87.0),
	('06', '01', 31.0),
	('06', '03', 34.0),
	('07', '02', 89.0),
	('07', '03', 98.0),
	('01', '01', 80.0),
	('01', '02', 90.0),
	('01', '03', 99.0),
	('02', '01', 70.0),
	('02', '02', 60.0),
	('02', '03', 80.0),
	('03', '01', 80.0),
	('03', '02', 80.0),
	('03', '03', 80.0),
	('04', '01', 50.0),
	('04', '02', 30.0),
	('04', '03', 20.0),
	('05', '01', 76.0),
	('05', '02', 87.0),
	('06', '01', 31.0),
	('06', '03', 34.0),
	('07', '02', 89.0),
	('07', '03', 98.0),
	('01', '01', 80.0),
	('01', '02', 90.0),
	('01', '03', 99.0),
	('02', '01', 70.0),
	('02', '02', 60.0),
	('02', '03', 80.0),
	('03', '01', 80.0),
	('03', '02', 80.0),
	('03', '03', 80.0),
	('04', '01', 50.0),
	('04', '02', 30.0),
	('04', '03', 20.0),
	('05', '01', 76.0),
	('05', '02', 87.0),
	('06', '01', 31.0),
	('06', '03', 34.0),
	('07', '02', 89.0),
	('07', '03', 98.0);
/*!40000 ALTER TABLE `sc` ENABLE KEYS */;

-- 导出  表 bubble.student 结构
CREATE TABLE IF NOT EXISTS `student` (
  `sid` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `sname` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `sage` datetime DEFAULT NULL,
  `ssex` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 正在导出表  bubble.student 的数据：24 rows
DELETE FROM `student`;
/*!40000 ALTER TABLE `student` DISABLE KEYS */;
INSERT INTO `student` (`sid`, `sname`, `sage`, `ssex`) VALUES
	('01', '赵雷', '1990-01-01 00:00:00', '男'),
	('02', '钱电', '1990-12-21 00:00:00', '男'),
	('03', '孙风', '1990-05-20 00:00:00', '男'),
	('04', '李云', '1990-08-06 00:00:00', '男'),
	('05', '周梅', '1991-12-01 00:00:00', '女'),
	('06', '吴兰', '1992-03-01 00:00:00', '女'),
	('07', '郑竹', '1989-07-01 00:00:00', '女'),
	('08', '王菊', '1990-01-20 00:00:00', '女'),
	('01', '赵雷', '1990-01-01 00:00:00', '男'),
	('02', '钱电', '1990-12-21 00:00:00', '男'),
	('03', '孙风', '1990-05-20 00:00:00', '男'),
	('04', '李云', '1990-08-06 00:00:00', '男'),
	('05', '周梅', '1991-12-01 00:00:00', '女'),
	('06', '吴兰', '1992-03-01 00:00:00', '女'),
	('07', '郑竹', '1989-07-01 00:00:00', '女'),
	('08', '王菊', '1990-01-20 00:00:00', '女'),
	('01', '赵雷', '1990-01-01 00:00:00', '男'),
	('02', '钱电', '1990-12-21 00:00:00', '男'),
	('03', '孙风', '1990-05-20 00:00:00', '男'),
	('04', '李云', '1990-08-06 00:00:00', '男'),
	('05', '周梅', '1991-12-01 00:00:00', '女'),
	('06', '吴兰', '1992-03-01 00:00:00', '女'),
	('07', '郑竹', '1989-07-01 00:00:00', '女'),
	('08', '王菊', '1990-01-20 00:00:00', '女');
/*!40000 ALTER TABLE `student` ENABLE KEYS */;

-- 导出  表 bubble.teacher 结构
CREATE TABLE IF NOT EXISTS `teacher` (
  `tid` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `tname` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 正在导出表  bubble.teacher 的数据：9 rows
DELETE FROM `teacher`;
/*!40000 ALTER TABLE `teacher` DISABLE KEYS */;
INSERT INTO `teacher` (`tid`, `tname`) VALUES
	('01', '张三'),
	('02', '李四'),
	('03', '王五'),
	('01', '张三'),
	('02', '李四'),
	('03', '王五'),
	('01', '张三'),
	('02', '李四'),
	('03', '王五');
/*!40000 ALTER TABLE `teacher` ENABLE KEYS */;

-- 导出  表 bubble.todos 结构
CREATE TABLE IF NOT EXISTS `todos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 正在导出表  bubble.todos 的数据：2 rows
DELETE FROM `todos`;
/*!40000 ALTER TABLE `todos` DISABLE KEYS */;
INSERT INTO `todos` (`id`, `title`, `status`) VALUES
	(1, 'hhh', 0),
	(2, 'jjj', 0);
/*!40000 ALTER TABLE `todos` ENABLE KEYS */;

-- 导出  表 bubble.user 结构
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(10) unsigned zerofill NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户注册表';

-- 正在导出表  bubble.user 的数据：1 rows
DELETE FROM `user`;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (`id`, `name`, `password`, `created_at`, `updated_at`) VALUES
	(0000000001, 'hello', '123456', '2021-05-03 09:38:17', '2021-05-03 09:38:18');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
