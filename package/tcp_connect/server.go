package tcp_connect

import (
	"MDA/package/file_function/file_storage/storage_pathCheck"
	"fmt"
	"net"
)

// Server 监听客户端及等待客户端连接
func Server() {
	fmt.Println("服务端启动中，请稍后...")

	// 服务端进行监听
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听失败，具体错误是：", err)
		return
	}
	defer listen.Close()
	fmt.Println("监听客户端成功，等待客户端连接...")

	// 使用for循环，以便不间断接收客户端连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接收客户端连接失败，具体错误是：", err)
			continue
		}
		fmt.Printf("客户端连接成功，客户端地址：%v\n", conn.RemoteAddr())

		// 为每个客户端启动独立的协程
		go storage_pathCheck.HandleClient(conn)
	}
}
