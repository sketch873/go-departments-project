CREATE TABLE `users` (
  `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) UNIQUE NOT NULL,
  `username` varchar(255) NOT NULL UNIQUE,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `uuid` varchar(255) UNIQUE NOT NULL,
  INDEX `idx_user_uuid`(`uuid`),
  INDEX `idx_user_email`(`email`)
) ENGINE=INNODB;