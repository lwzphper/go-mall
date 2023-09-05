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
	"github.com/lwzphper/go-mall/server/member/until/hash"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

var _ memberpb.MemberServiceServer = (*MemberService)(nil)

var (
	InternalError       = status.Errorf(codes.Internal, "用户模块内部错误")
	memberExistError    = status.Errorf(codes.AlreadyExists, "用户已存在")
	memberNotFoundError = status.Errorf(codes.NotFound, "用户不存在")
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
	pbMember.Username = member.Username
	pbMember.Phone = member.Phone
	pbMember.Icon = member.Icon
	pbMember.Status = memberpb.MemberStatus(member.Status)
	pbMember.Gender = memberpb.MemberGender(member.Gender)
	pbMember.City = member.City
	pbMember.Job = member.Job
	pbMember.Growth = member.Growth
	pbMember.CreatedAt = until.TimeToDateTime(member.CreatedAt)
	if member.Birthday != nil {
		pbMember.Birthday = timestamppb.New(*member.Birthday)
	}
	return pbMember
}

// CreateMember 创建会员
func (s *MemberService) CreateMember(ctx context.Context, req *memberpb.CreateRequest) (*memberpb.CreateResponse, error) {
	// 校验用户是否存在
	member, err := s.MemberDao.GetItemByWhere(ctx, &entity.Member{Phone: req.Phone})
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Logger.Errorf("create member error:%v", err)
		return nil, InternalError
	}

	// 用户已存在
	if member.Id > 0 {
		return nil, memberExistError
	}

	// 创建用户
	pwdHash, err := hash.HashAndSalt([]byte(req.Password))
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

	return &memberpb.CreateResponse{
		Id: m.Id,
	}, nil
}

// GetMemberById 通过用户id获取会员信息
func (s *MemberService) GetMemberById(ctx context.Context, req *memberpb.IdRequest) (*memberpb.MemberEntity, error) {
	memberRecord, err := s.MemberDao.GetItemById(ctx, id.MemberID(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, memberNotFoundError
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
			return nil, memberNotFoundError
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
			return nil, memberNotFoundError
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	birthday := req.Birthday.AsTime()
	member.Birthday = &birthday
	member.Username = req.Username
	member.MemberLevelId = req.MemberLevelId
	member.Icon = req.Icon
	member.Status = entity.MemberStatus(req.Status)
	member.Gender = entity.Gender(req.Gender)
	member.City = req.City
	member.Job = req.Job
	member.Growth = req.Growth
	err = s.MemberDao.UpdateByEntity(ctx, member)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

// CheckPassWord 检查密码
func (s *MemberService) CheckPassWord(ctx context.Context, req *memberpb.PasswordCheckInfo) (*memberpb.CheckResponse, error) {
	return &memberpb.CheckResponse{
		Success: hash.ComparePwd(req.EncryptedPassword, []byte(req.Password)),
	}, nil
}
