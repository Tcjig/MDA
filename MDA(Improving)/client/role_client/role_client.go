package role_client

import (
	"MDA/client/file_transfer"
	"MDA/client/folder_transfer"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

	// 获取用户输入的传输文件路径
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入需要传输的文件或文件夹路径：")
	filePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入传输文件路径错误，具体错误是：", err)
		return
	}

	filePath = strings.TrimSpace(filePath)

	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("路径无效：%v", err)
		return
	}

	// 根据路径选择传输模式
	// 调用文件夹传输模块
	if info.IsDir() {
		err = folder_transfer.TransferFolderZip(filePath, "127.0.0.1:8080", conn)
		if err != nil {
			fmt.Printf("文件夹传输失败，路径：%s，错误：%v\n", filePath, err)
			return
		}
		fmt.Println("文件夹传输成功")
	} else {
		// 调用单个文件传输模块
		err = file_transfer.TransferSmallFile(filePath, "127.0.0.1:8080", conn)
		if err != nil {
			fmt.Printf("文件传输失败，文件路径：%s，错误：%v\n", filePath, err)
			return
		}
		fmt.Println("文件夹传输成功")
	}

	// 等待服务器确认接收完毕
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer) // 阻塞等待服务器的确认消息
	if err != nil {
		fmt.Println("等待服务器确认失败，具体错误是：", err)
		return
	}

	// 输出服务器的确认消息
	fmt.Println("服务器确认收到文件：", string(buffer))
	fmt.Println("文件传输成功，服务器已接收完毕，连接即将关闭")
}
