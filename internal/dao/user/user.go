package user

import (
	"context"
	"gorm.io/gorm"
	"ship/internal/dao"
	"ship/internal/entities/po"
	"ship/pkg/reply"
)

type RepoUser struct {
	Handle *dao.DbClient
}

func NewRepoUser() IUserDao {
	return &RepoUser{
		Handle: dao.NewDbClient(),
	}
}

func (r *RepoUser) Add(ctx context.Context, user *po.User) error {
	return r.Handle.MysqlClient.WithContext(ctx).Create(user).Error
}

func (r *RepoUser) Delete(ctx context.Context, id int64) error {
	return r.Handle.MysqlClient.WithContext(ctx).Model(po.User{}).Delete("id = ?", id).Error
}

func (r *RepoUser) condition(ctx context.Context, opt *po.UserOption) (*gorm.DB, error) {
	if opt == nil {
		return nil, reply.ErrSQLCondition
	}
	db := r.Handle.MysqlClient.WithContext(ctx)
	if opt.ID > 0 {
		db = db.Where("id = ?", opt.ID)
	}
	return db, nil
}
