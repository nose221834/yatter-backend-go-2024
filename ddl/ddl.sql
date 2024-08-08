CREATE TABLE IF NOT EXISTS `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE,
  `password_hash` varchar(255) NOT NULL,
  `display_name` varchar(255),
  `avatar` text,
  `header` text,
  `note` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `status` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `AccountID` INT NOT NULL,
  `URL` VARCHAR(255),
  `Content` TEXT NOT NULL,
  `CreatedAt` DATETIME,
  PRIMARY KEY (`ID`)
);
