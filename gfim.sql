-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- 主机： mysql:3306
-- 生成日期： 2020-04-23 15:09:53
-- 服务器版本： 8.0.13
-- PHP 版本： 7.2.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `gfim`
--

-- --------------------------------------------------------

--
-- 表的结构 `gf_apply`
--

CREATE TABLE `gf_apply` (
  `id` int(10) UNSIGNED NOT NULL,
  `type` enum('friend','group') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'friend' COMMENT '类型，friend:好友申请,group:群组申请',
  `from_user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发起人',
  `to_user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '接收人',
  `target_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '目标ID，好友申请时为好友分组ID，群组申请时为群组ID',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '验证信息',
  `state` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,1:待处理,2:已同意,3:已拒绝,4:已忽略',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `handle_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '处理时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='申请表';

-- --------------------------------------------------------

--
-- 表的结构 `gf_apply_remind`
--

CREATE TABLE `gf_apply_remind` (
  `id` int(10) UNSIGNED NOT NULL,
  `apply_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '申请ID',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
  `is_read` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已读,1是0否',
  `read_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '已读时间',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='申请提醒';

-- --------------------------------------------------------

--
-- 表的结构 `gf_friend`
--

CREATE TABLE `gf_friend` (
  `id` int(11) UNSIGNED NOT NULL COMMENT 'ID',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
  `friend_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '好友ID',
  `friend_group_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '好友分组ID',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '好友备注'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户朋友表';

-- --------------------------------------------------------

--
-- 表的结构 `gf_friend_group`
--

CREATE TABLE `gf_friend_group` (
  `id` int(11) UNSIGNED NOT NULL COMMENT 'ID',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID,0:系统默认',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '好友分组名称',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='好友组';

--
-- 转存表中的数据 `gf_friend_group`
--

INSERT INTO `gf_friend_group` (`id`, `user_id`, `name`, `create_time`, `update_time`) VALUES
(1, 0, '我的好友', 0, 0),
(2, 1, '同学', 0, 0),
(3, 6, '同学123', 0, 0);

-- --------------------------------------------------------

--
-- 表的结构 `gf_group`
--

CREATE TABLE `gf_group` (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '群名',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `create_user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创始人ID',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '所属用户ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='聊天群';

-- --------------------------------------------------------

--
-- 表的结构 `gf_group_record`
--

CREATE TABLE `gf_group_record` (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'ID',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发送人',
  `group_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '群ID',
  `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '聊天内容',
  `create_time` int(10) UNSIGNED NOT NULL COMMENT '创建时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='群聊天记录';

-- --------------------------------------------------------

--
-- 表的结构 `gf_group_user`
--

CREATE TABLE `gf_group_user` (
  `id` int(11) UNSIGNED NOT NULL COMMENT 'ID',
  `group_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '群ID',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
  `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户群昵称',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='群用户表';

-- --------------------------------------------------------

--
-- 表的结构 `gf_user`
--

CREATE TABLE `gf_user` (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'ID',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(30) NOT NULL DEFAULT '' COMMENT '密码盐',
  `email` varchar(100) NOT NULL DEFAULT '' COMMENT '电子邮箱',
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '性别',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `sign` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '签名',
  `prev_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '上次登录时间',
  `login_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '登录时间',
  `login_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登录IP',
  `join_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '加入IP',
  `join_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '加入时间',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `status` enum('normal','disable') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'normal' COMMENT '用户状态:normal:正常,disable:禁用',
  `im_status` enum('offline','online','hide') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'offline' COMMENT 'im状态;offlie:离线,online:在线,hide:隐身'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会员表' ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `gf_user_record`
--

CREATE TABLE `gf_user_record` (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'ID',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发送人',
  `friend_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '好友ID',
  `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '聊天内容',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_read` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已读;1:是,0:否',
  `is_notify` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已通知 1是0否'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户聊天记录';

-- --------------------------------------------------------

--
-- 表的结构 `gf_user_token`
--

CREATE TABLE `gf_user_token` (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'ID',
  `token` varchar(50) NOT NULL COMMENT 'Token',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '会员ID',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `expire_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '过期时间',
  `is_valid` tinyint(3) UNSIGNED DEFAULT '1' COMMENT '是否有效 1有效 0无效'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会员Token表' ROW_FORMAT=COMPACT;

--
-- 转储表的索引
--

--
-- 表的索引 `gf_apply`
--
ALTER TABLE `gf_apply`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `gf_apply_remind`
--
ALTER TABLE `gf_apply_remind`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `gf_friend`
--
ALTER TABLE `gf_friend`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `gf_friend_group`
--
ALTER TABLE `gf_friend_group`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `gf_group`
--
ALTER TABLE `gf_group`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `gf_group_record`
--
ALTER TABLE `gf_group_record`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `gf_group_user`
--
ALTER TABLE `gf_group_user`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `gf_user`
--
ALTER TABLE `gf_user`
  ADD PRIMARY KEY (`id`),
  ADD KEY `username` (`username`),
  ADD KEY `email` (`email`),
  ADD KEY `mobile` (`mobile`);

--
-- 表的索引 `gf_user_record`
--
ALTER TABLE `gf_user_record`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `gf_user_token`
--
ALTER TABLE `gf_user_token`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `gf_apply`
--
ALTER TABLE `gf_apply`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `gf_apply_remind`
--
ALTER TABLE `gf_apply_remind`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `gf_friend`
--
ALTER TABLE `gf_friend`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `gf_friend_group`
--
ALTER TABLE `gf_friend_group`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `gf_group`
--
ALTER TABLE `gf_group`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `gf_group_record`
--
ALTER TABLE `gf_group_record`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `gf_group_user`
--
ALTER TABLE `gf_group_user`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `gf_user`
--
ALTER TABLE `gf_user`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `gf_user_record`
--
ALTER TABLE `gf_user_record`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `gf_user_token`
--
ALTER TABLE `gf_user_token`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID';
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
