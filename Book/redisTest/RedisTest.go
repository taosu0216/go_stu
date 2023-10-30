package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Inited Redis!")
	err = rdb.Set(ctx, "name", "123", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	fmt.Println("name", val)

	//一下设置很多个键和值,不能设置过期时间,只有用set单独设置时才能设置过期时间
	err = rdb.MSet(ctx, "user_token", 25416531562, "user_cookie", "a6d5415as3zXcsd").Err()
	if err != nil {
		log.Fatalln(err)
	}
	//这里的value是个键值的切片,也就是说循环打印时只能打印键值而获取不到键名
	value, err := rdb.MGet(ctx, "user_token", "user_cookie").Result()
	if err != nil {
		log.Fatalln(err)
	}
	for i, v := range value {
		fmt.Printf("%v is : %v\n", i, v)
	}
	//左侧是顺序,Result是当前队列中元素数量
	Result, err := rdb.RPush(ctx, "list", "test2", "test3").Result()
	if err != nil {
		log.Fatalln(err)
	}
	//右侧是逆序
	fmt.Println("当前元素数量 : ", Result)
	Result, err = rdb.LPush(ctx, "list", "test5", "test6").Result()
	if err != nil {
		log.Fatalln(err)
	}
	//0是起始索引序号,-1是结束索引序号,也就是全部打印
	fmt.Println("当前元素数量 : ", Result)
	elements, err := rdb.LRange(ctx, "list", 0, -1).Result()
	if err != nil {
		log.Fatalln(err)
	}

	// 打印元素
	fmt.Println("list 列表的元素 :")
	for _, element := range elements {
		fmt.Println(element)
	}

	/*
		test6
		test5
		test2
		test3
	*/
}
