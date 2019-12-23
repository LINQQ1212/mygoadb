package utils

import "os"

// PathExists 文件或文件是否存在
func PathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// DirStart 文件夹 是否存在不存在则创建
func DirStart(path string) (err error) {
	exist := PathExists(path)
	if !exist {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return
}
