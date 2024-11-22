package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	// 服务端地址
	serverAddr := "localhost:8443" // 根据实际服务端地址和端口修改

	// 配置TLS
	config := &tls.Config{ // 用于配置TLS连接的参数
		InsecureSkipVerify: true, // 跳过证书验证
	}

	// 建立TLS连接
	conn, err := tls.Dial("tcp", serverAddr, config)
	if err != nil {
		fmt.Println("连接服务器发生错误:", err)
		return
	}
	defer conn.Close()

	fmt.Println("连接成功")
}
