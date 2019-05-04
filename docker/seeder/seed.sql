SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

USE `go-hr`;

SET NAMES utf8mb4;

CREATE TABLE `companies` (
                             `id` varchar(36) NOT NULL,
                             `name` varchar(100) NOT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `companies` (`id`, `name`) VALUES
('a0b8dc73-67e9-11e9-91f4-0242ac120002',	'go');

CREATE TABLE `employees` (
                             `id` varchar(36) NOT NULL,
                             `name` varchar(100) NOT NULL,
                             `company_id` varchar(36) NOT NULL,
                             `role` enum('admin','normal') NOT NULL,
                             `employee_no` varchar(10) NOT NULL,
                             `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                             `failed_login_count` int(11) NOT NULL DEFAULT '0',
                             `failed_login_time` datetime DEFAULT NULL,
                             `jail_time` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             KEY `company_id` (`company_id`),
                             CONSTRAINT `employees_ibfk_1` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `employees` (`id`, `name`, `company_id`, `role`, `employee_no`, `password`, `failed_login_count`, `failed_login_time`, `jail_time`) VALUES
('badbc49e-67e9-11e9-91f4-0242ac120002',	'azby',	'a0b8dc73-67e9-11e9-91f4-0242ac120002',	'normal',	'10001',	'ZDMyMjQ4ZWQwNzI3YTJlMmYzMGViNjEwNDliNmFhNzI1MWI5ODVhY2E2NjE1ZTE2MGNlNzI5MDBkMjllNGY0Ny52YXpvNW9xZXM5PQ==',	0,	NULL,	NULL),
('f5c1eb13-697e-11e9-a31d-0242ac120003',	'admin-hr',	'a0b8dc73-67e9-11e9-91f4-0242ac120002',	'admin',	'10000',	'YmQzMzBmZGJkMzM3ZTk1OGU3NWNmMWM1M2E0ZDU4ZTNkYWI2NzRiMDQ0NDhlMzVkZmQ2MDdiNTMxYWUwNTQzNC5sMzF3Y25tbWR5PQ==',	0,	NULL,	NULL);

CREATE TABLE `notices` (
                           `id` varchar(36) NOT NULL,
                           `employee_id` varchar(36) NOT NULL,
                           `type` enum('sick','remote','vacation') NOT NULL,
                           `visibility` enum('public','private') NOT NULL,
                           `period_start` date NOT NULL,
                           `period_end` date NOT NULL,
                           KEY `employee_id` (`employee_id`),
                           CONSTRAINT `notices_ibfk_2` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `notices` (`id`, `employee_id`, `type`, `visibility`, `period_start`, `period_end`) VALUES
('f13b775f-6987-11e9-a31d-0242ac120003',	'badbc49e-67e9-11e9-91f4-0242ac120002',	'vacation',	'public',	'2019-04-28',	'2019-04-28'),
('d624769d-6989-11e9-a31d-0242ac120003',	'badbc49e-67e9-11e9-91f4-0242ac120002',	'remote',	'public',	'2019-04-25',	'2019-04-27'),
('f941bd51-6989-11e9-a31d-0242ac120003',	'badbc49e-67e9-11e9-91f4-0242ac120002',	'vacation',	'public',	'2019-04-29',	'2019-04-30'),
('37e434c0-698a-11e9-a31d-0242ac120003',	'badbc49e-67e9-11e9-91f4-0242ac120002',	'vacation',	'private',	'2019-04-24',	'2019-04-24'),
('bcada520-698c-11e9-a31d-0242ac120003',	'f5c1eb13-697e-11e9-a31d-0242ac120003',	'sick',	'private',	'2019-04-20',	'2019-04-20');