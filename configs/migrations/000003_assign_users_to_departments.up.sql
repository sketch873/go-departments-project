CREATE TABLE `users_in_departments` (
  `id` bigint UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `department_id` bigint UNSIGNED NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
  FOREIGN KEY (`department_id`) REFERENCES `departments`(`id`),
  CONSTRAINT `unique_user_department` UNIQUE (user_id, department_id)
) ENGINE=INNODB;