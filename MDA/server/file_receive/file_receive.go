package file_receive

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func ReceiveSmallFile(conn net.Conn, savePath string) error {
	// 读取文件大小，固定长度为10字节
	sizeBuffer := make([]byte, 10)
	_, err := conn.Read(sizeBuffer)
	if err != nil {
		fmt.Println("读取文件大小失败，具体错误是：", err)
	}
	fileSize, err := strconv.ParseInt(string(sizeBuffer), 10, 64)
	if err != nil {
		fmt.Println("解析文件大小失败，具体错误是：%v", err)
	}

	// 检查文件大小是否有效
	if fileSize == 0 {
		fmt.Println("接收到的文件大小为0，无法继续")
	}

	// 创建目标文件
	file, err := os.Create(savePath)
	if err != nil {
		fmt.Println("创建文件失败，具体错误是：", err)
	}
	defer file.Close()

	// 接收文件数据并写入文件
	buffer := make([]byte, 1024) // 设置每次读取的缓冲区大小
	totalReceived := int64(0)

	for {
		n, err := conn.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("接收文件数据失败，具体错误是：", err)
		}

		// 写入数据到文件
		byteWritten := 0
		for byteWritten < n {
			written, err := file.Write(buffer[byteWritten:n])
			if err != nil {
				fmt.Println("写入文件失败，具体错误是：", err)
			}
			byteWritten += written
		}

		// 更新总接收字节数
		totalReceived += int64(n)

		// 显示文件接收进度
		progress := float64(totalReceived) / float64(fileSize) * 100
		fmt.Printf("\r接收进度：%.2f%%", progress)
	}
	fmt.Println("\n文件接收完成")
	return nil
}
