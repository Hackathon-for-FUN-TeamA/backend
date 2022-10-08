-- CREATE DATABASE
DROP DATABASE IF EXISTS `asayake`;
CREATE DATABASE `asayake`;

-- map指定
USE `asayake`;

-- CREATE TABLE IN user
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `token` varchar(64) NOT NULL,
  `username` varchar(32) NOT NULL,
  `password` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token` (`token`)
);

-- CREATE TABLE IN drivelog
DROP TABLE IF EXISTS `drivelogs`;
CREATE TABLE `drivelogs` (
    `log_id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL,
    `date` varchar(32) NOT NULL,
    `speed` real NOT NULL,
    `acceleration` real NOT NULL,
    `latitude` real NOT NULL,
    `longtitude` real NOT NULL,
    PRIMARY KEY (`log_id`)
);
