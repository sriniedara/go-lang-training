package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBConnection() (interface{}, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/mygodb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
