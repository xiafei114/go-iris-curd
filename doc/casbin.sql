/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50714
Source Host           : localhost:3306
Source Database       : casbin

Target Server Type    : MYSQL
Target Server Version : 50714
File Encoding         : 65001

Date: 2019-04-04 17:23:00
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
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

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('65', 'p', '90', '', '/*', 'ANY', '.*', '超级用户', null);
INSERT INTO `casbin_rule` VALUES ('66', 'g', '90', 'admin', 'a', '', '', '', null);
INSERT INTO `casbin_rule` VALUES ('67', 'g', '90', 'demo', 'a', '', '', '', null);
INSERT INTO `casbin_rule` VALUES ('68', 'p', 'admin', 'a', '/admin*', 'GET|POST|DELETE|PUT', '.*', '角色管理', null);
INSERT INTO `casbin_rule` VALUES ('69', 'p', 'demo', 'a', '/demo*', 'GET|POST|DELETE|PUT', '.*', 'demo角色12', null);
INSERT INTO `casbin_rule` VALUES ('71', 'p', 't1', 'a', '/*', 'PUT|DELETE|GET|POST', '.*', '测试1', null);
INSERT INTO `casbin_rule` VALUES ('72', 'p', 'user', 'a', '/*', 'PUT|DELETE|GET|POST', '.*', '普通用户', null);
INSERT INTO `casbin_rule` VALUES ('73', 'p', 'components', 'a', '/components*', 'GET|POST|DELETE|PUT', '.*', '组件管理', null);
INSERT INTO `casbin_rule` VALUES ('74', 'g', '90', 'components', 'a', '', '', '', null);

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES ('1', 'PD-001', 'asc1', '2','2', '2019-01-08 10:30:33', '2019-01-08 10:30:33',null);
INSERT INTO `product` VALUES ('2', 'PD-002', 'asc2', '2','2', '2019-01-08 10:38:21', '2019-01-08 10:30:33',null);
INSERT INTO `product` VALUES ('3', 'PD-003', 'asc3', '2','2', '2019-01-08 10:38:51', '2019-01-08 10:30:33',null);
INSERT INTO `product` VALUES ('4', 'PD-005', 'asc4', '2','2', '2019-01-08 10:39:33', '2019-01-08 10:30:33',null);
INSERT INTO `product` VALUES ('5', 'PD-006', 'asc5', '2','2', '2019-01-08 10:44:21', '2019-01-08 10:30:33',null);

