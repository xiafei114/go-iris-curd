/*
MariaDB Backup
Database: casbin
Backup Time: 2019-04-15 12:37:13
*/

SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `casbin`.`casbin_rule`;
DROP TABLE IF EXISTS `casbin`.`demo`;
DROP TABLE IF EXISTS `casbin`.`dep`;
DROP TABLE IF EXISTS `casbin`.`dep_user`;
DROP TABLE IF EXISTS `casbin`.`menu`;
DROP TABLE IF EXISTS `casbin`.`post`;
DROP TABLE IF EXISTS `casbin`.`product`;
DROP TABLE IF EXISTS `casbin`.`product_category`;
DROP TABLE IF EXISTS `casbin`.`role_menu`;
DROP TABLE IF EXISTS `casbin`.`user`;
CREATE TABLE `casbin_rule` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_casbin_rule_p_type` (`p_type`),
  KEY `IDX_casbin_rule_v2` (`v2`),
  KEY `IDX_casbin_rule_v3` (`v3`),
  KEY `IDX_casbin_rule_v4` (`v4`),
  KEY `IDX_casbin_rule_v5` (`v5`),
  KEY `IDX_casbin_rule_v0` (`v0`),
  KEY `IDX_casbin_rule_v1` (`v1`)
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8;
CREATE TABLE `demo` (
  `pid` int(10) NOT NULL AUTO_INCREMENT,
  `product_code` varchar(255) DEFAULT '',
  `product_name` varchar(255) DEFAULT '',
  `number` int(11) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
CREATE TABLE `dep` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dep_name` varchar(64) DEFAULT '' COMMENT '部门名称',
  `leader` varchar(64) NOT NULL COMMENT '部门领导人uid',
  `phone` varchar(64) DEFAULT '' COMMENT '部门电话',
  `email` varchar(64) DEFAULT '' COMMENT '部门邮箱',
  `disabled` tinyint(1) DEFAULT NULL COMMENT '0:可用 否则:不可用',
  `parent_id` int(10) DEFAULT NULL,
  `dep_desc` varchar(255) DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;
CREATE TABLE `dep_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dep_ip` int(10) DEFAULT NULL COMMENT '部门id',
  `uid` int(10) DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `path` varchar(64) DEFAULT '',
  `name` varchar(64) DEFAULT '' COMMENT '必须是唯一的',
  `modular` varchar(64) DEFAULT '',
  `component` varchar(64) DEFAULT '',
  `meta` text DEFAULT NULL,
  `parent_id` int(10) DEFAULT NULL,
  `is_sub` tinyint(1) NOT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8;
CREATE TABLE `post` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `code` varchar(64) DEFAULT '' COMMENT '职位编码',
  `name` varchar(64) DEFAULT '' COMMENT '职位名称',
  `enable` tinyint(1) DEFAULT NULL COMMENT '是否启用',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `product` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `product_code` varchar(255) DEFAULT '',
  `product_name` varchar(255) DEFAULT '',
  `price` int(11) DEFAULT NULL,
  `number` int(11) DEFAULT NULL,
  `created_At` datetime DEFAULT NULL,
  `updated_At` datetime DEFAULT NULL,
  `deleted_At` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
CREATE TABLE `product_category` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `num_code` varchar(50) DEFAULT '',
  `chn_name` varchar(200) DEFAULT '',
  `created_At` datetime DEFAULT NULL,
  `updated_At` datetime DEFAULT NULL,
  `deleted_At` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
CREATE TABLE `role_menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `rid` int(10) NOT NULL COMMENT '角色id',
  `mid` int(10) NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;
CREATE TABLE `user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `gender` varchar(2) DEFAULT '' COMMENT '性别',
  `enable` tinyint(1) DEFAULT 0 COMMENT '0：启用用户 1：禁用用户',
  `name` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `userface` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=127 DEFAULT CHARSET=utf8;
