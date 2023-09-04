package global

import (
	"github.com/lwzphper/go-mall/bff/config"
	"github.com/lwzphper/go-mall/pkg/logger"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
)

var (
	C *config.Config
	L *logger.Logger

	MemberSrvClient memberpb.MemberServiceClient
)
