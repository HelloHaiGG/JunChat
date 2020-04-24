package igorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
var DbClient *gorm.DB

func Init(option *IOption) {
	option.init()
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		option.User,
		option.Password,
		option.Host,
		option.Port,
		option.DB))
	if err != nil {
		log.Fatal("connect mysql err")
	}
	db.LogMode(option.IsDebug)
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(50)  //最大空闲连接数
	db.DB().SetMaxIdleConns(120) //连接池最大连接数
	//自动建表
	db.AutoMigrate()
	DbClient = db
}

type IOption struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
	IsDebug  bool
}

func (p *IOption) init() {
	if p.Host == "" {
		p.Host = "127.0.0.1"
	}
	if p.Port == 0 {
		p.Port = 3306
	}
}
