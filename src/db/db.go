package db

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	RestDB struct{}
)

var (
	err     error
	connect string
	db      *gorm.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	connect = os.Getenv("CONNECT")
}

func (r *RestDB) Sql() (con *gorm.DB) {
	db, err = gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	sqlConf, _ := db.DB()
	sqlConf.SetMaxIdleConns(1)
	sqlConf.SetMaxOpenConns(1)
	sqlConf.SetConnMaxLifetime(10 * time.Minute)

	con = db
	return
}
