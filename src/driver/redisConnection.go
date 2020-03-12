package driver

import(
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func MakeRedisConnection() redis.Conn {
	pool := newPool()
	conn := pool.Get()
	// defer conn.Close()
	err := ping(conn)
	if err != nil {
		fmt.Println(err)
	}
	return conn
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func ping(c redis.Conn) error {
	// Send PING command to Redis
	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}