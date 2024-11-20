package role_client

import (
	"MDA/client/file_transfer"
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
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
	fmt.Print("请输入需要传输的文件路径：")
	filePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入传输文件路径错误，具体错误是：", err)
		return
	}

	/*
		// 校验文件路径是否存在
		filePath, err = validateFilePath(filePath)
		if err != nil {
			fmt.Printf("文件路径无效：%v\n", err)
			return
		}
	*/

	// 调用文件传输模块
	err = file_transfer.TransferSmallFile(filePath, "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("文件传输失败，文件路径：%s，错误：%v\n", filePath, err)
		return
	}

	fmt.Println("文件传输成功")
}

// 校验文件
func validateFilePath(path string) (string, error) {
	// 解析绝对路径
	cleanPath := filepath.Clean(path)
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return "", fmt.Errorf("无法解析路径：%w", err)
	}

	// 检查文件是否存在
	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("文件不存在：%s", absPath)
	}
	if err != nil {
		return "", fmt.Errorf("无法读取文件信息：%w", err)
	}

	// 检查是否为文件
	if info.IsDir() {
		return "", fmt.Errorf("路径不是文件：%s", absPath)
	}

	return absPath, nil
}
