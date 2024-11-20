package file_receive

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

// ReceiveSmallFile 接收单个文件函数
func ReceiveSmallFile(conn net.Conn, savePath string) error {
	// 读取文件大小，固定长度为10字节
	sizeBuffer := make([]byte, 10)  // 创建一个包含10个字节的数组作为一个缓冲区，注意这也意味着只能传10^10大小的文件
	_, err := conn.Read(sizeBuffer) // 从网络读取10个字节的数据
	if err != nil {
		fmt.Println("读取文件大小失败，具体错误是：", err)
	}
	fileSize, err := strconv.ParseInt(string(sizeBuffer), 10, 64) // 将字节数组转换为字符串
	if err != nil {
		fmt.Println("解析文件大小失败，具体错误是：", err)
	}

	// 检查文件大小是否有效
	if fileSize == 0 {
		fmt.Println("接收到的文件大小为0，无法继续")
	}

	// 创建目标文件
	file, err := os.Create(savePath) // 创建一个空文件，保存接收到的文件数据，文件由savePath指定
	if err != nil {
		fmt.Println("创建文件失败，具体错误是：", err)
	}
	defer file.Close()

	// 接收文件数据并写入文件
	buffer := make([]byte, 1024) // 每次读取的缓冲区大小为1024
	totalReceived := int64(0)

	for {
		n, err := conn.Read(buffer) // 从网络中读取数据存入buffer
		if err == io.EOF {          // EOF：文件结束标志
			break
		}
		if err != nil {
			fmt.Println("接收文件数据失败，具体错误是：", err)
		}

		// 写入数据到文件
		byteWritten := 0
		for byteWritten < n {
			written, err := file.Write(buffer[byteWritten:n]) // 将读到的数据写入目标文件
			if err != nil {
				fmt.Println("写入文件失败，具体错误是：", err)
			}
			byteWritten += written
		}

		// 更新总接收字节数
		totalReceived += int64(n) // 累加接收到的总字节数

		// 显示文件接收进度
		progress := float64(totalReceived) / float64(fileSize) * 100
		fmt.Printf("\r接收进度：%.2f%%", progress)
	}
	fmt.Println("\n文件接收完成")
	return nil
}
