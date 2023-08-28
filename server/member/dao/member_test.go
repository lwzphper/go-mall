package dao

import (
	"context"
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/lwzphper/go-mall/pkg/until"
	"github.com/lwzphper/go-mall/server/member/entity"
	"testing"
)

// 测试会员创建和查询
func TestCreateAndQueryMember(t *testing.T) {
	// 创建数据表
	if err := initTable(); err != nil {
		t.Errorf("create table error: %v", err)
	}

	dao := NewMember(mysqltesting.GormDB)
	ctx := context.Background()

	testCase := []struct {
		caseName  string // 测试名称
		username  string // 会员名称
		wantName  string // 期望值
		wantEqual bool   // 是否期望相等
	}{
		{
			caseName:  "test equal",
			username:  "张三",
			wantName:  "张三",
			wantEqual: true,
		},
		{
			caseName:  "test not equal",
			username:  "李四",
			wantName:  "张三",
			wantEqual: false,
		},
	}

	for _, c := range testCase {
		t.Run(c.caseName, func(t *testing.T) {
			member := entity.Member{
				Username: c.username,
				Password: until.RandomString(32),
			}
			err := dao.CreateMember(ctx, &member)
			if err != nil {
				t.Errorf("[%s]:create member error:%v", c.caseName, err)
			}

			memberRecord, err := dao.GetItem(ctx, &entity.Member{Username: c.username})
			if err != nil {
				t.Errorf("[%s]:get member info error: %v", c.caseName, err)
			}

			if c.wantEqual && c.wantName != memberRecord.Username {
				t.Errorf("[%s]:want %s but get %s", c.caseName, c.wantName, memberRecord.Username)
			}

			if c.wantEqual == false && c.wantName == memberRecord.Username {
				t.Errorf("[%s]:value equal %s", c.caseName, c.wantName)
			}
		})
	}
}

// 初始化数据表
func initTable() error {
	sql := "CREATE TABLE `member`\n(\n    `id`                     bigint unsigned   NOT NULL AUTO_INCREMENT,\n    `member_level_id`        bigint unsigned   NOT NULL DEFAULT 0 COMMENT '会员等级',\n    `username`               varchar(64)       NOT NULL COMMENT '用户名',\n    `password`               varchar(64)       NOT NULL COMMENT '密码',\n    `nickname`               varchar(64)       NOT NULL DEFAULT '' COMMENT '昵称',\n    `phone`                  char(11)          NOT NULL DEFAULT '' COMMENT '手机号码',\n    `status`                 tinyint unsigned  NOT NULL DEFAULT 1 COMMENT '帐号启用状态:0->禁用；1->启用',\n    `icon`                   varchar(255)      NOT NULL DEFAULT '' COMMENT '头像',\n    `gender`                 tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '性别：0->未知；1->男；2->女',\n    `birthday`               date                       DEFAULT NULL COMMENT '生日',\n    `city`                   varchar(64)       NOT NULL DEFAULT '' COMMENT '所做城市',\n    `job`                    varchar(100)      NOT NULL DEFAULT '' COMMENT '职业',\n    `personalized_signature` varchar(200)      NOT NULL DEFAULT '' COMMENT '个性签名',\n    `source_type`            tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '用户来源',\n    `integration`            smallint unsigned NOT NULL DEFAULT 0 COMMENT '积分',\n    `growth`                 smallint unsigned NOT NULL DEFAULT 0 COMMENT '成长值',\n    `lucky_count`            smallint unsigned NOT NULL DEFAULT 0 COMMENT '剩余抽奖次数',\n    `history_integration`    smallint unsigned NOT NULL DEFAULT 0 COMMENT '历史积分数量',\n    `created_at`             timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n    `updated_at`             timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',\n    `is_delete`             tinyint unsigned  NOT NULL DEFAULT 0 COMMENT '删除时间',\n    PRIMARY KEY (`id`),\n    UNIQUE KEY `idx_username` (`username`),\n    INDEX `idx_phone` (`phone`)\n) ENGINE = InnoDB\n  DEFAULT CHARSET = utf8mb4\n  COLLATE = utf8mb4_0900_ai_ci COMMENT ='会员表';"
	return mysqltesting.GormDB.Exec(sql).Error
}

func TestMain(m *testing.M) {
	mysqltesting.RunMysqlInDocker(m)
}
