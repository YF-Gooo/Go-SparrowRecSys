/*
 Navicat MySQL Data Transfer

 Source Server         : looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : tutuy_user

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 10/03/2022 17:14:49
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint                                                        NOT NULL AUTO_INCREMENT,
    `create_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `is_admin`    tinyint                                                       NOT NULL DEFAULT '0' COMMENT '超级管理员 1表示是 0:表示不是',
    `nickname`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `password`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `googleauth`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `salt`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `mobile`      char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci     NOT NULL DEFAULT '',
    `email`       char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci    NOT NULL DEFAULT '',
    `sex`         tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0:男 1:女',
    `age`         int                                                           NOT NULL DEFAULT 0,
    `level`       int                                                           NOT NULL DEFAULT 0,
    `avatar`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `info`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `status`      tinyint                                                       NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_nickname` (`nickname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
    `id`          bigint                                                        NOT NULL AUTO_INCREMENT,
    `create_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `status`      tinyint                                                       NOT NULL DEFAULT '1',
    `role_name`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `role_desc`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色表';

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`
(
    `id`          bigint   NOT NULL AUTO_INCREMENT,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `status`      tinyint  NOT NULL DEFAULT '1',
    `version`     bigint   NOT NULL DEFAULT '0' COMMENT '版本号',
    `user_id`     bigint   NOT NULL DEFAULT '0',
    `role_id`     bigint   NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户角色表';

DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`
(
    `id`          bigint        NOT NULL AUTO_INCREMENT,
    `create_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `status`      tinyint       NOT NULL DEFAULT '1',
    `info`        varchar(1000) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='权限详情表';

DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission`
(
    `id`            bigint   NOT NULL AUTO_INCREMENT,
    `create_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `status`        tinyint  NOT NULL DEFAULT '1',
    `role_id`       bigint   NOT NULL DEFAULT '0',
    `permission_id` bigint   NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色权限表';

SET
FOREIGN_KEY_CHECKS = 1;
