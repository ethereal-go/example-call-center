
-- +migrate Up
CREATE TABLE `attachment_types` (
	`id` int NOT NULL AUTO_INCREMENT,
	`name` varchar(255) NOT NULL,
	`text` TEXT NOT NULL,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
);


-- +migrate Down
DROP TABLE attachment_types;