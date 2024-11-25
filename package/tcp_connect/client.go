package tcp_connect

import (
	"MDA/package/file_function/file_delivery/delivery_pathCheck"
	"fmt"
	"net"
)

// Client 请求与服务端函数，并调用传输文件函数
func Client() {
	fmt.Println("客户端启动中，请稍后...")

	// 连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接服务端失败，具体错误是：", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接服务端成功，通信通道为：", conn.RemoteAddr())

	delivery_pathCheck.HandleServer(conn)
}
