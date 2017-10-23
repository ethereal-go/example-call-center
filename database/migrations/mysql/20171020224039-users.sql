
-- +migrate Up
CREATE TABLE `users` (
	`id` int NOT NULL AUTO_INCREMENT,
	`email` varchar(255) NOT NULL UNIQUE,
	`name` varchar(255) NOT NULL,
	`phone` varchar(255) NOT NULL,
	`password` varchar(255) NOT NULL,
	`token` varchar(255) NOT NULL,
	`role_id` int NULL DEFAULT NULL,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
);
ALTER TABLE `users` ADD CONSTRAINT `users_fk0` FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`);
-- +migrate Down
DROP TABLE users;