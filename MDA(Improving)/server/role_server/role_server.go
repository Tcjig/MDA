package role_server

import (
	"MDA/server/file_receive"
	"MDA/server/folder_receive"
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// Server 监听客户端及等待客户端连接，并调用文件接收函数
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

		// 为每个客户端启动独立的协程：选择接收方式
		reader := bufio.NewReader(conn)
		choiceStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败：", err)
			return
		}

		// 去掉多余的空格和换行符
		choiceStr = strings.TrimSpace(choiceStr)

		// 转换为整数
		serverChoice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("无法将输入解析为数字：", err)
			return
		}
		switch serverChoice {
		case 1:
			go handleClientFile(conn)
		case 2:
			go handleClientFolder(conn)
		}
	}
}

// 处理客户端的连接，并进行单文件接收
func handleClientFile(conn net.Conn) {
	defer conn.Close()

	savePath, err := getSavePath()
	if err != nil {
		fmt.Println("获取保存路径失败，具体错误是：", err)
		return
	}

	// 调用接收文件模块
	err = file_receive.ReceiveSmallFile(conn, savePath)
	if err != nil {
		fmt.Printf("文件接收失败，保存文件路径为：%s,具体错误是：%v\n", savePath, err)
		return
	}

	fmt.Printf("文件接收成功，已保存至：%s\n", savePath)

	// 发送确认消息给客户端
	confirmationMessage := "文件接收成功"
	_, err = conn.Write([]byte(confirmationMessage)) // 向客户端发送确认消息
	if err != nil {
		fmt.Printf("发送确认消息失败，具体错误是：%v\n", err)
		return
	}
	fmt.Println("已向客户端发送文件接收成功的确认消息")
}

// 处理客户端的连接，并进行文件夹的接收
func handleClientFolder(conn net.Conn) {
	defer conn.Close()

	savePath, err := getSavePath()
	if err != nil {
		fmt.Println("获取保存路径失败，具体错误是：", err)
		return
	}

	// 调用接收文件模块
	err = folder_receive.ReceiveFolder(conn, savePath)
	if err != nil {
		fmt.Printf("文件夹接收失败，保存文件路径为：%s,具体错误是：%v\n", savePath, err)
		return
	}

	fmt.Printf("文件接收成功，已保存至：%s\n", savePath)

	// 发送确认消息给客户端
	confirmationMessage := "文件接收成功"
	_, err = conn.Write([]byte(confirmationMessage)) // 向客户端发送确认消息
	if err != nil {
		fmt.Printf("发送确认消息失败，具体错误是：%v\n", err)
		return
	}
	fmt.Println("已向客户端发送文件接收成功的确认消息")
}

// 获取用户的保存路径
func getSavePath() (string, error) {
	// 获取用户输入的保存文件地址
	fmt.Print("请输入需要保存文件的路径：")
	reader := bufio.NewReader(os.Stdin)

	// 读取用户输入
	savePath, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// 去掉多余的空格和换行符
	savePath = strings.TrimSpace(savePath)
	if savePath == "" {
		return "", fmt.Errorf("保存路径不能为空")
	}

	// 效验路径合法性
	if !isValidPath(savePath) {
		return "", fmt.Errorf("保存路径无效：%s", savePath)
	}

	return savePath, nil
}

// 校验路径是否合法
func isValidPath(path string) bool {
	illegalChars := []string{"*", "?", "<", ">", "|"}
	for _, char := range illegalChars {
		if strings.Contains(path, char) {
			return false
		}
	}
	return true
}
