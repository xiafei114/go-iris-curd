/*
 Navicat Premium Data Transfer

 Source Server         : casbin
 Source Server Type    : MariaDB
 Source Server Version : 100214
 Source Host           : localhost:3306
 Source Schema         : casbin

 Target Server Type    : MariaDB
 Target Server Version : 100214
 File Encoding         : 65001

 Date: 22/04/2019 14:09:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for demo
-- ----------------------------
DROP TABLE IF EXISTS `demo`;
CREATE TABLE `demo`  (
  `pid` int(10) NOT NULL AUTO_INCREMENT,
  `product_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `product_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `number` int(11) NULL DEFAULT NULL,
  `create_date` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`pid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
