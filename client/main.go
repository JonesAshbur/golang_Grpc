package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 使用rpc.Dial和为服务端建立连接
	coon, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("rpc.Dial:", err)
		return
	}
	defer coon.Close()

	// 调用远程方法
	var reply string
	// 调用远程方法，HelloService.Hello是远程方法名，world是参数，&reply是返回值
	err = coon.Call("HelloService.Hello", "我是client", &reply)
	if err != nil {
		log.Fatal("coon.Call:", err)
		return
	}

	// 打印返回值
	fmt.Println(reply)
}
