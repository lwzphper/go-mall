package global

import (
	"github.com/lwzphper/go-mall/bff/config"
	"github.com/lwzphper/go-mall/pkg/logger"
	addresspb "github.com/lwzphper/go-mall/server/member/api/gen/v1/address"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1/member"
)

var (
	C *config.Config
	L *logger.Logger

	MemberSrvClient  memberpb.MemberServiceClient
	AddressSrvClient addresspb.AddressServiceClient
)
