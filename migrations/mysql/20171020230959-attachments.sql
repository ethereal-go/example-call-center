
-- +migrate Up

CREATE TABLE `attachments` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`type_id` INT NOT NULL,
	`link` varchar(255) NOT NULL,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
);



ALTER TABLE `attachments` ADD CONSTRAINT `attachment_fk0` FOREIGN KEY (`type_id`) REFERENCES `attachment_types`(`id`);

-- +migrate Down
DROP TABLE attachments;
