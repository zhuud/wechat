CREATE TABLE `tb_user_open` (
  `open_site` int(10)  NOT NULL DEFAULT '0' COMMENT '固定值3',
  `open_id` varchar(50) NOT NULL DEFAULT '' COMMENT '微信unionid',
  `uid` int(10)  NOT NULL DEFAULT '0' COMMENT '业务uid',
  PRIMARY KEY (`open_id`)
) ENGINE = InnoDB;

CREATE TABLE `tb_private_domain_user` (
  `from` int(10)  NOT NULL DEFAULT '0' COMMENT '企微平台',
  `qywx_user_id` varchar(50) NOT NULL DEFAULT '' COMMENT '外部联系人的userid',
  `external_userid` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人的userid',
  `add_time` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '添加时间',
  `chat_agree_status` int(10)  NOT NULL DEFAULT '0' COMMENT '会话存档状态 (0:不同意 / 1:同意)',
  `last_chat_time` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '最近沟通时间',
  `blacklist_time` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '拉黑时间',
  `blacklist_type` varchar(50) NOT NULL DEFAULT '' COMMENT '拉黑删除类型：external_user|staff',
  `status` int(10)  NOT NULL DEFAULT '1' COMMENT '拉黑删除类型：external_user|staff',
  PRIMARY KEY (`external_userid`)
) ENGINE = InnoDB;