package service

import (
	"github.com/google/wire"
	"ship/internal/service/user"
)

var ProviderSet = wire.NewSet(wire.Struct(new(AllService), "*"))

type AllService struct {
	UserSrv *user.ServiceUser
}
