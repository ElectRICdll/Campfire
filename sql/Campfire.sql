-- Active: 1709392296239@@127.0.0.1@3306@campfire
DROP DATABASE Campfire;
CREATE DATABASE IF NOT EXISTS `Campfire`;
USE Campfire;
-- ---------------------------------------------------------------------------------------------------------------------

-- 用户信息表
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

-- ---------------------------------------------------------------------------------------------------------------------

-- Projects
create table projects
(
    project_id      int          not null        primary key AUTO_INCREMENT,
    title           varchar(32)  not null,
    description     text,
    leader          int          not null,
    foreign key     (leader)     REFERENCES      user_info(user_id),
    state           int          not null,
    files_url   varchar(255)    not null
);


-- ---------------------------------------------------------------------------------------------------------------------

-- Campsite
create table camp
(
    project_id      int          not null,
    foreign key     (project_id)    REFERENCES      projects(project_id),
    leader          int,
    foreign key     (leader)     REFERENCES      user_info(user_id),
    camp_id     int          not null        primary key AUTO_INCREMENT,
    name            varchar(32)    not null
);


-- ---------------------------------------------------------------------------------------------------------------------

-- Member
create table member
(
    user_id         int          not null,
    camp_id     int          not null,
    project_id     int          not null,
    foreign key     (user_id)       REFERENCES      user_info(user_id),
    foreign key     (camp_id)   REFERENCES      camp(camp_id),
    foreign key     (project_id)   REFERENCES      projects(project_id),
    nickname        varchar(32),
    designation     varchar(32),
    primary key     (`user_id`,`camp_id`)
);


-- -----------------------------------------------------------------------------------------

-- Message
create table message
(
    user_id         int          not null,
    project_id      int          not null,
    accept_id       int          not null,-- 对方id 如果是私聊消息才有意义
    camp_id     int          not null,-- 群聊id 如果是群聊消息才有意义
    foreign key     (user_id)       REFERENCES      user_info(user_id),
    foreign key     (project_id)    REFERENCES      projects(project_id),
    foreign key     (accept_id)     REFERENCES      user_info(user_id),
    foreign key     (camp_id)   REFERENCES      camp(camp_id),
    timestamp    datetime        not null,
    message_id      bigint          not null        primary key AUTO_INCREMENT,
    isPrivateChat   bool,-- 标记是否为私聊消息
    reply           int,
    content         text    not null
);


-- -----------------------------------------------------------------------------------------

-- Task
create table task
(
    launch_id         int          not null,
    accept_id         int          not null,
    project_id     int          not null,
    foreign key     (launch_id)     REFERENCES      user_info(user_id),
    foreign key     (accept_id)     REFERENCES      user_info(user_id),
    foreign key     (project_id)   REFERENCES      projects(project_id),
    task_id         int          not null        primary key AUTO_INCREMENT,
    title           varchar(32) not null,
    start_time      datetime        not null,
    end_time        datetime        not null,
    content         text         not null,
    state           int          not null
);


-- -----------------------------------------------------------------------------------------

-- Announcement
create table announcement
(
    user_id         int          not null,
    camp_id     int          not null,
    foreign key     (user_id)     REFERENCES        user_info(user_id),
    foreign key     (camp_id)   REFERENCES      camp(camp_id),
    announcement_id int          not null        primary key AUTO_INCREMENT,
    title           varchar(32) not null,
    start_time      datetime        not null,
    content         text    not null
)


-- -----------------------------------------------------------------------------------------