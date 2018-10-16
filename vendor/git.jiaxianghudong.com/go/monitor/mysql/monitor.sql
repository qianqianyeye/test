# Host: 192.168.67.83  (Version 5.7.18-log)
# Date: 2018-05-10 15:46:23
# Generator: MySQL-Front 6.0  (Build 2.20)


#
# Structure for table "t_trace"
#

DROP TABLE IF EXISTS `t_trace`;
CREATE TABLE `t_trace` (
  `trace_id` bigint(13) NOT NULL DEFAULT '0' COMMENT '追踪链id',
  `span_id` bigint(13) NOT NULL DEFAULT '0' COMMENT '调用id',
  `p_span_id` bigint(13) NOT NULL DEFAULT '0' COMMENT '父级调用id',
  `func_id` int(11) DEFAULT '0' COMMENT '功能id',
  `s_addr` varchar(30) NOT NULL DEFAULT '' COMMENT '源地址',
  `d_addr` varchar(30) DEFAULT NULL COMMENT '目标地址',
  `data_size` int(11) NOT NULL DEFAULT '0' COMMENT '数据包大小',
  `req_datetime` datetime DEFAULT NULL COMMENT '发出请求时间',
  `spend_times` int(11) NOT NULL DEFAULT '0' COMMENT '消耗时间（毫秒）',
  `rsp_code` int(11) DEFAULT NULL COMMENT '响应码',
  `rsp_msg` varchar(255) DEFAULT NULL COMMENT '响应消息',
  PRIMARY KEY (`trace_id`,`span_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Structure for table "wuid"
#

DROP TABLE IF EXISTS `wuid`;
CREATE TABLE `wuid` (
  `h` int(10) NOT NULL AUTO_INCREMENT,
  `x` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`x`),
  UNIQUE KEY `h` (`h`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=latin1;
