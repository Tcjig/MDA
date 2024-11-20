// 程序的主入口，让用户选择服务端还是客户端
package main

import (
	"MDAtest/client/role_client"
	"MDAtest/server/role_server"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // 读取用户输入

	fmt.Println("欢迎来到MDA，请输入您的选项，共有四个选项，如下所示")
	fmt.Println("选项1：作为服务端接收文件")
	fmt.Println("选项2：作为客户端传输文件")
	fmt.Println("选项3：退出MDA")
	fmt.Print("请输入您的选择：")

	choiceStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入失败，具体错误是：", err)
		return
	}

	choiceStr = strings.TrimSpace(choiceStr) // 去掉用户输入两侧的换行符和多余空格
	choice, err := strconv.Atoi(choiceStr)
	if err != nil {
		fmt.Println("无效输入，具体错误是：", err)
	}

	// 根据用户的不同选项进入到不同的功能当中
	if choice == 1 {
		// 作为服务器端
		role_server.Server()
	} else if choice == 2 {
		// 作为客户端
		role_client.Client()
	} else if choice == 3 {
		// 退出程序
		fmt.Println("正在退出MDA中...")
		os.Exit(0)
	} else {
		fmt.Println("无效选择，自动退出程序中...")
		os.Exit(0)
	}
}
