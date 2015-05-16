-- MySQL dump 10.13  Distrib 5.6.17, for osx10.9 (x86_64)
--
-- Host: localhost    Database: dad_tree
-- ------------------------------------------------------
-- Server version	5.6.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `members`
--

DROP TABLE IF EXISTS `members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `members` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `member_type_id` int(11) NOT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `lft` int(11) DEFAULT NULL,
  `rght` int(11) DEFAULT NULL,
  `first_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `last_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1304 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `members`
--

LOCK TABLES `members` WRITE;
/*!40000 ALTER TABLE `members` DISABLE KEYS */;
INSERT INTO `members` VALUES (1000,1,1000,41,223,'Marc','Schneider'),(1100,2,1000,41,41,'Marc','Schneider'),(1101,1,1000,111,126,'Hein','Ludick'),(1102,1,1000,2,3,'Julian','Du Plessis'),(1103,1,1000,4,5,'Dale','Swanepoel'),(1104,1,1000,6,11,'Dwayn','Smith'),(1105,1,1000,12,13,'Wayne','Laubscher'),(1106,1,1000,14,44,'Sean','De Beer'),(1116,1,1000,45,46,'Kay','Cheytanov'),(1118,2,1101,112,113,'bruce','van der west huizen'),(1124,1,1101,114,115,'Ferdinand ','Groenewald'),(1125,1,1101,116,117,'Sasha','Janse van Vuuren'),(1126,1,1101,118,119,'Lucan','Drummond'),(1128,1,1000,47,48,'Noxolo','Skomolo'),(1129,1,1000,49,50,'Clarissa','Britz'),(1130,2,1000,51,52,'ayabonga','booi'),(1132,1,1000,53,54,'Cindie','Scheepers'),(1133,1,1101,120,121,'Jolene','Thompson'),(1135,2,1000,55,56,'Lusanda',' Yose'),(1137,2,1000,57,58,'nicolene','buitendag'),(1138,1,1104,7,8,'Thabang','Pataka'),(1142,1,1000,59,60,'Val','Haysom'),(1145,2,1000,61,62,'Marc','Schneider'),(1146,1,1000,63,64,'GJ','Wentzel'),(1150,1,1000,65,70,'Stefan','Schneider'),(1154,1,1000,71,72,'Anelda','Jonker'),(1156,1,1106,15,18,'Zibele','Xhayimpi'),(1157,1,1156,16,17,'Patrick','Lethulwe'),(1159,1,1106,19,20,'Brian','De Beer'),(1160,1,1106,21,22,'Zukile','Gwala'),(1162,1,1000,73,74,'Seyi ','Olamide'),(1163,1,1000,75,76,'Nazeema','Van rooi'),(1165,2,1000,77,78,'malizole','mkosi'),(1167,1,1150,66,67,'yvette','herrmann'),(1168,2,1000,79,80,'Jerry','Makgale'),(1170,1,1000,81,82,'Luzuko','Ngqakayi'),(1173,1,1150,68,69,'Nomsa','Nhlapo'),(1174,2,1000,83,84,'Akhona Dibbsy ','Dibela '),(1175,1,1101,122,125,'Nadia','Ludick'),(1176,1,1000,85,86,'Emmanuel Themba','Mkhwanazi'),(1177,2,1000,87,88,'Lindokuhle','Nkambule'),(1178,1,1000,89,90,'xolisile','king'),(1180,1,1000,91,92,'Simon','Mashala'),(1181,2,1000,93,94,'Fransie','van der Hoven'),(1182,2,1000,95,96,'Mary','Modupo'),(1183,1,1106,23,24,'Junaid','Bayat'),(1184,2,1000,97,98,'Bianca','Heunes'),(1185,1,1000,99,100,'Denzil','Klaasen'),(1186,1,1106,25,26,'Darren','Lurie'),(1187,1,1106,27,28,'Tammy','Potgieter'),(1188,1,1106,29,30,'Shumani Shylock','Maziya'),(1189,2,1000,101,102,'Phillicity','Cosmo'),(1190,1,1000,103,104,'Chucks','Moses'),(1191,1,1000,105,106,'John Charles','Murray'),(1192,1,1106,41,41,'Louis','Mokonyane '),(1194,2,1106,39,40,'Babalwa Gloria',' Mdladlamba'),(1196,2,1106,41,41,'Fredricka (Rika)','Pretoruis'),(1199,2,1106,41,41,'Kwame','Attafuah'),(1205,1,1000,41,41,'GJ','Wentzel'),(1208,2,1142,41,41,'Motlalepule','Baloyi'),(1230,1,1000,107,108,'Jacob Ntsholeng','Segale'),(1233,2,1000,109,110,'stella','nel'),(1238,2,1000,127,128,'olwethu','mhlungu'),(1239,2,1000,129,130,'Carmelitha',' Rogers'),(1240,2,1000,131,132,'rochelle','nell'),(1241,2,1104,9,10,'BOIKHUTSO PORTIA','KOLOI'),(1242,2,1000,133,134,'thokozile','ndlovu'),(1244,2,1000,135,136,'Annika','Saunders'),(1245,2,1000,137,138,'Winile','Conny'),(1247,2,1000,139,140,'Gugu','Nkosi'),(1248,2,1000,141,142,'Ramunenyiwa','Thiambiwi'),(1252,2,1000,143,144,'refilwe','nerriea'),(1254,2,1000,145,146,'Belinda','Noxabiso'),(1255,1,1000,147,152,'Foundation of Leadership and Development',''),(1258,2,1106,42,43,'Evidence','Moyo'),(1259,2,1000,153,154,'nokwanda','nokuthula'),(1261,2,1000,155,156,'Botshelo','Moreki'),(1263,1,1000,157,158,'khumoetsile','morubudi'),(1264,2,1000,159,160,'Joy','Petrus'),(1265,2,1000,161,162,'agreement ','mopai'),(1267,2,1000,163,164,'Christian','Kazadi'),(1268,2,1000,165,166,'Adedayo','Adeboyejo'),(1269,2,1255,148,149,'Tadios','Munyimani'),(1271,2,1000,167,168,'Caroline','Witbooi'),(1274,2,1000,169,170,'Rose','Kekana'),(1275,2,1000,171,172,'Amy','Laurence'),(1276,2,1255,150,151,'Dimakatso Joyce','Motaung'),(1277,2,1000,173,174,'paradise ','skosana '),(1278,2,1000,175,176,'Amanda','Russon'),(1279,2,1000,177,178,'Vuyelwa','Chikosi'),(1280,2,1000,179,180,'Amanda','Russon'),(1281,2,1000,181,182,'Mohamed Abukar','Hassan'),(1282,2,1000,183,184,'Danisile','Nkosi'),(1283,1,1175,123,124,'Michel','Ranck'),(1284,2,1000,185,186,'Nomfundiso','Lehlakane'),(1285,2,1000,187,188,'Noluvo','Zembe'),(1286,1,1000,189,190,'Percy','Phillipson'),(1287,2,1000,191,192,'Sibonisile ','Sithole'),(1288,1,1000,193,194,'Siduduzo','Ntshingila'),(1289,1,1000,195,196,'Vessel','Mukhawana'),(1290,1,1000,197,198,'Mathe','Molusi'),(1292,2,1000,199,200,'Zwelikhanyile','Sikhotha'),(1293,2,1000,201,202,'Landi','Stoltz'),(1294,2,1000,203,204,'mandy ayanda','ncube'),(1295,1,1000,205,206,'Absolom','Sithole'),(1296,2,1000,207,208,'Tebelelo','Phasha'),(1297,2,1000,209,210,'cowan','carlo'),(1298,2,1000,211,212,'Thandeka','Moureen'),(1299,2,1000,213,214,'Robert','makopo'),(1300,2,1000,215,216,'Titus','rathaga'),(1301,2,1000,217,218,'moshe','modise'),(1302,2,1000,219,220,'Ruth','Momotshi'),(1303,2,1000,221,222,'Mellisa ','Creswell');
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2015-05-16 17:16:48
