CREATE TABLE `docs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `hash` char(32) DEFAULT NULL COMMENT '文档绝对路径哈希值',
  `path` varchar(200) DEFAULT NULL COMMENT '访问绝对路径',
  `title` varchar(200) DEFAULT NULL COMMENT '标题',
  `desc` varchar(200) DEFAULT NULL COMMENT 'SEO 描述',
  `intro` varchar(200) DEFAULT NULL COMMENT '文章描述',
  `status` tinyint(1) DEFAULT '0' COMMENT '0 正常，1 禁用',
  `parent` varchar(20) DEFAULT '' COMMENT '文件夹，标签',
  `parent_hash` char(32) DEFAULT '' COMMENT '文件夹，标签',
  `img` varchar(200) DEFAULT '' COMMENT '配图',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `edited_at` int(11) DEFAULT NULL COMMENT '编辑时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 ;

CREATE TABLE `clicks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `hash` char(32) DEFAULT NULL COMMENT '文档绝对路径哈希值',
  `counter` int(11) DEFAULT NULL COMMENT '访问绝对路径',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ;

CREATE TABLE `artlogs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(200) NOT NULL COMMENT '客户端ip地址',
  `articleid` int(11) NOT NULL COMMENT '文章id',
  `sessid` varchar(200) NOT NULL COMMENT 'PHP session id',
  `created_at` int(11) NOT NULL COMMENT '添加时间',
  `useragent` text COMMENT '浏览器信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='文章浏览记录表';