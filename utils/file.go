package utils

import "os"

// CheckFolderExists 检查文件夹是否存在，如果不存在则创建
func CheckFolderExists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

// CheckPathExists 判断文件/文件夹是否存在
func CheckPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
