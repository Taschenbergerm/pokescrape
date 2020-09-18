-- pokescraper.pokemon_references definition
CREATE DATABASE IF NOT EXISTS pokescraper; 

USE pokescraper; 

CREATE TABLE IF NOT EXISTS `pokemon_references` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `pokemon_references_UN` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=152 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `pokemon` (
`id` int PRIMARY KEY,
`Name` varchar(50) NOT NULL,
`BaseHappiness` int ,
`CaptureRate` int,
`Color` varchar(50) ,
`EvolvesFrom` varchar(50) ,
`GenderRate` int ,
`Generation` varchar(50) ,
`GrowthRate` varchar(50) ,
`HasGenderDifference` int(4) ,
`HatchCounter` int
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;