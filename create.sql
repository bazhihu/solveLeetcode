CREATE TABLE `test_test` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `a` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'a',
  `b` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'b',
  `c` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'c',
  `d` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'd',
  `e` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'e',
  `f` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'f',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试表';