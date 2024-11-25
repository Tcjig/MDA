package storage_pathCheck

import (
	"MDA/package/file_function/file_storage/file_storage"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// HandleClient 输入保存文件路径，并调用文件接收函数
func HandleClient(conn net.Conn) {
	defer conn.Close()

	savePath, err := getSavePath()
	if err != nil {
		fmt.Println("获取保存路径失败，具体错误是：", err)
		return
	}

	// 调用接收文件模块
	err = file_storage.ReceiveSmallFile(conn, savePath)
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
