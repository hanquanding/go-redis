package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

func main() {
	cli, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	reply, err := cli.Do("SET", "hello", "hello  golang")
	log.Println(reply, err)

	r, e := redis.Bool(cli.Do("DEL", "hello"))
	log.Println(r, e)
}
