CREATE TABLE `tb_external_user` (
  `external_userid` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人的userid | 2020-09-10',
  `unionid` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人在微信开放平台的唯一身份标识（联系人类型是微信用户且企业绑定了微信开发者ID有此字段 第三方应用和代开发应用均不可获取 上游企业不可获取下游企业客户该字段） | 2020-09-10',
  `type` int(3)  NOT NULL DEFAULT '0' COMMENT '外部联系人的类型 (1:微信用户 / 2:企业微信用户) | 2020-09-10',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人的名称(微信用户返回其微信昵称 企业微信联系人返回其设置对外展示的别名或实名) | 2020-09-10',
  `avatar` varchar(100) NOT NULL DEFAULT '' COMMENT '外部联系人头像(代开发自建应用需要管理员授权才可以获取 第三方不可获取 上游企业不可获取下游企业客户该字段) | 2020-09-10',
  `gender` int(3)  NOT NULL DEFAULT '0' COMMENT '外部联系人性别 (0:未知 / 1:男性 / 2:女性)(第三方不可获取 上游企业不可获取下游企业客户该字段 返回值为0) | 2020-09-10',
  `corp_name` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人所在企业的简称(仅当联系人类型是企业微信用户时有此字段) | 2020-09-10',
  `corp_full_name` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人所在企业的主体名称(仅当联系人类型是企业微信用户时有此字段) | 2020-09-10',
  `position` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人的职位(仅当联系人类型是企业微信用户时有此字段) | 2020-09-10',
  `status` int(3) NOT NULL DEFAULT '1' COMMENT '状态 (0:删除 / 1:正常) | 2020-09-10',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间 | 2020-09-10',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间 | 2020-09-10',
  PRIMARY KEY (`external_userid`),
  KEY `idx_unionid` (`unionid`)
) ENGINE = InnoDB COMMENT = '外部联系人信息表 | 2020-09-10';
CREATE TABLE `tb_external_user_attribute` (
  `seq` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 | 2020-09-10',
  `external_userid` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人的userid | 2020-09-10',
  `attribute_type` int(3) NOT NULL DEFAULT '0' COMMENT '类型 (0:文本 / 1:网页 / 2:小程序) | 2020-09-10',
  `attribute_value` varchar(200) NOT NULL DEFAULT '' COMMENT '类型值 | 2020-09-10',
  `extension` varchar(500) NOT NULL DEFAULT '' COMMENT '扩展信息 | 2020-09-10',
  `status` int(3) NOT NULL DEFAULT '1' COMMENT '状态 (0:删除 / 1:正常) | 2020-09-10',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间 | 2020-09-10',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间 | 2020-09-10',
  PRIMARY KEY (`seq`),
  KEY `idx_external_userid` (`external_userid`)
) ENGINE = InnoDB COMMENT = '外部联系人属性表(attribute_type为业务枚举) | 2020-09-10';
CREATE TABLE `tb_external_user_follow` (
  `seq` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 | 2020-09-10',
  `external_userid` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人的userid | 2020-09-10',
  `unionid` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人在微信开放平台的唯一身份标识（联系人类型是微信用户且企业绑定了微信开发者ID有此字段 第三方应用和代开发应用均不可获取 上游企业不可获取下游企业客户该字段） | 2020-09-10',
  `userid` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人的userid | 2020-09-10',
  `crop` varchar(50) NOT NULL DEFAULT '' COMMENT '企微平台(多企微情况) | 2020-09-10',
  `oper_userid` varchar(50) NOT NULL DEFAULT '' COMMENT '发起添加的userid(成员主动添加为成员的userid 客户主动添加为客户的外部联系人userid 内部成员共享/管理员分配为对应的成员/管理员userid) | 2020-09-10',
  `add_way` int(11) NOT NULL DEFAULT 0 COMMENT '添加外部联系人的方式(https://developer.work.weixin.qq.com/document/path/92114#%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89) | 2020-09-10',
  `state` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人添加渠道 | 2020-09-10',
  `state_channel` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人添加渠道 | 2020-09-10',
  `state_channel_value` varchar(500) NOT NULL DEFAULT '' COMMENT '外部联系人添加渠道透传参数 | 2020-09-10',
  `remark` varchar(50) NOT NULL DEFAULT '' COMMENT '对外部联系人的备注 | 2020-09-10',
  `remark_mobiles` varchar(500) NOT NULL DEFAULT '[]' COMMENT '备注的手机号码 | 2020-09-10',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '对外部联系人的描述 | 2020-09-10',
  `remark_corp_name` varchar(50) NOT NULL DEFAULT '' COMMENT '对外部联系人备注的所属公司名称 | 2020-09-10',
  `remark_pic_mediaid` varchar(200) NOT NULL DEFAULT '' COMMENT '对外部联系人备注的图片ID | 2020-09-10',
  `chat_agree_status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '会话存档状态 (0:不同意 / 1:同意)',
  `last_chat_time` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '最近沟通时间 | 2020-09-10',
  `status` int(3) NOT NULL DEFAULT 1 COMMENT '状态 (0:删除 / 1:正常) | 2020-09-10',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间 | 2020-09-10',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间 | 2020-09-10',
  `deleted_at` datetime NOT NULL DEFAULT '0001-01-01 00:00:00' COMMENT '删除时间 | 2020-09-10',
  PRIMARY KEY (`seq`),
  KEY `idx_unionid` (`unionid`),
  KEY `idx_external_userid` (`external_userid`),
  KEY `idx_userid` (`userid`)
) ENGINE = InnoDB COMMENT = '外部联系人属性表(attribute_type为业务枚举) | 2020-09-10';
CREATE TABLE `tb_external_user_follow_attribute` (
  `seq` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 | 2020-09-10',
  `external_userid` varchar(100) NOT NULL DEFAULT '' COMMENT '外部联系人的userid | 2020-09-10',
  `userid` varchar(100) NOT NULL DEFAULT '' COMMENT '联系人的userid | 2020-09-10',
  `crop` varchar(50) NOT NULL DEFAULT '' COMMENT '企微平台(多企微情况) | 2020-09-10',
  `attribute_type` int(3) NOT NULL DEFAULT '0' COMMENT '类型 (1:备注标签 / 2:视频号信息) | 2020-09-10',
  `attribute_value` varchar(200) NOT NULL DEFAULT '' COMMENT '类型值 | 2020-09-10',
  `extension` varchar(500) NOT NULL DEFAULT '' COMMENT '扩展信息 | 2020-09-10',
  `status` int(3) NOT NULL DEFAULT '1' COMMENT '状态 (0:删除 / 1:正常) | 2020-09-10',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间 | 2020-09-10',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间 | 2020-09-10',
PRIMARY KEY (`seq`),
KEY `idx_userid_attribute_type` (`userid`, `attribute_type`),
KEY `idx_external_userid_attribute_type` (`external_userid`, `attribute_type`),
KEY `idx_attribute_type_value` (`attribute_type`, `attribute_value`)
) ENGINE = InnoDB COMMENT = '外部联系人添加员工信息属性表(attribute_type为业务枚举) | 2020-09-10';
CREATE TABLE `tb_external_user_tag` (
  `tag_id` varchar(50) NOT NULL DEFAULT '' COMMENT '标签id | 2020-09-10',
  `group_id` varchar(50) NOT NULL DEFAULT '' COMMENT '标签组id | 2020-09-10',
  `group_name` varchar(50) NOT NULL DEFAULT '' COMMENT '标签组名字 | 2020-09-10',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '标签名字 | 2020-09-10',
  `weight` int(11) NOT NULL DEFAULT '0' COMMENT '排序的次序值，order值大的排序靠前 微信为order关键字 | 2020-09-10',
  `status` int(3) NOT NULL DEFAULT '1' COMMENT '状态 (0:删除 / 1:正常) | 2020-09-10',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间 | 2020-09-10',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间 | 2020-09-10',
  PRIMARY KEY (`tag_id`),
  KEY `idx_group_id` (`group_id`),
  KEY `idx_weight` (`weight`)
) ENGINE = InnoDB COMMENT = '外部联系人进行标记和分类的标签 | 2020-09-10';
CREATE TABLE `tb_user` (
  `userid` varchar(50) NOT NULL DEFAULT '' COMMENT '员工的userid | 2020-09-10',
  `unionid` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人在微信开放平台的唯一身份标识（联系人类型是微信用户且企业绑定了微信开发者ID有此字段 第三方应用和代开发应用均不可获取 上游企业不可获取下游企业客户该字段） | 2020-09-10',
  `type` int(3) NOT NULL DEFAULT '0' COMMENT '外部联系人的类型 (1:微信用户 / 2:企业微信用户) | 2020-09-10',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人的名称(微信用户返回其微信昵称 企业微信联系人返回其设置对外展示的别名或实名) | 2020-09-10',
  `avatar` varchar(100) NOT NULL DEFAULT '' COMMENT '外部联系人头像(代开发自建应用需要管理员授权才可以获取 第三方不可获取 上游企业不可获取下游企业客户该字段) | 2020-09-10',
  `gender` int(3) NOT NULL DEFAULT '0' COMMENT '外部联系人性别 (0:未知 / 1:男性 / 2:女性)(第三方不可获取 上游企业不可获取下游企业客户该字段 返回值为0) | 2020-09-10',
  `corp_name` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人所在企业的简称(仅当联系人类型是企业微信用户时有此字段) | 2020-09-10',
  `corp_full_name` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人所在企业的主体名称(仅当联系人类型是企业微信用户时有此字段) | 2020-09-10',
  `position` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人的职位(仅当联系人类型是企业微信用户时有此字段) | 2020-09-10',
  `status` int(3) NOT NULL DEFAULT '1' COMMENT '状态 (0:删除 / 1:正常) | 2020-09-10',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间 | 2020-09-10',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间 | 2020-09-10',
  PRIMARY KEY (`userid`),
  KEY `idx_unionid` (`unionid`)
) ENGINE = InnoDB COMMENT = '外部联系人属性表(attribute_type为业务枚举) | 2020-09-10';
CREATE TABLE `tb_operation_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 | 柴利瑶 | 2020-05-15',
  `related_id` int(11) NOT NULL DEFAULT '0' COMMENT '关联表的id |柴利瑶 | 2020-06-15 ',
  `admin_id` int(11) NOT NULL DEFAULT '0' COMMENT '操作人id | 柴利瑶 | 2020-05-15',
  `type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '操作类型 | 柴利瑶 | 2020-05-15',
  `log_type` int(11) NOT NULL DEFAULT '0' COMMENT '日志详情的类型 | 柴利瑶 | 2020-05-15',
  `log_desc` text  COMMENT '日志内容 | 柴利瑶 |2020-05-15',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_type_related_id` (`type`, `related_id`)
) ENGINE = InnoDB COMMENT = '操作记录表| 柴利瑶 | 2020-05-13'