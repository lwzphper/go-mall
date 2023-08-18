CREATE TABLE `cart_item`
(
    `id`                bigint         NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `product_id`        bigint unsigned NOT NULL COMMENT '商品SPU ID',
    `product_sku_id`    bigint unsigned NOT NULL COMMENT '商品SKU ID',
    `member_id`         bigint unsigned NOT NULL COMMENT '用户ID',
    `product_quantity`  smallint unsigned NOT NULL DEFAULT 0 COMMENT '加购物车数量',
    `product_price`     decimal(10, 2) NOT NULL COMMENT '商品价格',
    `product_pic`       varchar(255)   NOT NULL DEFAULT '' COMMENT '商品图',
    `product_name`      varchar(255)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `product_sub_title` varchar(255)   NOT NULL DEFAULT '' COMMENT '商品副标题（卖点）',
    `product_brand`     varchar(255)   NOT NULL DEFAULT '' COMMENT '商品品牌',
    `product_attribute` varchar(1024)  NOT NULL DEFAULT '' COMMENT '商品规格，JSON 格式：[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]',
    `product_sku_code`  varchar(200)   NOT NULL DEFAULT '' COMMENT '商品sku条码',
    `product_sn`        varchar(128)   NOT NULL DEFAULT '' COMMENT '商品编码',
    `created_at`        timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `deleted_at`        int unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='购物车表';

CREATE TABLE `order`
(
    `id`                      bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `member_id`               bigint unsigned NOT NULL COMMENT '用户ID',
    `coupon_id`               bigint unsigned NOT NULL DEFAULT 0 COMMENT '优惠券id',
    `order_sn`                varchar(64)    NOT NULL COMMENT '订单编号',
    `total_amount`            decimal(10, 2) NOT NULL COMMENT '订单总金额',
    `pay_amount`              decimal(10, 2) NOT NULL COMMENT '支付金额',
    `freight_amount`          decimal(10, 2) NOT NULL COMMENT '运费金额',
    `promotion_amount`        decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '促销优化金额（促销价、满减、阶梯价）',
    `integration_amount`      decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '积分抵扣金额',
    `use_integration`         smallint       NOT NULL DEFAULT 0 COMMENT '下单时使用的积分',
    `coupon_amount`           decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '优惠券抵扣金额',
    `discount_amount`         decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '管理员后台调整订单使用的折扣金额',
    `pay_type`                tinyint unsigned NOT NULL DEFAULT '' COMMENT '支付方式：0->未支付；1->支付宝；2->微信',
    `pay_time`                int unsigned NOT NULL DEFAULT '' COMMENT '支付时间',
    `source`                  tinyint unsigned NOT NULL DEFAULT '' COMMENT '订单来源：1->PC订单；2->app订单',
    `status`                  tinyint unsigned NOT NULL DEFAULT '' COMMENT '订单状态。1->待付款；2->待发货；3->已发货；4->已完成；5->已关闭；6->无效订单',
    `type`                    tinyint unsigned NOT NULL DEFAULT 1 COMMENT '订单类型 1：正常订单 2：秒杀订单 3：促销订单',
    `delivery_company`        varchar(64)    NOT NULL DEFAULT '' COMMENT '物流公司',
    `delivery_sn`             varchar(64)    NOT NULL DEFAULT '' COMMENT '物流单号',
    `auto_confirm_day`        tinyint unsigned NOT NULL DEFAULT '' COMMENT '自动确认天数',
    `integration`             smallint unsigned NOT NULL DEFAULT 0 COMMENT '可以获得的积分',
    `growth`                  smallint unsigned NOT NULL DEFAULT 0 COMMENT '可以活动的成长值',
    `receive_address_id`      bigint unsigned NOT NULL COMMENT '收货地址id',
    `receiver_name`           varchar(64)    NOT NULL COMMENT '收货人',
    `receiver_phone`          varchar(32)    NOT NULL COMMENT '收货人电话',
    `receiver_post_code`      varchar(32)    NOT NULL COMMENT '收货人邮编',
    `receiver_province`       varchar(32)    NOT NULL COMMENT '收货人所在省',
    `receiver_city`           varchar(32)    NOT NULL COMMENT '收货人所在市',
    `receiver_region`         varchar(32)    NOT NULL COMMENT '收货人所在区',
    `receiver_detail_address` varchar(255)   NOT NULL COMMENT '收货人详细地址',
    `receive_time`            int unsigned NOT NULL DEFAULT '' COMMENT '收货时间',
    `confirm_status`          tinyint unsigned DEFAULT '0' COMMENT '收货状态 0：未接收 1：已接收',
    `delivery_time`           int unsigned NOT NULL DEFAULT 0 COMMENT '发货时间',
    `remark`                  varchar(255)   NOT NULL DEFAULT '' COMMENT '订单备注信息',
    `created_at`              timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`              timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `deleted_at`              int unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='订单表';

CREATE TABLE `order_item`
(
    `id`                 bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `order_id`           bigint unsigned NOT NULL COMMENT '订单ID',
    `order_sn`           varchar(64)    NOT NULL COMMENT '订单编号',
    `product_id`         bigint unsigned NOT NULL COMMENT '商品SPU ID',
    `product_sku_id`     bigint unsigned NOT NULL COMMENT '商品SKU ID',
    `product_pic`        varchar(255)   NOT NULL DEFAULT '' COMMENT '商品图',
    `product_name`       varchar(255)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `product_brand`      varchar(100)   NOT NULL DEFAULT '' COMMENT '商品品牌',
    `product_price`      decimal(10, 2) NOT NULL DEFAULT '' COMMENT '商品价格',
    `product_quantity`   smallint unsigned NOT NULL DEFAULT '' COMMENT '商品购买数量',
    `product_attribute`  varchar(1024)  NOT NULL DEFAULT '' COMMENT '规格，JSON 格式。[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]',
    `promotion_name`     varchar(200)   NOT NULL DEFAULT '' COMMENT '商品促销名称',
    `promotion_amount`   decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '商品促销分解金额',
    `coupon_amount`      decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '优惠券优惠分解金额',
    `integration_amount` decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '积分优惠分解金额',
    `real_amount`        decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '该商品经过优惠后的分解金额',
    `gift_integration`   smallint unsigned NOT NULL DEFAULT '0' COMMENT '活得积分',
    `gift_growth`        smallint unsigned NOT NULL DEFAULT '0' COMMENT '成长值',
    `created_at`         timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`         timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `deleted_at`         int unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY                  `idx_order_id` (`order_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='订单明细表';

CREATE TABLE `order_return_reason`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(100) NOT NULL DEFAULT '' COMMENT '退货类型',
    `sort`       tinyint unsigned NOT NULL DEFAULT 0 COMMENT '排序',
    `status`     tinyint unsigned NOT NULL DEFAULT '' COMMENT '状态：0->不启用；1->启用',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='退货原因表';

CREATE TABLE `order_operate_history`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT,
    `order_id`     bigint unsigned NOT NULL DEFAULT '' COMMENT '订单id',
    `operate_type` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '操作人：1：用户；2：系统；3：后台管理员',
    `order_status` int unsigned NOT NULL DEFAULT '' COMMENT '订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单',
    `remark`       varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
    `created_at`   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='订单操作历史记录';

CREATE TABLE `order_setting`
(
    `id`                    bigint NOT NULL AUTO_INCREMENT,
    `flash_order_overtime`  tinyint unsigned NOT NULL DEFAULT 0 COMMENT '秒杀订单超时关闭时间(分)',
    `normal_order_overtime` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '正常订单超时时间(分)',
    `confirm_overtime`      tinyint unsigned NOT NULL DEFAULT 0 COMMENT '发货后自动确认收货时间（天）',
    `finish_overtime`       tinyint unsigned NOT NULL DEFAULT 0 COMMENT '自动完成交易时间，不能申请售后（天）',
    `comment_overtime`      tinyint unsigned NOT NULL DEFAULT 0 COMMENT '订单完成后自动好评时间（天）',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='订单设置表';

CREATE TABLE `order_return_apply`
(
    `id`                 bigint unsigned NOT NULL AUTO_INCREMENT,
    `order_id`           bigint unsigned NOT NULL COMMENT '订单id',
    `receive_address_id` bigint unsigned NOT NULL COMMENT '收货地址id',
    `product_id`         bigint unsigned NOT NULL COMMENT '退货商品id',
    `order_sn`           varchar(64)    NOT NULL COMMENT '订单编号',
    `member_username`    varchar(64)    NOT NULL COMMENT '会员用户名',
    `return_amount`      decimal(10, 2) NOT NULL COMMENT '退款金额',
    `return_name`        varchar(32)    NOT NULL DEFAULT '' COMMENT '退货人姓名',
    `return_phone`       varchar(32)    NOT NULL DEFAULT '' COMMENT '退货人电话',
    `status`             tinyint unsigned NOT NULL DEFAULT 1 COMMENT '申请状态：1->待处理；2->退货中；3->已完成；4->已拒绝',
    `handle_time`        int unsigned NOT NULL DEFAULT 0 COMMENT '处理时间',
    `product_pic`        varchar(255)   NOT NULL DEFAULT '' COMMENT '商品图片',
    `product_name`       varchar(200)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `product_brand`      varchar(200)   NOT NULL DEFAULT '' COMMENT '商品品牌',
    `product_attr`       varchar(500)   NOT NULL DEFAULT '' COMMENT '商品销售属性：[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]',
    `product_count`      int unsigned NOT NULL COMMENT '退货数量',
    `product_price`      decimal(10, 2) NOT NULL COMMENT '商品单价',
    `product_real_price` decimal(10, 2) NOT NULL COMMENT '商品实际支付单价',
    `reason`             varchar(200)   NOT NULL DEFAULT '' COMMENT '原因',
    `description`        varchar(500)   NOT NULL DEFAULT '' COMMENT '描述',
    `proof_pics`         varchar(1000)  NOT NULL DEFAULT '' COMMENT '凭证图片，以逗号隔开',
    `handler_remark`     varchar(500)   NOT NULL DEFAULT '' COMMENT '处理人备注',
    `handler_name`       varchar(100)   NOT NULL DEFAULT '' COMMENT '处理人员',
    `receiver_name`      varchar(64)    NOT NULL COMMENT '收货人',
    `receive_time`       int unsigned NOT NULL DEFAULT '' COMMENT '收货时间',
    `receive_note`       varchar(500)   NOT NULL DEFAULT '' COMMENT '收货备注',
    `created_at`         timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='订单退货申请';



