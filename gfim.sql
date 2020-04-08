-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- 主机： mysql:3306
-- 生成日期： 2020-04-08 09:54:58
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

--
-- 转存表中的数据 `gf_friend`
--

INSERT INTO `gf_friend` (`id`, `user_id`, `friend_id`, `friend_group_id`, `create_time`, `remark`) VALUES
(1, 1, 2, 1, 0, '11'),
(2, 2, 1, 1, 0, '11'),
(3, 2, 3, 1, 0, '11'),
(4, 1, 3, 1, 0, '11'),
(5, 3, 2, 1, 0, '11'),
(6, 3, 1, 1, 0, '11');

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
(2, 1, '同学', 0, 0);

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

--
-- 转存表中的数据 `gf_group`
--

INSERT INTO `gf_group` (`id`, `name`, `avatar`, `create_time`, `update_time`, `create_user_id`, `user_id`) VALUES
(1, 'go学习群', 'https://zhiyu.web.xmchuangyi.com/static/img/no_img.jpg', 0, 0, 1, 1);

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

--
-- 转存表中的数据 `gf_group_record`
--

INSERT INTO `gf_group_record` (`id`, `user_id`, `group_id`, `content`, `create_time`, `delete_time`) VALUES
(1, 1, 1, '123', 1585818916, 0),
(2, 1, 1, '2123', 1585819000, 0),
(3, 1, 1, '出来····', 1585819053, 0),
(4, 1, 1, '123', 1585819079, 0),
(5, 1, 1, '312', 1585819086, 0),
(6, 1, 1, '321', 1585819093, 0),
(7, 1, 1, 'res', 1585819256, 0),
(8, 1, 1, 'zzz', 1585819704, 0),
(9, 1, 1, 'asasd', 1585820063, 0),
(10, 1, 1, 'lai ren ', 1585820486, 0),
(11, 1, 1, '三大类', 1585820532, 0),
(12, 2, 1, '22', 1585820564, 0),
(13, 2, 1, '222', 1585820573, 0),
(14, 1, 1, 'as', 1585820669, 0),
(15, 1, 1, 'zcx', 1585820701, 0),
(16, 2, 1, 'asdasd', 1585820707, 0),
(17, 2, 1, '321', 1585820721, 0),
(18, 2, 1, '321', 1585820745, 0),
(19, 1, 1, 'ok', 1585820829, 0),
(20, 1, 1, 'ok', 1585820839, 0),
(21, 1, 1, 'ok', 1585820875, 0),
(22, 1, 1, '123', 1585820878, 0),
(23, 2, 1, '321', 1585820890, 0),
(24, 2, 1, '43', 1585820895, 0),
(25, 1, 1, '我知道了', 1585820942, 0),
(26, 2, 1, '完美解决', 1585820946, 0),
(27, 1, 1, 'naskl', 1585821022, 0),
(28, 3, 1, 'sda', 1585821026, 0),
(29, 3, 1, '321', 1585821041, 0),
(30, 2, 1, '123', 1585821043, 0),
(31, 1, 1, '312312', 1585821053, 0);

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

--
-- 转存表中的数据 `gf_group_user`
--

INSERT INTO `gf_group_user` (`id`, `group_id`, `user_id`, `nickname`, `create_time`, `update_time`) VALUES
(1, 1, 1, 'aa', 0, 0),
(2, 1, 2, 'bb', 0, 0),
(3, 1, 3, 'bb', 0, 0);

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

--
-- 转存表中的数据 `gf_user`
--

INSERT INTO `gf_user` (`id`, `username`, `nickname`, `password`, `salt`, `email`, `mobile`, `avatar`, `gender`, `birthday`, `sign`, `prev_time`, `login_time`, `login_ip`, `join_ip`, `join_time`, `create_time`, `update_time`, `status`, `im_status`) VALUES
(1, 'admin', 'hhhjx', 'c13f62012fd6a8fdf06b3452a94430e5', 'rpR6Bv', 'admin@163.com', '13888888888', 'http://tva3.sinaimg.cn/crop.0.0.512.512.180/8693225ajw8f2rt20ptykj20e80e8weu.jpg', 0, '2017-04-15', '没有....11122', 1585820689, 1586324230, '172.20.0.1', '127.0.0.1', 1491461418, 0, 1586324230, 'normal', 'online'),
(2, 'admin2', 'admin2', 'c13f62012fd6a8fdf06b3452a94430e5', 'rpR6Bv', 'admin@163.com', '13888882888', 'http://tva3.sinaimg.cn/crop.0.0.512.512.180/8693225ajw8f2rt20ptykj20e80e8weu.jpg', 0, '2017-04-15', '坚持学习666', 1585820689, 1586324230, '172.20.0.1', '127.0.0.1', 1491461418, 0, 1586324230, 'normal', 'offline'),
(3, 'echome', 'echome', 'c13f62012fd6a8fdf06b3452a94430e5', 'rpR6Bv', 'admin@163.com', '13888882888', 'http://tva3.sinaimg.cn/crop.0.0.512.512.180/8693225ajw8f2rt20ptykj20e80e8weu.jpg', 0, '2017-04-15', '1111', 1585820689, 1586324230, '172.20.0.1', '127.0.0.1', 1491461418, 0, 1586324230, 'normal', 'offline');

