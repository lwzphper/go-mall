package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lwzphper/go-mall/pkg/common/id"
	"github.com/lwzphper/go-mall/pkg/logger"
	"github.com/lwzphper/go-mall/pkg/until"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"github.com/lwzphper/go-mall/server/member/dao"
	"github.com/lwzphper/go-mall/server/member/entity"
	"github.com/lwzphper/go-mall/server/member/global"
	memberUntil "github.com/lwzphper/go-mall/server/member/until"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type MemberService struct {
	memberpb.UnimplementedMemberServiceServer
	MemberDao *dao.Member
	Logger    *logger.Logger
}

func modelToResponse(member *entity.Member) *memberpb.MemberEntity {
	pbMember := &memberpb.MemberEntity{}
	pbMember.Id = member.Id
	pbMember.MemberLevelId = member.MemberLevelId
	pbMember.Password = member.Password
	pbMember.Nickname = member.Nickname
	pbMember.Phone = member.Phone
	pbMember.Icon = member.Icon
	pbMember.Status = memberpb.MemberStatus(member.Status)
	pbMember.Gender = memberpb.MemberGender(member.Gender)
	pbMember.Birthday = timestamppb.New(*member.Birthday)
	pbMember.City = member.City
	pbMember.Job = member.Job
	pbMember.Growth = member.Growth
	pbMember.CreatedAt = until.TimeToYmdHis(&member.CreatedAt)
	return pbMember
}

// CreateMember 创建会员
func (s *MemberService) CreateMember(ctx context.Context, req *memberpb.CreateRequest) (*memberpb.MemberEntity, error) {
	// 校验用户是否存在
	_, err := s.MemberDao.GetItemByWhere(ctx, &entity.Member{Phone: req.Phone})
	if err == nil || (err != nil && err == gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.Unavailable, "用户已存在")
	}

	// 创建用户
	pwdHash, err := memberUntil.HashAndSalt([]byte(req.Password))
	if err != nil {
		global.Logger.Errorf("密码加密失败:%s", err)
		return nil, status.Errorf(codes.Internal, "密码加密失败")
	}
	m := &entity.Member{
		Username: req.Username,
		Phone:    req.Phone,
		Password: pwdHash,
	}
	err = s.MemberDao.CreateMember(ctx, m)
	if err != nil {
		return nil, err
	}

	return &memberpb.MemberEntity{
		Id:       m.Id,
		Username: m.Username,
		Nickname: m.Nickname,
	}, nil
}

// GetMemberById 通过用户id获取会员信息
func (s *MemberService) GetMemberById(ctx context.Context, req *memberpb.IdRequest) (*memberpb.MemberEntity, error) {
	memberRecord, err := s.MemberDao.GetItemById(ctx, id.MemberID(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "用户不存在")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return modelToResponse(memberRecord), nil
}

// GetMemberByPhone 通过手机号码获取会员信息
func (s *MemberService) GetMemberByPhone(ctx context.Context, req *memberpb.PhoneRequest) (*memberpb.MemberEntity, error) {
	item, err := s.MemberDao.GetItemByWhere(ctx, &entity.Member{Phone: req.Phone})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "用户不存在")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return modelToResponse(item), nil
}

// UpdateMember 更新会员
func (s *MemberService) UpdateMember(ctx context.Context, req *memberpb.MemberEntity) (*empty.Empty, error) {
	// 获取需要更新的数据
	member, err := s.MemberDao.GetItemById(ctx, id.MemberID(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "用户不存在")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	birthday := req.Birthday.AsTime()
	member.Birthday = &birthday
	member.MemberLevelId = req.MemberLevelId
	member.Nickname = req.Nickname
	member.Icon = req.Icon
	member.Status = entity.MemberStatus(req.Status)
	member.Gender = entity.Gender(req.Gender)
	member.City = req.City
	member.Job = req.Job
	member.Growth = req.Growth
	err = s.MemberDao.Update(ctx, member)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

// CheckPassWord 检查密码
func (s *MemberService) CheckPassWord(ctx context.Context, req *memberpb.PasswordCheckInfo) (*memberpb.CheckResponse, error) {
	return &memberpb.CheckResponse{
		Success: memberUntil.ComparePwd(req.EncryptedPassword, []byte(req.Password)),
	}, nil
}
