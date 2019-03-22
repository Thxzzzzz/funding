package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

var (
	db *gorm.DB
)

func init() {
	InitDB()
}

func InitDB() {
	dbStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlurls"),
		beego.AppConfig.String("mysqldb"))
	var err error
	db, err = gorm.Open("mysql", dbStr)
	if err != nil {
		fmt.Println(err)
		//这里先忽略生产环境的错误，方便测试
		if beego.BConfig.RunMode != "prod" {
			panic("failed to connect database")
		}
		return
	}
	db.LogMode(true)
}

func CloseDB() {
	defer db.Close()
}
