package file_storage

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
	if err != nil && err != io.EOF {
		return fmt.Errorf("读取文件大小失败: %v", err)
	}

	fmt.Printf("接收到的文件大小数据：%s\n", string(sizeBuffer)) // 打印接收到的文件大小数据

	fileSize, err := strconv.ParseInt(string(sizeBuffer), 10, 64)
	if err != nil {
		return fmt.Errorf("解析文件大小失败：%v", err)
	}

	// 检查文件大小是否有效
	if fileSize <= 0 {
		return fmt.Errorf("文件大小无效：%d", fileSize)
	}

	// 创建目标文件
	file, err := os.Create(savePath) // 创建一个空文件，保存接收到的文件数据，文件由savePath指定
	if err != nil {
		return fmt.Errorf("创建文件失败：%v", err)
	}
	defer file.Close()

	// 接收文件内容
	fmt.Printf("开始接收文件，目标路径：%s，文件大小：%d\n", savePath, fileSize)
	buffer := make([]byte, 1024) // 每次读取的缓冲区大小为1024
	var totalReceived int64 = 0

	for totalReceived < fileSize {
		n, err := conn.Read(buffer) // 从网络中读取数据存入buffer
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("接收文件数据失败：%v", err)
			} else {
				break // 若读取到EOF说明数据已完成
			}
		}

		// 将读取的数据写入文件
		_, err = file.Write(buffer[:n])
		if err != nil {
			return fmt.Errorf("写入文件失败：%v", err)
		}

		// 更新总接收字节数
		totalReceived += int64(n) // 累加接收到的总字节数

		// 显示文件接收进度
		progress := float64(totalReceived) / float64(fileSize) * 100
		fmt.Printf("\r接收进度：%.2f%%", progress)

		//处理EOF
		if err == io.EOF {
			break
		}
	}
	fmt.Println("\n文件接收完成")

	return nil
}
