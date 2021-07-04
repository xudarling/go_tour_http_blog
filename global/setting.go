package global

import (
	"github.com/jinzhu/gorm"
	"go_tour/blog_service/pkg/logger"
	"go_tour/blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)

var (
	DBEngine	 *gorm.DB
	Logger       *logger.Logger
)
