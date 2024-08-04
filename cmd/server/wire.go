package main

import (
	"github.com/google/wire"
	"ship/internal/biz"
	"ship/internal/dao"
	userDao "ship/internal/dao/user"
	"ship/internal/service"
	userService "ship/internal/service/user"
)

func initApp2() (*service.AllService, func(), error) {
	panic(wire.Build(
		dao.NewDbClient,
		userDao.ProviderSet,
		biz.ProviderSet,
		userService.ProviderSet,
		service.ProviderSet,
	))
}
