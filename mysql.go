package main

import (
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysql *gorm.DB

func InitMysql() {
	dialect := "mysql"
	host := "127.0.0.1"
	port := "13306"
	user := "root"
	password := "password"
	database := "debug"
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", user, password, host, port, database)

	db, err := gorm.Open(dialect, url)
	if err != nil {
		log.Fatal("MYSQL ERROR: ", err)
	}
	db.SingularTable(true)
	db.BlockGlobalUpdate(true)
	db.LogMode(true)

	mysql = db
}
