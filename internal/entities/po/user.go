package po

import (
	"gorm.io/gorm"
	"time"
)

type Users []*User
type User struct {
	ID   int64  `gorm:"column:id"`   // 自增PK
	Name string `gorm:"column:name"` // 用户名
	Model
}

type UserOption struct {
	ID int64
}

func (u *User) BeforeSave(*gorm.DB) error {
	u.UpdatedAt = time.Now().UnixMilli()
	return nil
}

func (u *User) TableName() string { return "user" }
