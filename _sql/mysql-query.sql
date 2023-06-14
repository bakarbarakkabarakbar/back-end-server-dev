CREATE SCHEMA `miniproject` DEFAULT CHARACTER SET utf8mb4 ;

USE miniproject;
CREATE TABLE actors(
	`id` BIGINT UNSIGNED,
    `username` VARCHAR(50),
    `password` VARCHAR(50),
    `role_id` INT UNSIGNED,
    `is_verified` ENUM('true','false'),
    `is_active` ENUM('true','false'),
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `modified_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `actorsPK` PRIMARY KEY (`id`),
    CONSTRAINT `role_idFK` FOREIGN KEY (`role_id`) REFERENCES actor_roles(`id`)
);
DROP TABLE actors;

CREATE TABLE customers(
	`id` BIGINT UNSIGNED,
    `first_name` VARCHAR(50),
    `last_name` VARCHAR(50),
    `email` VARCHAR(50),
    `avatar` VARCHAR(200),
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `modified_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `customersPK` PRIMARY KEY (`id`)
);
DROP TABLE customers;

CREATE TABLE actor_roles(
	`id` INT UNSIGNED,
    `role_name` VARCHAR(50),
    CONSTRAINT `actor_rolesPK` PRIMARY KEY (`id`)
);
DROP TABLE actor_roles;

CREATE TABLE register_approvals(
	`id` INT UNSIGNED,
    `admin_id` BIGINT UNSIGNED,
    `super_admin_id` BIGINT UNSIGNED,
	`status` VARCHAR(50),
    CONSTRAINT `register_approvalsPK` PRIMARY KEY (`id`)
);
DROP TABLE register_approvals;

CREATE TABLE users (
   `id` INT(11) UNSIGNED PRIMARY KEY,
   `name` VARCHAR(255) NOT NULL,
   `email` VARCHAR(255) NOT NULL,
   `password` VARCHAR(255) NOT NULL,
   `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE actor_sessions (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL,
  `token` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `expires_at` TIMESTAMP,
  CONSTRAINT `register_approvalsPK` PRIMARY KEY (`id`));
DROP TABLE actor_sessions;

INSERT INTO actor_roles(`id`,`role_name`) VALUES(1, 'super-admin'),(2, 'admin'), (3, 'customer');
INSERT INTO actors(`id`,`username`, `password`, `role_id`, `is_verified`, `is_active`) VALUES (1, "super-admin", "7fbe1732f8b44c15b88f0c1e4fe94fcd0c60ccec", 1, true, true);
SELECT * FROM customers WHERE CONCAT(first_name, ' ', last_name) LIKE "%lana%";
