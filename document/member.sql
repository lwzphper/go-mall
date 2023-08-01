CREATE TABLE `member_level`
(
    `id`                     bigint       NOT NULL AUTO_INCREMENT,
    `name`                   varchar(100) NOT NULL COMMENT '等级名称',
    `growth_point`           int          NOT NULL COMMENT '积分',
    `default_status`         tinyint      NOT NULL DEFAULT 0 COMMENT '是否为默认等级：0->不是；1->是',
    `comment_growth_point`   tinyint      NOT NULL DEFAULT 0 COMMENT '每次评价获取的成长值',
    `privilege_free_freight` tinyint      NOT NULL DEFAULT 0 COMMENT '是否有免邮特权',
    `privilege_sign_in`      tinyint      NOT NULL DEFAULT 0 COMMENT '是否有签到特权',
    `privilege_comment`      tinyint      NOT NULL DEFAULT 0 COMMENT '是否有评论获奖励特权',
    `privilege_promotion`    tinyint      NOT NULL DEFAULT 0 COMMENT '是否有专享活动特权',
    `privilege_member_price` tinyint      NOT NULL DEFAULT 0 COMMENT '是否有会员价格特权',
    `privilege_birthday`     tinyint      NOT NULL DEFAULT 0 COMMENT '是否有生日特权',
    `created_at`             int          NOT NULL COMMENT '创建时间',
    `updated_at`             int          NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`             int          NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='会员等级表';

CREATE TABLE `member`
(
    `id`                     bigint       NOT NULL AUTO_INCREMENT,
    `member_level_id`        bigint       NOT NULL DEFAULT 0 COMMENT '会员等级',
    `username`               varchar(64)  NOT NULL COMMENT '用户名',
    `password`               varchar(64)  NOT NULL COMMENT '密码',
    `nickname`               varchar(64)  NOT NULL DEFAULT '' COMMENT '昵称',
    `phone`                  varchar(64)  NOT NULL DEFAULT '' COMMENT '手机号码',
    `status`                 tinyint      NOT NULL DEFAULT 1 COMMENT '帐号启用状态:0->禁用；1->启用',
    `icon`                   varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `gender`                 tinyint      NOT NULL COMMENT '性别：0->未知；1->男；2->女',
    `birthday`               date                  DEFAULT NULL COMMENT '生日',
    `city`                   varchar(64)  NOT NULL DEFAULT '' COMMENT '所做城市',
    `job`                    varchar(100) NOT NULL DEFAULT '' COMMENT '职业',
    `personalized_signature` varchar(200) NOT NULL DEFAULT '' COMMENT '个性签名',
    `source_type`            tinyint      NOT NULL DEFAULT 0 COMMENT '用户来源',
    `integration`            smallint     NOT NULL DEFAULT 0 COMMENT '积分',
    `growth`                 smallint     NOT NULL DEFAULT 0 COMMENT '成长值',
    `lucky_count`            smallint     NOT NULL DEFAULT 0 COMMENT '剩余抽奖次数',
    `history_integration`    smallint     NOT NULL DEFAULT 0 COMMENT '历史积分数量',
    `created_at`             int          NOT NULL COMMENT '创建时间',
    `updated_at`             int          NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`             int          NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`),
    UNIQUE KEY `idx_phone` (`phone`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='会员表';

-- 收货地址
CREATE TABLE `member_receive_address`
(
    `id`             bigint       NOT NULL AUTO_INCREMENT,
    `name`           varchar(100) NOT NULL COMMENT '收货人名称',
    `phone_number`   varchar(64)  NOT NULL COMMENT '手机号码',
    `is_default`     tinyint      NOT NULL DEFAULT 0 COMMENT '是否为默认。0：否；1：是',
    `post_code`      varchar(100) NOT NULL DEFAULT '' COMMENT '邮政编码',
    `province`       varchar(100) NOT NULL COMMENT '省份/直辖市',
    `city`           varchar(100) NOT NULL COMMENT '城市',
    `region`         varchar(100) NOT NULL COMMENT '区',
    `detail_address` varchar(128) NOT NULL COMMENT '详细地址(街道)',
    `member_id`      bigint       NOT NULL COMMENT '会员id',
    `created_at`     int          NOT NULL COMMENT '创建时间',
    `updated_at`     int          NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`     int          NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='会员收货地址表';

-- ========================================= 扩展 =================================================================
CREATE TABLE `member_login_log`
(
    `id`         bigint      NOT NULL AUTO_INCREMENT,
    `login_type` tinyint     NOT NULL COMMENT '登录类型：1：PC；2：android;3：ios;4：小程序',
    `ip`         varchar(64) NOT NULL DEFAULT '' COMMENT 'ip',
    `province`   varchar(64) NOT NULL DEFAULT '' COMMENT '省份',
    `city`       varchar(64) NOT NULL DEFAULT '' COMMENT '城市',
    `member_id`  bigint      NOT NULL COMMENT '会员id',
    `created_at` int         NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='会员登录记录';

CREATE TABLE `member_rule_setting`
(
    `id`                  bigint NOT NULL AUTO_INCREMENT,
    `type`                tinyint            NOT NULL DEFAULT 0 COMMENT '类型：1->积分规则；2->成长值规则',
    `continue_sign_day`   tinyint            NOT NULL DEFAULT 0 COMMENT '连续签到天数',
    `continue_sign_point` int            NOT NULL DEFAULT 0 COMMENT '连续签到赠送数量',
    `consume_per_point`   decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '每消费多少元获取1个点',
    `low_order_amount`    decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '最低获取点数的订单金额',
    `max_point_per_order` int            NOT NULL DEFAULT 0 COMMENT '每笔订单最高获取点数',
    `created_at`     int          NOT NULL COMMENT '创建时间',
    `updated_at`     int          NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`     int          NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='会员积分成长规则表';

CREATE TABLE `member_task`
(
    `id`           bigint NOT NULL AUTO_INCREMENT,
    `name`         varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
    `growth`       int                                                     DEFAULT NULL COMMENT '赠送成长值',
    `intergration` int                                                     DEFAULT NULL COMMENT '赠送积分',
    `type`         int                                                     DEFAULT NULL COMMENT '任务类型：0->新手任务；1->日常任务',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='会员任务表';