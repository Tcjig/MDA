package file_transfer

import (
	"fmt"
	"io"
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
	// 传输空文件提示
	if fileInfo.Size() == 0 {
		return fmt.Errorf("传输的文件内容为空，无法进行传输")
	}
	// 连接到服务器
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("连接服务器失败，具体错误是：", err)
	}
	defer conn.Close()

	// 发送文件大小给服务器
	sizeBuffer := fmt.Sprintf("%010d", fileInfo.Size()) // 固定10位长
	_, err = conn.Write([]byte(sizeBuffer))
	if err != nil {
		fmt.Println("发送文件大小失败，具体错误是：", err)
	}

	// 发送文件内容
	buffer := make([]byte, 1024)
	totalSent := 0
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				break
			}
			fmt.Printf("读取文件失败,具体错误是：%v", err)
		}

		bytesSent := 0
		for bytesSent < n {
			sent, err := conn.Write(buffer[bytesSent:n])
			if err != nil {
				fmt.Println("发送文件块失败，具体错误是：%v", err)
			}
			bytesSent += sent
		}
		totalSent += n
		// 显示进度
		progress := float64(totalSent) / float64(fileInfo.Size()) * 100
		fmt.Printf("\r传输进度：%.2f%%", progress)
	}

	// 确保最后刷新到100%
	fmt.Printf("\r传输进度：100.00%%\n")
	fmt.Println("\n文件传输已完成")
	return nil
}
