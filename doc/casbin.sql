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

 Date: 16/04/2019 11:05:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `p_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `IDX_casbin_rule_p_type`(`p_type`) USING BTREE,
  INDEX `IDX_casbin_rule_v2`(`v2`) USING BTREE,
  INDEX `IDX_casbin_rule_v3`(`v3`) USING BTREE,
  INDEX `IDX_casbin_rule_v4`(`v4`) USING BTREE,
  INDEX `IDX_casbin_rule_v5`(`v5`) USING BTREE,
  INDEX `IDX_casbin_rule_v0`(`v0`) USING BTREE,
  INDEX `IDX_casbin_rule_v1`(`v1`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 75 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (65, 'p', '90', '', '/*', 'ANY', '.*', '超级用户', NULL);
INSERT INTO `casbin_rule` VALUES (66, 'g', '90', 'admin', 'a', '', '', NULL, NULL);
INSERT INTO `casbin_rule` VALUES (67, 'g', '90', 'demo', 'a', '', '', NULL, NULL);
INSERT INTO `casbin_rule` VALUES (68, 'p', 'admin', 'a', '/admin*', 'GET|POST|DELETE|PUT', '.*', NULL, NULL);
INSERT INTO `casbin_rule` VALUES (69, 'p', 'demo', 'a', '/demo*', 'GET|POST|DELETE|PUT', '.*', NULL, NULL);
INSERT INTO `casbin_rule` VALUES (71, 'p', 't1', 'a', '/*', 'PUT|DELETE|GET|POST', '.*', NULL, NULL);
INSERT INTO `casbin_rule` VALUES (72, 'p', 'user', 'a', '/*', 'PUT|DELETE|GET|POST', '.*', NULL, NULL);
INSERT INTO `casbin_rule` VALUES (73, 'p', 'components', 'a', '/components*', 'GET|POST|DELETE|PUT', '.*', NULL, NULL);
INSERT INTO `casbin_rule` VALUES (74, 'g', '90', 'components', 'a', '', '', NULL, NULL);

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

-- ----------------------------
-- Records of demo
-- ----------------------------
INSERT INTO `demo` VALUES (1, 'PD-001', 'asc', 2, '2019-01-08 10:30:33');
INSERT INTO `demo` VALUES (2, 'PD-002', 'asc', 2, '2019-01-08 10:38:21');
INSERT INTO `demo` VALUES (3, 'PD-003', 'asc', 2, '2019-01-08 10:38:51');
INSERT INTO `demo` VALUES (4, 'PD-005', 'asc', 2, '2019-01-08 10:39:33');
INSERT INTO `demo` VALUES (5, 'PD-006', 'asc', 2, '2019-01-08 10:44:21');

-- ----------------------------
-- Table structure for dep
-- ----------------------------
DROP TABLE IF EXISTS `dep`;
CREATE TABLE `dep`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dep_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `leader` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '部门领导人uid',
  `phone` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门电话',
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门邮箱',
  `disabled` tinyint(1) NULL DEFAULT NULL COMMENT '0:可用 否则:不可用',
  `parent_id` int(10) NULL DEFAULT NULL,
  `dep_desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dep
