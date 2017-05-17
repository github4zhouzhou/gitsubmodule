package main

import (
	"fmt"
	"time"

	"dsp.ohko.cn/zhouzhou/vcommon/goroutinepool"
)

func main() {
	goPool := goroutinepool.NewPool(10, 100)
	goPool.Start()

	goPool.AddTask(func() {
		fmt.Println("test")
	})

	fmt.Println("test")

	time.Sleep(time.Second)
}
