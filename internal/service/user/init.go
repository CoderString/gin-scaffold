package user

import (
	"github.com/google/wire"
	"ship/internal/biz/user"
)

var ProviderSet = wire.NewSet(wire.Struct(new(ServiceUser), "*"))

type ServiceUser struct {
	BizUser user.IUserBiz
}
