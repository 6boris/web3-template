-- MySQL dump 10.13  Distrib 8.0.33, for macos13.3 (x86_64)
--
-- Host: 127.0.0.1    Database: web3
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `block`
--

use web3;
DROP TABLE IF EXISTS `block`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `block` (
  `chain_unique_id` bigint NOT NULL DEFAULT '0',
  `number` bigint NOT NULL DEFAULT '0' COMMENT 'block number',
  `chain_id` bigint NOT NULL DEFAULT '0',
  `hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `parent_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `timestamp` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Block Time',
  `difficulty` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `extra_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `gas_limit` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `gas_used` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `base_fee_per_gas` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `miner` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `mix_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `nonce` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `receipts_root` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `sha3_uncles` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `size` bigint NOT NULL DEFAULT '0',
  `state_root` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `transactions_root` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `transactions_count` bigint NOT NULL DEFAULT '0',
  `uncles_count` bigint NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time',
  PRIMARY KEY (`number`,`chain_unique_id`),
  KEY `block_idx_hash` (`hash`),
  KEY `block_idx_timestamp` (`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Block Info';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `block`
--

LOCK TABLES `block` WRITE;
/*!40000 ALTER TABLE `block` DISABLE KEYS */;
/*!40000 ALTER TABLE `block` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `block_transaction`
--

DROP TABLE IF EXISTS `block_transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `block_transaction` (
  `type` bigint NOT NULL DEFAULT '0',
  `status` bigint NOT NULL DEFAULT '0',
  `block_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `block_number` bigint NOT NULL DEFAULT '0' COMMENT 'block number',
  `block_timestamp` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Block Time',
  `transaction_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `transaction_index` bigint NOT NULL DEFAULT '0',
  `from_address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `to_address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `input` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `nonce` bigint NOT NULL DEFAULT '0',
  `contract_address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `gas` bigint NOT NULL DEFAULT '0',
  `gas_price` bigint NOT NULL DEFAULT '0',
  `gas_used` bigint NOT NULL DEFAULT '0',
  `effective_gas_price` bigint NOT NULL DEFAULT '0',
  `cumulative_gas_used` bigint NOT NULL DEFAULT '0',
  `max_fee_per_gas` bigint NOT NULL DEFAULT '0',
  `max_priority_fee_per_gas` bigint NOT NULL DEFAULT '0',
  `r` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `s` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `v` bigint NOT NULL DEFAULT '0',
  `logs_count` bigint NOT NULL DEFAULT '0',
  `chain_id` bigint NOT NULL DEFAULT '0',
  `chain_unique_id` bigint DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time',
  PRIMARY KEY (`transaction_hash`),
  KEY `block_idx_block_timestamp` (`block_timestamp`),
  KEY `transaction_block_timestamp_index` (`block_timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Transaction Info';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `block_transaction`
--

LOCK TABLES `block_transaction` WRITE;
/*!40000 ALTER TABLE `block_transaction` DISABLE KEYS */;
/*!40000 ALTER TABLE `block_transaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chain_client`
--

DROP TABLE IF EXISTS `chain_client`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chain_client` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `chain_unique_id` bigint NOT NULL DEFAULT '0',
  `chain_id` bigint NOT NULL DEFAULT '0',
  `provider_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0',
  `provider_id` bigint NOT NULL DEFAULT '0',
  `transport_schema` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `transport_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `priority` bigint NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Chain metrics';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chain_client`
--

LOCK TABLES `chain_client` WRITE;
/*!40000 ALTER TABLE `chain_client` DISABLE KEYS */;
/*!40000 ALTER TABLE `chain_client` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chain_info`
--

DROP TABLE IF EXISTS `chain_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chain_info` (
  `unique_id` bigint NOT NULL DEFAULT '0',
  `token_symbol` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `order_seq` bigint NOT NULL DEFAULT '0',
  `chain_id` bigint NOT NULL DEFAULT '0',
  `chain_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `chain_env` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `logo_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `website_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `explorer_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`unique_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chain_info`
--

LOCK TABLES `chain_info` WRITE;
/*!40000 ALTER TABLE `chain_info` DISABLE KEYS */;
INSERT INTO `chain_info` VALUES (1000100,'ETH',99,1,'Ethereum','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ethereum.jpg??w=48&h=48','https://ethereum.org/','https://etherscan.io/','2023-05-09 01:04:56','2023-06-23 13:49:36'),(1000101,'ETH',29,11155111,'Ethereum','Sepolia','https://icons.llamao.fi/icons/chains/rsz_ethereum.jpg??w=48&h=48','https://ethereum.org/','https://sepolia.etherscan.io/','2023-05-09 01:04:56','2023-06-23 14:03:26'),(1000102,'ETH',29,5,'Ethereum','Goerli','https://icons.llamao.fi/icons/chains/rsz_ethereum.jpg??w=48&h=48','https://ethereum.org/','https://goerli.etherscan.io/','2023-05-09 01:04:56','2023-06-23 14:03:26'),(1000200,'',0,0,'Tron','Mainnet','https://icons.llamao.fi/icons/chains/rsz_tron?w=48&h=48','','','2023-06-20 16:38:48','2023-06-20 16:47:07'),(1000201,'',0,0,'Tron','Testnet','https://icons.llamao.fi/icons/chains/rsz_tron?w=48&h=48','','','2023-06-20 16:38:48','2023-06-20 16:47:07'),(1000300,'BNB',97,56,'BSC','Mainnet','https://icons.llamao.fi/icons/chains/rsz_bsc?w=48&h=48','','','2023-06-20 16:38:49','2023-06-23 14:04:10'),(1000301,'BNB',29,97,'BSC','Testnet','https://icons.llamao.fi/icons/chains/rsz_bsc?w=48&h=48','','','2023-06-20 16:38:49','2023-06-23 14:04:10'),(1000400,'',10,42161,'Arbitrum','Mainnet','https://icons.llamao.fi/icons/chains/rsz_arbitrum?w=48&h=48','','','2023-06-20 16:38:49','2023-06-21 13:32:14'),(1000401,'',3,200,'Arbitrum','xDai','https://icons.llamao.fi/icons/chains/rsz_arbitrum?w=48&h=48','','','2023-06-20 16:38:49','2023-06-20 17:38:40'),(1000402,'',3,42170,'Arbitrum','Nova','https://icons.llamao.fi/icons/chains/rsz_arbitrum?w=48&h=48','','','2023-06-20 16:38:49','2023-06-20 17:38:40'),(1000403,'',3,421613,'Arbitrum','Goerli','https://icons.llamao.fi/icons/chains/rsz_arbitrum?w=48&h=48','','','2023-06-20 16:38:49','2023-06-20 17:38:40'),(1000500,'MATIC',98,137,'Polygon','Mainnet','https://icons.llamao.fi/icons/chains/rsz_polygon?w=48&h=48','','','2023-06-20 16:38:49','2023-06-23 14:03:54'),(1000501,'MATIC',3,80001,'Polygon','Mumbai','https://icons.llamao.fi/icons/chains/rsz_polygon?w=48&h=48','','','2023-06-20 16:38:49','2023-06-23 14:03:54'),(1000502,'',87,1442,'Polygon zkEVM','Mainnet','https://icons.llamao.fi/icons/chains/rsz_polygon%20zkevm?w=48&h=48','','','2023-06-20 16:38:49','2023-06-20 19:07:34'),(1000601,'',10,10,'Optimism','Mainnet','https://icons.llamao.fi/icons/chains/rsz_optimism?w=48&h=48','','','2023-06-20 16:38:50','2023-06-21 13:32:14'),(1000602,'',3,69,'Optimism','Kovan','https://icons.llamao.fi/icons/chains/rsz_optimism?w=48&h=48','','','2023-06-20 16:38:50','2023-06-20 17:38:40'),(1000603,'',3,420,'Optimism','Goerli','https://icons.llamao.fi/icons/chains/rsz_optimism?w=48&h=48','','','2023-06-20 16:38:50','2023-06-20 17:38:40'),(1000700,'AVAX',96,43114,'Avalanche','Mainnet','https://icons.llamao.fi/icons/chains/rsz_avalanche?w=48&h=48','','','2023-06-20 16:38:50','2023-06-23 14:04:23'),(1000701,'AVAX',25,43113,'Avalanche','Testnet','https://icons.llamao.fi/icons/chains/rsz_avalanche?w=48&h=48','','','2023-06-20 16:38:50','2023-06-23 14:04:23'),(1000800,'',90,73927,'Mixin','Mainnet','https://icons.llamao.fi/icons/chains/rsz_mixin?w=48&h=48','','','2023-06-20 16:38:50','2023-06-20 17:43:21'),(1000900,'',89,25,'Cronos','Mainnet','https://icons.llamao.fi/icons/chains/rsz_cronos?w=48&h=48','','','2023-06-20 16:38:50','2023-06-20 17:44:45'),(1001000,'',88,369,'Pulse','Mainnet','https://icons.llamao.fi/icons/chains/rsz_pulse?w=48&h=48','','','2023-06-20 16:38:51','2023-06-20 17:45:23'),(1001100,'',0,0,'Solana','Mainnet','https://icons.llamao.fi/icons/chains/rsz_solana?w=48&h=48','','','2023-06-20 16:38:51','2023-06-20 16:38:51'),(1001200,'',0,0,'DefiChain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_defichain?w=48&h=48','','','2023-06-20 16:38:51','2023-06-20 16:38:51'),(1001300,'',0,0,'Kava','Mainnet','https://icons.llamao.fi/icons/chains/rsz_kava?w=48&h=48','','','2023-06-20 16:38:52','2023-06-20 16:38:52'),(1001400,'FTM',95,250,'Fantom','Mainnet','https://icons.llamao.fi/icons/chains/rsz_fantom?w=48&h=48','','','2023-06-20 16:38:52','2023-06-23 14:04:48'),(1001401,'FTM',24,4002,'Fantom','Testnet','https://icons.llamao.fi/icons/chains/rsz_fantom?w=48&h=48','','','2023-06-20 16:38:52','2023-06-23 14:04:48'),(1001500,'',0,0,'Bitcoin','Mainnet','https://icons.llamao.fi/icons/chains/rsz_bitcoin?w=48&h=48','','','2023-06-20 16:38:52','2023-06-20 16:38:52'),(1001600,'',0,0,'Fusion','Mainnet','https://icons.llamao.fi/icons/chains/rsz_fusion?w=48&h=48','','','2023-06-20 16:38:52','2023-06-20 16:38:52'),(1001700,'ETH',1,324,'zkSync Era','Mainnet','https://icons.llamao.fi/icons/chains/rsz_zksync era?w=48&h=48','','','2023-06-20 16:38:53','2023-06-23 14:09:13'),(1001800,'',0,0,'Klaytn','Mainnet','https://icons.llamao.fi/icons/chains/rsz_klaytn?w=48&h=48','','','2023-06-20 16:38:53','2023-06-20 16:38:53'),(1001900,'',0,0,'Cardano','Mainnet','https://icons.llamao.fi/icons/chains/rsz_cardano?w=48&h=48','','','2023-06-20 16:38:53','2023-06-20 16:38:53'),(1002000,'',0,0,'Osmosis','Mainnet','https://icons.llamao.fi/icons/chains/rsz_osmosis?w=48&h=48','','','2023-06-20 16:38:54','2023-06-20 16:38:54'),(1002100,'',0,0,'Algorand','Mainnet','https://icons.llamao.fi/icons/chains/rsz_algorand?w=48&h=48','','','2023-06-20 16:38:54','2023-06-20 16:38:54'),(1002200,'',0,0,'Rootstock','Mainnet','https://icons.llamao.fi/icons/chains/rsz_rootstock?w=48&h=48','','','2023-06-20 16:38:54','2023-06-20 16:38:54'),(1002300,'',0,0,'Celo','Mainnet','https://icons.llamao.fi/icons/chains/rsz_celo?w=48&h=48','','','2023-06-20 16:38:54','2023-06-20 16:38:54'),(1002400,'',0,0,'Ronin','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ronin?w=48&h=48','','','2023-06-20 16:38:55','2023-06-20 16:38:55'),(1002500,'',85,100,'Gnosis','Mainnet','https://icons.llamao.fi/icons/chains/rsz_gnosis?w=48&h=48','','','2023-06-20 16:38:55','2023-06-20 19:09:40'),(1002600,'',0,0,'EOS','Mainnet','https://icons.llamao.fi/icons/chains/rsz_eos?w=48&h=48','','','2023-06-20 16:38:55','2023-06-20 16:38:55'),(1002700,'',0,0,'Canto','Mainnet','https://icons.llamao.fi/icons/chains/rsz_canto?w=48&h=48','','','2023-06-20 16:38:56','2023-06-20 16:38:56'),(1002800,'',0,0,'Waves','Mainnet','https://icons.llamao.fi/icons/chains/rsz_waves?w=48&h=48','','','2023-06-20 16:38:56','2023-06-20 16:38:56'),(1002900,'',0,0,'Tezos','Mainnet','https://icons.llamao.fi/icons/chains/rsz_tezos?w=48&h=48','','','2023-06-20 16:38:56','2023-06-20 16:38:56'),(1003000,'',0,0,'MultiversX','Mainnet','https://icons.llamao.fi/icons/chains/rsz_multiversx?w=48&h=48','','','2023-06-20 16:38:56','2023-06-20 16:38:56'),(1003100,'',0,0,'KCC','Mainnet','https://icons.llamao.fi/icons/chains/rsz_kcc?w=48&h=48','','','2023-06-20 16:38:57','2023-06-20 16:38:57'),(1003200,'',0,0,'WEMIX','Mainnet','https://icons.llamao.fi/icons/chains/rsz_wemix?w=48&h=48','','','2023-06-20 16:38:57','2023-06-20 16:38:57'),(1003300,'',0,0,'Moonbeam','Mainnet','https://icons.llamao.fi/icons/chains/rsz_moonbeam?w=48&h=48','','','2023-06-20 16:38:57','2023-06-20 16:38:57'),(1003400,'',0,0,'Thorchain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_thorchain?w=48&h=48','','','2023-06-20 16:38:58','2023-06-20 16:38:58'),(1003500,'',90,1,'Aptos','Mainnet','https://icons.llamao.fi/icons/chains/rsz_aptos?w=48&h=48','','','2023-06-20 16:38:58','2023-06-21 13:14:49'),(1003501,'',3,2,'Aptos','Testnet','https://icons.llamao.fi/icons/chains/rsz_aptos?w=48&h=48','','','2023-06-20 16:38:58','2023-06-21 13:14:50'),(1003600,'',0,0,'Metis','Mainnet','https://icons.llamao.fi/icons/chains/rsz_metis?w=48&h=48','','','2023-06-20 16:38:58','2023-06-20 16:38:58'),(1003700,'',0,0,'Near','Mainnet','https://icons.llamao.fi/icons/chains/rsz_near?w=48&h=48','','','2023-06-20 16:38:58','2023-06-20 16:38:58'),(1003800,'',0,0,'NEO','Mainnet','https://icons.llamao.fi/icons/chains/rsz_neo?w=48&h=48','','','2023-06-20 16:38:59','2023-06-20 16:38:59'),(1003900,'',0,0,'OKTChain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_oktchain?w=48&h=48','','','2023-06-20 16:38:59','2023-06-20 16:38:59'),(1004100,'',0,0,'Conflux','Mainnet','https://icons.llamao.fi/icons/chains/rsz_conflux?w=48&h=48','','','2023-06-20 16:39:00','2023-06-20 16:39:00'),(1004200,'',0,0,'Telos','Mainnet','https://icons.llamao.fi/icons/chains/rsz_telos?w=48&h=48','','','2023-06-20 16:39:00','2023-06-20 16:39:00'),(1004300,'',0,0,'Astar','Mainnet','https://icons.llamao.fi/icons/chains/rsz_astar?w=48&h=48','','','2023-06-20 16:39:00','2023-06-20 16:39:00'),(1004400,'',86,128,'Heco','Mainnet','https://icons.llamao.fi/icons/chains/rsz_heco?w=48&h=48','','','2023-06-20 16:39:00','2023-06-20 19:08:28'),(1004500,'',0,0,'Stacks','Mainnet','https://icons.llamao.fi/icons/chains/rsz_stacks?w=48&h=48','','','2023-06-20 16:39:01','2023-06-20 16:39:01'),(1004600,'',0,0,'Hedera','Mainnet','https://icons.llamao.fi/icons/chains/rsz_hedera?w=48&h=48','','','2023-06-20 16:39:01','2023-06-20 16:39:01'),(1004700,'',0,0,'Acala','Mainnet','https://icons.llamao.fi/icons/chains/rsz_acala?w=48&h=48','','','2023-06-20 16:39:01','2023-06-20 16:39:01'),(1004800,'',0,0,'Parallel','Mainnet','https://icons.llamao.fi/icons/chains/rsz_parallel?w=48&h=48','','','2023-06-20 16:39:02','2023-06-20 16:39:02'),(1004900,'',0,0,'Stellar','Mainnet','https://icons.llamao.fi/icons/chains/rsz_stellar?w=48&h=48','','','2023-06-20 16:39:02','2023-06-20 16:39:02'),(1005000,'',0,0,'Injective','Mainnet','https://icons.llamao.fi/icons/chains/rsz_injective?w=48&h=48','','','2023-06-20 16:39:02','2023-06-20 16:39:02'),(1005100,'',0,0,'IoTeX','Mainnet','https://icons.llamao.fi/icons/chains/rsz_iotex?w=48&h=48','','','2023-06-20 16:39:03','2023-06-20 16:39:03'),(1005200,'',0,0,'TON','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ton?w=48&h=48','','','2023-06-20 16:39:03','2023-06-20 16:39:03'),(1005300,'',0,0,'Secret','Mainnet','https://icons.llamao.fi/icons/chains/rsz_secret?w=48&h=48','','','2023-06-20 16:39:03','2023-06-20 16:39:03'),(1005400,'',0,0,'Aurora','Mainnet','https://icons.llamao.fi/icons/chains/rsz_aurora?w=48&h=48','','','2023-06-20 16:39:03','2023-06-20 16:39:03'),(1005500,'',0,0,'Neutron','Mainnet','https://icons.llamao.fi/icons/chains/rsz_neutron?w=48&h=48','','','2023-06-20 16:39:04','2023-06-20 16:39:04'),(1005600,'',0,0,'Ontology','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ontology?w=48&h=48','','','2023-06-20 16:39:04','2023-06-20 16:39:04'),(1005700,'',0,0,'Starknet','Mainnet','https://icons.llamao.fi/icons/chains/rsz_starknet?w=48&h=48','','','2023-06-20 16:39:04','2023-06-20 16:39:04'),(1005800,'',0,0,'Bittorrent','Mainnet','https://icons.llamao.fi/icons/chains/rsz_bittorrent?w=48&h=48','','','2023-06-20 16:39:05','2023-06-20 16:39:05'),(1005900,'',0,0,'Kujira','Mainnet','https://icons.llamao.fi/icons/chains/rsz_kujira?w=48&h=48','','','2023-06-20 16:39:05','2023-06-20 16:39:05'),(1006000,'',80,0,'Sui','Mainnet','https://icons.llamao.fi/icons/chains/rsz_sui?w=48&h=48','','','2023-06-20 16:39:05','2023-06-21 13:20:53'),(1006001,'',3,0,'Sui','Testnet','https://icons.llamao.fi/icons/chains/rsz_sui?w=48&h=48','','','2023-06-20 16:39:05','2023-06-21 13:20:53'),(1006100,'',0,0,'Icon','Mainnet','https://icons.llamao.fi/icons/chains/rsz_icon?w=48&h=48','','','2023-06-20 16:39:05','2023-06-20 16:39:05'),(1006200,'',0,0,'Moonriver','Mainnet','https://icons.llamao.fi/icons/chains/rsz_moonriver?w=48&h=48','','','2023-06-20 16:39:06','2023-06-20 16:39:06'),(1006300,'',0,0,'Vision','Mainnet','https://icons.llamao.fi/icons/chains/rsz_vision?w=48&h=48','','','2023-06-20 16:39:06','2023-06-20 16:39:06'),(1006400,'',0,0,'Umee','Mainnet','https://icons.llamao.fi/icons/chains/rsz_umee?w=48&h=48','','','2023-06-20 16:39:06','2023-06-20 16:39:06'),(1006500,'',0,0,'ThunderCore','Mainnet','https://icons.llamao.fi/icons/chains/rsz_thundercore?w=48&h=48','','','2023-06-20 16:39:07','2023-06-20 16:39:07'),(1006600,'',0,0,'Proton','Mainnet','https://icons.llamao.fi/icons/chains/rsz_proton?w=48&h=48','','','2023-06-20 16:39:07','2023-06-20 16:39:07'),(1006700,'',0,0,'CORE','Mainnet','https://icons.llamao.fi/icons/chains/rsz_core?w=48&h=48','','','2023-06-20 16:39:07','2023-06-20 16:39:07'),(1006800,'',0,0,'Terra2','Mainnet','https://icons.llamao.fi/icons/chains/rsz_terra2?w=48&h=48','','','2023-06-20 16:39:07','2023-06-20 16:39:07'),(1006900,'',0,0,'Karura','Mainnet','https://icons.llamao.fi/icons/chains/rsz_karura?w=48&h=48','','','2023-06-20 16:39:08','2023-06-20 16:39:08'),(1007000,'',0,0,'Wanchain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_wanchain?w=48&h=48','','','2023-06-20 16:39:08','2023-06-20 16:39:08'),(1007100,'',0,0,'Ultron','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ultron?w=48&h=48','','','2023-06-20 16:39:08','2023-06-20 16:39:08'),(1007200,'',0,0,'Oraichain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_oraichain?w=48&h=48','','','2023-06-20 16:39:09','2023-06-20 16:39:09'),(1007300,'',0,0,'Kardia','Mainnet','https://icons.llamao.fi/icons/chains/rsz_kardia?w=48&h=48','','','2023-06-20 16:39:09','2023-06-20 16:39:09'),(1007400,'',0,0,'DFK','Mainnet','https://icons.llamao.fi/icons/chains/rsz_dfk?w=48&h=48','','','2023-06-20 16:39:09','2023-06-20 16:39:09'),(1007500,'',0,0,'Vite','Mainnet','https://icons.llamao.fi/icons/chains/rsz_vite?w=48&h=48','','','2023-06-20 16:39:09','2023-06-20 16:39:09'),(1007600,'',0,0,'Arbitrum Nova','Mainnet','https://icons.llamao.fi/icons/chains/rsz_arbitrum nova?w=48&h=48','','','2023-06-20 16:39:10','2023-06-20 16:39:10'),(1007700,'',0,0,'Flow','Mainnet','https://icons.llamao.fi/icons/chains/rsz_flow?w=48&h=48','','','2023-06-20 16:39:10','2023-06-20 16:39:10'),(1007800,'',0,0,'Songbird','Mainnet','https://icons.llamao.fi/icons/chains/rsz_songbird?w=48&h=48','','','2023-06-20 16:39:10','2023-06-20 16:39:10'),(1007900,'',0,0,'Everscale','Mainnet','https://icons.llamao.fi/icons/chains/rsz_everscale?w=48&h=48','','','2023-06-20 16:39:11','2023-06-20 16:39:11'),(1008000,'',0,0,'Filecoin','Mainnet','https://icons.llamao.fi/icons/chains/rsz_filecoin?w=48&h=48','','','2023-06-20 16:39:11','2023-06-20 16:39:11'),(1008100,'',0,0,'Crescent','Mainnet','https://icons.llamao.fi/icons/chains/rsz_crescent?w=48&h=48','','','2023-06-20 16:39:11','2023-06-20 16:39:11'),(1008200,'',0,0,'Oasis','Mainnet','https://icons.llamao.fi/icons/chains/rsz_oasis?w=48&h=48','','','2023-06-20 16:39:11','2023-06-20 16:39:11'),(1008300,'',0,0,'Ergo','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ergo?w=48&h=48','','','2023-06-20 16:39:12','2023-06-20 16:39:12'),(1008400,'',0,0,'smartBCH','Mainnet','https://icons.llamao.fi/icons/chains/rsz_smartbch?w=48&h=48','','','2023-06-20 16:39:12','2023-06-20 16:39:12'),(1008500,'',0,0,'Nuls','Mainnet','https://icons.llamao.fi/icons/chains/rsz_nuls?w=48&h=48','','','2023-06-20 16:39:12','2023-06-20 16:39:12'),(1008600,'',0,0,'GodwokenV1','Mainnet','https://icons.llamao.fi/icons/chains/rsz_godwokenv1?w=48&h=48','','','2023-06-20 16:39:13','2023-06-20 16:39:13'),(1008700,'',0,0,'Zilliqa','Mainnet','https://icons.llamao.fi/icons/chains/rsz_zilliqa?w=48&h=48','','','2023-06-20 16:39:13','2023-06-20 16:39:13'),(1008800,'',0,0,'Interlay','Mainnet','https://icons.llamao.fi/icons/chains/rsz_interlay?w=48&h=48','','','2023-06-20 16:39:13','2023-06-20 16:39:13'),(1008900,'',0,0,'SXnetwork','Mainnet','https://icons.llamao.fi/icons/chains/rsz_sxnetwork?w=48&h=48','','','2023-06-20 16:39:13','2023-06-20 16:39:13'),(1009000,'',0,0,'Boba','Mainnet','https://icons.llamao.fi/icons/chains/rsz_boba?w=48&h=48','','','2023-06-20 16:39:14','2023-06-20 16:39:14'),(1009100,'',0,0,'Carbon','Mainnet','https://icons.llamao.fi/icons/chains/rsz_carbon?w=48&h=48','','','2023-06-20 16:39:14','2023-06-20 16:39:14'),(1009200,'',0,0,'Dogechain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_dogechain?w=48&h=48','','','2023-06-20 16:39:14','2023-06-20 16:39:14'),(1009300,'',0,0,'XDC','Mainnet','https://icons.llamao.fi/icons/chains/rsz_xdc?w=48&h=48','','','2023-06-20 16:39:15','2023-06-20 16:39:15'),(1009400,'',0,0,'Persistence','Mainnet','https://icons.llamao.fi/icons/chains/rsz_persistence?w=48&h=48','','','2023-06-20 16:39:15','2023-06-20 16:39:15'),(1009500,'',0,0,'Bifrost','Mainnet','https://icons.llamao.fi/icons/chains/rsz_bifrost?w=48&h=48','','','2023-06-20 16:39:15','2023-06-20 16:39:15'),(1009600,'',0,0,'Meter','Mainnet','https://icons.llamao.fi/icons/chains/rsz_meter?w=48&h=48','','','2023-06-20 16:39:15','2023-06-20 16:39:15'),(1009700,'',0,0,'Harmony','Mainnet','https://icons.llamao.fi/icons/chains/rsz_harmony?w=48&h=48','','','2023-06-20 16:39:16','2023-06-20 16:39:16'),(1009800,'',0,0,'Onus','Mainnet','https://icons.llamao.fi/icons/chains/rsz_onus?w=48&h=48','','','2023-06-20 16:39:16','2023-06-20 16:39:16'),(1009900,'',0,0,'Velas','Mainnet','https://icons.llamao.fi/icons/chains/rsz_velas?w=48&h=48','','','2023-06-20 16:39:16','2023-06-20 16:39:16'),(1010000,'',0,0,'FunctionX','Mainnet','https://icons.llamao.fi/icons/chains/rsz_functionx?w=48&h=48','','','2023-06-20 16:39:17','2023-06-20 16:39:17'),(1010100,'',0,0,'Libre','Mainnet','https://icons.llamao.fi/icons/chains/rsz_libre?w=48&h=48','','','2023-06-20 16:39:17','2023-06-20 16:39:17'),(1010200,'',0,0,'Oasys','Mainnet','https://icons.llamao.fi/icons/chains/rsz_oasys?w=48&h=48','','','2023-06-20 16:39:17','2023-06-20 16:39:17'),(1010300,'',0,0,'OntologyEVM','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ontologyevm?w=48&h=48','','','2023-06-20 16:39:18','2023-06-20 16:39:18'),(1010400,'',0,0,'Terra Classic','Mainnet','https://icons.llamao.fi/icons/chains/rsz_terra classic?w=48&h=48','','','2023-06-20 16:39:18','2023-06-20 16:39:18'),(1010500,'',0,0,'Zeniq','Mainnet','https://icons.llamao.fi/icons/chains/rsz_zeniq?w=48&h=48','','','2023-06-20 16:39:18','2023-06-20 16:39:18'),(1010600,'',0,0,'Bitindi','Mainnet','https://icons.llamao.fi/icons/chains/rsz_bitindi?w=48&h=48','','','2023-06-20 16:39:18','2023-06-20 16:39:18'),(1010700,'',0,0,'Evmos','Mainnet','https://icons.llamao.fi/icons/chains/rsz_evmos?w=48&h=48','','','2023-06-20 16:39:19','2023-06-20 16:39:19'),(1010800,'',0,0,'Heiko','Mainnet','https://icons.llamao.fi/icons/chains/rsz_heiko?w=48&h=48','','','2023-06-20 16:39:19','2023-06-20 16:39:19'),(1010900,'',0,0,'Sora','Mainnet','https://icons.llamao.fi/icons/chains/rsz_sora?w=48&h=48','','','2023-06-20 16:39:19','2023-06-20 16:39:19'),(1011000,'',0,0,'Milkomeda C1','Mainnet','https://icons.llamao.fi/icons/chains/rsz_milkomeda c1?w=48&h=48','','','2023-06-20 16:39:20','2023-06-20 16:39:20'),(1011100,'',0,0,'Juno','Mainnet','https://icons.llamao.fi/icons/chains/rsz_juno?w=48&h=48','','','2023-06-20 16:39:20','2023-06-20 16:39:20'),(1011200,'',0,0,'Grove','Mainnet','https://icons.llamao.fi/icons/chains/rsz_grove?w=48&h=48','','','2023-06-20 16:39:20','2023-06-20 16:39:20'),(1011300,'',0,0,'Litecoin','Mainnet','https://icons.llamao.fi/icons/chains/rsz_litecoin?w=48&h=48','','','2023-06-20 16:39:20','2023-06-20 16:39:20'),(1011400,'',0,0,'Hydra','Mainnet','https://icons.llamao.fi/icons/chains/rsz_hydra?w=48&h=48','','','2023-06-20 16:39:21','2023-06-20 16:39:21'),(1011500,'',0,0,'Comdex','Mainnet','https://icons.llamao.fi/icons/chains/rsz_comdex?w=48&h=48','','','2023-06-20 16:39:21','2023-06-20 16:39:21'),(1011600,'',0,0,'Elastos','Mainnet','https://icons.llamao.fi/icons/chains/rsz_elastos?w=48&h=48','','','2023-06-20 16:39:21','2023-06-20 16:39:21'),(1011700,'',0,0,'Fuse','Mainnet','https://icons.llamao.fi/icons/chains/rsz_fuse?w=48&h=48','','','2023-06-20 16:39:22','2023-06-20 16:39:22'),(1011800,'',0,0,'Obyte','Mainnet','https://icons.llamao.fi/icons/chains/rsz_obyte?w=48&h=48','','','2023-06-20 16:39:22','2023-06-20 16:39:22'),(1011900,'',0,0,'Nahmii','Mainnet','https://icons.llamao.fi/icons/chains/rsz_nahmii?w=48&h=48','','','2023-06-20 16:39:22','2023-06-20 16:39:22'),(1012000,'',0,0,'Theta','Mainnet','https://icons.llamao.fi/icons/chains/rsz_theta?w=48&h=48','','','2023-06-20 16:39:22','2023-06-20 16:39:22'),(1012100,'',0,0,'Step','Mainnet','https://icons.llamao.fi/icons/chains/rsz_step?w=48&h=48','','','2023-06-20 16:39:23','2023-06-20 16:39:23'),(1012200,'',0,0,'Map','Mainnet','https://icons.llamao.fi/icons/chains/rsz_map?w=48&h=48','','','2023-06-20 16:39:23','2023-06-20 16:39:23'),(1012300,'',0,0,'Equilibrium','Mainnet','https://icons.llamao.fi/icons/chains/rsz_equilibrium?w=48&h=48','','','2023-06-20 16:39:23','2023-06-20 16:39:23'),(1012400,'',0,0,'Syscoin','Mainnet','https://icons.llamao.fi/icons/chains/rsz_syscoin?w=48&h=48','','','2023-06-20 16:39:24','2023-06-20 16:39:24'),(1012500,'',0,0,'Kadena','Mainnet','https://icons.llamao.fi/icons/chains/rsz_kadena?w=48&h=48','','','2023-06-20 16:39:24','2023-06-20 16:39:24'),(1012600,'',0,0,'Bitgert','Mainnet','https://icons.llamao.fi/icons/chains/rsz_bitgert?w=48&h=48','','','2023-06-20 16:39:24','2023-06-20 16:39:24'),(1012700,'',0,0,'VeChain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_vechain?w=48&h=48','','','2023-06-20 16:39:24','2023-06-20 16:39:24'),(1012800,'',0,0,'Kintsugi','Mainnet','https://icons.llamao.fi/icons/chains/rsz_kintsugi?w=48&h=48','','','2023-06-20 16:39:25','2023-06-20 16:39:25'),(1012900,'',0,0,'Nova Network','Mainnet','https://icons.llamao.fi/icons/chains/rsz_nova network?w=48&h=48','','','2023-06-20 16:39:25','2023-06-20 16:39:25'),(1013000,'',0,0,'EOS EVM','Mainnet','https://icons.llamao.fi/icons/chains/rsz_eos evm?w=48&h=48','','','2023-06-20 16:39:25','2023-06-20 16:39:25'),(1013100,'',0,0,'Doge','Mainnet','https://icons.llamao.fi/icons/chains/rsz_doge?w=48&h=48','','','2023-06-20 16:39:26','2023-06-20 16:39:26'),(1013200,'',0,0,'Godwoken','Mainnet','https://icons.llamao.fi/icons/chains/rsz_godwoken?w=48&h=48','','','2023-06-20 16:39:26','2023-06-20 16:39:26'),(1013300,'',0,0,'Bitcoincash','Mainnet','https://icons.llamao.fi/icons/chains/rsz_bitcoincash?w=48&h=48','','','2023-06-20 16:39:26','2023-06-20 16:39:26'),(1013400,'',0,0,'Crab','Mainnet','https://icons.llamao.fi/icons/chains/rsz_crab?w=48&h=48','','','2023-06-20 16:39:26','2023-06-20 16:39:26'),(1013500,'',0,0,'Tombchain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_tombchain?w=48&h=48','','','2023-06-20 16:39:27','2023-06-20 16:39:27'),(1013600,'',0,0,'Starcoin','Mainnet','https://icons.llamao.fi/icons/chains/rsz_starcoin?w=48&h=48','','','2023-06-20 16:39:27','2023-06-20 16:39:27'),(1013700,'',0,0,'Callisto','Mainnet','https://icons.llamao.fi/icons/chains/rsz_callisto?w=48&h=48','','','2023-06-20 16:39:27','2023-06-20 16:39:27'),(1013800,'',0,0,'Europa','Mainnet','https://icons.llamao.fi/icons/chains/rsz_europa?w=48&h=48','','','2023-06-20 16:39:28','2023-06-20 16:39:28'),(1013900,'',0,0,'CosmosHub','Mainnet','https://icons.llamao.fi/icons/chains/rsz_cosmoshub?w=48&h=48','','','2023-06-20 16:39:28','2023-06-20 16:39:28'),(1014000,'',0,0,'Genshiro','Mainnet','https://icons.llamao.fi/icons/chains/rsz_genshiro?w=48&h=48','','','2023-06-20 16:39:28','2023-06-20 16:39:28'),(1014100,'',0,0,'Wax','Mainnet','https://icons.llamao.fi/icons/chains/rsz_wax?w=48&h=48','','','2023-06-20 16:39:28','2023-06-20 16:39:28'),(1014200,'',0,0,'EnergyWeb','Mainnet','https://icons.llamao.fi/icons/chains/rsz_energyweb?w=48&h=48','','','2023-06-20 16:39:29','2023-06-20 16:39:29'),(1014300,'',0,0,'Energi','Mainnet','https://icons.llamao.fi/icons/chains/rsz_energi?w=48&h=48','','','2023-06-20 16:39:29','2023-06-20 16:39:29'),(1014400,'',0,0,'Boba_Bnb','Mainnet','https://icons.llamao.fi/icons/chains/rsz_boba_bnb?w=48&h=48','','','2023-06-20 16:39:29','2023-06-20 16:39:29'),(1014500,'',0,0,'CSC','Mainnet','https://icons.llamao.fi/icons/chains/rsz_csc?w=48&h=48','','','2023-06-20 16:39:30','2023-06-20 16:39:30'),(1014600,'',0,0,'TomoChain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_tomochain?w=48&h=48','','','2023-06-20 16:39:30','2023-06-20 16:39:30'),(1014700,'',0,0,'ICP','Mainnet','https://icons.llamao.fi/icons/chains/rsz_icp?w=48&h=48','','','2023-06-20 16:39:30','2023-06-20 16:39:30'),(1014800,'',0,0,'ENULS','Mainnet','https://icons.llamao.fi/icons/chains/rsz_enuls?w=48&h=48','','','2023-06-20 16:39:30','2023-06-20 16:39:30'),(1014900,'',0,0,'EthereumPoW','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ethereumpow?w=48&h=48','','','2023-06-20 16:39:31','2023-06-20 16:39:31'),(1015000,'',0,0,'Findora','Mainnet','https://icons.llamao.fi/icons/chains/rsz_findora?w=48&h=48','','','2023-06-20 16:39:31','2023-06-20 16:39:31'),(1015100,'',0,0,'EthereumClassic','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ethereumclassic?w=48&h=48','','','2023-06-20 16:39:31','2023-06-20 16:39:31'),(1015200,'',0,0,'Loop','Mainnet','https://icons.llamao.fi/icons/chains/rsz_loop?w=48&h=48','','','2023-06-20 16:39:32','2023-06-20 16:39:32'),(1015300,'',0,0,'Ubiq','Mainnet','https://icons.llamao.fi/icons/chains/rsz_ubiq?w=48&h=48','','','2023-06-20 16:39:32','2023-06-20 16:39:32'),(1015400,'',0,0,'Empire','Mainnet','https://icons.llamao.fi/icons/chains/rsz_empire?w=48&h=48','','','2023-06-20 16:39:32','2023-06-20 16:39:32'),(1015500,'',0,0,'Rangers','Mainnet','https://icons.llamao.fi/icons/chains/rsz_rangers?w=48&h=48','','','2023-06-20 16:39:32','2023-06-20 16:39:32'),(1015600,'',0,0,'Cube','Mainnet','https://icons.llamao.fi/icons/chains/rsz_cube?w=48&h=48','','','2023-06-20 16:39:33','2023-06-20 16:39:33'),(1015700,'',0,0,'ZYX','Mainnet','https://icons.llamao.fi/icons/chains/rsz_zyx?w=48&h=48','','','2023-06-20 16:39:33','2023-06-20 16:39:33'),(1015800,'',0,0,'Milkomeda A1','Mainnet','https://icons.llamao.fi/icons/chains/rsz_milkomeda a1?w=48&h=48','','','2023-06-20 16:39:33','2023-06-20 16:39:33'),(1015900,'',0,0,'Shiden','Mainnet','https://icons.llamao.fi/icons/chains/rsz_shiden?w=48&h=48','','','2023-06-20 16:39:34','2023-06-20 16:39:34'),(1016000,'',0,0,'Lachain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_lachain?w=48&h=48','','','2023-06-20 16:39:34','2023-06-20 16:39:34'),(1016100,'',0,0,'Stargaze','Mainnet','https://icons.llamao.fi/icons/chains/rsz_stargaze?w=48&h=48','','','2023-06-20 16:39:34','2023-06-20 16:39:34'),(1016200,'',0,0,'Bone','Mainnet','https://icons.llamao.fi/icons/chains/rsz_bone?w=48&h=48','','','2023-06-20 16:39:34','2023-06-20 16:39:34'),(1016300,'',0,0,'Goerli','Mainnet','https://icons.llamao.fi/icons/chains/rsz_goerli?w=48&h=48','','','2023-06-20 16:39:35','2023-06-20 16:39:35'),(1016400,'',0,0,'REIchain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_reichain?w=48&h=48','','','2023-06-20 16:39:35','2023-06-20 16:39:35'),(1016500,'',0,0,'Flare','Mainnet','https://icons.llamao.fi/icons/chains/rsz_flare?w=48&h=48','','','2023-06-20 16:39:35','2023-06-20 16:39:35'),(1016600,'',0,0,'MultiVAC','Mainnet','https://icons.llamao.fi/icons/chains/rsz_multivac?w=48&h=48','','','2023-06-20 16:39:36','2023-06-20 16:39:36'),(1016700,'',0,0,'Boba_Avax','Mainnet','https://icons.llamao.fi/icons/chains/rsz_boba_avax?w=48&h=48','','','2023-06-20 16:39:36','2023-06-20 16:39:36'),(1016800,'',0,0,'Lamden','Mainnet','https://icons.llamao.fi/icons/chains/rsz_lamden?w=48&h=48','','','2023-06-20 16:39:36','2023-06-20 16:39:36'),(1016900,'',0,0,'HPB','Mainnet','https://icons.llamao.fi/icons/chains/rsz_hpb?w=48&h=48','','','2023-06-20 16:39:36','2023-06-20 16:39:36'),(1017000,'',0,0,'GoChain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_gochain?w=48&h=48','','','2023-06-20 16:39:37','2023-06-20 16:39:37'),(1017100,'',0,0,'Sifchain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_sifchain?w=48&h=48','','','2023-06-20 16:39:37','2023-06-20 16:39:37'),(1017200,'',0,0,'Polis','Mainnet','https://icons.llamao.fi/icons/chains/rsz_polis?w=48&h=48','','','2023-06-20 16:39:37','2023-06-20 16:39:37'),(1017300,'',0,0,'REI','Mainnet','https://icons.llamao.fi/icons/chains/rsz_rei?w=48&h=48','','','2023-06-20 16:39:37','2023-06-20 16:39:37'),(1017400,'',0,0,'MUUCHAIN','Mainnet','https://icons.llamao.fi/icons/chains/rsz_muuchain?w=48&h=48','','','2023-06-20 16:39:38','2023-06-20 16:39:38'),(1017500,'',0,0,'CLV','Mainnet','https://icons.llamao.fi/icons/chains/rsz_clv?w=48&h=48','','','2023-06-20 16:39:38','2023-06-20 16:39:38'),(1017600,'',0,0,'Lung','Mainnet','https://icons.llamao.fi/icons/chains/rsz_lung?w=48&h=48','','','2023-06-20 16:39:38','2023-06-20 16:39:38'),(1017700,'',0,0,'Palm','Mainnet','https://icons.llamao.fi/icons/chains/rsz_palm?w=48&h=48','','','2023-06-20 16:39:39','2023-06-20 16:39:39'),(1017800,'',0,0,'Kekchain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_kekchain?w=48&h=48','','','2023-06-20 16:39:39','2023-06-20 16:39:39'),(1017900,'',0,0,'Omax','Mainnet','https://icons.llamao.fi/icons/chains/rsz_omax?w=48&h=48','','','2023-06-20 16:39:39','2023-06-20 16:39:39'),(1018000,'',0,0,'zkSync Lite','Mainnet','https://icons.llamao.fi/icons/chains/rsz_zksync lite?w=48&h=48','','','2023-06-20 16:39:39','2023-06-20 16:39:39'),(1018100,'',0,0,'Polkadot','Mainnet','https://icons.llamao.fi/icons/chains/rsz_polkadot?w=48&h=48','','','2023-06-20 16:39:40','2023-06-20 16:39:40'),(1018200,'',0,0,'Kusama','Mainnet','https://icons.llamao.fi/icons/chains/rsz_kusama?w=48&h=48','','','2023-06-20 16:39:40','2023-06-20 16:39:40'),(1018300,'',0,0,'Dexit','Mainnet','https://icons.llamao.fi/icons/chains/rsz_dexit?w=48&h=48','','','2023-06-20 16:39:40','2023-06-20 16:39:40'),(1018400,'',0,0,'Stride','Mainnet','https://icons.llamao.fi/icons/chains/rsz_stride?w=48&h=48','','','2023-06-20 16:39:41','2023-06-20 16:39:41'),(1018500,'',0,0,'Stafi','Mainnet','https://icons.llamao.fi/icons/chains/rsz_stafi?w=48&h=48','','','2023-06-20 16:39:41','2023-06-20 16:39:41'),(1018600,'',0,0,'Migaloo','Mainnet','https://icons.llamao.fi/icons/chains/rsz_migaloo?w=48&h=48','','','2023-06-20 16:39:41','2023-06-20 16:39:41'),(1018700,'',0,0,'Quicksilver','Mainnet','https://icons.llamao.fi/icons/chains/rsz_quicksilver?w=48&h=48','','','2023-06-20 16:39:41','2023-06-20 16:39:41'),(1018800,'',0,0,'Hoo','Mainnet','https://icons.llamao.fi/icons/chains/rsz_hoo?w=48&h=48','','','2023-06-20 16:39:42','2023-06-20 16:39:42'),(1018900,'',0,0,'Pokt','Mainnet','https://icons.llamao.fi/icons/chains/rsz_pokt?w=48&h=48','','','2023-06-20 16:39:42','2023-06-20 16:39:42'),(1019000,'',0,0,'Echelon','Mainnet','https://icons.llamao.fi/icons/chains/rsz_echelon?w=48&h=48','','','2023-06-20 16:39:42','2023-06-20 16:39:42'),(1019100,'',0,0,'Coti','Mainnet','https://icons.llamao.fi/icons/chains/rsz_coti?w=48&h=48','','','2023-06-20 16:39:43','2023-06-20 16:39:43'),(1019200,'',0,0,'Tlchain','Mainnet','https://icons.llamao.fi/icons/chains/rsz_tlchain?w=48&h=48','','','2023-06-20 16:39:43','2023-06-20 16:39:43'),(1019300,'',0,0,'XPLA','Mainnet','https://icons.llamao.fi/icons/chains/rsz_xpla?w=48&h=48','','','2023-06-20 16:39:43','2023-06-20 16:39:43');
/*!40000 ALTER TABLE `chain_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chain_metrics`
--

DROP TABLE IF EXISTS `chain_metrics`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chain_metrics` (
  `chain_unique_id` bigint NOT NULL DEFAULT '0',
  `latest_block_number` bigint NOT NULL DEFAULT '0',
  `gas_price` bigint NOT NULL DEFAULT '0',
  `tx_tps` double NOT NULL DEFAULT '0',
  `tx_tpd` double NOT NULL DEFAULT '0',
  `tx_pending` bigint NOT NULL DEFAULT '0',
  `token_price_usd` double NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time',
  PRIMARY KEY (`chain_unique_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Chain metrics';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chain_metrics`
--

LOCK TABLES `chain_metrics` WRITE;
/*!40000 ALTER TABLE `chain_metrics` DISABLE KEYS */;
/*!40000 ALTER TABLE `chain_metrics` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint NOT NULL DEFAULT '0' COMMENT 'User ID',
  `nick_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'Nick Name',
  `avatar_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Avatar URL',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Email Address',
  `password` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time',
  PRIMARY KEY (`id`),
  KEY `user_idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User Data Table';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'Angelique Morse','https://api-dev-minimal-v5.vercel.app/assets/images/avatar/avatar_4.jpg','admin@web3.com','web3','2023-05-11 23:16:23','2023-05-11 23:16:23'),(17363249093103683,'fde5990e-755d-4800-bccd-edc9fd470f03','','','','2023-06-22 10:14:55','2023-06-22 10:14:55'),(524755722288019732,'Ariana Lang','https://api-dev-minimal-v5.vercel.app/assets/images/avatar/avatar_1.jpg','aditya_greenfelder31@gmail.com','web3','2023-05-11 23:21:30','2023-05-11 23:21:30'),(586819483167142030,'Aspen Schmitt','https://api-dev-minimal-v5.vercel.app/assets/images/avatar/avatar_2.jpg','zella_hickle4@yahoo.com','web3','2023-05-11 23:21:23','2023-05-11 23:21:23'),(894225568368303173,'92684121-bb6e-47ef-9dbd-9a63500050e6','','','','2023-06-22 10:14:57','2023-06-22 10:14:57'),(1151889042857807099,'Brycen Jimenez','https://api-dev-minimal-v5.vercel.app/assets/images/avatar/avatar_3.jpg','ashlynn_ohara62@gmail.com','web3','2023-05-11 23:20:54','2023-05-11 23:20:54'),(1269792913506349309,'Chase Day','https://api-dev-minimal-v5.vercel.app/assets/images/avatar/avatar_5.jpg','vergie_block82@hotmail.com','web3','2023-05-11 23:21:10','2023-05-11 23:21:10'),(1477442411851155581,'7fb7fea5-f9e4-48fc-8eff-69006f43834f','051ea566-4c4b-47f5-a660-afc8b05ff7dd','fbd048cc-88e3-4a3a-a20a-93edd22dce1e','','2023-06-22 10:19:28','2023-06-22 10:19:28'),(1683228382167360180,'Giana Brandt','https://api-dev-minimal-v5.vercel.app/assets/images/avatar/avatar_6.jpg','marjolaine_white94@gmail.com','web3','2023-05-11 23:21:26','2023-05-11 23:21:26'),(1795904323381936605,'c081ed60-1f8a-47a0-891f-e15e34efabd3','','','','2023-06-22 10:14:56','2023-06-22 10:14:56'),(2201400527445200681,'Deja Brady','https://api-dev-minimal-v5.vercel.app/assets/images/avatar/avatar_8.jpg','letha_lubowitz24@yahoo.com','web3','2023-05-11 23:21:32','2023-05-11 23:21:32');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-23 15:00:37

/*



*/
CREATE database dex DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
CREATE TABLE dex.user (
   `id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
   `nick_name` varchar(50) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '用户昵称',
   `real_name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '真实名称',
   `avatar_url` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像URL',
   `email` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱地址',
   `password` varchar(500) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
   `country` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '国家',
   `company` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '公司',
   `address` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '地址',
   `state` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '街道',
   `city` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '城市',
   `zip_code` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '邮编',
   `about` varchar(200) COLLATE utf8mb4_general_ci DEFAULT '自我介绍',
   `phone_number` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
   `phone_area_code` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机区域号',
   `roles` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '角色',
   `status` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '状态：0，启用；1，禁用；',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY (`id`),
   KEY `user_idx_email` (`email`),
   KEY `user_idx_phone` (`phone_area_code`,`phone_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表'

CREATE TABLE `account` (
   `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
   `user_id` bigint NOT NULL COMMENT '用户ID',
   `coin_id` bigint NOT NULL COMMENT '币种ID',
   `status` tinyint(1) NOT NULL COMMENT '账号状态：1，正常；2，冻结；',
   `balance_amount` decimal(40,20) NOT NULL COMMENT '币种可用金额',
   `freeze_amount` decimal(40,20) NOT NULL COMMENT '币种冻结金额',
   `recharge_amount` decimal(40,20) NOT NULL COMMENT '累计充值金额',
   `withdrawals_amount` decimal(40,20) NOT NULL COMMENT '累计提现金额',
   `rec_addr` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '充值地址',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY (`id`) USING BTREE,
   UNIQUE KEY `userid_coinid_unique` (`user_id`,`coin_id`) USING BTREE,
   KEY `account_coin_id_ref` (`coin_id`) USING BTREE,
   KEY `inx_platform_account` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户资产记录';