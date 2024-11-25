package delivery_pathCheck

import (
	"MDA/package/file_function/file_delivery/file_delivery"
	"bufio"
	"fmt"
	"net"
	"os"
)

func HandleServer(conn net.Conn) {
	// 获取用户输入的传输文件路径
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入需要传输的文件路径：")
	filePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入传输文件路径错误，具体错误是：", err)
		return
	}

	// 调用文件传输模块
	err = file_delivery.TransferSmallFile(filePath, "127.0.0.1:8080", conn)
	if err != nil {
		fmt.Printf("文件传输失败，文件路径：%s，错误：%v\n", filePath, err)
		return
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
