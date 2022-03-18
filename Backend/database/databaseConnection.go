package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName     string = "root"
	Password     string = "admin"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "URL"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func CreateDatabase() *gorm.DB {
	log.Println("Database connecting ..........")
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Addr, Port, Database)
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		log.Println("Database connection Failed", err)
		return nil
	}
	return db
}