BEGIN;
LOCK TABLES `casbin`.`casbin_rule` WRITE;
DELETE FROM `casbin`.`casbin_rule`;
INSERT INTO `casbin`.`casbin_rule` (`id`,`p_type`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`create_time`) VALUES (65, 'p', '90', '', '/*', 'ANY', '.*', '超级用户', NULL),(66, 'g', '90', 'admin', 'a', '', '', NULL, NULL),(67, 'g', '90', 'demo', 'a', '', '', NULL, NULL),(68, 'p', 'admin', 'a', '/admin*', 'GET|POST|DELETE|PUT', '.*', NULL, NULL),(69, 'p', 'demo', 'a', '/demo*', 'GET|POST|DELETE|PUT', '.*', NULL, NULL),(71, 'p', 't1', 'a', '/*', 'PUT|DELETE|GET|POST', '.*', NULL, NULL),(72, 'p', 'user', 'a', '/*', 'PUT|DELETE|GET|POST', '.*', NULL, NULL),(73, 'p', 'components', 'a', '/components*', 'GET|POST|DELETE|PUT', '.*', NULL, NULL),(74, 'g', '90', 'components', 'a', '', '', NULL, NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `casbin`.`demo` WRITE;
DELETE FROM `casbin`.`demo`;
INSERT INTO `casbin`.`demo` (`pid`,`product_code`,`product_name`,`number`,`create_date`) VALUES (1, 'PD-001', 'asc', 2, '2019-01-08 10:30:33'),(2, 'PD-002', 'asc', 2, '2019-01-08 10:38:21'),(3, 'PD-003', 'asc', 2, '2019-01-08 10:38:51'),(4, 'PD-005', 'asc', 2, '2019-01-08 10:39:33'),(5, 'PD-006', 'asc', 2, '2019-01-08 10:44:21');
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `casbin`.`dep` WRITE;
DELETE FROM `casbin`.`dep`;
INSERT INTO `casbin`.`dep` (`id`,`dep_name`,`leader`,`phone`,`email`,`disabled`,`parent_id`,`dep_desc`,`create_time`,`update_time`) VALUES (1, '稷下之学', '荀子', '17830895300', '614143260@qq.com', 0, 0, '', '2019-03-27 15:19:39', '2019-04-24 17:13:59'),(2, '儒家', '孔子', '1000', 'laozi@163.com', 0, 1, '崇尚\"礼乐\"和\"仁义\"，主张\"德治\"和\"仁政\"，教育了很多人为人处世所遵循的准则规范。', '2019-04-04 17:13:43', '2019-04-08 15:07:03'),(3, '道家', '老子', '1001', 'laozi@163.com', 0, 1, '天道运行的原理。天道轮回各种哲学上的思想都能在道家找到答案。世界上很多发生的事情都有其规律，人力不可更改。', '2019-04-30 17:13:45', '2019-04-24 17:13:59'),(4, '法家', '李斯', '1002', 'lisi@163.com', 0, 1, '依法治国，天子犯法与庶民同罪。法律视为一种有利于社会统治的强制性工具，体现法制建设的思想，一直被沿用至今。', '2019-04-16 17:13:49', '2019-04-24 17:13:59'),(5, '墨家', '墨子', '1003', 'mozi@163.com', 0, 1, '兼爱非攻，反对强大的王国去攻击弱小的王国，在思想上尊天事鬼，一切以保持社会现状的稳定为主。', '2019-04-10 17:13:52', '2019-04-24 17:13:59'),(6, '名家', '公孙龙', '1004', 'gongsl@163.com', 0, 1, '以辩论出名，著名的白马非马也是名家的思想。以逻辑推理来证明事物的好与坏、真实与否。', '2019-04-09 17:13:55', '2019-04-24 17:13:59'),(7, '阴阳家', '邹衍', '1005', 'zhouyan@163.com', 0, 1, '五行学说，从天文和地理方面来判断事物的凶吉。', '2019-04-24 17:13:59', '2019-04-24 17:13:59'),(8, '纵横家', '苏秦|张仪', '1006', 'sz@163.com', 0, 1, '合纵连横，捭阖之术，阴阳之变化也。', '2019-04-24 17:13:59', '2019-04-24 17:13:59'),(9, '农家', '许行', '1007', 'xuxing@163.com', 0, 1, '注重生产，研究植物生长和产出的学派。', '2019-04-24 17:13:59', '2019-04-24 17:13:59'),(10, '兵家', '孙膑', '1008', 'sunbing@163.com', 0, 1, '讲究利用武力，最大化的夺取敌方的利益从而赢得战争的胜利。', '2019-04-24 17:13:59', '2019-04-24 17:13:59'),(11, '医家', '扁鹊', '1009', 'bianque@163.com', 0, 1, '医者仁心', '2019-04-24 17:13:59', '2019-04-24 17:13:59'),(12, '礼乐', '孟子', '1010', 'mengzi@163.com', 0, 2, '', '2019-04-24 17:13:59', '2019-04-24 17:13:59'),(13, '武当', '张三丰', '1011', 'zsf@163.com', 0, 3, '', '2019-04-24 17:13:59', '2019-04-24 17:13:59'),(14, '庄子游', '庄子', '1012', 'zz@163.com', 0, 3, '', '2019-04-24 17:13:59', '2019-04-24 17:13:59'),(24, 'test', '', '', '', 1, 5, '', '2019-04-08 14:53:05', NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `casbin`.`dep_user` WRITE;
DELETE FROM `casbin`.`dep_user`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `casbin`.`menu` WRITE;
DELETE FROM `casbin`.`menu`;
INSERT INTO `casbin`.`menu` (`id`,`path`,`name`,`modular`,`component`,`meta`,`parent_id`,`is_sub`,`create_time`,`update_time`) VALUES (1, '/', '所有', '', NULL, NULL, NULL, 0, NULL, NULL),(2, 'sys_mgr', 'sys_mgr', NULL, NULL, '{\"title\":\"系统管理\", \"icon\":\"ios-settings\"}', 1, 0, NULL, NULL),(3, 'user_mgr', 'user_mgr', '/admin/users', '/users', '{\"title\":\"用户管理\", \"icon\":\"ios-settings\"}', 2, 1, NULL, NULL),(4, 'role_mgr', 'role_mgr', '/admin', '/role', '{\"title\":\"角色管理\", \"icon\":\"ios-settings\"}', 2, 1, NULL, NULL),(5, 'menu_mgr', 'menu_mgr', '/admin', '/menu', '{\"title\":\"菜单管理\", \"icon\":\"ios-settings\"}', 2, 1, NULL, NULL),(6, 'dep_mgr', 'dep_mgr', '/admin/dep', '/dep', '{\"title\":\"部门管理\", \"icon\":\"ios-settings\"}', 2, 1, NULL, NULL),(7, 'demo_mgr', 'demo_mgr', NULL, NULL, '{\"title\":\"demo管理\", \"icon\":\"ios-settings\"}', 1, 0, NULL, NULL),(8, 'product_mgr', 'product_mgr', '/demo/product', '/product', '{\"title\":\"产品管理\", \"icon\":\"ios-settings\"}', 7, 1, NULL, NULL),(9, 'product_cate_mgr', 'product_cate_mgr', '/demo/product_cate', '/product_cate', '{\"title\":\"产品类别\", \"icon\":\"ios-settings\"}', 7, 1, '2018-07-25 14:00:27', NULL),(10, '/demo3', 'demo3', '/demo', '/demo3', NULL, 7, 1, '2018-07-25 14:14:05', NULL),(11, 'doc', 'doc', '', '', '{\"title\":\"文档\", \"href\":\"https://lison16.github.io/iview-admin-doc/#/\", \"icon\":\"ios-book\"}', 1, 0, NULL, NULL),(12, 'join', 'join', '', '', '{\"hideInBread\": true}', 1, 0, NULL, NULL),(13, 'join_page', 'join_page', '', '/join_page', '{\"icon\":\"_qq\", \"title\":\"QQ群\"}', 12, 1, NULL, NULL),(14, 'message', 'message', '', '', '{\"hideInBread\":true, \"hideInMenu\":true}', 1, 0, NULL, NULL),(15, 'message_page', 'message_page', '/single-page/message', '/index', '{\"icon\":\"md-notifications\", \"title\":\"消息中心\"}', 14, 1, NULL, NULL),(16, 'components', 'components', NULL, NULL, '{\"icon\":\"logo-buffer\", \"title\":\"组件\"}', 1, 0, NULL, NULL),(17, 'tree_select_page', 'tree_select_page', '/components/tree-select', '/index', '{\"icon\":\"md-arrow-dropdown-circle\", \"title\":\"树状下拉选择器\"}', 16, 1, NULL, NULL),(18, 'count_to_page', 'count_to_page', '/components/count-to', '/count-to', '{\"icon\":\"md-trending-up\", \"title\":\"数字渐变\"}', 16, 1, NULL, NULL),(19, 'drag_list_page', 'drag_list_page', '/components/drag-list', '/drag-list', '{\"icon\":\"ios-infinite\", \"title\":\"拖拽列表\"}', 16, 1, NULL, NULL),(20, 'drag_drawer_page', 'drag_drawer_page', '/components/drag-drawer', '/index', '{\"icon\":\"md-list\", \"title\":\"可拖拽抽屉\"}', 16, 1, NULL, NULL),(21, 'org_tree_page', 'org_tree_page', '/components/org-tree', '/index', '{\"icon\":\"ios-people\", \"title\":\"组织结构树\"}', 16, 1, NULL, NULL),(22, 'tree_table_page', 'tree_table_page', '/components/tree-table', '/index', '{\"icon\":\"md-git-branch\", \"title\":\"树状表格\"}', 16, 1, NULL, NULL),(23, 'cropper_page', 'cropper_page', '/components/cropper', '/cropper', '{\"icon\":\"md-crop\", \"title\":\"图片裁剪\"}', 16, 1, NULL, NULL),(24, 'tables_page', 'tables_page', '/components/tables', '/tables', '{\"icon\":\"md-grid\", \"title\":\"多功能表格\"}', 16, 1, NULL, NULL),(25, 'split_pane_page', 'split_pane_page', '/components/split-pane', '/split-pane', '{\"icon\":\"md-pause\", \"title\":\"分割窗口\"}', 16, 1, NULL, NULL),(26, 'markdown_page', 'markdown_page', '/components/markdown', '/markdown', '{\"icon\":\"logo-markdown\", \"title\":\"Markdown编辑器\"}', 16, 1, NULL, NULL),(27, 'editor_page', 'editor_page', '/components/editor', '/editor', '{\"icon\":\"ios-create\", \"title\":\"富文本编辑器\"}', 16, 1, NULL, NULL),(28, 'icons_page', 'icons_page', '/components/icons', '/icons', '{\"icon\":\"_bear\", \"title\":\"自定义图标\"}', 16, 1, NULL, NULL),(29, 'update', 'update', '', '', '{\"icon\":\"md-cloud-upload\", \"title\":\"数据上传\"}', 1, 0, NULL, NULL),(30, 'update_table_page', 'update_table_page', '/update', '/update-table', '{\"icon\":\"ios-document\", \"title\":\"上传Csv\"}', 29, 1, NULL, NULL),(31, 'update_paste_page', 'update_paste_page', '/update', '/update-paste', '{\"icon\":\"md-clipboard\", \"title\":\"粘贴表格数据\"}', 29, 1, NULL, NULL),(32, 'excel', 'excel', '', '', '{\"icon\":\"ios-stats\", \"title\":\"EXCEL导入导出\"}', 1, 0, NULL, NULL),(33, 'upload-excel', 'upload-excel', '/excel', '/upload-excel', '{\"icon\":\"md-add\", \"title\":\"导入EXCEL\"}', 32, 1, NULL, NULL),(34, 'export-excel', 'export-excel', '/excel', '/export-excel', '{\"icon\":\"md-download\", \"title\":\"导出EXCEL\"}', 32, 1, NULL, NULL),(35, 'tools_methods', 'tools_methods', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL),(36, 'tools_methods_page', 'tools_methods_page', '/tools-methods', '/tools-methods', '{\"icon\":\"ios-hammer\", \"title\":\"工具方法\", \"beforeCloseName\":\"before_close_normal\"}', 35, 1, NULL, NULL),(37, 'i18n', 'i18n', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL),(38, 'i18n_page', 'i18n_page', '/i18n', '/i18n-page', '{\"icon\":\"md-planet\", \"title\":\"i18n - {{ i18n_page }}\"}', 37, 1, NULL, NULL),(39, 'error_store', 'error_store', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL),(40, 'error_store_page', 'error_store_page', '/error-store', '/error-store', '{\"icon\":\"ios-bug\", \"title\":\"错误收集\"}', 39, 1, NULL, NULL),(41, 'error_logger', 'error_logger', '', '', '{\"hideInBread\":true, \"hideInMenu\":true}', 1, 0, NULL, NULL),(42, 'error_logger_page', 'error_logger_page', '/single-page', '/error-logger', '{\"icon\":\"ios-bug\", \"title\":\"错误收集\"}', 41, 1, NULL, NULL),(43, 'directive', 'directive', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL),(44, 'directive_page', 'directive_page', '/directive', '/directive', '{\"icon\":\"ios-navigate\", \"title\":\"指令\"}', 43, 1, NULL, NULL),(45, 'argu', 'argu', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL),(46, 'params/:id', 'params', '/argu-page', '/params', '{\"icon\":\"md-flower\", \"title\":\"route => `{{ params }}-${route.params.id}`\", \"notCache\":true, \"beforeCloseName\":true}', 45, 1, NULL, NULL),(47, 'query', 'query', '/argu-page', '/query', '{\"icon\":\"md-flower\", \"title\":\"route => `{{ params }}-${route.params.id}`\", \"notCache\":true}', 45, 1, NULL, NULL),(48, '401', 'error_401', '/error-page', '/401', '{\"hideInBread\":true}', 1, 0, NULL, NULL),(49, '500', 'error_500', '/error-page', '/500', '{\"hideInBread\":true}', 1, 0, NULL, NULL),(50, '*', 'error_404', '/error-page', '/404', '{\"hideInBread\":true}', 1, 0, NULL, NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `casbin`.`post` WRITE;
DELETE FROM `casbin`.`post`;
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `casbin`.`product` WRITE;
DELETE FROM `casbin`.`product`;
INSERT INTO `casbin`.`product` (`id`,`product_code`,`product_name`,`price`,`number`,`created_At`,`updated_At`,`deleted_At`) VALUES (1, 'PD-001', 'asc1', 2, 2, '2019-01-08 10:30:33', '2019-01-08 10:30:33', NULL),(20, 'aaaaaa', 'bbb', 222, 333, '2019-04-15 09:56:47', NULL, NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `casbin`.`product_category` WRITE;
DELETE FROM `casbin`.`product_category`;
INSERT INTO `casbin`.`product_category` (`id`,`num_code`,`chn_name`,`created_At`,`updated_At`,`deleted_At`) VALUES (21, '1', 'aaa', '2019-04-15 10:44:37', NULL, NULL),(25, '2', '3', '2019-04-15 12:19:41', NULL, NULL),(26, '3', '3', '2019-04-15 12:19:54', NULL, NULL);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `casbin`.`role_menu` WRITE;
DELETE FROM `casbin`.`role_menu`;
INSERT INTO `casbin`.`role_menu` (`id`,`rid`,`mid`) VALUES (1, 68, 2),(2, 68, 3),(3, 68, 4),(4, 68, 5),(5, 68, 6),(6, 73, 16),(7, 73, 17),(8, 73, 18),(9, 73, 19),(10, 73, 20),(11, 73, 21),(12, 73, 22),(13, 73, 23),(14, 73, 24),(15, 73, 25),(16, 73, 26),(17, 73, 27),(18, 73, 28),(19, 69, 7),(20, 69, 8),(39, 69, 9);
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `casbin`.`user` WRITE;
DELETE FROM `casbin`.`user`;
INSERT INTO `casbin`.`user` (`id`,`username`,`password`,`gender`,`enable`,`name`,`phone`,`email`,`userface`,`create_time`,`update_time`) VALUES (76, 'yhm1', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, '1', '', '', '', '2019-01-14 11:54:11', NULL),(90, 'root', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, '超级用户', '', '', '', '2019-01-16 10:42:34', NULL),(108, '4', '6Xf9dl8Cv7OVFdc45vR7ig==', '', 1, '2', '', '', '', '2019-01-18 13:34:56', NULL),(111, 'root1', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, '超级用户', '', '', '', '2019-01-16 10:42:34', NULL),(112, 'yhm21', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, 'yy', '3213123', 'qq.com', '', '2019-01-18 10:08:12', NULL),(114, '22-2012341', '123456', '', 1, 'fsdvcdxcvx', 'ffffffffffffffffffffff', '.,.mn,nm,hgntfrgffffffffffff', '', '2019-01-18 12:54:54', NULL),(115, '3-121', '32qewqewqe', '', 0, '3ewqdd343444', '', 'sadsad1fs,..', '', '2019-01-18 12:56:31', NULL),(116, '41', '6Xf9dl8Cv7OVFdc45vR7ig==', '', 1, '77', '', '', '', '2019-01-18 13:34:56', NULL),(123, '3-127', '32qewqewqe', '', 0, '3ewqdd343444', '', 'sadsad1fs,..', '', '2019-01-18 12:56:31', NULL),(124, '48', '6Xf9dl8Cv7OVFdc45vR7ig==', '', 1, '4', '', '', '', '2019-01-18 13:34:56', NULL),(125, '9', 'CBpaVgtlw7hA/zD1hFbcdw==', '', 1, '', '', '', '', '2019-03-25 14:06:35', NULL),(126, 'yhm11', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, '266', '', '', '', '2019-01-14 11:54:11', NULL);
UNLOCK TABLES;
COMMIT;
