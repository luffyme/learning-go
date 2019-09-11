package models

import (
	"fmt"

	"gin_blog/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedTime  int `json:"created_time"`
	ModifiedTime int `json:"modified_time"`
	DeletedTime  int `json:"deleted_time"`
}

func Setup() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.Config.Mysql.User,
		setting.Config.Mysql.Pass,
		setting.Config.Mysql.Host,
		setting.Config.Mysql.DB))
	
	if err != nil {
		panic("models.Setup error")
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}