
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
    `project_type` char(32) default "" NOT NULL,
    `project_name` char(50) NOT NULL,                          #存储项目的名称
    `project_detail` text default "" NOT NULL,                 #存储项目的详细说明
    `manager` char(32) default "" NOT NULL,                    #存储`user_code`，项目负责人的usercode
    `start_date` datetime  default CURRENT_TIMESTAMP,          #存储项目计划开始的时间
    `due_date` datetime default CURRENT_TIMESTAMP,             #存储项目计划结束的时间
    `real_start_date` datetime default CURRENT_TIMESTAMP,      #存储项目实际开始的时间
    `real_due_date` datetime default CURRENT_TIMESTAMP,        #存储项目实际结束的时间
    `status` int default 0 NOT NULL,                           #0未开始，1已经完成,2进行中
    `win_path` varchar default "" NOT NULL,
    `linux_path` varchar default "" NOT NULL,
    `macos_path` varchar default "" NOT NULL,
    `picture` mediumtext,                                      #直接往mysql中写入照片的base64编码
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
    `priority` int default 5,                              #优先级
    `degree_of_difficulty` int default 0,                  #难度系数
    `project_code` char(32) NOT NULL,                      #任务所属project
    `has_child` tinyint default 0 NOT NULL,                #是否包含子任务，0不是，1是
    `parent_code` char(32) NOT NULL,                       #父级mission_code
    `child_index` int default 0 NOT NULL,                  #作为子任务，其所在父下的顺序
    `start_date` datetime default CURRENT_TIMESTAMP,              #计划开始时间
    `due_date` datetime default CURRENT_TIMESTAMP,                #计划结束时间
    `real_start_date` datetime default CURRENT_TIMESTAMP,         #实际开始时间
    `real_due_date` datetime default CURRENT_TIMESTAMP,           #实际结束时间
    `assign_to` char(32) default "" NOT NULL,              #存储`user_code`，任务的负责人
    `status` int default 0 NOT NULL,                       #0指定了人但未开始
    `picture` mediumtext,                                  #照片的base64编码
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`mission_id`),
    INDEX(`mission_code`),
    INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;





