CREATE TABLE `admin_user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `username`   varchar(64)  NOT NULL COMMENT '用户名称',
    `password`   varchar(64)  NOT NULL COMMENT '密码',
    `icon`       varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `email`      varchar(30)  NOT NULL DEFAULT '' COMMENT '邮箱',
    `nick_name`  varchar(30)  NOT NULL DEFAULT '' COMMENT '昵称',
    `remark`     varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
    `status`     tinyint unsigned NOT NULL DEFAULT '1' COMMENT '帐号启用状态：0->禁用；1->启用',
    `login_time` int unsigned NOT NULL DEFAULT 0 COMMENT '最后登录时间',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY          `idx_username` (`username`, `status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='后台用户表';

CREATE TABLE `admin_login_log`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `admin_id`   bigint unsigned NOT NULL COMMENT '管理员id',
    `ip`         varchar(64)  NOT NULL DEFAULT '' COMMENT '登录ip',
    `address`    varchar(100) NOT NULL DEFAULT '' COMMENT 'ip地址',
    `user_agent` varchar(100) NOT NULL DEFAULT '' COMMENT '浏览器登录类型',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY          `idx_admin_id` (`admin_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='后台用户登录日志表';

CREATE TABLE `ums_permission`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `pid`        bigint unsigned NOT NULL DEFAULT 0 COMMENT '父级权限id',
    `name`       varchar(100) NOT NULL COMMENT '名称',
    `status`     tinyint unsigned NOT NULL DEFAULT 1 COMMENT '启用状态；0->禁用；1->启用',
    `value`      varchar(200) NOT NULL DEFAULT '' COMMENT '权限值',
    `icon`       varchar(255) NOT NULL DEFAULT '' COMMENT '图标',
    `type`       int unsigned NOT NULL DEFAULT '' COMMENT '权限类型：1->目录；2->菜单；3->按钮（接口绑定权限）',
    `uri`        varchar(200) NOT NULL DEFAULT '' COMMENT '前端资源路径',
    `sort`       smallint unsigned NOT NULL DEFAULT '' COMMENT '排序',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='后台用户权限表';

CREATE TABLE `admin_permission_relation`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `type`          tinyint unsigned NOT NULL DEFAULT 0 COMMENT '类型',
    `admin_id`      bigint unsigned NOT NULL COMMENT '管理员id',
    `permission_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '权限id',
    PRIMARY KEY (`id`),
    KEY             `idx_admin_id` (`admin_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='后台用户和权限关系表(除角色中定义的权限以外的加减权限)';

CREATE TABLE `admin_role`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(30)  NOT NULL DEFAULT '' COMMENT '名称',
    `status`      tinyint unsigned DEFAULT '1' COMMENT '启用状态：0->禁用；1->启用',
    `admin_count` smallint unsigned NOT NULL DEFAULT 0 COMMENT '后台用户数量',
    `sort`        int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
    `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
    `created_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='后台用户角色表';

CREATE TABLE `admin_role_permission_relation`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `role_id`       bigint unsigned NOT NULL DEFAULT 0 COMMENT '角色id',
    `permission_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '权限id',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='后台用户角色和权限关系表';

CREATE TABLE `admin_role_relation`
(
    `id`       bigint unsigned NOT NULL AUTO_INCREMENT,
    `admin_id` bigint unsigned NOT NULL COMMENT '管理员id',
    `role_id`  bigint unsigned NOT NULL DEFAULT 0 COMMENT '角色id',
    PRIMARY KEY (`id`),
    KEY        `idx_admin_id` (`admin_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='后台用户和角色关系表';

CREATE TABLE `menu`
(
    `id`         bigint       NOT NULL AUTO_INCREMENT,
    `name`       varchar(100) NOT NULL DEFAULT '' COMMENT '菜单名称',
    `parent_id`  bigint unsigned NOT NULL DEFAULT 0 COMMENT '父级ID',
    `title`      varchar(100) NOT NULL COMMENT '菜单名称',
    `level`      smallint unsigned NOT NULL DEFAULT 0 COMMENT '菜单级数',
    `sort`       smallint unsigned NOT NULL DEFAULT '' COMMENT '菜单排序',
    `icon`       varchar(200) NOT NULL DEFAULT '' COMMENT '图标',
    `is_hidden`  int unsigned NOT NULL DEFAULT '' COMMENT '是否隐藏。0：否；1：是',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='后台菜单表';

CREATE TABLE `admin_role_menu_relation`
(
    `id`      bigint unsigned NOT NULL AUTO_INCREMENT,
    `role_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '角色ID',
    `menu_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '菜单ID',
    PRIMARY KEY (`id`),
    KEY       `idx_role_id` (`role_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='后台角色菜单关系表';