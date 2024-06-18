-- MariaDB dump 10.19  Distrib 10.6.16-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: crowdfunding
-- ------------------------------------------------------
-- Server version	8.4.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `campaign_images`
--

DROP TABLE IF EXISTS `campaign_images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `campaign_images` (
  `id` int NOT NULL AUTO_INCREMENT,
  `campaign_id` int NOT NULL,
  `file_name` varchar(255) NOT NULL,
  `is_primary` tinyint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fkCampaignImagesCampaigns` (`campaign_id`),
  CONSTRAINT `fkCampaignImagesCampaigns` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `campaign_images`
--

LOCK TABLES `campaign_images` WRITE;
/*!40000 ALTER TABLE `campaign_images` DISABLE KEYS */;
INSERT INTO `campaign_images` VALUES (1,1,'images/1-jacob-pretorius-BKm7pliIO-s-unsplash.jpg',0,'2024-06-02 16:50:38','2024-06-02 16:53:11'),(2,1,'images/1-dao-hi-u-rMr58Ek_cZM-unsplash.jpg',0,'2024-06-02 16:51:25','2024-06-02 16:53:11'),(3,1,'images/1-paul-esch-laurent-oZMUrWFHOB4-unsplash.jpg',1,'2024-06-02 16:53:11','2024-06-02 16:53:11'),(4,2,'images/2-mediamodifier-jDldhJmvOe8-unsplash.jpg',0,'2024-06-04 12:01:19','2024-06-04 12:01:35'),(5,2,'images/2-sebastian-bednarek-Lxcn2wrM6UY-unsplash.jpg',0,'2024-06-04 12:01:27','2024-06-04 12:01:35'),(6,2,'images/2-syful-islam-f-h3nbAi45E-unsplash.jpg',1,'2024-06-04 12:01:35','2024-06-04 12:01:35'),(7,3,'images/2-sven-mieke-7NTc2bMDHHg-unsplash.jpg',0,'2024-06-04 12:10:40','2024-06-04 12:10:56'),(8,3,'images/2-abel-y-costa-jb-SMviXCjI-unsplash.jpg',0,'2024-06-04 12:10:49','2024-06-04 12:10:56'),(9,3,'images/2-john-mark-arnold-ti4kGLkGgmU-unsplash.jpg',1,'2024-06-04 12:10:56','2024-06-04 12:10:56');
/*!40000 ALTER TABLE `campaign_images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `campaigns`
--

DROP TABLE IF EXISTS `campaigns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `campaigns` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `short_description` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `perks` text NOT NULL,
  `backer_count` int NOT NULL,
  `goal_amount` int NOT NULL,
  `current_amount` int NOT NULL,
  `slug` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fkCampaignsUsers` (`user_id`),
  CONSTRAINT `fkCampaignsUsers` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `campaigns`
--

LOCK TABLES `campaigns` WRITE;
/*!40000 ALTER TABLE `campaigns` DISABLE KEYS */;
INSERT INTO `campaigns` VALUES (1,1,'Aestetic keyboard','Aesthetic keyboard is a keyboard that can enhance the beauty and productivity of your work.','Keyboard estetik adalah keyboard yang memiliki desain yang indah dan menarik. Keyboard ini dapat menjadi pelengkap yang sempurna untuk meja kerja Anda, dan dapat meningkatkan produktivitas dan kreativitas Anda.','keuntungan satu,kemudian dua,dan tiga',1,1000000,50000,'testing-campaign-1','2023-08-12 14:48:35','2023-09-02 15:30:22'),(2,2,'Custom Ergonomic Mouse','For Precision,Control and Unleash Your Creativity','Elevate your efficiency with a mouse crafted for flawless command. Experience pinpoint precision, seamless navigation, and a design that empowers you to take charge of your digital workday.Ignite your creative spark with a mouse that becomes an extension of your imagination. Intuitive control and fluid motion let you bring your visions to life on the screen with unmatched ease.','Command Your Workflow,recision at Your Fingertips,Efficiency in the Palm of Your Hand',1,1200000,500000,'custom-ergonomic-mouse-2','2024-02-15 14:54:12','2024-02-15 18:24:57'),(3,2,'The Desk Den','Level Up At Your Desk','Are you ready to conquer your daily tasks and achieve your goals? Join our desk campaign and transform your workspace into a productivity powerhouse! Get ready to optimize your organization, streamline your workflow, and ignite your focus.Upgrade your workday with our desk campaign! Discover the power of an organized and ergonomic workspace. Expect increased productivity, reduced stress, and a surge of creativity. It\'s time to unleash your potential!','Improved Productivity,Better Physical Comfort,Increased Efficiency,Healthier Workforce',3,3000000,700000,'the-desk-den-2','2024-02-15 15:05:18','2024-02-15 18:32:52');
/*!40000 ALTER TABLE `campaigns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `campaign_id` int NOT NULL,
  `user_id` int NOT NULL,
  `amount` int NOT NULL,
  `status` varchar(255) NOT NULL,
  `payment_url` varchar(255) DEFAULT NULL,
  `code` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fkTransactionsUsers` (`user_id`),
  KEY `fkTransactionsCampaigns` (`campaign_id`),
  CONSTRAINT `fkTransactionsCampaigns` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `fkTransactionsUsers` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (1,1,6,50000,'paid','https://app.sandbox.midtrans.com/snap/v3/redirection/b2623027-906c-46b6-88c1-82a841d6f238','','2023-09-02 15:30:21','2023-09-02 15:30:22'),(25,2,3,500000,'paid','random.xyz/123','','2024-02-15 18:24:57','2024-02-15 18:24:57'),(26,3,3,100000,'paid','random.xyz/123','','2024-02-15 18:25:30','2024-02-15 18:25:30'),(27,3,3,100000,'paid','random.xyz/123','','2024-02-15 18:28:32','2024-02-15 18:28:33'),(28,3,3,500000,'paid','random.xyz/123','','2024-02-15 18:32:52','2024-02-15 18:32:52');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `occupation` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `avatar_file_name` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'John Doe','postman tes','admin@gmail.com','$2a$04$IGJbRAqiKLfNkPR3q/uiiOcQxsreOsy5Mi9YzUfso/9SEPNzU7Fzq','images/1-cesar-rincon-XHVpWcr5grQ-unsplash.jpg','admin','2024-01-25 10:46:46','2024-06-04 13:48:25'),(2,'Anne Marine','Designer','anne.marine@gmail.com','$2a$04$7AuBpuvPyf3ExPY4cackj.r9Tj3vLRW7Zq987x2743ZnGP4WDRe4.','images/2-6c6e516b-949c-4ac2-bb0a-7fec49c65eef.jpeg','user','2024-02-15 14:46:07','2024-06-04 12:59:14'),(3,'Donatur','donatur','donatur@gmail.com','$2a$04$OBOKNl3scHtuXE2U0XPD6eP1bffXZWlakQfRS8OaFmvqP6X0QiT1m','images/3-oladimeji-odunsi-n5aE6hOY6do-unsplash.jpg','user','2024-02-15 15:10:32','2024-06-04 14:18:19');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'crowdfunding'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-16 14:58:05
