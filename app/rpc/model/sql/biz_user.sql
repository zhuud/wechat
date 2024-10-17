CREATE TABLE `tb_user_open` (
  `open_site` int(10)  NOT NULL DEFAULT '0' COMMENT '固定值3',
  `open_id` varchar(50) NOT NULL DEFAULT '' COMMENT '微信unionid',
  `uid` int(10)  NOT NULL DEFAULT '0' COMMENT '业务uid',
  PRIMARY KEY (`open_id`)
) ENGINE = InnoDB;