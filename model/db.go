package model

import (
	"blog/utils"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() {
	conn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DBuser, utils.DBpassword, utils.DBhost, utils.DBport, utils.DBdatabase)
	fmt.Println(conn)
	db, err = gorm.Open(utils.DBType, conn)
	if err != nil {
		log.Error("gorm open failed!")
	}

	db.SingularTable(true) // 不设置复数
	db.AutoMigrate(&User{}, &Article{}, &Category{})
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
