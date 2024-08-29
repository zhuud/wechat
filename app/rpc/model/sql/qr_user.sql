CREATE TABLE user_service_qrcode (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `config_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '新增联系方式的配置id',
    `type` INT NOT NULL DEFAULT 1 COMMENT '联系方式类型，1-单人，2-多人',
    `scene` INT NOT NULL DEFAULT 1 COMMENT '场景，1-在小程序中联系，2-通过二维码联系',
    `style` INT NOT NULL DEFAULT 2 COMMENT '小程序中联系按钮的样式，仅在scene为1时返回，详见附录',
    `remark` VARCHAR(255) NOT NULL DEFAULT 'test remark' COMMENT '联系方式的备注信息，用于助记',
    `skip_verify` BOOLEAN NOT NULL DEFAULT true COMMENT '外部客户添加时是否无需验证',
    `state` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '企业自定义的state参数，用于区分不同的添加渠道',
    `qr_code` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '联系二维码的URL，仅在scene为2时返回',
    `user` VARCHAR(255) NOT NULL DEFAULT 'zhangsan,lisi,wangwu' COMMENT '使用该联系方式的用户userID列表',
    `party` VARCHAR(255) NOT NULL DEFAULT '2,3' COMMENT '使用该联系方式的部门id列表',
    `is_temp` tinyint(1) NOT NULL DEFAULT true COMMENT '是否临时会话模式0 不是 1 是',
    `expires_in` INT NOT NULL DEFAULT 86400 COMMENT '临时会话二维码有效期，以秒为单位',
    `chat_expires_in` INT NOT NULL DEFAULT 86400 COMMENT '临时会话有效期，以秒为单位',
    `unionid` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '可进行临时会话的客户unionid',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态 (0:删除,1:正常) | 2020-09-10',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间 | 2020-09-10',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间 | 2020-09-10',
    PRIMARY KEY (`id`)
)ENGINE = InnoDB COMMENT = '企业微信客服二维码信息| 2020-09-10';;

