CREATE TABLE IF NOT EXISTS `pokemon` (
`id` int(11) NOT NULL,
`Name` varchar(50) NOT NULL,
`BaseHappiness` int ,
`CaptureRate` int,
`Color` varchar(50) ,
`EvolvesFrom` varchar(50) ,
`GenderRate` int ,
`Generation` varchar(50) ,
`GrowthRate` varchar(50) ,
`HasGenderDifference` int(4) ,
`HatchCounter` int, 
PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;