-- --------------------------------------------------------

--
-- 表的结构 `gf_user_record`
--

CREATE TABLE `gf_user_record` (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'ID',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发送人',
  `friend_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '好友ID',
  `content` varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '聊天内容',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_read` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已读;1:是,0:否'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户聊天记录';

--
-- 转存表中的数据 `gf_user_record`
--

INSERT INTO `gf_user_record` (`id`, `user_id`, `friend_id`, `content`, `create_time`, `delete_time`, `is_read`) VALUES
(1, 1, 2, '你说呢', 1585729065, 0, 0),
(2, 1, 2, '1123232', 1585729301, 0, 0),
(3, 2, 1, 'zzz', 1585729363, 0, 0),
(4, 2, 1, '222', 1585730294, 0, 0),
(5, 1, 2, '3333', 1585730315, 0, 0),
(6, 3, 2, '!!!!', 1585730603, 0, 0),
(7, 2, 3, 'dsd', 1585730612, 0, 0),
(8, 1, 2, '你说呢', 1585815714, 0, 0),
(9, 2, 1, '!!', 1585816200, 0, 0),
(10, 1, 2, '22', 1585817731, 0, 0),
(11, 1, 1, '222', 1585818718, 0, 0),
(12, 1, 3, 'zxczcx', 1585884334, 0, 0),
(13, 2, 1, '123', 1586324238, 0, 0),
(14, 1, 2, '321', 1586324242, 0, 0),
(15, 1, 2, 'asd', 1586324246, 0, 0);

-- --------------------------------------------------------

--
-- 表的结构 `gf_user_token`
--

CREATE TABLE `gf_user_token` (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'ID',
  `token` varchar(50) NOT NULL COMMENT 'Token',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '会员ID',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `expire_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '过期时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会员Token表' ROW_FORMAT=COMPACT;

--
-- 转存表中的数据 `gf_user_token`
--

INSERT INTO `gf_user_token` (`id`, `token`, `user_id`, `create_time`, `expire_time`) VALUES
(1, '634e3a487fd745a6b69c1e2de5a80376', 1, 1584951101, 1587543101),
(2, 'cb13646f474bfcb222d6e4374ccdc835', 1, 1584951412, 1587543412),
(3, 'a2a69c254ccd3942c6b74c017fe96f3c', 1, 1584951417, 1587543417),
(4, '1aad10efa613235320ee67c8889933be', 1, 1585056969, 1587648969),
(5, '5ffab5dfcf5928abdb0ee9f7b4a559e3', 1, 1585057015, 1587649015),
(6, '57be7e8149d437b96297b7b744ed7ab0', 1, 1585098973, 1587690973),
(7, '38e7430b827168b3932b545a50e6a64a', 1, 1585128029, 1587720029),
(8, 'e7f791e8c895defa21e5269b4021945c', 1, 1585128029, 1587720029),
(9, 'e4a28408cb75b28f17eb201c339ec8e2', 2, 1585558191, 1588150191),
(10, '1768164d2d6a2dec5b004fdbe57c569e', 2, 1585617094, 1588209094),
(11, '57464c4383a27fc3ebedaa11e50e788f', 2, 1585645761, 1588237761),
(12, 'd8584e50304ac023f2429f7ad7bf5eef', 2, 1585712345, 1588304345),
(13, '243bdc08f5fe8fec3e2d3657686e7479', 3, 1585712365, 1588304365),
(14, '7b21ae565395be24218da3c0eef6c872', 2, 1585727076, 1588319076),
(15, '7d5f93dc0a52682df2773dcb195725c6', 2, 1585816194, 1588408194),
(16, '3d7e8b9bec0fd30c8d859dbd97046d81', 2, 1585820689, 1588412689),
(17, '506a5b2d87ae6e43c504893ddc0e5771', 2, 1586324230, 1588916230);

--
-- 转储表的索引
--

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
-- 使用表AUTO_INCREMENT `gf_friend`
--
ALTER TABLE `gf_friend`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID', AUTO_INCREMENT=7;

--
-- 使用表AUTO_INCREMENT `gf_friend_group`
--
ALTER TABLE `gf_friend_group`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `gf_group`
--
ALTER TABLE `gf_group`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `gf_group_record`
--
ALTER TABLE `gf_group_record`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID', AUTO_INCREMENT=32;

--
-- 使用表AUTO_INCREMENT `gf_group_user`
--
ALTER TABLE `gf_group_user`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `gf_user`
--
ALTER TABLE `gf_user`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `gf_user_record`
--
ALTER TABLE `gf_user_record`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID', AUTO_INCREMENT=16;

--
-- 使用表AUTO_INCREMENT `gf_user_token`
--
ALTER TABLE `gf_user_token`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID', AUTO_INCREMENT=18;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
