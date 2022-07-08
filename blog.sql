/*
 Navicat Premium Data Transfer

 Source Server         : 局域网k8s
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : 192.168.1.227:30306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 08/07/2022 20:49:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `author` varchar(255) NOT NULL DEFAULT '' COMMENT '作者',
  `cover` varchar(255) NOT NULL DEFAULT '' COMMENT '文章封面',
  `title` varchar(255) NOT NULL COMMENT '文章标题',
  `summary` varchar(255) NOT NULL DEFAULT '' COMMENT '文章摘要',
  `content` text COMMENT '文章内容',
  `visit` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章阅读量',
  `support` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章点赞量',
  `category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章分类id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  FULLTEXT KEY `idx_title_content` (`title`,`content`) /*!50100 WITH PARSER `ngram` */ 
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='文章';

-- ----------------------------
-- Records of blog_article
-- ----------------------------
BEGIN;
INSERT INTO `blog_article` (`id`, `author`, `cover`, `title`, `summary`, `content`, `visit`, `support`, `category_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '路人甲', '“”', '美国制裁香港和阿联酋公司，加大对伊朗施压', '美国政府星期三宣布制裁被控帮助伊朗石油和石化产品向东亚地区交付和销售的一个由香港、阿联酋和其它公司组成的网络，在美国寻求恢复2015年伊朗核协定时向伊朗施压。\n', '美国政府星期三宣布制裁被控帮助伊朗石油和石化产品向东亚地区交付和销售的一个由香港、阿联酋和其它公司组成的网络，在美国寻求恢复2015年伊朗核协定时向伊朗施压。\n\n美国财政部在一份声明中说，这个由个人与实体组成的网络利用一个波斯湾的幌子公司的网络从伊朗公司购买价值数亿美元的产品并运往东亚地区。\n\n伊朗和美国上星期在多哈举行的间接谈判没有出现突破。\n\n财政部负责反恐与金融情报事务的副部长布莱恩·纳尔逊（Brian Nelson）说，“美国尽管致力于与伊朗达成一项协议，寻求相互重返联合全面行动计划，但我们将继续利用所有我们的权力执行对伊朗石油与石化产品销售的制裁。”\n\n星期三的制裁对象包括伊朗的扎木石化公司（Jam Petrochemical），华盛顿指责扎木石化向东亚各地的公司出口石化产品，很多是向受美国制裁的伊朗石化商务公司销售，并运往中国。\n\n扎木没有对此立即置评。\n\n受制裁的还有埃德加商务解决方案（Edgar Commercial Solutions）。财政部说这家公司从受制裁的伊朗公司购买并出口了石化产品。华盛顿说这家公司利用香港的幌子公司粲泽工业有限公司（Lustro Industry Limited）来掩盖大宗购买石化产品的角色。粲泽工业有限公司星期三受到了制裁。\n\n阿联酋的阿里阿木塔瓦石油石化贸易公司（Ali Almutawa Petroleum and Petrochemical Trading）等几家公司也受到了制裁。华盛顿说，阿里阿木塔瓦石油石化贸易公司是受美国制裁的众祥石化有限公司（Triliance Petrochemical Co）的幌子公司。\n\n路透社无法立即得到埃德加商务解决方案、粲泽工业有限公司和阿里阿木塔瓦石油石化贸易公司的评论。\n\n在越南和新加坡的公司也受到了制裁。\n\n制裁措施冻结了这些被制裁公司的任何美国资产，通常禁止美国人与他们交往。那些与受制裁人员和实体交往的人也可能受到制裁。\n\n（本文依据了路透社的报道）\n\n', 11, 114, 1, '2022-07-07 23:17:49', '2022-07-08 20:35:00', NULL);
INSERT INTO `blog_article` (`id`, `author`, `cover`, `title`, `summary`, `content`, `visit`, `support`, `category_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '', '', '如果', '', '今生今世 永不再将你想起 除了\n除了在有些个\n因落泪而湿润的夜里 如果\n如果你愿意', 0, 102, 2, '2022-07-07 23:17:49', '2022-07-08 19:16:01', NULL);
INSERT INTO `blog_article` (`id`, `author`, `cover`, `title`, `summary`, `content`, `visit`, `support`, `category_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, '', '', '爱情', '', '有一天路标迁了希望你能从容\n有一天桥墩断了希望你能渡越\n有一天栋梁倒了希望你能坚强\n有一天期待蔫了希望你能理解', 0, 103, 2, '2022-07-07 23:17:49', '2022-07-08 19:16:05', NULL);
INSERT INTO `blog_article` (`id`, `author`, `cover`, `title`, `summary`, `content`, `visit`, `support`, `category_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, '', '', '远和近', '', '你 一会看我\n一会看云\n我觉得\n你看我时很远\n你看云时很近', 0, 104, 2, '2022-07-07 23:17:49', '2022-07-08 19:16:02', NULL);
INSERT INTO `blog_article` (`id`, `author`, `cover`, `title`, `summary`, `content`, `visit`, `support`, `category_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, '', '', '断章', '', '你站在桥上看风景，\n看风景人在楼上看你。\n明月装饰了你的窗子，\n你装饰了别人的梦。', 0, 105, 2, '2022-07-07 23:17:49', '2022-07-08 19:16:04', NULL);
INSERT INTO `blog_article` (`id`, `author`, `cover`, `title`, `summary`, `content`, `visit`, `support`, `category_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, '', '', '独语', '', '我向你倾吐思念\n你如石像\n沉默不应\n如果沉默是你的悲抑\n你知道这悲抑\n最伤我心', 0, 106, 2, '2022-07-07 23:17:49', '2022-07-08 19:16:07', NULL);
INSERT INTO `blog_article` (`id`, `author`, `cover`, `title`, `summary`, `content`, `visit`, `support`, `category_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'admin', 'string', 'string', 'string', 'string', 0, 107, 0, '2022-07-08 19:10:20', '2022-07-08 19:16:04', NULL);
INSERT INTO `blog_article` (`id`, `author`, `cover`, `title`, `summary`, `content`, `visit`, `support`, `category_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 'admin', 'http::/www.baidu.com1', '内容1', '摘要1', '标题1', 0, 108, 2, '2022-07-08 19:10:58', '2022-07-08 19:33:55', NULL);
COMMIT;

