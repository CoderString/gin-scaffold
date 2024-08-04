package user

import (
	"context"
	"github.com/google/wire"
	"ship/internal/entities/dto"
)

type IUserBiz interface {
	Add(context.Context, *dto.UserDTO) error
	Delete(context.Context, int64) error
}

var (
	ProviderSet          = wire.NewSet(wire.Struct(new(BizUser), "*"), wire.Bind(new(IUserBiz), new(*BizUser)))
	_           IUserBiz = &BizUser{}
)
