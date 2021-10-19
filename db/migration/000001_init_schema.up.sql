CREATE TABLE `target_account` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `atm_bank_code` varchar(10) DEFAULT NULL,
  `bank_detail` bigint(20) DEFAULT NULL,
  `account_type` varchar(100) NOT NULL,
  `account_number` varchar(34) DEFAULT NULL,
  `bank` varchar(40) NOT NULL,
  `bank_branch` varchar(30) DEFAULT NULL,
  `description` varchar(100) DEFAULT NULL,
  `target_type` bigint(20) NOT NULL,
  `status` bigint(20) NOT NULL,
  `fourth_digit` varchar(1) DEFAULT "0",
  `customer_id` bigint(20) NOT NULL,
  `currency` varchar(10) DEFAULT NULL,
  `account_type_code` varchar(6) DEFAULT NULL,
  `amount` decimal(20,4) DEFAULT NULL,
  `is_favorite` varchar(10) DEFAULT NULL
);

CREATE INDEX `target_account_index_0` ON `target_account` (`bank_detail`);

CREATE INDEX `target_account_index_1` ON `target_account` (`target_type`);

CREATE INDEX `target_account_index_2` ON `target_account` (`status`);

CREATE INDEX `target_account_index_3` ON `target_account` (`customer_id`);
