-- Active: 1709392296239@@127.0.0.1@3306@campfire
CREATE DATABASE IF NOT EXISTS `Campfire`;

-- ---------------------------------------------------------------------------------------------------------------------

-- 用户信息表
DROP TABLE IF EXISTS `user_info`;
create table user_info
(
    user_id         int             not null        primary key AUTO_INCREMENT,
    email           varchar(32)                    not null,
    name            varchar(16)                    not null,
    password        varchar(16)                    not null,
    signature       text,
    avatar_url      varchar(255)                    not null
);
-- ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- Records of user_info
INSERT INTO `user_info`(`email`, `name`, `password`) VALUES ('1234567890@qq.com', 'sa', '123456');

-- ---------------------------------------------------------------------------------------------------------------------

-- Projects
DROP TABLE IF EXISTS `projects`;
create table projects
(
    project_id      int          not null        primary key AUTO_INCREMENT,
    leader          int          not null,
    foreign key     (leader)     REFERENCES      user_info(user_id),
    state           int          not null,
    codespace_url   varchar(255)    not null
);


-- ---------------------------------------------------------------------------------------------------------------------

-- Campsite
DROP TABLE IF EXISTS `campsite`;
create table campsite
(
    project_id      int          not null,     
    foreign key     (project_id)    REFERENCES      projects(project_id),
    leader          int          not null,
    foreign key     (leader)     REFERENCES      user_info(user_id),
    campsite_id     int          not null        primary key AUTO_INCREMENT,
    name            varchar(32)    not null
);


-- ---------------------------------------------------------------------------------------------------------------------

-- Member
DROP TABLE IF EXISTS `member`;
create table member
(
    user_id         int          not null,
    campsite_id     int          not null,
    project_id     int          not null,
    foreign key     (user_id)       REFERENCES      user_info(user_id),
    foreign key     (campsite_id)   REFERENCES      campsite(campsite_id),
    foreign key     (project_id)   REFERENCES      projects(project_id),
    nickname        varchar(32),
    designation     varchar(32),
    primary key     (`user_id`,`campsite_id`)
);


-- -----------------------------------------------------------------------------------------

-- Message
DROP TABLE IF EXISTS `message`;
create table message
(
    user_id         int          not null,
    project_id      int          not null,
    accept_id       int          not null,-- 对方id 如果是私聊消息才有意义
    campsite_id     int          not null,-- 群聊id 如果是群聊消息才有意义
    foreign key     (user_id)       REFERENCES      user_info(user_id),
    foreign key     (project_id)    REFERENCES      projects(project_id),
    foreign key     (accept_id)     REFERENCES      user_info(user_id),
    foreign key     (campsite_id)   REFERENCES      campsite(campsite_id),
    message_time    datetime        not null,
    message_id      bigint          not null        primary key AUTO_INCREMENT,
    privateChat     int,-- 标记是否为私聊消息
    reply           int,
    content         text    not null
);


-- -----------------------------------------------------------------------------------------

-- Task
DROP TABLE IF EXISTS `task`;
create table task
(
    launch_id         int          not null,
    accept_id         int          not null,
    project_id     int          not null,
    foreign key     (launch_id)     REFERENCES      user_info(user_id),
    foreign key     (accept_id)     REFERENCES      user_info(user_id),
    foreign key     (project_id)   REFERENCES      projects(project_id),
    task_id         int          not null        primary key AUTO_INCREMENT,
    start_time      datetime        not null,
    end_time        datetime        not null,
    content         text         not null,
    state           int          not null
);


-- -----------------------------------------------------------------------------------------

-- Announcement
DROP TABLE IF EXISTS `Announcement`;
create table announcement
(
    user_id         int          not null,
    campsite_id     int          not null,
    foreign key     (user_id)     REFERENCES        user_info(user_id),
    foreign key     (campsite_id)   REFERENCES      campsite(campsite_id),
    announcement_id int          not null        primary key AUTO_INCREMENT,
    start_time      datetime        not null,
    content         text    not null
)


-- -----------------------------------------------------------------------------------------