-- ----------------------------
-- Table structure for dep
-- ----------------------------
DROP TABLE IF EXISTS `dep`;
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
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of dep
-- ----------------------------
INSERT INTO `dep` VALUES ('1', '稷下之学', '荀子', '17830895300', '614143260@qq.com', '0', '0', '', '2019-03-27 15:19:39', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('2', '儒家', '孔子', '1000', 'laozi@163.com', '0', '1', '崇尚\"礼乐\"和\"仁义\"，主张\"德治\"和\"仁政\"，教育了很多人为人处世所遵循的准则规范。', '2019-04-04 17:13:43', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('3', '道家', '老子', '1001', 'laozi@163.com', '0', '1', '天道运行的原理。天道轮回各种哲学上的思想都能在道家找到答案。世界上很多发生的事情都有其规律，人力不可更改。', '2019-04-30 17:13:45', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('4', '法家', '李斯', '1002', 'lisi@163.com', '0', '1', '依法治国，天子犯法与庶民同罪。法律视为一种有利于社会统治的强制性工具，体现法制建设的思想，一直被沿用至今。', '2019-04-16 17:13:49', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('5', '墨家', '墨子', '1003', 'mozi@163.com', '0', '1', '兼爱非攻，反对强大的王国去攻击弱小的王国，在思想上尊天事鬼，一切以保持社会现状的稳定为主。', '2019-04-10 17:13:52', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('6', '名家', '公孙龙', '1004', 'gongsl@163.com', '0', '1', '以辩论出名，著名的白马非马也是名家的思想。以逻辑推理来证明事物的好与坏、真实与否。', '2019-04-09 17:13:55', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('7', '阴阳家', '邹衍', '1005', 'zhouyan@163.com', '0', '1', '五行学说，从天文和地理方面来判断事物的凶吉。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('8', '纵横家', '苏秦|张仪', '1006', 'sz@163.com', '0', '1', '合纵连横，捭阖之术，阴阳之变化也。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('9', '农家', '许行', '1007', 'xuxing@163.com', '0', '1', '注重生产，研究植物生长和产出的学派。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('10', '兵家', '孙膑', '1008', 'sunbing@163.com', '0', '1', '讲究利用武力，最大化的夺取敌方的利益从而赢得战争的胜利。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('11', '医家', '扁鹊', '1009', 'bianque@163.com', '0', '1', '医者仁心', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('12', '礼乐', '孟子', '1010', 'mengzi@163.com', '0', '2', '', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('13', '武当', '张三丰', '1011', 'zsf@163.com', '0', '3', '', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('14', '庄子游', '庄子', '1012', 'zz@163.com', '0', '3', '', '2019-04-24 17:13:59', '2019-04-24 17:13:59');

-- ----------------------------
-- Table structure for dep_user
-- ----------------------------
DROP TABLE IF EXISTS `dep_user`;
CREATE TABLE `dep_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dep_ip` int(10) DEFAULT NULL COMMENT '部门id',
  `uid` int(10) DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of dep_user
-- ----------------------------

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `path` varchar(64) DEFAULT '',
  `name` varchar(64) DEFAULT '' COMMENT '必须是唯一的',
  `modular` varchar(64) DEFAULT '',
  `component` varchar(64) DEFAULT '',
  `meta` text,
  `parent_id` int(10) DEFAULT NULL,
  `is_sub` tinyint(1) NOT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES ('1', '/', '所有', '', null, null, null, '0', null, null);
INSERT INTO `menu` VALUES ('2', 'sys_mgr', 'sys_mgr', '', '', '{\"title\":\"系统管理\", \"icon\":\"ios-settings\"}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('3', 'user_mgr', 'user_mgr', '/admin/user', '/user', '{\"title\":\"用户管理\", \"icon\":\"ios-settings\"}', '2', '1', null, null);
INSERT INTO `menu` VALUES ('4', 'role_mgr', 'role_mgr', '/admin', '/role', '{\"title\":\"角色管理\", \"icon\":\"ios-settings\"}', '2', '1', null, null);
INSERT INTO `menu` VALUES ('5', 'menu_mgr', 'menu_mgr', '/admin', '/menu', '{\"title\":\"菜单管理\", \"icon\":\"ios-settings\"}', '2', '1', null, null);
INSERT INTO `menu` VALUES ('6', 'dep_mgr', 'dep_mgr', '/admin/dep', '/dep', '{\"title\":\"部门管理\", \"icon\":\"ios-settings\"}', '2', '1', null, null);
INSERT INTO `menu` VALUES ('7', '/demo', 'demo管理', '', '', null, '1', '0', '2018-07-25 13:59:04', null);
INSERT INTO `menu` VALUES ('8', '/demo1', 'demo1', '/demo', '/demo1', null, '7', '1', '2018-07-25 13:59:57', null);
INSERT INTO `menu` VALUES ('9', '/demo2', 'demo2', '/demo', '/demo2', null, '7', '1', '2018-07-25 14:00:27', null);
INSERT INTO `menu` VALUES ('10', '/demo3', 'demo3', '/demo', '/demo3', null, '7', '1', '2018-07-25 14:14:05', null);
INSERT INTO `menu` VALUES ('11', 'doc', 'doc', '', '', '{\"title\":\"文档\", \"href\":\"https://lison16.github.io/iview-admin-doc/#/\", \"icon\":\"ios-book\"}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('12', 'join', 'join', '', '', '{\"hideInBread\": true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('13', 'join_page', 'join_page', '', '/join_page', '{\"icon\":\"_qq\", \"title\":\"QQ群\"}', '12', '1', null, null);
INSERT INTO `menu` VALUES ('14', 'message', 'message', '', '', '{\"hideInBread\":true, \"hideInMenu\":true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('15', 'message_page', 'message_page', '/single-page/message', '/index', '{\"icon\":\"md-notifications\", \"title\":\"消息中心\"}', '14', '1', null, null);
INSERT INTO `menu` VALUES ('16', 'components', 'components', '', '', '{\"icon\":\"logo-buffer\", \"title\":\"组件\"}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('17', 'tree_select_page', 'tree_select_page', '/components/tree-select', '/index', '{\"icon\":\"md-arrow-dropdown-circle\", \"title\":\"树状下拉选择器\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('18', 'count_to_page', 'count_to_page', '/components/count-to', '/count-to', '{\"icon\":\"md-trending-up\", \"title\":\"数字渐变\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('19', 'drag_list_page', 'drag_list_page', '/components/drag-list', '/drag-list', '{\"icon\":\"ios-infinite\", \"title\":\"拖拽列表\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('20', 'drag_drawer_page', 'drag_drawer_page', '/components/drag-drawer', '/index', '{\"icon\":\"md-list\", \"title\":\"可拖拽抽屉\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('21', 'org_tree_page', 'org_tree_page', '/components/org-tree', '/index', '{\"icon\":\"ios-people\", \"title\":\"组织结构树\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('22', 'tree_table_page', 'tree_table_page', '/components/tree-table', '/index', '{\"icon\":\"md-git-branch\", \"title\":\"树状表格\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('23', 'cropper_page', 'cropper_page', '/components/cropper', '/cropper', '{\"icon\":\"md-crop\", \"title\":\"图片裁剪\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('24', 'tables_page', 'tables_page', '/components/tables', '/tables', '{\"icon\":\"md-grid\", \"title\":\"多功能表格\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('25', 'split_pane_page', 'split_pane_page', '/components/split-pane', '/split-pane', '{\"icon\":\"md-pause\", \"title\":\"分割窗口\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('26', 'markdown_page', 'markdown_page', '/components/markdown', '/markdown', '{\"icon\":\"logo-markdown\", \"title\":\"Markdown编辑器\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('27', 'editor_page', 'editor_page', '/components/editor', '/editor', '{\"icon\":\"ios-create\", \"title\":\"富文本编辑器\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('28', 'icons_page', 'icons_page', '/components/icons', '/icons', '{\"icon\":\"_bear\", \"title\":\"自定义图标\"}', '16', '1', null, null);
INSERT INTO `menu` VALUES ('29', 'update', 'update', '', '', '{\"icon\":\"md-cloud-upload\", \"title\":\"数据上传\"}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('30', 'update_table_page', 'update_table_page', '/update', '/update-table', '{\"icon\":\"ios-document\", \"title\":\"上传Csv\"}', '29', '1', null, null);
INSERT INTO `menu` VALUES ('31', 'update_paste_page', 'update_paste_page', '/update', '/update-paste', '{\"icon\":\"md-clipboard\", \"title\":\"粘贴表格数据\"}', '29', '1', null, null);
INSERT INTO `menu` VALUES ('32', 'excel', 'excel', '', '', '{\"icon\":\"ios-stats\", \"title\":\"EXCEL导入导出\"}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('33', 'upload-excel', 'upload-excel', '/excel', '/upload-excel', '{\"icon\":\"md-add\", \"title\":\"导入EXCEL\"}', '32', '1', null, null);
INSERT INTO `menu` VALUES ('34', 'export-excel', 'export-excel', '/excel', '/export-excel', '{\"icon\":\"md-download\", \"title\":\"导出EXCEL\"}', '32', '1', null, null);
INSERT INTO `menu` VALUES ('35', 'tools_methods', 'tools_methods', '', '', '{\"hideInBread\":true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('36', 'tools_methods_page', 'tools_methods_page', '/tools-methods', '/tools-methods', '{\"icon\":\"ios-hammer\", \"title\":\"工具方法\", \"beforeCloseName\":\"before_close_normal\"}', '35', '1', null, null);
INSERT INTO `menu` VALUES ('37', 'i18n', 'i18n', '', '', '{\"hideInBread\":true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('38', 'i18n_page', 'i18n_page', '/i18n', '/i18n-page', '{\"icon\":\"md-planet\", \"title\":\"i18n - {{ i18n_page }}\"}', '37', '1', null, null);
INSERT INTO `menu` VALUES ('39', 'error_store', 'error_store', '', '', '{\"hideInBread\":true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('40', 'error_store_page', 'error_store_page', '/error-store', '/error-store', '{\"icon\":\"ios-bug\", \"title\":\"错误收集\"}', '39', '1', null, null);
INSERT INTO `menu` VALUES ('41', 'error_logger', 'error_logger', '', '', '{\"hideInBread\":true, \"hideInMenu\":true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('42', 'error_logger_page', 'error_logger_page', '/single-page', '/error-logger', '{\"icon\":\"ios-bug\", \"title\":\"错误收集\"}', '41', '1', null, null);
INSERT INTO `menu` VALUES ('43', 'directive', 'directive', '', '', '{\"hideInBread\":true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('44', 'directive_page', 'directive_page', '/directive', '/directive', '{\"icon\":\"ios-navigate\", \"title\":\"指令\"}', '43', '1', null, null);
INSERT INTO `menu` VALUES ('45', 'argu', 'argu', '', '', '{\"hideInBread\":true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('46', 'params/:id', 'params', '/argu-page', '/params', '{\"icon\":\"md-flower\", \"title\":\"route => `{{ params }}-${route.params.id}`\", \"notCache\":true, \"beforeCloseName\":true}', '45', '1', null, null);
INSERT INTO `menu` VALUES ('47', 'query', 'query', '/argu-page', '/query', '{\"icon\":\"md-flower\", \"title\":\"route => `{{ params }}-${route.params.id}`\", \"notCache\":true}', '45', '1', null, null);
INSERT INTO `menu` VALUES ('48', '401', 'error_401', '/error-page', '/401', '{\"hideInBread\":true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('49', '500', 'error_500', '/error-page', '/500', '{\"hideInBread\":true}', '1', '0', null, null);
INSERT INTO `menu` VALUES ('50', '*', 'error_404', '/error-page', '/404', '{\"hideInBread\":true}', '1', '0', null, null);

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `code` varchar(64) DEFAULT '' COMMENT '职位编码',
  `name` varchar(64) DEFAULT '' COMMENT '职位名称',
  `enable` tinyint(1) DEFAULT NULL COMMENT '是否启用',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of post
-- ----------------------------

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `rid` int(10) NOT NULL COMMENT '角色id',
  `mid` int(10) NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES ('1', '68', '2');
INSERT INTO `role_menu` VALUES ('2', '68', '3');
INSERT INTO `role_menu` VALUES ('3', '68', '4');
INSERT INTO `role_menu` VALUES ('4', '68', '5');
INSERT INTO `role_menu` VALUES ('5', '68', '6');
INSERT INTO `role_menu` VALUES ('6', '73', '16');
INSERT INTO `role_menu` VALUES ('7', '73', '17');
INSERT INTO `role_menu` VALUES ('8', '73', '18');
INSERT INTO `role_menu` VALUES ('9', '73', '19');
INSERT INTO `role_menu` VALUES ('10', '73', '20');
INSERT INTO `role_menu` VALUES ('11', '73', '21');
INSERT INTO `role_menu` VALUES ('12', '73', '22');
INSERT INTO `role_menu` VALUES ('13', '73', '23');
INSERT INTO `role_menu` VALUES ('14', '73', '24');
INSERT INTO `role_menu` VALUES ('15', '73', '25');
INSERT INTO `role_menu` VALUES ('16', '73', '26');
INSERT INTO `role_menu` VALUES ('17', '73', '27');
INSERT INTO `role_menu` VALUES ('18', '73', '28');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `gender` varchar(2) DEFAULT '' COMMENT '性别',
  `enable` tinyint(1) DEFAULT '0' COMMENT '0：启用用户 1：禁用用户',
  `name` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `userface` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('76', 'yhm1', 'x04jpoIrc8/mvNRqAG59Wg==', '', '1', '1', '', '', '', '2019-01-14 11:54:11', null);
INSERT INTO `user` VALUES ('90', 'root', 'x04jpoIrc8/mvNRqAG59Wg==', '', '1', '超级用户', '', '', '', '2019-01-16 10:42:34', null);
INSERT INTO `user` VALUES ('92', 'yhm2', 'x04jpoIrc8/mvNRqAG59Wg==', '', '1', 'yy', '3213123', 'qq.com', '', '2019-01-18 10:08:12', null);
INSERT INTO `user` VALUES ('104', '1-01', '1234', '', '0', 'mhjmhghf', 'nnvbn432423', '98089089', '', '2019-01-18 11:58:53', null);
INSERT INTO `user` VALUES ('105', '22-201234', '123456', '', '1', 'fsdvcdxcvx', 'ffffffffffffffffffffff', '.,.mn,nm,hgntfrgffffffffffff', '', '2019-01-18 12:54:54', null);
INSERT INTO `user` VALUES ('106', '3-12', '32qewqewqe', '', '0', '3ewqdd343444', '', 'sadsad1fs,..', '', '2019-01-18 12:56:31', null);
INSERT INTO `user` VALUES ('108', '4', '6Xf9dl8Cv7OVFdc45vR7ig==', '', '1', '2', '', '', '', '2019-01-18 13:34:56', null);
INSERT INTO `user` VALUES ('109', '', 'CBpaVgtlw7hA/zD1hFbcdw==', '', '1', '', '', '', '', '2019-03-25 14:06:35', null);