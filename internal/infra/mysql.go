package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql() *gorm.DB {
	dsn := "user:password@tcp(go-auth-template-mysql-1)/go-auth-template?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
