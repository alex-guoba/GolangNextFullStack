CREATE DATABASE blog_service;

CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '',
  `created_by` varchar(100) DEFAULT '' COMMENT '',
  `modified_by` varchar(100) DEFAULT '' COMMENT '',
  `ts_create` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `ts_update` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '',
  PRIMARY KEY (`id`)
);

CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  `created_by` varchar(100) DEFAULT '' COMMENT '',
  `modified_by` varchar(100) DEFAULT '' COMMENT '',
  `ts_create` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `ts_update` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '',
  PRIMARY KEY (`id`) USING BTREE
);

CREATE TABLE `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '',
  `created_by` varchar(100) DEFAULT '' COMMENT '',
  `modified_by` varchar(100) DEFAULT '' COMMENT '',
  `ts_create` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `ts_update` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '',
  PRIMARY KEY (`id`)
);

CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '',
  `desc` varchar(255) DEFAULT '' COMMENT '',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '',
  `content` longtext COMMENT '',
  `created_by` varchar(100) DEFAULT '' COMMENT '',
  `modified_by` varchar(100) DEFAULT '' COMMENT '',
  `ts_create` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `ts_update` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '',
  PRIMARY KEY (`id`)
) ;
