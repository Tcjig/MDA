package role_client

import (
	"MDA/client/file_transfer"
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client() {
	fmt.Println("客户端启动中，请稍后...")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接服务端失败，具体错误是：", err)
	}
	fmt.Println("连接服务端成功，通信通道为：", conn.RemoteAddr())

	// 获取用户输入的传输文件路径
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入需要传输的文件路径：")
	filePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入传输文件路径错误，具体错误是：", err)
	}
	// 去除掉路径末尾的换行符
	filePath = filePath[:len(filePath)-1]

	err = file_transfer.TransferSmallFile(filePath, conn.RemoteAddr().String())
	if err != nil {
		fmt.Println("文件传输失败，具体错误是：", err)
	}
}
