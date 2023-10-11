package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/lwzphper/go-mall/admin/config"
	"github.com/lwzphper/go-mall/pkg/logger"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1/member"
)

var (
	C *config.Config
	L *logger.Logger
	T ut.Translator

	JwtSecret       []byte
	MemberSrvClient memberpb.MemberServiceClient
)
