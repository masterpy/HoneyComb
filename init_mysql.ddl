
# MySQL DB Honeycomb

CREATE DATABASE IF NOT EXISTS Honeycomb DEFAULT CHARSET utf8;


# Table: user
# 存储用户名、密码、组别、APIKEY等信息，后续改为用email登录

CREATE TABLE `user` (
    `user_id` int unsigned NOT NULL AUTO_INCREMENT,
    `user_code` char(32) NOT NULL unique,
    `company_code` char(32) NOT NULL,
    `email` char(30) NOT NULL,
    `password` char(32) NOT NULL,
    `group` varchar(20) NOT NULL,
    `display_name` char(20) NOT NULL,
    `position` varchar(50) NOT NULL,
    `picture` mediumtext NOT NULL,
    `phone` char(20) NOT NULL,
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    INDEX(`email`),
    INDEX(`group`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


# Table: email


# Table: phone




# Table: project
# 存储项目的相关信息

CREATE TABLE `project`(
    `project_id` int unsigned NOT NULL AUTO_INCREMENT,
    `project_code` char(32) NOT NULL,
    `project_type` char(32) default '' NOT NULL,
    `project_name` char(50) NOT NULL,
    `project_detail` text NOT NULL,
    `manager` char(32) default '' NOT NULL,
    `start_date` datetime  default CURRENT_TIMESTAMP,
    `due_date` datetime default CURRENT_TIMESTAMP,
    `real_start_date` datetime default CURRENT_TIMESTAMP,
    `real_due_date` datetime default CURRENT_TIMESTAMP,
    `status` int default 0 NOT NULL,
    `win_path` varchar(255) default '' NOT NULL,
    `linux_path` varchar(255) default '' NOT NULL,
    `macos_path` varchar(255) default '' NOT NULL,
    `picture` mediumtext,
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`project_id`),
    INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


# Table: Mission Type

CREATE TABLE `missiontype` (
    `mission_type_id` int unsigned NOT NULL AUTO_INCREMENT,
    `mission_type_code` char(32) NOT NULL unique,
    `mission_type_name` char(50) NOT NULL,
    `mission_type_detail` varchar(200) NOT NULL,
)


# Table: mission
# 存储任务相关信息

CREATE TABLE `mission` (
    `mission_id` int unsigned NOT NULL AUTO_INCREMENT,
    `mission_code` char(32) NOT NULL unique,
    `mission_name` char(50) NOT NULL,                             #任务的名称
    `mission_type` char(32) NOT NULL,                             #存储mission_type_code，任务的类型
    `priority` int default 5,                                     #优先级
    `degree_of_difficulty` int default 0,                         #难度系数
    `project_code` char(32) NOT NULL,                             #任务所属project
    `has_child` tinyint default 0 NOT NULL,                       #是否包含子任务，0不是，1是
    `parent_code` char(32) NOT NULL,                              #父级mission_code
    `child_index` int default 0 NOT NULL,                         #作为子任务，其所在父下的顺序
    `start_date` datetime default CURRENT_TIMESTAMP,              #计划开始时间
    `due_date` datetime default CURRENT_TIMESTAMP,                #计划结束时间
    `real_start_date` datetime default CURRENT_TIMESTAMP,         #实际开始时间
    `real_due_date` datetime default CURRENT_TIMESTAMP,           #实际结束时间
    `assign_to` char(32) default "" NOT NULL,                     #存储`user_code`，任务的负责人
    `status` int default 0 NOT NULL,                              #0指定了人但未开始
    `picture` mediumtext,                                         #照片的base64编码
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`mission_id`),
    INDEX(`mission_code`),
    INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


# Table: mission_require
# 存储任务需求相关信息
CREATE TABLE `mission_require` (
    `require_id` int unsigned NOT NULL AUTO_INCREMENT,
    `require_code` char(32) NOT NULL unique,
    `mission_code` char(32) NOT NULL,
    `require_type` int default 0 NOT NULL,                             #需求的类型，0文本 1图片
    `content` text NOT NULL,
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`require_id`),
    INDEX(`require_code`),
    INDEX(`mission_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


# Table: mission_target
# 存储任务目标文件相关信息
CREATE TABLE `mission_target` (
    `target_id` int unsigned NOT NULL AUTO_INCREMENT,
    `target_code` char(32) NOT NULL unique,
    `mission_code` char(32) NOT NULL,
    `target_name` char(32) NOT NULL,
    `target_type` int default 0 NOT NULL,                             #需求的类型，0文本 1图片
    `comment` text NOT NULL,
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`target_id`),
    INDEX(`target_code`),
    INDEX(`mission_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




