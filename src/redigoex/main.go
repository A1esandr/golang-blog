package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "hello", "world")
	if err != nil {
		fmt.Println(err)
		return
	}
	s, err := redis.String(c.Do("GET", "hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", s)
}
