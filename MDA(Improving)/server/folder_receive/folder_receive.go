package folder_receive

import (
	"MDA/server/file_receive"
	"archive/zip"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

// 解压缩文件
func Unzip(src, dest string) error {
	zipReader, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("解压缩文件失败: %s", err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		// 拼接目标路径
		targetPath := filepath.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			// 创建目录
			os.MkdirAll(targetPath, os.ModePerm)
		} else {
			// 创建文件并写入
			os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("创建文件失败：%v", err)
			}

			rc, err := file.Open()
			if err != nil {
				return fmt.Errorf("%v", err)
			}

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			outFile.Close()
			rc.Close()
		}
	}

	return nil
}

// 服务端接收文件夹
func ReceiveFolder(conn net.Conn, saveDir string) error {
	zipFilePath := filepath.Join(saveDir, "folder.zip")

	// 调用接收单文件模块
	err := file_receive.ReceiveSmallFile(conn, zipFilePath)
	if err != nil {
		return fmt.Errorf("解压失败：%v", err)
	}

	// 删除压缩文件
	os.RemoveAll(zipFilePath)

	return nil
}
