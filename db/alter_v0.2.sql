CREATE TABLE IF NOT EXISTS `node` (
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

CREATE TABLE IF NOT EXISTS `node_relation` (
                                 `id` int unsigned NOT NULL AUTO_INCREMENT,
                                 `parent` int unsigned NOT NULL,
                                 `child` int unsigned NOT NULL,
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `file` (
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