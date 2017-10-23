
-- +migrate Up


CREATE TABLE `calls` (
	`id` int NOT NULL AUTO_INCREMENT,
	`from_id` int NOT NULL,
	`to_id` int NOT NULL,
	`outgoing_incoming` varchar(255) NOT NULL,
	`coordinates_incoming` varchar(255) NOT NULL,
	`duration` INT NOT NULL,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`)
);

ALTER TABLE `calls` ADD CONSTRAINT `calls_fk0` FOREIGN KEY (`from_id`) REFERENCES `users`(`id`);
ALTER TABLE `calls` ADD CONSTRAINT `calls_fk1` FOREIGN KEY (`to_id`) REFERENCES `users`(`id`);

-- +migrate Down
DROP TABLE calls;