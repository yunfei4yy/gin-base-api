package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"yunfei/internal/config"
	"yunfei/internal/crontab"
	"yunfei/internal/event"
)

var (
	Config          *config.Config
	DB              *gorm.DB
	Logger          *zap.Logger
	EventDispatcher *event.Dispatcher
	Crontab         *crontab.Crontab
)
