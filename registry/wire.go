//+build wireinject

package registry

import (
	"github.com/Mekawy5/chatserv/pkg/application"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitApplicationApi(db *gorm.DB) *application.ApplicationApi {
	wire.Build(application.NewApplicationApi, application.NewApplicationService, application.NewApplicationRepository)
	return &application.ApplicationApi{}
}
