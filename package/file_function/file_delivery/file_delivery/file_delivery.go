package file_delivery

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// TransferSmallFile 传输单个文件函数
func TransferSmallFile(filePath, serverAddr string, conn net.Conn) error {

	// 去除掉两边的空格和换行符
	filePath = strings.TrimSpace(filePath)

	// 打开文件
	file, err := os.Open(filePath) // 根据路径打开文件
	if err != nil {
		return fmt.Errorf("打开文件失败：%v", err)
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("获取文件信息失败：%v", err)
	}

	// 检查文件是否为空
	if fileInfo.Size() == 0 { // 若文件大小为0
		return fmt.Errorf("传输的文件内容为空，无法进行传输")
	}

	// 发送文件大小
	sizeBuffer := fmt.Sprintf("%010d", fileInfo.Size()) // 固定文件大小格式化为10位的字符串
	_, err = conn.Write([]byte(sizeBuffer))             // 将文件大小发给服务器
	if err != nil {
		return fmt.Errorf("发送文件大小失败：%v", err)
	}

	// 发送文件内容
	buffer := make([]byte, 1024) // 创建容量为1024字节的数组作为缓冲区
	var totalSent int64 = 0
	for {
		// 读取文件数据
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("读取文件数据失败：%v", err)
		}
		// 读取完退出
		if n == 0 {
			break
		}

		// 发送数据库
		bytesSent := 0
		for bytesSent < n { // 字节数不为0
			sent, err := conn.Write(buffer[bytesSent:n]) // 从缓冲区取出数据块
			if err != nil {
				return fmt.Errorf("发送文件数据失败：%v", err)
			}
			bytesSent += sent // 更新已发送的字节数
		}

		totalSent += int64(n) // 累计已发送的字节数

		// 显示传输进度
		progress := float64(totalSent) / float64(fileInfo.Size()) * 100 // 将已发送的字节数转为浮点数，并设置成百分之的形式
		fmt.Printf("\r传输进度：%.2f%%", progress)                           // 覆盖掉之前的输入，更新传输进度
	}

	// 确保最后刷新到100%
	fmt.Printf("\r传输进度：100.00%%\n")
	fmt.Println("\n文件传输已完成")

	return nil // 表示文件传输成功
}
