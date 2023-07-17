-- 创建数据库
CREATE DATABASE team_todo;

-- 使用数据库
USE team_todo;

-- 创建用户表
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
    email VARCHAR(30),
    group_ids VARCHAR(255),
    --group用字符串存储，逗号分隔，需要时进行解析
    avatar VARCHAR(255)
);

-- 创建群组表
CREATE TABLE groups (
    id INT PRIMARY KEY AUTO_INCREMENT,
    group_name VARCHAR(30),
    group_owner_id INT,
    group_description TEXT,
    member_ids VARCHAR(255),
    task_ids VARCHAR(255)
);

-- 创建任务表
CREATE TABLE tasks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    task_name VARCHAR(30),
    task_content TEXT,
    task_description TEXT,
    assignee VARCHAR(255),
    ddl DATETIME,
    status VARCHAR(10)
);

-- 创建提醒表
CREATE TABLE reminders (
    reminder_id INT PRIMARY KEY AUTO_INCREMENT,
    task_id INT,
    reminder_time DATETIME
);
