package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	// 加载证书和私钥
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem") // 从两个文件中加载证书和私钥
	if err != nil {
		fmt.Println("加载证书错误:", err)
		return
	}

	// 配置TLS参数
	config := &tls.Config{Certificates: []tls.Certificate{cert}} // 仅仅加载一个证书对应的私钥

	// 创建TLS监听器
	listener, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		fmt.Println("创建错误监听器:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TLS监听的端口号为8443...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接错误:", err)
			continue
		}
		fmt.Println("连接成功", conn.RemoteAddr())
	}
}
