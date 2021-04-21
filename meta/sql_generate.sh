#!/bin/bash

# Don't exceed 16
URL_TABLECNT=1
ACCESS_TABLECNT=1

function generateShorten() {
    local tableCnt=$(expr $URL_TABLECNT - 1)
    for i in $(seq 0 $tableCnt)
    do
        local incre=$(((50+$i)*10000*10000))
        echo '
CREATE TABLE IF NOT EXISTS `url_shorten_'$i'` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `url` varchar(256) NOT NULL DEFAULT "" COMMENT "url",
  `ip` varchar(41) NOT NULL DEFAULT "" COMMENT "creater ip",
  `ctime` int(10) unsigned NOT NULL DEFAULT "0" COMMENT "create time",
  `extra` varchar(256) NOT NULL DEFAULT "" COMMENT "extra info",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT='$incre' DEFAULT CHARSET=utf8;'
    done
}

function generateAccessHistory() {
    local tableCnt=$(expr $ACCESS_TABLECNT - 1)
    for i in $(seq 0 $tableCnt)
    do
        local incre=$(((50+$i)*10000*10000))
        echo '
CREATE TABLE IF NOT EXISTS `access_history_'$i'` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` char(6) NOT NULL COMMENT "short code",
  `ip` varchar(41) NOT NULL DEFAULT "" COMMENT "access user ip",
  `ctime` int(10) unsigned NOT NULL DEFAULT "0" COMMENT "access time",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;'
    done
}

generateShorten
generateAccessHistory
