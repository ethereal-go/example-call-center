
-- +migrate Up
CREATE TABLE `autoresponders` (
	`id` int NOT NULL AUTO_INCREMENT,
	`call_id` int NOT NULL,
	`attachment_id` int NOT NULL,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
);


ALTER TABLE `autoresponders` ADD CONSTRAINT `autoresponder_fk0` FOREIGN KEY (`call_id`) REFERENCES `calls`(`id`);

ALTER TABLE `autoresponders` ADD CONSTRAINT `autoresponder_fk1` FOREIGN KEY (`attachment_id`) REFERENCES `attachments`(`id`);

-- +migrate Down
DROP TABLE autoresponders;