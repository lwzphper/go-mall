package member

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lwzphper/go-mall/pkg/common/id"
	"github.com/lwzphper/go-mall/pkg/logger"
	"github.com/lwzphper/go-mall/pkg/until"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1/member"
	"github.com/lwzphper/go-mall/server/member/dao/member"
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
	MemberDao *member.Member
	Logger    *logger.Logger
}

func modelToResponse(member *entity.Member) *memberpb.MemberEntity {
	if member.Username == "" {
		member.Username = "默认名称"
	}

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
		global.Logger.Errorf("get member error:%v", err)
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
		global.Logger.Errorf("create member error:%v", err)
		return nil, InternalError
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
		global.Logger.Errorf("get member by id error:%v", err)
		return nil, InternalError
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
		global.Logger.Errorf("get member by phone error:%v", err)
		return nil, InternalError
	}
	return modelToResponse(item), nil
}

// UpdateMember 更新会员
func (s *MemberService) UpdateMember(ctx context.Context, req *memberpb.MemberEntity) (*empty.Empty, error) {
	// 获取需要更新的数据
	detail, err := s.MemberDao.GetItemById(ctx, id.MemberID(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, memberNotFoundError
		}
		global.Logger.Errorf("get member by id error:%v", err)
		return nil, InternalError
	}

	if until.TimeToDate(req.Birthday.AsTime()) == "0001-01-01" {
		detail.Birthday = nil
	} else {
		birthday := req.Birthday.AsTime()
		detail.Birthday = &birthday
	}
	detail.Username = req.Username
	detail.MemberLevelId = req.MemberLevelId
	detail.Icon = req.Icon
	detail.Status = entity.MemberStatus(req.Status)
	detail.Gender = entity.Gender(req.Gender)
	detail.City = req.City
	detail.Job = req.Job
	detail.Growth = req.Growth
	err = s.MemberDao.UpdateByEntity(ctx, detail)
	if err != nil {
		global.Logger.Errorf("update member error:%v", err)
		return nil, InternalError
	}
	return &empty.Empty{}, nil
}

// CheckPassWord 检查密码
func (s *MemberService) CheckPassWord(ctx context.Context, req *memberpb.PasswordCheckInfo) (*memberpb.CheckResponse, error) {
	return &memberpb.CheckResponse{
		Success: hash.ComparePwd(req.EncryptedPassword, []byte(req.Password)),
	}, nil
}
