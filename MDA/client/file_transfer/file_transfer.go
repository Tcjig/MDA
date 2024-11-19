package file_transfer

import (
	"fmt"
	"net"
	"os"
)

func TransferSmallFile(filePath, serverAddr string) error {
	// 进行打开文件操作
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件失败，具体错误是：", err)
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败，具体错误是：", err)
	}
	// 连接到服务器
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("连接服务器失败，具体错误是：", err)
	}
	defer conn.Close()

	// 发送文件大小给服务器
	n, err := conn.Write([]byte(fmt.Sprintf("%d", fileInfo.Size())))
	if err != nil {
		fmt.Println("发送文件大小失败，具体错误是：", err)
	}
	fmt.Printf("发送的文件大小为%v", n)
	buffer := make([]byte, 1024)
	totalSent := 0
	for {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Printf("读取文件失败：%s", err)
		}
		if n > 0 {
			_, err = conn.Write(buffer[:n])
			if err != nil {
				fmt.Printf("发送文件块失败：%s", err)
			}
			totalSent += n

			// 显示进度
			progress := float64(totalSent) / float64(fileInfo.Size()) * 100
			fmt.Printf("传输进度：%.2f%%\r", progress)
		}

		if err != nil {
			break
		}
	}

	fmt.Println("\n文件传输已完成")
	return nil
}
