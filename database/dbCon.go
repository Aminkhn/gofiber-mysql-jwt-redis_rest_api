package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func MysqlConnectDb() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Faied to Connect db \n", err.Error())
		os.Exit(2)
	}
	log.Println("connection to Database established")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migratons")
	// adding migrations
	Database = DbInstance{Db: db}
}
