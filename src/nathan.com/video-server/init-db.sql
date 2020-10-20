CREATE TABLE IF NOT EXISTS users(
  id int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  username varchar(50) NOT NULL COMMENT '用户登录名',
  passwd text NOT NULL COMMENT '用户密码',

  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  modify_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',

  primary key(id),
  unique key(login_name)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息';

CREATE TABLE IF NOT EXISTS videos(
  id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'video id',
  author_id int unsigned NOT NULL COMMENT '用户id',
  vname text NOT NULL COMMENT 'video name',

  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  modify_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',

  primary key(id),
  unique key(vname)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频信息';

CREATE TABLE IF NOT EXISTS comments(
  id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'comment id',
  video_id bigint unsigned NOT NULL COMMENT 'video id',
  author_id int unsigned NOT NULL COMMENT '用户id',
  content text NOT NULL COMMENT 'content',

  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  modify_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',

  primary key(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户评论信息';

CREATE TABLE IF NOT EXISTS sessions(
  session_id text NOT NULL COMMENT 'session id',
  ttl text COMMENT 'session 失效时间',
  login_name varchar(50) NOT NULL COMMENT '用户登录名',

  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  modify_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',

  primary key(session_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='session信息';