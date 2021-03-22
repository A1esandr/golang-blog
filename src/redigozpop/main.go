package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// dial wraps DialDefaultServer() with a more suitable function name for examples.
func dial() (redis.Conn, error) {
	return redis.Dial("tcp", ":6379")
}

func main() {
	c, err := dial()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	_, err = c.Do("SADD", "set_with_integers", 4, 5, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	ints, err := redis.Ints(c.Do("SMEMBERS", "set_with_integers"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", ints)
	// Output: []int{4, 5, 6}
}
