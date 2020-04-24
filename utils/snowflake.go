package utils

import (
	"log"
	"sync"
	"time"
)

// 1位 空位 | 41位 时间戳 | 10位 机器ID | 12位 一毫秒生成的编号

var SFIdTool SnowFlake

var once sync.Once //保证改生成器只被实例化一次

type SnowFlake struct {
	machineId int64 //机器ID
	curTS     int64 //当前时间戳 毫秒
	lastTS    int64 //上一毫秒
	order     int64
	lock      sync.Mutex //保证序列号
}

func InitSF(machineId int64) {
	if machineId > 1023 {
		log.Fatal("SnowFlake Id Tool machineId max 1 << 10 .")
	}
	once.Do(func() {
		SFIdTool = SnowFlake{
			machineId: machineId,
			lock:      sync.Mutex{},
			lastTS:    time.Now().UnixNano() / 1e6,
		}
	})
}

func (p *SnowFlake) GetID() int64 {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.curTS = time.Now().UnixNano() / 1e6
	if p.curTS == p.lastTS { //同一毫秒
		p.order++
		if p.order > 4095 {
			time.Sleep(time.Millisecond)
			p.curTS = time.Now().UnixNano() / 1e6
			p.lastTS = p.curTS
			p.order = 0
		}
	} else if p.curTS > p.lastTS {
		p.lastTS = p.curTS
		p.order = 0
	}
	return p.curTS&0x1FFFFFFFFFF << 22 | p.machineId<<12 | p.order
}
