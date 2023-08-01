CREATE TABLE `product_category`
(
    `id`         bigint       NOT null AUTO_INCREMENT COMMENT 'ID',
    `name`       varchar(64)  NOT NULL COMMENT '分类名称',
    `status`     tinyint      NOT NULL DEFAULT 0 COMMENT '状态。0：隐藏；1：展示',
    `parent_id`  bigint       NOT NULL DEFAULT 0 COMMENT '父级ID',
    `level`      tinyint      NOT NULL DEFAULT 0 COMMENT '层级',
    `icon_url`   varchar(255) NOT NULL DEFAULT '' COMMENT '图标',
    `link`       varchar(255) NOT NULL DEFAULT '' COMMENT '跳转地址',
    `sort`       int          NOT NULL DEFAULT 0 COMMENT '排序',
    `deleted_at` int          NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '商品分类表';

CREATE TABLE `product_brand`
(
    `id`          bigint       NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`        varchar(64)  NOT NULL COMMENT '品牌名称',
    `logo`        varchar(255) NOT NULL DEFAULT '' COMMENT '品牌图',
    `sort`        int          NOT NULL DEFAULT 0 COMMENT '排序',
    `description` varchar(512) NOT NULL DEFAULT '' COMMENT '品牌介绍',
    `deleted_at`  int          NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品品牌表';

CREATE TABLE `product_spu`
(
    `id`                   bigint         NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `category_id`          bigint         NOT NULL COMMENT '商品类型ID',
    `brand_id`             bigint         NOT NULL COMMENT '商品品牌ID',
    `name`                 varchar(200)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `sub_title`            varchar(255)   NOT NULL DEFAULT '' COMMENT '副标题',
    `product_sn`           varchar(128)   NOT NULL DEFAULT '' COMMENT '商品编码',
    `brand_name`           varchar(64)    NOT NULL DEFAULT '' COMMENT '品牌名称',
    `category_name`        varchar(64)    NOT NULL DEFAULT '' COMMENT '商品分类名称',
    `pic`                  varchar(255)   NOT NULL DEFAULT '' COMMENT '商品主图',
    `album_pics`           varchar(500)   NOT NULL DEFAULT '' COMMENT '商品图集，逗号分割',
    `price`                decimal(10, 2) NOT NULL COMMENT '商品价格',
    `original_price`       decimal(10, 2) NOT NULL COMMENT '市场价',
    `promotion_price`      decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '促销价格',
    `promotion_limit`      int            NOT NULL DEFAULT 0 COMMENT '活动限购数量',
    `promotion_start_time` int            NOT NULL DEFAULT DEFAULT 0 COMMENT '促销开始时间',
    `promotion_end_time`   int            NOT NULL DEFAULT DEFAULT 0 COMMENT '促销结束时间',
    `gift_growth`          int            NOT NULL DEFAULT 0 COMMENT '赠送的成长值',
    `gift_point`           int            NOT NULL DEFAULT 0 COMMENT '赠送的积分',
    `sort`                 int            NOT NULL DEFAULT 0 COMMENT '排序',
    `sales_count`          int            NOT NULL DEFAULT '0' COMMENT '销量',
    `unit`                 varchar(32)    NOT NULL DEFAULT '' COMMENT '单位',
    `weight`               decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '商品重量，默认为克',
    `service_ids`          varchar(64)    NOT NULL DEFAULT '' COMMENT '以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮',
    `publish_status`       tinyint        NOT NULL DEFAULT 0 COMMENT '发布状态 0：未发布；1：发布',
    `new_status`           tinyint        NOT NULL DEFAULT 0 COMMENT '新品状态 0：非新品；1：新品',
    `recommend_status`     tinyint        NOT NULL DEFAULT 0 COMMENT '推荐状态 0：非推荐；1：推荐',
    `verify_status`        tinyint        NOT NULL DEFAULT 0 COMMENT '审核状态 0：未审核；1：审核通过；3：审核不通过',
    `description`          text COMMENT '商品详情',
    `created_at`           int            NOT NULL COMMENT '创建时间',
    `updated_at`           int            NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`           int            NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY                    `idx_brand_id` (`brand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品表';

CREATE TABLE `product_sku`
(
    `id`          bigint         NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `category_id` bigint         NOT NULL DEFAULT 0 COMMENT '商品类型ID',
    `brand_id`    bigint         NOT NULL DEFAULT 0 COMMENT '商品品牌ID',
    `product_id`  bigint         NOT NULL COMMENT '商品ID',
    `price`       decimal(10, 2) NOT NULL COMMENT '价格',
    `stock`       int            NOT NULL DEFAULT '0' COMMENT '库存',
    `lock_stock`  int            NOT NULL DEFAULT '0' COMMENT '锁定库存',
    `pic`         varchar(255)   NOT NULL DEFAULT '' COMMENT '图片',
    `sp_data`     varchar(1024)  NOT NULL DEFAULT '' COMMENT '商品属性，JSON 格式', -- 如：[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]
    `created_at`  int            NOT NULL COMMENT '创建时间',
    `updated_at`  int            NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`  int            NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY           `idx_brand_id` (`brand_id`,`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品库存表';

CREATE TABLE `product_attribute_category` -- 如：电脑办公-笔记本
(
    `id`         bigint      NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`       varchar(64) NOT NULL COMMENT '分类名称',
    `created_at` int         NOT NULL COMMENT '创建时间',
    `updated_at` int         NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at` int         NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品属性分类表';

CREATE TABLE `product_attribute` -- 如：颜色
(
    `id`                            bigint        NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `product_attribute_category_id` bigint        NOT NULL DEFAULT 0 COMMENT '商品属性分类 ID',
    `name`                          varchar(64)   NOT NULL COMMENT '名称',
    `select_type`                   tinyint       NOT NULL COMMENT '熟悉选择类型。1：唯一；2：单选；3：多选',
    `option_status`                 tinyint       NOT NULL COMMENT '录入方式 1：手动录入 2：列表中选取',
    `option_list`                   varchar(1024) NOT NULL COMMENT '可选值列表，以逗号隔开',
    `sort`                          int           NOT NULL COMMENT '排序',
    `type`                          tinyint       NOT NULL COMMENT '类型 1：规则 2：参数',
    `created_at`                    int           NOT NULL COMMENT '创建时间',
    `updated_at`                    int           NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`                    int           NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品属性表';

CREATE TABLE `product_attribute_value` -- 如：黑色,白色
(
    `id`                   bigint      NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `product_id`           bigint      NOT NULL COMMENT '商品 ID',
    `product_attribute_id` bigint      NOT NULL COMMENT '商品属性 ID',
    `value`                varchar(64) NOT NULL COMMENT '手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开',
    `created_at`           int         NOT NULL COMMENT '创建时间',
    `updated_at`           int         NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`           int         NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品属性值表';

CREATE TABLE `product_comment`
(
    `id`               bigint        NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `product_id`       bigint        NOT NULL DEFAULT 0 COMMENT '商品ID',
    `product_sku_id`   bigint        NOT NULL DEFAULT 0 COMMENT '商品SKU ID',
    `customer_user_id` bigint        NOT NULL DEFAULT 0 COMMENT '用户ID',
    `member_nick_name` varchar(100)  NOT NULL COMMENT '用户昵称',
    `member_icon`      varchar(255)  NOT NULL COMMENT '用户头像',
    `like_count`       int           NOT NULL DEFAULT '0' COMMENT '点赞数',
    `reply_count`      int           NOT NULL DEFAULT '0' COMMENT '回复数',
    `star`             tinyint       NOT NULL DEFAULT 0 COMMENT '评星数。1-5',
    `content`          varchar(1000) NOT NULL DEFAULT '' COMMENT '评论',
    `hide_flag`        tinyint       NOT NULL DEFAULT 0 COMMENT '匿名标识 1：匿名 2：不匿名',
    `resource`         varchar(1000) NOT NULL DEFAULT '' COMMENT '评论图片/视频，JSON格式',
    `created_at`       int           NOT NULL COMMENT '创建时间',
    `updated_at`       int           NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`       int           NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY                `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品评论表';

CREATE TABLE `product_comment_replay`
(
    `id`               bigint        NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `type`             tinyint       NOT NULL DEFAULT 0 COMMENT '评论人员类型。 1：会员 2：管理员',
    `comment_id`       bigint        NOT NULL DEFAULT 0 COMMENT '评论id',
    `customer_user_id` bigint        NOT NULL DEFAULT 0 COMMENT '用户ID',
    `append_flag`      tinyint       NOT NULL DEFAULT 0 COMMENT '追加标识 0：否 1：是',
    `member_nick_name` varchar(100)  NOT NULL COMMENT '用户昵称',
    `member_icon`      varchar(255)  NOT NULL COMMENT '用户头像',
    `content`          varchar(1000) NOT NULL DEFAULT '' COMMENT '评论',
    `created_at`       int           NOT NULL COMMENT '创建时间',
    `deleted_at`       int           NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY                `idx_comment_id` (`comment_id`, `deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品评论回复表';

CREATE TABLE `product_operate_log`
(
    `id`                 bigint         NOT NULL AUTO_INCREMENT,
    `product_id`         bigint         NOT NULL COMMENT '商品id',
    `price_old`          decimal(10, 2) NOT NULL COMMENT '旧的商品价格',
    `price_new`          decimal(10, 2) NOT NULL COMMENT '新的商品价格',
    `original_price_old` decimal(10, 2) NOT NULL COMMENT '旧的市场价',
    `original_price_new` decimal(10, 2) NOT NULL COMMENT '新的市场价',
    `gift_point_old`     int            NOT NULL COMMENT '旧的赠送积分',
    `gift_point_new`     int            NOT NULL COMMENT '新的赠送积分',
    `operate_man`        varchar(64)    NOT NULL COMMENT '操作人',
    `created_at`         int            NOT NULL COMMENT '创建时间',
    `updated_by`         int            NOT NULL DEFAULT 0 COMMENT '操作人id',
    `updated_at`         int            NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deleted_at`         int            NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品修改记录表';