-- ----------------------------
-- Table structure for blog_category
-- ----------------------------
DROP TABLE IF EXISTS `blog_category`;
CREATE TABLE `blog_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '分类名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='文章分类';

-- ----------------------------
-- Records of blog_category
-- ----------------------------
BEGIN;
INSERT INTO `blog_category` (`id`, `name`) VALUES (1, '政治');
INSERT INTO `blog_category` (`id`, `name`) VALUES (2, '抒情');
COMMIT;

-- ----------------------------
-- Table structure for blog_comment
-- ----------------------------
DROP TABLE IF EXISTS `blog_comment`;
CREATE TABLE `blog_comment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `content` text COMMENT '评论内容',
  `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '评论者名称',
  `email` varchar(255) NOT NULL COMMENT '评论者邮箱',
  `article_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_article_id` (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='博客评论';

-- ----------------------------
-- Records of blog_comment
-- ----------------------------
BEGIN;
INSERT INTO `blog_comment` (`id`, `content`, `nickname`, `email`, `article_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'string', 'string', '7256316367@qq.com', 1, '2022-07-08 20:39:54', '2022-07-08 12:49:34', NULL);
INSERT INTO `blog_comment` (`id`, `content`, `nickname`, `email`, `article_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'string', '路人甲', '7256316367@qq.com', 1, '2022-07-08 20:40:04', '2022-07-08 12:49:37', NULL);
COMMIT;

-- ----------------------------
-- Table structure for blog_user
-- ----------------------------
DROP TABLE IF EXISTS `blog_user`;
CREATE TABLE `blog_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `username` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '用户状态（0：关闭、1：开启）',
  `salt` char(50) NOT NULL COMMENT '用户密码加盐',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `role_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户角色 0普通用户 1管理员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='网站用户';

-- ----------------------------
-- Records of blog_user
-- ----------------------------
BEGIN;
INSERT INTO `blog_user` (`id`, `username`, `password`, `status`, `salt`, `created_at`, `updated_at`, `role_type`) VALUES (1, 'admin', '6451a7174083172893f5497dbd1bdec1', 1, 'ajhg', '2022-07-08 10:30:27', '2022-07-08 11:09:48', 1);
COMMIT;

-- ----------------------------
-- Table structure for blog_website
-- ----------------------------
DROP TABLE IF EXISTS `blog_website`;
CREATE TABLE `blog_website` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '网站id',
  `name` varchar(32) NOT NULL COMMENT '网站名称',
  `logo` varchar(512) NOT NULL COMMENT '网站LOGO',
  `copyright` varchar(512) NOT NULL COMMENT '网站版权',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网站信息';

-- ----------------------------
-- Records of blog_website
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
