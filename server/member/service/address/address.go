package address

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lwzphper/go-mall/pkg/common/id"
	"github.com/lwzphper/go-mall/pkg/logger"
	addresspb "github.com/lwzphper/go-mall/server/member/api/gen/v1/address"
	"github.com/lwzphper/go-mall/server/member/dao/address"
	"github.com/lwzphper/go-mall/server/member/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ addresspb.AddressServiceServer = (*Service)(nil)

type Service struct {
	addresspb.UnimplementedAddressServiceServer

	AddressDao *address.Address
	Logger     *logger.Logger
}

func modelToResponse(addr entity.Address) *addresspb.Entity {
	var pb *addresspb.Entity
	pb.Id = addr.Id
	pb.MemberId = addr.MemberId
	pb.Name = addr.Name
	pb.IsDefault = addr.IsDefault
	pb.PostCode = addr.PostCode
	pb.Province = addr.Province
	pb.City = addr.City
	pb.Region = addr.Region
	pb.Address = addr.Address
	return pb
}

// Create 创建
func (s *Service) Create(ctx context.Context, req *addresspb.CreateRequest) (*addresspb.CreateResponse, error) {
	addr := &entity.Address{
		Name:      req.Name,
		Phone:     req.Phone,
		IsDefault: req.IsDefault,
		PostCode:  req.PostCode,
		Province:  req.Province,
		City:      req.City,
		Region:    req.Region,
		Address:   req.Address,
		MemberId:  req.MemberId,
	}
	err := s.AddressDao.Create(ctx, addr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &addresspb.CreateResponse{
		Id: addr.Id,
	}, nil
}

// Update 更新
func (s *Service) Update(ctx context.Context, req *addresspb.Entity) (*empty.Empty, error) {
	// todo 更新
	return &empty.Empty{}, nil
}

// Delete 删除
func (s *Service) Delete(ctx context.Context, req *addresspb.DeleteRequest) (*empty.Empty, error) {
	// todo 删除
	return &empty.Empty{}, nil
}

// GetList 获取列表
func (s *Service) GetList(ctx context.Context, req *addresspb.ListRequest) (*addresspb.ListResponse, error) {
	list, err := s.AddressDao.GetList(ctx, id.MemberID(req.MemberId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var result *addresspb.ListResponse
	for _, addr := range list {
		result.List = append(result.List, modelToResponse(addr))
	}
	return result, nil
}
