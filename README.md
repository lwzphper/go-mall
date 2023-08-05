## go-mall

go商城系统（开发中...）

策划：

前台商城系统包括：用户登录、注册、商品搜索、商品详情、购物车、订单、秒杀活动等模块。

后台管理系统包括：系统管理、商品系统、优惠营销、库存系统、订单系统、用户系统、内容管理等七大模块。

## 依赖
protoc
grpc-go：https://github.com/grpc/grpc-go

日志切割：
+ 按大小切割：github.com/natefinch/lumberjack
+ 按日期切割：
  + 方式1：使用 Linux 系统自带的 logrotate
  + 方式2：github.com/lestrrat-go/file-rotatelogs （不再维护）

## 技术栈

1. 熔断、限流：sentinel
2. 分布式事务：seata
3. 注册中心：nacos
4. 远程调用&负载均衡：Feign
5. 分布式id：ksuid
6. 链路追踪
7. MQ
8. redis
9. mysql
10. MongoDB
11. elasticsearch
12. canal



## 服务划分

1. 商品服务：分类管理、属性分组、
2. 用户服务
3. 仓储服务
4. 秒杀服务
5. 订单服务
6. 购物车服务
7. 检索服务
8. 中央认证服务
9. 支付服务
10. 优惠服务
11. IM客服服务

第三方服务

1. 物流
2. 短信
3. 金融
4. 身份认证

## 目录划分

document：文档，如：sql、docker-comopse、shell

config：配置文件

middleware：中间件

routere：路由
