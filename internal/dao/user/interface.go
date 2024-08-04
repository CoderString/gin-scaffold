package user

import (
	"context"
	"github.com/google/wire"
	"ship/internal/entities/po"
)

type IUserDao interface {
	Add(context.Context, *po.User) error
	Delete(context.Context, int64) error
}

var (
	_           IUserDao = &RepoUser{}
	ProviderSet          = wire.NewSet(wire.Struct(new(RepoUser), "*"), wire.Bind(new(IUserDao), new(*RepoUser)))
)
