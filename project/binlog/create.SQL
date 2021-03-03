CREATE TABLE `test` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `a` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'a',
  `b` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'b',
  `c` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'c',
  `d` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'd',
  `e` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'e',
  `f` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'f',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '默认 1 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试表';