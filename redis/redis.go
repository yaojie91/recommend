package redis

import (
	"fmt"
	"recommend/config"

	"github.com/gomodule/redigo/redis"
)

var Conn redis.Conn

func InitRedis() {
	fmt.Println("111")
	Conn = redisConnect()
}

func redisConnect() redis.Conn{
	server := fmt.Sprintf("%s:%s", config.Conf.Host, config.Conf.Port)
	conn, err := redis.Dial("tcp", server)
	if err != nil {
		panic(err)
	}
	return conn
}
