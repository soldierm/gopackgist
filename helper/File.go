package helper

import (
	"os"
	"os/exec"
	"path/filepath"
)

//获取当前可执行文件路径
//Golang的相对路径是相对于执行命令时的目录
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

//判断目录或者文件是否存在
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}