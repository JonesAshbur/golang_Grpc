package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// 定义一个远程调用的方法

type HelloService struct {
}

// func只能有两个可序列化参数，第二个参数是指针类型，具体数据类型根据实际需求
// req: 请求参数，res: 响应参数
// func 要求公开，方法名首字母大写
// req和res的类型不能是channel，complex，func，均不能进行序列化
func (this HelloService) Hello(req string, res *string) error {
	fmt.Println(req)
	*res = "Hello " + req
	return nil
}
func main() {
	// 注册rpc服务
	err1 := rpc.RegisterName("HelloService", new(HelloService))
	if err1 != nil {
		log.Fatal("rpc.RegisterName:", err1)
	}

	// 监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		log.Fatal("net.Listen:", err2)
	}

	// 退出时关闭
	defer listener.Close()

	// 建立客户端连接
	for {
		fmt.Println("wait for connection...")
		conn, err3 := listener.Accept()
		if err3 != nil {
			log.Fatal("listener.Accept:", err3)
		}
		// 绑定服务
		rpc.ServeConn(conn)
	}
}
