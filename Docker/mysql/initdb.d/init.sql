-- CREATE DATABASE
DROP DATABASE IF EXISTS `asayake`;
CREATE DATABASE `asayake`;

-- map指定
USE `asayake`;

-- CREATE TABLE IN user
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL,
  `password` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
);
