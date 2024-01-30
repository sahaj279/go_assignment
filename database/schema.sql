--
-- Table structure for table `item`
--

DROP TABLE IF EXISTS `items`;
CREATE TABLE `items` (
  `Name` varchar(255) DEFAULT NULL,
  `Price` decimal(9,2) DEFAULT NULL,
  `Quantity` int DEFAULT NULL,
  `Type` enum('Raw','Manufactured','Imported') DEFAULT NULL,
  `ID` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;