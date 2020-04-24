package iredis

import (
	"errors"
	"time"
)

func RLock(key string, duration int) error {
	if RedisCli == nil {
		return errors.New("redis cli nil")
	}
	if b, _ := RedisCli.SetNX(key, "redis-lock", time.Duration(duration)*time.Second).Result(); !b {
		return errors.New("redis lock")
	}
	return nil
}

func RUnlock(key string) error {
	if RedisCli == nil{
		return errors.New("redis cli nil")
	}
	_, _ = RedisCli.Del(key).Result()
	return nil
}
