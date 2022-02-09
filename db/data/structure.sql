-- Adminer 4.7.7 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `repository`;
CREATE TABLE `repository` (
                              `id` int unsigned NOT NULL AUTO_INCREMENT,
                              `user_id` int unsigned NOT NULL,
                              `type` int unsigned NOT NULL,
                              `private` tinyint unsigned NOT NULL,
                              `service_address` varchar(1024) NOT NULL,
                              `insecure` tinyint unsigned NOT NULL,
                              `username` varchar(1024) NOT NULL,
                              `secret` varchar(1024) NOT NULL,
                              `name` varchar(1024) NOT NULL,
                              `reference` varchar(1024) NOT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
                        `id` int unsigned NOT NULL AUTO_INCREMENT,
                        `user_id` int unsigned NOT NULL,
                        `repository_id` int unsigned NOT NULL,
                        `type` int unsigned NOT NULL,
                        `origin_url` varchar(4096) NOT NULL,
                        `status` int unsigned NOT NULL,
                        `archive_password` varchar(1024) NOT NULL,
                        `filename` varchar(1024) NOT NULL,
                        `digest` varchar(1024) NOT NULL,
                        `size` int NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` int unsigned NOT NULL AUTO_INCREMENT,
                        `username` varchar(255) NOT NULL,
                        `password` varchar(255) NOT NULL,
                        `is_admin` tinyint unsigned NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- 2021-01-13 16:58:52

DROP TABLE IF EXISTS `node`;
CREATE TABLE `node` (
                        `id` int unsigned NOT NULL AUTO_INCREMENT,
                        `index` int unsigned NOT NULL,
                        `name` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        `node_type` tinyint unsigned NOT NULL,
                        `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        `content_id` int unsigned NOT NULL,
                        `content_type` int unsigned NOT NULL,
                        `leaf_type` tinyint unsigned NOT NULL,
                        `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `node_relation`;
CREATE TABLE `node_relation` (
                                 `id` int unsigned NOT NULL AUTO_INCREMENT,
                                 `parent` int unsigned NOT NULL,
                                 `child` int unsigned NOT NULL,
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `file`;
CREATE TABLE `file` (
                        `id` int unsigned NOT NULL AUTO_INCREMENT,
                        `filename` varchar(1024) NOT NULL,
                        `digest` varchar(1024) NOT NULL,
                        `size` int NOT NULL,
                        `user_id` int unsigned NOT NULL,
                        `repository_id` int unsigned NOT NULL,
                        `origin_url` varchar(10240) NOT NULL,
                        `archive_password` varchar(1024) NOT NULL,
                        `last_modified` datetime NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;