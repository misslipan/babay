package main

import (
	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func Initredis() {
	Pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

}
