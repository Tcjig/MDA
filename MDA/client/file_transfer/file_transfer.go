package file_transfer

import (
	"fmt"
	"io"
	"net"
	"os"
)

func TransferSmallFile(filePath, serverAddr string) error {
	// 进行打开文件操作
	file, err := os.Open(filePath) // 根据路径打开文件
	if err != nil {
		fmt.Println("打开文件失败，具体错误是：", err)
		return err
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败，具体错误是：", err)
		return err
	}
	// 传输空文件提示
	if fileInfo.Size() == 0 { // 若文件大小为0
		return fmt.Errorf("传输的文件内容为空，无法进行传输")
	}
	// 连接到服务器
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("连接服务器失败，具体错误是：", err)
	}
	defer conn.Close()

	// 发送文件大小给服务器
	sizeBuffer := fmt.Sprintf("%010d", fileInfo.Size()) // 固定文件大小格式化为10位的字符串
	_, err = conn.Write([]byte(sizeBuffer))             // 将文件大小发给服务器
	if err != nil {
		fmt.Println("发送文件大小失败，具体错误是：", err)
	}

	// 发送文件内容
	buffer := make([]byte, 1024) // 创建容量为1024字节的数组作为缓冲区
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
		for bytesSent < n { // 字节数不为0
			sent, err := conn.Write(buffer[bytesSent:n]) // 从缓冲区取出数据块
			if err != nil {
				fmt.Printf("发送文件块失败，具体错误是：%v", err)
			}
			bytesSent += sent // 更新已发送的字节数
		}
		totalSent += n // 累计已发送的字节数
		// 显示进度
		progress := float64(totalSent) / float64(fileInfo.Size()) * 100 // 将已发送的字节数转为浮点数，并设置成百分之的形式
		fmt.Printf("\r传输进度：%.2f%%", progress)                           // 覆盖掉之前的输入，更新传输进度
	}

	// 确保最后刷新到100%
	fmt.Printf("\r传输进度：100.00%%\n")
	fmt.Println("\n文件传输已完成")
	return nil // 表示文件传输成功
}
