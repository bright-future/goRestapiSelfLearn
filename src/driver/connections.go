package driver

import(
	"github.com/gomodule/redigo/redis"
	"fmt"
)

var redisConn redis.Conn

func MakeConnections() {
	redisConn = MakeRedisConnection()
	fmt.Println(redisConn, "this is object")
}

func GetConnections() redis.Conn {
	fmt.Println(redisConn, "this is return object")
	return redisConn
}