CREATE TABLE `tb_user_service_qrcode` (
	`id` int(11)  NOT NULL AUTO_INCREMENT COMMENT '主键id | 2020-09-10',
	`config_id` varchar(255) NOT NULL DEFAULT '' COMMENT '新增联系方式的配置id | 2020-09-10',
	`type` int(11)  NOT NULL DEFAULT 1 COMMENT '联系方式类型(1:单人 / 2:多人) | 2020-09-10',
	`scene` int(11)  NOT NULL DEFAULT 1 COMMENT '场景(1:在小程序中联系 / 2:通过二维码联系) | 2020-09-10',
	`style` int (1)  NOT NULL DEFAULT 2 COMMENT '小程序中联系按钮的样式, 仅在scene为1时返回 | 2020-09-10',
	`remark` varchar(255) NOT NULL DEFAULT 'test remark' COMMENT '联系方式的备注信息用于助记 | 2020-09-10',
	`skip_verify` int (1) NOT NULL DEFAULT 1 COMMENT '外部客户添加时是否无需验证(0:否 / 1:是) | 2020-09-10',
	`state` varchar(255) NOT NULL DEFAULT '' COMMENT '企业自定义的state参数, 用于区分不同的添加渠道 | 2020-09-10',
	`qr_code` varchar(255) NOT NULL DEFAULT '' COMMENT '联系二维码的URL, 仅在scene为2时返回 | 2020-09-10',
	`user` varchar(255) NOT NULL DEFAULT '' COMMENT '使用该联系方式的用户userid, eg:zhangsan,lisi,wangwu | 2020-09-10',
	`party` varchar(255) NOT NULL DEFAULT '' COMMENT '使用该联系方式的部门id列表, 只在type为2时有效, eg:2,3 | 2020-09-10',
	`is_temp` int (1)  NOT NULL DEFAULT 0 COMMENT '是否临时会话模式(0:否 / 1:是) | 2020-09-10',
	`expires_in` int(11)  NOT NULL DEFAULT 86400 COMMENT '临时会话二维码有效期, 以秒为单位 | 2020-09-10',
	`chat_expires_in` int(11)  NOT NULL DEFAULT 86400 COMMENT '临时会话有效期, 以秒为单位 | 2020-09-10',
	`unionid` varchar(255) NOT NULL DEFAULT '' COMMENT '可进行临时会话的客户unionid | 2020-09-10',
	`is_exclusive` int (1)  NOT NULL DEFAULT 0 COMMENT '开启后同一个企业的客户会优先添加到同一个跟进人(0:否 / 1:是) | 2020-09-10',
	`status` int (3)  NOT NULL DEFAULT '1' COMMENT '状态 (0:删除 / 1:正常) | 2020-09-10',
	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间 | 2020-09-10',
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间 | 2020-09-10',
	PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '企业微信客服二维码信息| 2020-09-10';

CREATE TABLE `tb_user_service_qrcode_conclusions` (
	`id` int(11)  NOT NULL AUTO_INCREMENT COMMENT '主键id',
	`user_service_qc_code_id` int(11)  NOT NULL DEFAULT 0 COMMENT '企业微信客服二维码信息表主键ID',
	`type` varchar(50) NOT NULL DEFAULT '' COMMENT '结束语类型 (text:文本 / image:图片 / link:图文 / miniprogram:小程序)',
	`content` varchar(2000) NOT NULL DEFAULT '' COMMENT '结束语内容, json字符串',
	`status` int (3)  NOT NULL DEFAULT '1' COMMENT '状态 (0:删除 / 1:正常)',
	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '企业微信客服二维码信息-临时会话结束语，会话结束时自动发送给客户，仅在is_temp为true时有效';