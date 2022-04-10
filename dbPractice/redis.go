package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func RedisSimple() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
		// handle error
	}
	defer func() {
		if err = c.Close(); err != nil {
			panic(err)
		}
	}()

	c.Do("SET", "k1", 1)
	n, _ := redis.Int(c.Do("GET", "k1"))
	fmt.Printf("%#v\n", n)
	n, _ = redis.Int(c.Do("INCR", "k1"))
	fmt.Printf("%#v\n", n)
}

func RedisString() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
		// handle error
	}
	defer c.Close()

	c.Do("SET", "strVal", "This is a string value")
	n, _ := redis.String(c.Do("GET", "strVal"))
	fmt.Printf("%#v\n", n)
}

func RedisNonExisting() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
		// handle error
	}
	defer c.Close()

	n, err := redis.String(c.Do("GET", "noVal"))
	fmt.Printf("%s with error %s\n", n, err)
}

func RedisScan() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
		// handle error
	}
	defer c.Close()

	// here we'll store our iterator value
	iter := 0

	// this will store the keys of each iteration
	var keys []string
	for {

		// we scan with our iter offset, starting at 0
		if arr, err := redis.Values(c.Do("SCAN", iter)); err != nil {
			panic(err)
		} else {
			iter, _ = redis.Int(arr[0], nil)
			keys, _ = redis.Strings(arr[1], nil)
		}

		fmt.Println(keys)

		if iter == 0 {
			break
		}
	}
}

func RedisListType() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
		// handle error
	}
	defer c.Close()

	c.Do("RPUSH", "keyOne", "String number One")
	c.Do("RPUSH", "keyOne", "String number Two")
	c.Do("RPUSH", "keyOne", "String number Three")

	vars, err := redis.Values(c.Do("LRANGE", "keyOne", 0, 2))
	fmt.Println(redis.String(vars[0], err))
	fmt.Println(redis.String(vars[1], err))
	fmt.Println(redis.String(vars[2], err))
	fmt.Println(redis.String(vars[3], err))
}

func main() {
	// RedisSimple()
	// RedisString()
	// RedisNonExisting()
	// RedisScan()
	RedisListType()
}
