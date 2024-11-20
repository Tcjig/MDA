package role_server

import (
	"MDA/server/file_receive"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

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

	// 服务端等待客户端连接
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("等待客户端连接失败，具体错误是：", err)
		return
	}
	defer conn.Close()
	fmt.Printf("客户端连接成功，通信通道是：%v，客户端信息是：%v\n", conn.RemoteAddr(), conn.LocalAddr())

	// 获取用户输入的保存文件路径
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入需要保存的文件路径：")
	savePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入保存文件路径错误，具体错误是：", err)
		return
	}
	savePath = strings.TrimSpace(savePath)
	if savePath == "" {
		fmt.Println("保存路径不能为空")
		return
	}

	// 调用接收文件模块
	err = file_receive.ReceiveSmallFile(conn, savePath)
	if err != nil {
		fmt.Printf("文件接受失败，保存文件路径为：%s，具体错误是：%v", savePath, err)
		return
	}

	fmt.Println("文件接收成功")
}
