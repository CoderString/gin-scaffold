package user

import (
	"context"
	"ship/internal/dao/user"
	"ship/internal/entities/dto"
	"ship/internal/entities/po"
)

type BizUser struct {
	UserDao user.IUserDao
}

func (b *BizUser) Add(ctx context.Context, req *dto.UserDTO) error {
	return b.UserDao.Add(ctx, &po.User{Name: req.Name})
}

func (b *BizUser) Delete(ctx context.Context, id int64) error {
	return b.UserDao.Delete(ctx, id)
}
