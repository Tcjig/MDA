package file_receive

import (
	"fmt"
	"io"
	"net"
	"os"
)

func ReceiveSmallFile(conn net.Conn, savePath string) error {
	// 获取文件大小
	var fileSize int64
	_, err := fmt.Fscanln(conn, &fileSize)
	if err != nil {
		fmt.Println("读取文件大小失败，具体错误是：", err)
	}

	// 创建目标文件
	file, err := os.Open(savePath)
	if err != nil {
		fmt.Println("创建文件失败，具体错误是：", err)
	}
	defer file.Close()

	// 接收文件数据并写入文件
	buffer := make([]byte, 1024) // 设置每次读取的缓冲区大小
	totalReceived := int64(0)

	for {
		n, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("接收文件数据失败，具体错误是：", err)
		}

		// 写入数据到文件
		_, err = file.Write(buffer[:n])
		if err != nil {
			fmt.Println("写入文件失败：", err)
		}
		totalReceived += int64(n)

		// 显示文件接收进度
		progress := float64(totalReceived) / float64(fileSize) * 100
		fmt.Printf("接受进度：%.2%%\r", progress)

		// 文件接收完成
		if err == io.EOF {
			break
		}
	}
	fmt.Println("\n文件接收完成")
	return nil
}
