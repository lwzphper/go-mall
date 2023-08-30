package global

import (
	"github.com/lwzphper/go-mall/pkg/logger"
	"github.com/lwzphper/go-mall/server/member/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config *config.Config
	Logger *logger.Logger
)
