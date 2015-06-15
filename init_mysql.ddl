
MySQL DB Honeycomb

CREATE DATABASE IF NOT EXISTS Honeycomb DEFAULT CHARSET utf8;


Table: user
存储用户名、密码、组别、APIKEY等信息，后续改为用email登录

CREATE TABLE `user` (
    `user_id` int unsigned NOT NULL AUTO_INCREMENT,
    `user_code` char(32) NOT NULL unique,               #计算生成的唯一识别符
    `company_code` char(32) NOT NULL ,                  #本人所属的公司
    `email` char(30) NOT NULL,                          #用户邮箱，用于登录
    `password` char(32) NOT NULL,                       #用户的密码
    `group` varchar(20) NOT NULL,                       #用户的组别，目前有系统管理员、统筹、八个部门的组
    `display_name` char(20) NOT NULL,                   #用于展示的名称
    `position` varchar(50) NOT NULL,                    #所在位置
    `picture` mediumtext NOT NULL,                      #头像照片的base64编码
    `phone` char(20) NOT NULL,                          #用户电话号码
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    INDEX(`email`),
    INDEX(`group`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


Table: email


Table: phone




Table: project
存储项目的相关信息

CREATE TABLE `project`(
    `project_id` int unsigned NOT NULL AUTO_INCREMENT,
    `project_code` char(32) NOT NULL,
    `project_name` char(50) NOT NULL,                  #存储项目的名称
    `project_detail` varchar(2000) NOT NULL,           #存储项目的详细说明
    #`plan_begin_datetime` datetime  NOT NULL,          #存储项目计划开始的时间
    #`plan_end_datetime` datetime NOT NULL,             #存储项目计划结束的时间
    #`real_begin_datetime` datetime NOT NULL,           #存储项目实际开始的时间
    #`real_end_datetime` datetime NOT NULL,             #存储项目实际结束的时间
    `person_in_charge` char(32) NOT NULL,              #存储`user_code`，项目负责人的usercode
    `company_code` char(32) NOT NULL ,                 #项目所属的公司
    `status` int default 0 NOT NULL,                   #0未开始，1已经完成,2进行中
    #`picture` mediumtext NOT NULL,                     #直接往mysql中写入照片的base64编码
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`project_id`),
    INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


Table: Mission Type

CREATE TABLE `missiontype` (
    `mission_type_id` int unsigned NOT NULL AUTO_INCREMENT,
    `mission_type_code` char(32) NOT NULL unique,
    `mission_type_name` char(50) NOT NULL,
    `mission_type_detail` varchar(200) NOT NULL,                #给任务一段更详细的文本说明
)


Table: mission
存储任务相关信息

CREATE TABLE `mission` (
    `mission_id` int unsigned NOT NULL AUTO_INCREMENT,
    `mission_code` char(32) NOT NULL unique,
    `mission_name` char(50) NOT NULL,                      #任务的名称
    `mission_type` char(32) NOT NULL,                      #存储mission_type_code，任务的类型
    `mission_detail` varchar(200) NOT NULL,                #给任务一段更详细的文本说明
    `project_code` char(32) NOT NULL,                      #任务所属project
    `has_child` tinyint default 0 NOT NULL,                #是否包含子任务，0不是，1是
    `parent_code` char(32) NOT NULL,                       #父级mission_code
    `child_index` int default 0 NOT NULL,                  #作为子任务，其所在父下的顺序
    `plan_begin_datetime` datetime NOT NULL,               #计划开始时间
    `plan_end_datetime` datetime NOT NULL,                 #计划结束时间
    `real_begin_datetime` datetime NOT NULL,               #实际开始时间
    `real_end_datetime` datetime NOT NULL,                 #实际结束时间
    `person_in_charge` char(32) NOT NULL,                  #存储`user_code`，任务的负责人
    `status` int default 0 NOT NULL,                       #0指定了人但未开始，1已经完成,2已经通过，3进行中，4未指定人
    `picture` mediumtext NOT NULL,                         #照片的base64编码
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`mission_id`),
    INDEX(`mission_code`),
    INDEX(`person_in_charge`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;





