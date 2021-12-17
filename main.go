package main

import (
	"fmt"
	"os"
	"week-09/client"
	"week-09/server"
)

func main() {
	// 从参数获取要开启的任务，分别为client，server
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("参数欠缺")
		return
	}

	if args[0] == "server" {
		server.Serve()
		return
	}

	if args[0] == "client" {
		client.Request()
		return
	}
}
