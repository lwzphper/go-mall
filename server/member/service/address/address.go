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
	"gorm.io/gorm"
)

var _ addresspb.AddressServiceServer = (*Service)(nil)

var (
	internalError      = status.Errorf(codes.Internal, "收货地址服务内部错误")
	notFoundError      = status.Errorf(codes.NotFound, "数据不存在")
	notPermissionError = status.Error(codes.Unauthenticated, "无权操作")
)

type Service struct {
	addresspb.UnimplementedAddressServiceServer

	AddressDao *address.Address
	Logger     *logger.Logger
}

func modelToResponse(addr entity.Address) addresspb.Entity {
	return addresspb.Entity{
		Id:        addr.Id,
		MemberId:  addr.MemberId,
		Name:      addr.Name,
		Phone:     addr.Phone,
		IsDefault: addr.IsDefault,
		PostCode:  addr.PostCode,
		Province:  addr.Province,
		City:      addr.City,
		Region:    addr.Region,
		Detail:    addr.Detail,
	}
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
		Detail:    req.Detail,
		MemberId:  req.MemberId,
	}
	err := s.AddressDao.Create(ctx, addr)
	if err != nil {
		s.Logger.Errorf("address create error：%s", err)
		return nil, internalError
	}
	return &addresspb.CreateResponse{
		Id: addr.Id,
	}, nil
}

// Update 更新
func (s *Service) Update(ctx context.Context, req *addresspb.Entity) (*empty.Empty, error) {
	uData := entity.Address{
		Name:      req.Name,
		Phone:     req.Phone,
		IsDefault: req.IsDefault,
		PostCode:  req.PostCode,
		Province:  req.Province,
		City:      req.City,
		Region:    req.Region,
		Detail:    req.Detail,
		MemberId:  req.MemberId,
	}
	err := s.AddressDao.UpdateUserItem(ctx, id.MemberID(req.MemberId), id.AddressID(req.Id), uData)
	if err != nil {
		s.Logger.Errorf("address update error：%s", err)
	}
	return &empty.Empty{}, err
}

// Delete 删除
func (s *Service) Delete(ctx context.Context, req *addresspb.DeleteRequest) (*empty.Empty, error) {
	pbEmpty := &empty.Empty{}
	// 权限校验
	data, err := s.checkCanEdit(ctx, id.AddressID(req.GetId()), id.MemberID(req.GetMemberId()))
	if err != nil {
		return pbEmpty, err
	}

	// 删除
	err = s.AddressDao.DeleteById(ctx, id.AddressID(data.Id))
	if err != nil {
		s.Logger.Errorf("delete address error：%s", err)
		return pbEmpty, internalError
	}
	return pbEmpty, nil
}

// GetList 获取列表
func (s *Service) GetList(ctx context.Context, req *addresspb.ListRequest) (*addresspb.ListResponse, error) {
	list, err := s.AddressDao.GetList(ctx, id.MemberID(req.MemberId))
	if err != nil {
		s.Logger.Errorf("address getList error：%s", err)
		return nil, internalError
	}

	result := &addresspb.ListResponse{}
	for _, addr := range list {
		respItem := modelToResponse(addr)
		result.List = append(result.List, &respItem)
	}
	return result, nil
}

// 检查是否可以编辑数据
func (s *Service) checkCanEdit(ctx context.Context, aid id.AddressID, mId id.MemberID) (*entity.Address, error) {
	// 数据不存在或异常
	data, err := s.AddressDao.GetItemById(ctx, aid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, notFoundError
		}
		s.Logger.Errorf("address get item by id error：%s", err)
		return nil, internalError
	}

	// 数据不属于当前用户
	if data.MemberId != mId.Uint64() {
		return nil, notPermissionError
	}
	return data, nil
}
