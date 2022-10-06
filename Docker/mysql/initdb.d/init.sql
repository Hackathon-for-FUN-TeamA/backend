-- CREATE DATABASE
DROP DATABASE IF EXISTS `asayake`;
CREATE DATABASE `asayake`;

-- map指定
USE `asayake`;

-- CREATE TABLE IN user
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(5) NOT NULL,
  `password` varchar(5) NOT NULL,
  PRIMARY KEY (`id`)
);
