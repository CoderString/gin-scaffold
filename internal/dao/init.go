package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *DbClient

type DbClient struct {
	MysqlClient *gorm.DB
}

func NewDbClient() *DbClient {
	dbClient = &DbClient{}
	dbClient.MysqlClient = InitDB()
	return dbClient
}

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/ship?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
