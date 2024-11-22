package folder_transfer

import (
	"MDA/client/file_transfer"
	"archive/zip"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func ZipFolder(srcDier string, zipFileName string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return fmt.Errorf("压缩文件夹失败: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.Walk(srcDier, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		// 获取相对路径
		relPath, err := filepath.Rel(srcDier, path)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		if info.IsDir() {
			return nil
		}

		// 打开文件
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		defer file.Close()

		// 创建压缩包中的文件
		zipFileWriter, err := zipWriter.Create(relPath)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		// 将文件内容写入到压缩包中
		_, err = io.Copy(zipFileWriter, file)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	})
}

// 客户端传输文件夹
func TransferFolderZip(folderPath, serverAddr string, conn net.Conn) error {
	zipFileName := "folder.zip"

	// 压缩文件夹
	err := ZipFolder(folderPath, zipFileName)
	if err != nil {
		return fmt.Errorf("压缩文件夹失败：%v", err)
	}

	// 调用单文件传输模块
	err = file_transfer.TransferSmallFile(zipFileName, serverAddr, conn)
	if err != nil {
		return fmt.Errorf("传输文件夹失败：%v", err)
	}

	// 删除临时压缩文件
	os.Remove(zipFileName)

	return nil
}
