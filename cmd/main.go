// 程序的主入口，让用户选择服务端还是客户端
package main

import (
	"MDA/package/tcp_connect"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		// 展示菜单
		displayMenu()

		// 获取用户输入
		choice, err := getUserChoice()
		if err != nil {
			fmt.Println("输入无效，请输入有效的数字选项 (1-3)：")
			continue
		}

		// 根据选项执行操作
		switch choice {
		case 1:
			fmt.Println("进入服务端模式...")
			tcp_connect.Server()
		case 2:
			fmt.Println("进入客户端模式...")
			tcp_connect.Client()
		case 3:
			fmt.Println("正在退出MDA...")
			os.Exit(0)
		default:
			fmt.Println("无效选择，请输入 1, 2 或 3")
		}
	}
}

// 展示菜单
func displayMenu() {
	fmt.Println("\n欢迎来到MDA，请输入您的选项，共有四个选项，如下所示：")
	fmt.Println("选项1：作为服务端接收文件")
	fmt.Println("选项2：作为客户端传输文件")
	fmt.Println("选项3：退出MDA")
	fmt.Print("请输入您的选择：")
}

// 获取用户输入的选择
func getUserChoice() (int, error) {
	reader := bufio.NewReader(os.Stdin) // 创建读取器

	// 读取输入内容
	choiceStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("读取输入失败：%v", err)
	}

	// 去掉多余的空格和换行符
	choiceStr = strings.TrimSpace(choiceStr)

	// 转换为整数
	choice, err := strconv.Atoi(choiceStr)
	if err != nil {
		return 0, fmt.Errorf("无法将输入解析为数字：%v", err)
	}

	return choice, nil
}