-- ----------------------------
INSERT INTO `dep` VALUES (1, '稷下之学', '荀子', '17830895300', '614143260@qq.com', 0, 0, '', '2019-03-27 15:19:39', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (2, '儒家', '孔子', '1000', 'laozi@163.com', 0, 1, '崇尚\"礼乐\"和\"仁义\"，主张\"德治\"和\"仁政\"，教育了很多人为人处世所遵循的准则规范。', '2019-04-04 17:13:43', '2019-04-08 15:07:03');
INSERT INTO `dep` VALUES (3, '道家', '老子', '1001', 'laozi@163.com', 0, 1, '天道运行的原理。天道轮回各种哲学上的思想都能在道家找到答案。世界上很多发生的事情都有其规律，人力不可更改。', '2019-04-30 17:13:45', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (4, '法家', '李斯', '1002', 'lisi@163.com', 0, 1, '依法治国，天子犯法与庶民同罪。法律视为一种有利于社会统治的强制性工具，体现法制建设的思想，一直被沿用至今。', '2019-04-16 17:13:49', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (5, '墨家', '墨子', '1003', 'mozi@163.com', 0, 1, '兼爱非攻，反对强大的王国去攻击弱小的王国，在思想上尊天事鬼，一切以保持社会现状的稳定为主。', '2019-04-10 17:13:52', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (6, '名家', '公孙龙', '1004', 'gongsl@163.com', 0, 1, '以辩论出名，著名的白马非马也是名家的思想。以逻辑推理来证明事物的好与坏、真实与否。', '2019-04-09 17:13:55', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (7, '阴阳家', '邹衍', '1005', 'zhouyan@163.com', 0, 1, '五行学说，从天文和地理方面来判断事物的凶吉。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (8, '纵横家', '苏秦|张仪', '1006', 'sz@163.com', 0, 1, '合纵连横，捭阖之术，阴阳之变化也。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (9, '农家', '许行', '1007', 'xuxing@163.com', 0, 1, '注重生产，研究植物生长和产出的学派。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (10, '兵家', '孙膑', '1008', 'sunbing@163.com', 0, 1, '讲究利用武力，最大化的夺取敌方的利益从而赢得战争的胜利。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (11, '医家', '扁鹊', '1009', 'bianque@163.com', 0, 1, '医者仁心', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (12, '礼乐', '孟子', '1010', 'mengzi@163.com', 0, 2, '', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (13, '武当', '张三丰', '1011', 'zsf@163.com', 0, 3, '', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (14, '庄子游', '庄子', '1012', 'zz@163.com', 0, 3, '', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES (24, 'test', '', '', '', 1, 5, '', '2019-04-08 14:53:05', NULL);

-- ----------------------------
-- Table structure for dep_user
-- ----------------------------
DROP TABLE IF EXISTS `dep_user`;
CREATE TABLE `dep_user`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dep_ip` int(10) NULL DEFAULT NULL COMMENT '部门id',
  `uid` int(10) NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `path` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '必须是唯一的',
  `modular` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `component` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `meta` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `parent_id` int(10) NULL DEFAULT NULL,
  `is_sub` tinyint(1) NOT NULL,
  `create_time` timestamp(0) NULL DEFAULT NULL,
  `update_time` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `key`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, '/', '所有', '', NULL, NULL, NULL, 0, NULL, NULL);
INSERT INTO `menu` VALUES (2, 'sys_mgr', 'sys_mgr', NULL, NULL, '{\"title\":\"系统管理\", \"icon\":\"ios-settings\"}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (3, 'user_mgr', 'user_mgr', '/admin/users', '/users', '{\"title\":\"用户管理\", \"icon\":\"ios-settings\"}', 2, 1, NULL, NULL);
INSERT INTO `menu` VALUES (4, 'role_mgr', 'role_mgr', '/admin', '/role', '{\"title\":\"角色管理\", \"icon\":\"ios-settings\"}', 2, 1, NULL, NULL);
INSERT INTO `menu` VALUES (5, 'menu_mgr', 'menu_mgr', '/admin', '/menu', '{\"title\":\"菜单管理\", \"icon\":\"ios-settings\"}', 2, 1, NULL, NULL);
INSERT INTO `menu` VALUES (6, 'dep_mgr', 'dep_mgr', '/admin/dep', '/dep', '{\"title\":\"部门管理\", \"icon\":\"ios-settings\"}', 2, 1, NULL, NULL);
INSERT INTO `menu` VALUES (7, 'demo_mgr', 'demo_mgr', NULL, NULL, '{\"title\":\"demo管理\", \"icon\":\"ios-settings\"}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (8, 'product_mgr', 'product_mgr', '/demo/product', '/product', '{\"title\":\"产品管理\", \"icon\":\"ios-settings\"}', 7, 1, NULL, NULL);
INSERT INTO `menu` VALUES (9, 'product_cate_mgr', 'product_cate_mgr', '/demo/product_cate', '/product_cate', '{\"title\":\"产品类别\", \"icon\":\"ios-settings\"}', 7, 1, '2018-07-25 14:00:27', NULL);
INSERT INTO `menu` VALUES (10, '/demo3', 'demo3', '/demo', '/demo3', NULL, 7, 1, '2018-07-25 14:14:05', NULL);
INSERT INTO `menu` VALUES (11, 'doc', 'doc', '', '', '{\"title\":\"文档\", \"href\":\"https://lison16.github.io/iview-admin-doc/#/\", \"icon\":\"ios-book\"}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (12, 'join', 'join', '', '', '{\"hideInBread\": true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (13, 'join_page', 'join_page', '', '/join_page', '{\"icon\":\"_qq\", \"title\":\"QQ群\"}', 12, 1, NULL, NULL);
INSERT INTO `menu` VALUES (14, 'message', 'message', '', '', '{\"hideInBread\":true, \"hideInMenu\":true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (15, 'message_page', 'message_page', '/single-page/message', '/index', '{\"icon\":\"md-notifications\", \"title\":\"消息中心\"}', 14, 1, NULL, NULL);
INSERT INTO `menu` VALUES (16, 'components', 'components', NULL, NULL, '{\"icon\":\"logo-buffer\", \"title\":\"组件\"}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (17, 'tree_select_page', 'tree_select_page', '/components/tree-select', '/index', '{\"icon\":\"md-arrow-dropdown-circle\", \"title\":\"树状下拉选择器\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (18, 'count_to_page', 'count_to_page', '/components/count-to', '/count-to', '{\"icon\":\"md-trending-up\", \"title\":\"数字渐变\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (19, 'drag_list_page', 'drag_list_page', '/components/drag-list', '/drag-list', '{\"icon\":\"ios-infinite\", \"title\":\"拖拽列表\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (20, 'drag_drawer_page', 'drag_drawer_page', '/components/drag-drawer', '/index', '{\"icon\":\"md-list\", \"title\":\"可拖拽抽屉\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (21, 'org_tree_page', 'org_tree_page', '/components/org-tree', '/index', '{\"icon\":\"ios-people\", \"title\":\"组织结构树\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (22, 'tree_table_page', 'tree_table_page', '/components/tree-table', '/index', '{\"icon\":\"md-git-branch\", \"title\":\"树状表格\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (23, 'cropper_page', 'cropper_page', '/components/cropper', '/cropper', '{\"icon\":\"md-crop\", \"title\":\"图片裁剪\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (24, 'tables_page', 'tables_page', '/components/tables', '/tables', '{\"icon\":\"md-grid\", \"title\":\"多功能表格\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (25, 'split_pane_page', 'split_pane_page', '/components/split-pane', '/split-pane', '{\"icon\":\"md-pause\", \"title\":\"分割窗口\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (26, 'markdown_page', 'markdown_page', '/components/markdown', '/markdown', '{\"icon\":\"logo-markdown\", \"title\":\"Markdown编辑器\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (27, 'editor_page', 'editor_page', '/components/editor', '/editor', '{\"icon\":\"ios-create\", \"title\":\"富文本编辑器\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (28, 'icons_page', 'icons_page', '/components/icons', '/icons', '{\"icon\":\"_bear\", \"title\":\"自定义图标\"}', 16, 1, NULL, NULL);
INSERT INTO `menu` VALUES (29, 'update', 'update', '', '', '{\"icon\":\"md-cloud-upload\", \"title\":\"数据上传\"}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (30, 'update_table_page', 'update_table_page', '/update', '/update-table', '{\"icon\":\"ios-document\", \"title\":\"上传Csv\"}', 29, 1, NULL, NULL);
INSERT INTO `menu` VALUES (31, 'update_paste_page', 'update_paste_page', '/update', '/update-paste', '{\"icon\":\"md-clipboard\", \"title\":\"粘贴表格数据\"}', 29, 1, NULL, NULL);
INSERT INTO `menu` VALUES (32, 'excel', 'excel', '', '', '{\"icon\":\"ios-stats\", \"title\":\"EXCEL导入导出\"}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (33, 'upload-excel', 'upload-excel', '/excel', '/upload-excel', '{\"icon\":\"md-add\", \"title\":\"导入EXCEL\"}', 32, 1, NULL, NULL);
INSERT INTO `menu` VALUES (34, 'export-excel', 'export-excel', '/excel', '/export-excel', '{\"icon\":\"md-download\", \"title\":\"导出EXCEL\"}', 32, 1, NULL, NULL);
INSERT INTO `menu` VALUES (35, 'tools_methods', 'tools_methods', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (36, 'tools_methods_page', 'tools_methods_page', '/tools-methods', '/tools-methods', '{\"icon\":\"ios-hammer\", \"title\":\"工具方法\", \"beforeCloseName\":\"before_close_normal\"}', 35, 1, NULL, NULL);
INSERT INTO `menu` VALUES (37, 'i18n', 'i18n', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (38, 'i18n_page', 'i18n_page', '/i18n', '/i18n-page', '{\"icon\":\"md-planet\", \"title\":\"i18n - {{ i18n_page }}\"}', 37, 1, NULL, NULL);
INSERT INTO `menu` VALUES (39, 'error_store', 'error_store', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (40, 'error_store_page', 'error_store_page', '/error-store', '/error-store', '{\"icon\":\"ios-bug\", \"title\":\"错误收集\"}', 39, 1, NULL, NULL);
INSERT INTO `menu` VALUES (41, 'error_logger', 'error_logger', '', '', '{\"hideInBread\":true, \"hideInMenu\":true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (42, 'error_logger_page', 'error_logger_page', '/single-page', '/error-logger', '{\"icon\":\"ios-bug\", \"title\":\"错误收集\"}', 41, 1, NULL, NULL);
INSERT INTO `menu` VALUES (43, 'directive', 'directive', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (44, 'directive_page', 'directive_page', '/directive', '/directive', '{\"icon\":\"ios-navigate\", \"title\":\"指令\"}', 43, 1, NULL, NULL);
INSERT INTO `menu` VALUES (45, 'argu', 'argu', '', '', '{\"hideInBread\":true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (46, 'params/:id', 'params', '/argu-page', '/params', '{\"icon\":\"md-flower\", \"title\":\"route => `{{ params }}-${route.params.id}`\", \"notCache\":true, \"beforeCloseName\":true}', 45, 1, NULL, NULL);
INSERT INTO `menu` VALUES (47, 'query', 'query', '/argu-page', '/query', '{\"icon\":\"md-flower\", \"title\":\"route => `{{ params }}-${route.params.id}`\", \"notCache\":true}', 45, 1, NULL, NULL);
INSERT INTO `menu` VALUES (48, '401', 'error_401', '/error-page', '/401', '{\"hideInBread\":true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (49, '500', 'error_500', '/error-page', '/500', '{\"hideInBread\":true}', 1, 0, NULL, NULL);
INSERT INTO `menu` VALUES (50, '*', 'error_404', '/error-page', '/404', '{\"hideInBread\":true}', 1, 0, NULL, NULL);

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '职位编码',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '职位名称',
  `enable` tinyint(1) NULL DEFAULT NULL COMMENT '是否启用',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `product_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `product_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `product_Cate_Name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `product_Cate_Id` int(10) NULL DEFAULT NULL,
  `price` int(11) NULL DEFAULT NULL,
  `number` int(11) NULL DEFAULT NULL,
  `created_At` datetime(0) NULL DEFAULT NULL,
  `updated_At` datetime(0) NULL DEFAULT NULL,
  `deleted_At` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES (1, 'PD-001', 'asc1', NULL, NULL, 2, 2, '2019-01-08 10:30:33', '2019-01-08 10:30:33', NULL);
INSERT INTO `product` VALUES (20, 'aaaaaa', 'bbb', NULL, NULL, 222, 333, '2019-04-15 09:56:47', NULL, NULL);

-- ----------------------------
-- Table structure for product_category
-- ----------------------------
DROP TABLE IF EXISTS `product_category`;
CREATE TABLE `product_category`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `num_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `chn_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `created_At` datetime(0) NULL DEFAULT NULL,
  `updated_At` datetime(0) NULL DEFAULT NULL,
  `deleted_At` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product_category
-- ----------------------------
INSERT INTO `product_category` VALUES (21, '1', 'aaa', '2019-04-15 10:44:37', NULL, NULL);
INSERT INTO `product_category` VALUES (25, '2', '3', '2019-04-15 12:19:41', NULL, NULL);
INSERT INTO `product_category` VALUES (26, '3', '3', '2019-04-15 12:19:54', NULL, NULL);

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `rid` int(10) NOT NULL COMMENT '角色id',
  `mid` int(10) NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES (1, 68, 2);
INSERT INTO `role_menu` VALUES (2, 68, 3);
INSERT INTO `role_menu` VALUES (3, 68, 4);
INSERT INTO `role_menu` VALUES (4, 68, 5);
INSERT INTO `role_menu` VALUES (5, 68, 6);
INSERT INTO `role_menu` VALUES (6, 73, 16);
INSERT INTO `role_menu` VALUES (7, 73, 17);
INSERT INTO `role_menu` VALUES (8, 73, 18);
INSERT INTO `role_menu` VALUES (9, 73, 19);
INSERT INTO `role_menu` VALUES (10, 73, 20);
INSERT INTO `role_menu` VALUES (11, 73, 21);
INSERT INTO `role_menu` VALUES (12, 73, 22);
INSERT INTO `role_menu` VALUES (13, 73, 23);
INSERT INTO `role_menu` VALUES (14, 73, 24);
INSERT INTO `role_menu` VALUES (15, 73, 25);
INSERT INTO `role_menu` VALUES (16, 73, 26);
INSERT INTO `role_menu` VALUES (17, 73, 27);
INSERT INTO `role_menu` VALUES (18, 73, 28);
INSERT INTO `role_menu` VALUES (19, 69, 7);
INSERT INTO `role_menu` VALUES (20, 69, 8);
INSERT INTO `role_menu` VALUES (39, 69, 9);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `gender` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '性别',
  `enable` tinyint(1) NULL DEFAULT 0 COMMENT '0：启用用户 1：禁用用户',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `userface` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 127 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (76, 'yhm1', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, '1', '', '', '', '2019-01-14 11:54:11', NULL);
INSERT INTO `user` VALUES (90, 'root', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, '超级用户', '', '', '', '2019-01-16 10:42:34', NULL);
INSERT INTO `user` VALUES (108, '4', '6Xf9dl8Cv7OVFdc45vR7ig==', '', 1, '2', '', '', '', '2019-01-18 13:34:56', NULL);
INSERT INTO `user` VALUES (111, 'root1', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, '超级用户', '', '', '', '2019-01-16 10:42:34', NULL);
INSERT INTO `user` VALUES (112, 'yhm21', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, 'yy', '3213123', 'qq.com', '', '2019-01-18 10:08:12', NULL);
INSERT INTO `user` VALUES (114, '22-2012341', '123456', '', 1, 'fsdvcdxcvx', 'ffffffffffffffffffffff', '.,.mn,nm,hgntfrgffffffffffff', '', '2019-01-18 12:54:54', NULL);
INSERT INTO `user` VALUES (115, '3-121', '32qewqewqe', '', 0, '3ewqdd343444', '', 'sadsad1fs,..', '', '2019-01-18 12:56:31', NULL);
INSERT INTO `user` VALUES (116, '41', '6Xf9dl8Cv7OVFdc45vR7ig==', '', 1, '77', '', '', '', '2019-01-18 13:34:56', NULL);
INSERT INTO `user` VALUES (123, '3-127', '32qewqewqe', '', 0, '3ewqdd343444', '', 'sadsad1fs,..', '', '2019-01-18 12:56:31', NULL);
INSERT INTO `user` VALUES (124, '48', '6Xf9dl8Cv7OVFdc45vR7ig==', '', 1, '4', '', '', '', '2019-01-18 13:34:56', NULL);
INSERT INTO `user` VALUES (125, '9', 'CBpaVgtlw7hA/zD1hFbcdw==', '', 1, '', '', '', '', '2019-03-25 14:06:35', NULL);
INSERT INTO `user` VALUES (126, 'yhm11', 'x04jpoIrc8/mvNRqAG59Wg==', '', 1, '266', '', '', '', '2019-01-14 11:54:11', NULL);

SET FOREIGN_KEY_CHECKS = 1;
