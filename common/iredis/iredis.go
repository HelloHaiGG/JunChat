package iredis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

//单点redis
var RedisCli *redis.Client

func Init(opt *IOptions) {
	//初始默认配置
	opt.init()
	RedisCli = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", opt.Host, opt.Port),
		Password:    opt.Password,
		DB:          opt.DB,
		MaxRetries:  opt.MaxRetry,
		DialTimeout: opt.DialTimeOut,
		MaxConnAge:  opt.MaxConnAge,
	})
	if RedisCli.Ping().Err() != nil {
		log.Fatal("Redis初始化失败.",RedisCli.Ping().Err())
	}
}

type IOptions struct {
	Host     string
	Port     int
	DB       int
	Password string
	MaxRetry int
	//秒
	DialTimeOut time.Duration
	//空闲连接的保活时常 分钟
	MaxConnAge time.Duration
}

//初始化默认值
func (p *IOptions) init() {
	if p.Host == "" {
		p.Host = "127.0.0.1"
	}
	if p.Port == 0 {
		p.Port = 6379
	}
	if p.DB == 0 {
		p.DB = 0
	}
	if p.Password == "" {
		p.Password = "root"
	}
	if p.MaxRetry == 0 {
		p.MaxRetry = 5
	}
	if p.DialTimeOut == 0 {
		p.DialTimeOut = 10 * time.Second
	} else {
		p.DialTimeOut = p.DialTimeOut * time.Second
	}
	//五分钟保活
	if p.MaxConnAge == 0 {
		p.MaxConnAge = time.Minute * 5
	} else {
		p.MaxConnAge = time.Minute * p.MaxConnAge
	}
}
