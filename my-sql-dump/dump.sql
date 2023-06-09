CREATE DATABASE  IF NOT EXISTS `miniproject` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `miniproject`;

CREATE USER 'golang-service-account' IDENTIFIED BY 'STRONG.password79';
GRANT ALL PRIVILEGES ON miniproject.* TO 'golang-service-account';

-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: 34.224.99.112    Database: miniproject
-- ------------------------------------------------------
-- Server version	8.0.33-0ubuntu0.20.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `actor_roles`
--

DROP TABLE IF EXISTS `actor_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actor_roles` (
                               `id` int unsigned NOT NULL,
                               `role_name` varchar(50) DEFAULT NULL,
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actor_roles`
--

LOCK TABLES `actor_roles` WRITE;
/*!40000 ALTER TABLE `actor_roles` DISABLE KEYS */;
INSERT INTO `actor_roles` VALUES (1,'super-admin'),(2,'admin'),(3,'customer');
/*!40000 ALTER TABLE `actor_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `actor_sessions`
--

DROP TABLE IF EXISTS `actor_sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actor_sessions` (
                                  `id` int unsigned NOT NULL AUTO_INCREMENT,
                                  `user_id` int unsigned NOT NULL,
                                  `token` varchar(255) NOT NULL,
                                  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  `expires_at` timestamp NULL DEFAULT NULL,
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actor_sessions`
--

LOCK TABLES `actor_sessions` WRITE;
/*!40000 ALTER TABLE `actor_sessions` DISABLE KEYS */;
INSERT INTO `actor_sessions` VALUES (1,4,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYwNzIwODksImlhdCI6MTY4NjA2ODQ4OSwiaXNzIjoiTWluaVByb2plY3QiLCJuYmYiOjE2ODYwNjg0ODksInN1YiI6IkJhY2tlbmREZXYiLCJ1c2VybmFtZSI6ImFkbWluIn0.3MXHS1i6ME5UTsh36trv0RkPPNUhoN7ragwq3dHh3cs','2023-06-06 16:21:30','2023-06-06 17:21:30'),(2,1,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYwNzIzNTMsImlhdCI6MTY4NjA2ODc1MywiaXNzIjoiTWluaVByb2plY3QiLCJuYmYiOjE2ODYwNjg3NTMsInN1YiI6IkJhY2tlbmREZXYiLCJ1c2VybmFtZSI6InN1cGVyLWFkbWluIn0.Ojm65tRjzjq97bCRMjlTOH2cYaQVwffdou8NCqoWLjo','2023-06-06 16:25:54','2023-06-06 17:25:54'),(3,1,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYyODYwNjUsImlhdCI6MTY4NjI4MjQ2NSwiaXNzIjoiTWluaVByb2plY3QiLCJuYmYiOjE2ODYyODI0NjUsInN1YiI6IkJhY2tlbmREZXYiLCJ1c2VybmFtZSI6InN1cGVyLWFkbWluIn0.uELo17Kb9Orc90GF8vF0g0DckkKclXVbcjd7J0w4B1I','2023-06-09 03:47:45','2023-06-09 04:47:45');
/*!40000 ALTER TABLE `actor_sessions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `actors`
--

DROP TABLE IF EXISTS `actors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actors` (
                          `id` bigint unsigned NOT NULL,
                          `username` varchar(50) DEFAULT NULL,
                          `password` varchar(50) DEFAULT NULL,
                          `role_id` int unsigned DEFAULT NULL,
                          `is_verified` enum('true','false') DEFAULT NULL,
                          `is_active` enum('true','false') DEFAULT NULL,
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `modified_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`),
                          KEY `role_idFK` (`role_id`),
                          CONSTRAINT `role_idFK` FOREIGN KEY (`role_id`) REFERENCES `actor_roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actors`
--

LOCK TABLES `actors` WRITE;
/*!40000 ALTER TABLE `actors` DISABLE KEYS */;
INSERT INTO `actors` VALUES (1,'super-admin','7fbe1732f8b44c15b88f0c1e4fe94fcd0c60ccec',1,'true','true','2023-06-05 09:57:31','2023-06-05 10:21:59'),(3,'akbar','d033e22ae348aeb5660fc2140aec35850c4da997',2,'false','true','2023-06-05 10:30:25','2023-06-06 11:28:31'),(4,'admin','d033e22ae348aeb5660fc2140aec35850c4da997',2,'false','false','2023-06-06 11:14:06','2023-06-06 11:14:38'),(5,'akbar','d033e22ae348aeb5660fc2140aec35850c4da997',2,'false','false','2023-06-06 11:21:51','2023-06-06 11:21:51');
/*!40000 ALTER TABLE `actors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
                             `id` bigint unsigned NOT NULL,
                             `first_name` varchar(50) DEFAULT NULL,
                             `last_name` varchar(50) DEFAULT NULL,
                             `email` varchar(50) DEFAULT NULL,
                             `avatar` varchar(200) DEFAULT NULL,
                             `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                             `modified_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (1,'akbar','maulana','akbar@example.com','katara','2023-06-05 10:24:10','2023-06-05 10:24:10'),(2,'annisa','fadhilah','annisa@example.com','kara','2023-06-05 10:24:29','2023-06-06 11:21:35'),(3,'annisa','fadhila','annisa@example.com','zuko','2023-06-05 14:09:56','2023-06-05 14:09:56'),(7,'Michael','Lawson','michael.lawson@reqres.in','https://reqres.in/img/faces/7-image.jpg','2023-06-06 11:20:40','2023-06-06 11:20:40'),(8,'Lindsay','Ferguson','lindsay.ferguson@reqres.in','https://reqres.in/img/faces/8-image.jpg','2023-06-05 14:38:43','2023-06-05 14:38:43'),(9,'Tobias','Funke','tobias.funke@reqres.in','https://reqres.in/img/faces/9-image.jpg','2023-06-05 14:38:44','2023-06-05 14:38:44'),(10,'Byron','Fields','byron.fields@reqres.in','https://reqres.in/img/faces/10-image.jpg','2023-06-05 14:38:45','2023-06-05 14:38:45'),(11,'George','Edwards','george.edwards@reqres.in','https://reqres.in/img/faces/11-image.jpg','2023-06-05 14:38:46','2023-06-05 14:38:46'),(12,'Rachel','Howell','rachel.howell@reqres.in','https://reqres.in/img/faces/12-image.jpg','2023-06-05 14:38:47','2023-06-05 14:38:47');
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `register_approvals`
--

DROP TABLE IF EXISTS `register_approvals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `register_approvals` (
                                      `id` int unsigned NOT NULL,
                                      `admin_id` bigint unsigned DEFAULT NULL,
                                      `super_admin_id` bigint unsigned DEFAULT NULL,
                                      `status` varchar(50) DEFAULT NULL,
                                      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `register_approvals`
--

LOCK TABLES `register_approvals` WRITE;
/*!40000 ALTER TABLE `register_approvals` DISABLE KEYS */;
INSERT INTO `register_approvals` VALUES (1,2,1,'approved');
/*!40000 ALTER TABLE `register_approvals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'miniproject'
--

--
-- Dumping routines for database 'miniproject'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-09 14:10:26
