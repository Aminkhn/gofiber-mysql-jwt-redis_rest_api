package database

import (
	"fmt"
	"log"
	"os"

	"github.com/aminkhn/golang-rest-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlDbInstance struct {
	Db *gorm.DB
}

var Database MysqlDbInstance

func MysqlConnectDb(config *config.Configuration) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUserame, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect db \n", err.Error())
		os.Exit(2)
	}
	log.Println("connection to Database established")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migratons")
	// adding migrations
	Database = MysqlDbInstance{Db: db}
}
