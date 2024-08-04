package biz

import (
	"github.com/google/wire"
	"ship/internal/biz/user"
)

var ProviderSet = wire.NewSet(
	user.ProviderSet,
)
