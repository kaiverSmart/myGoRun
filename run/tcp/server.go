package main

import "github.com/garyburd/redigo/redis"
import "time"
import (
	"fmt"
	"net"
)

func main() {
	//connectBegin()
	initRedis("localhost:13324",100,100,time.Second*300)
}

var pool *redis.Pool

func initRedis(addr string, idleConn, maxConn int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     idleConn,
		MaxActive:   maxConn,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			fmt.Println("redis")
			return redis.Dial("tcp", addr)
		},
	}
	return
}

func GetConn() redis.Conn {
	return pool.Get()
}

func PutConn(conn redis.Conn) {
	conn.Close()
}

//链接
func connectBegin()  {
	lis,err := net.Listen("tcp","localhost:12345")
	if err != nil {
		fmt.Println("listen err:",err)
	}
	//接受数据
	for {
		connect,conErr := lis.Accept()
		if conErr != nil{
			fmt.Println("con err",conErr)
		}
		go conPrecess(connect)
	}
}

func conPrecess(con net.Conn) {
	defer con.Close()
	buf := make([]byte,512)
	_,err := con.Read(buf)
	if err != nil {
		fmt.Println("conprecess",err)
	}
	fmt.Println("buf:",string(buf))
}
//处理链接