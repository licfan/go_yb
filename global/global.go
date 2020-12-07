package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go_yb/config"
	"gorm.io/gorm"
)

var (
	YB_DB *gorm.DB
	YB_REDIS *redis.Client
	YB_VP *viper.Viper
	YB_LOG *zap.Logger
	YB_CONFIG config.Server
)