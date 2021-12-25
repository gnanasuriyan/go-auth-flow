-- +goose Up
CREATE TABLE IF NOT EXISTS `user`(
     `id` INT NOT NULL AUTO_INCREMENT,
     `email` VARCHAR(255) NOT NULL,
     `password` VARCHAR(255) NOT NULL,
     `active` BOOLEAN NOT NULL DEFAULT TRUE,
     `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     `update_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY(`id`),
     UNIQUE KEY `user_email_active`(`email`, `active`)
);
