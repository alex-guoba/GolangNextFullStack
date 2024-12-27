CREATE TABLE `blog_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '',
  `desc` varchar(255) DEFAULT '',
  `cover_image_url` varchar(255) DEFAULT '',
  `content` longtext,
  `created_by` varchar(100) DEFAULT '',
  `modified_by` varchar(100) DEFAULT '',
  `ts_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `ts_update` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint unsigned DEFAULT '0' COMMENT ' 01',
  `state` tinyint unsigned DEFAULT '1' COMMENT ' 01',
  PRIMARY KEY (`id`)
);

CREATE TABLE `blog_article_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int NOT NULL COMMENT 'ID',
  `tag_id` int unsigned NOT NULL DEFAULT '0' COMMENT 'ID',
  `created_by` varchar(100) DEFAULT '',
  `modified_by` varchar(100) DEFAULT '',
  `ts_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `ts_update` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint unsigned DEFAULT '0' COMMENT ' 01',
  PRIMARY KEY (`id`)
);

CREATE TABLE `blog_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '',
  `created_by` varchar(100) DEFAULT '',
  `modified_by` varchar(100) DEFAULT '',
  `ts_create` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `ts_update` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_del` tinyint unsigned DEFAULT '0' COMMENT ' 01',
  `state` tinyint unsigned DEFAULT '1' COMMENT ' 01',
  PRIMARY KEY (`id`)
);