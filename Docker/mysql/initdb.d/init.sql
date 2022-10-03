-- CREATE DATABASE
DROP DATABASE IF EXISTS `map`;
CREATE DATABASE `map`;

-- map指定
USE `map`;

-- CREATE TABLE IN map
DROP TABLE IF EXISTS `gps`;
CREATE TABLE `gps` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `x` real NOT NULL,
  `y` real NOT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `route`;
CREATE TABLE `route` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `left_x` real NOT NULL,
  `left_y` real NOT NULL,
  `right_x` real NOT NULL,
  `right_y` real NOT NULL,
    `passed_count` int DEFAULT NULL,
  PRIMARY KEY (`id`)
);
