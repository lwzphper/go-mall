package service

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/logger"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"github.com/lwzphper/go-mall/server/member/dao"
)

type MemberService struct {
	memberpb.UnimplementedMemberServiceServer
	MemberDao *dao.Member
	Logger    *logger.Logger
}

// GetMember 获取会员详情
func (m *MemberService) GetMember(ctx context.Context, req *memberpb.GetMemberRequest) (*memberpb.Member, error) {
	return nil, nil
}

// CreateMember 创建会员
func (m *MemberService) CreateMember(ctx context.Context, req *memberpb.SaveRequest) (*memberpb.BasicInfo, error) {
	return nil, nil
}
