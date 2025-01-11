CREATE TABLE `departments` (
  `id` bigint UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) UNIQUE NOT NULL,
  `uuid` varchar(255) UNIQUE NOT NULL,
  `description` text,
  `flags` tinyint UNSIGNED,
  INDEX `idx_user_uuid`(`uuid`)
) ENGINE=INNODB;