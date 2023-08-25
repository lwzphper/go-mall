package service

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/logger"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"github.com/lwzphper/go-mall/server/member/dao"
	"github.com/lwzphper/go-mall/server/member/entity"
	"github.com/pkg/errors"
)

type MemberService struct {
	memberpb.UnimplementedMemberServiceServer
	MemberDao *dao.Member
	Logger    *logger.Logger
}

// CreateMember 创建会员
func (s *MemberService) CreateMember(ctx context.Context, req *memberpb.CreateRequest) (*memberpb.BasicInfo, error) {
	m := &entity.Member{
		Username: req.Member.Username,
		Password: req.Password,
	}

	err := s.MemberDao.CreateMember(ctx, m)
	if err != nil {
		return nil, err
	}

	return &memberpb.BasicInfo{
		Id:       m.Id,
		Username: m.Username,
		Nickname: m.Nickname,
	}, nil
}

// GetMember 获取会员详情
func (s *MemberService) GetMember(ctx context.Context, req *memberpb.GetMemberRequest) (*memberpb.Member, error) {
	if req.Username == "" && req.Phone == "" {
		return nil, errors.New("username or phone empty")
	}

	//memberRecord, err := s.MemberDao.GetItem(ctx, entity.Member{})
	return nil, nil
}
