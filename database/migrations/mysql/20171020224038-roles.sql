
-- +migrate Up

CREATE TABLE `roles` (
	`id` int NOT NULL AUTO_INCREMENT,
	`name` varchar(255) NOT NULL,
	`display_name` varchar(255) NOT NULL,
	`description` varchar(255) NOT NULL,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
);


-- +migrate Down
DROP TABLE roles;