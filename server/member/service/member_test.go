package service

import (
	"context"
	"fmt"
	"github.com/lwzphper/go-mall/pkg/common/id"
	mysqltesting "github.com/lwzphper/go-mall/pkg/db/mysql/testing"
	"github.com/lwzphper/go-mall/pkg/db/mysql/testing/init_table"
	"github.com/lwzphper/go-mall/pkg/logger"
	"github.com/lwzphper/go-mall/pkg/until"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"github.com/lwzphper/go-mall/server/member/dao"
	"github.com/lwzphper/go-mall/server/member/global"
	"github.com/lwzphper/go-mall/server/member/until/hash"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"testing"
	"time"
)

var srv *MemberService

// 创建会员测试
func TestMemberService_CreateMember(t *testing.T) {
	initDB()

	// 添加数据
	member, err := srv.CreateMember(context.Background(), &memberpb.CreateRequest{
		Username: "张三",
		Phone:    "15800000001",
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, id.MemberID(1).Uint64(), member.Id)

	// 手机号码重复添加测试
	member, err = srv.CreateMember(context.Background(), &memberpb.CreateRequest{
		Username: "张三",
		Phone:    "15800000001",
		Password: "123456",
	})
	if err != nil {
		assert.EqualError(t, err, memberExistError.Error())
	} else {
		t.Errorf("手机号码重复添加数据验证不生效")
	}
}

// 通过用户id获取会员信息
func TestMemberService_GetMemberById(t *testing.T) {
	initDB()

	ctx := context.Background()
	cReq := &memberpb.CreateRequest{
		Username: "张三",
		Phone:    "15800000001",
		Password: "123456",
	}
	member, err := srv.CreateMember(ctx, cReq)
	if err != nil {
		t.Error(err)
	}

	req := &memberpb.IdRequest{Id: member.Id}
	getVal, err := srv.GetMemberById(ctx, req)
	if err != nil {
		t.Errorf("get member error:%s", err)
	}
	assert.Equal(t, cReq.Phone, getVal.Phone)
}

// 通过手机号码获取会员信息
func TestMemberService_GetMemberByPhone(t *testing.T) {
	initDB()

	ctx := context.Background()
	phone := "15800000001"
	cReq := &memberpb.CreateRequest{
		Username: "张三",
		Phone:    phone,
		Password: "123456",
	}
	_, err := srv.CreateMember(ctx, cReq)
	if err != nil {
		t.Error(err)
	}

	req := &memberpb.PhoneRequest{Phone: phone}
	getVal, err := srv.GetMemberByPhone(ctx, req)
	if err != nil {
		t.Errorf("get member error:%s", err)
	}
	assert.Equal(t, cReq.Phone, getVal.Phone)
}

// 更新会员
func TestMemberService_UpdateMember(t *testing.T) {
	initDB()

	ctx := context.Background()
	cReq := &memberpb.CreateRequest{
		Username: "张三",
		Phone:    "15800000001",
		Password: "123456",
	}
	member, err := srv.CreateMember(ctx, cReq)
	if err != nil {
		t.Error(err)
	}

	testCases := []struct {
		name     string
		birthday *timestamppb.Timestamp
		want     string
	}{
		{
			name:     "正常时间格式更新",
			birthday: timestamppb.Now(),
			want:     until.PbTimeToDate(timestamppb.Now()),
		},
		{
			name:     "零值时间格式更新",
			birthday: timestamppb.New(time.Time{}),
			want:     "1970-01-02",
		},
	}

	var (
		username = "王五"
		job      = "go开发工程师"
		city     = "广州"
		gender   = memberpb.MemberGender_MAN
		icon     = "https://avatars.githubusercontent.com/u/35757691?v=4"
	)

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			_, err = srv.UpdateMember(ctx, &memberpb.MemberEntity{
				Id:       member.Id,
				Birthday: c.birthday,
				Username: username,
				Job:      job,
				City:     city,
				Gender:   gender,
				Icon:     icon,
			})
			if err != nil {
				t.Errorf("update member error:%v", err)
			}

			newMember, err := srv.GetMemberById(ctx, &memberpb.IdRequest{Id: member.Id})
			if err != nil {
				t.Errorf("get member error:%v", err)
			}
			assert.Equal(t, username, newMember.Username)
			assert.Equal(t, job, newMember.Job)
			assert.Equal(t, city, newMember.City)
			assert.Equal(t, c.want, until.PbTimeToDate(newMember.Birthday)) // 零值测试
			assert.Equal(t, gender, newMember.Gender)
			assert.Equal(t, icon, newMember.Icon)
		})
	}
}

// 检查密码
func TestMemberService_CheckPassWord(t *testing.T) {
	oriPwd := "123456"
	encryptPwd, err := hash.HashAndSalt([]byte(oriPwd))
	if err != nil {
		t.Errorf("password hash error:%v", err)
	}

	fmt.Printf(encryptPwd)

	cases := []struct {
		name       string
		oriPwd     string
		encryptPwd string
		want       bool
	}{
		{
			name:       "success test",
			oriPwd:     oriPwd,
			encryptPwd: encryptPwd,
			want:       true,
		},
		{
			name:       "fail test",
			oriPwd:     oriPwd,
			encryptPwd: "$2a$04$7o1PxRoELncS9oubwUYQ5utdl/ZYS5ywgSWOYHuV/vPPA.GnGCZAG",
			want:       false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			checkRet, _ := srv.CheckPassWord(context.Background(), &memberpb.PasswordCheckInfo{
				Password:          c.oriPwd,
				EncryptedPassword: c.encryptPwd,
			})
			assert.Equal(t, c.want, checkRet.Success)
		})
	}
}

// 数据库初始化
func initDB() {
	// 创建数据表
	if err := init_table.Member(); err != nil {
		log.Panicf("create table error: %v", err)
	}

	global.DB = mysqltesting.GormDB

	srv = &MemberService{
		MemberDao: dao.NewMember(),
		Logger:    logger.NewDefaultLogger(),
	}
}

func TestMain(m *testing.M) {
	mysqltesting.RunMysqlInDocker(m)
}
