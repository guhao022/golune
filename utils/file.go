package utils

import (
	"errors"
	"fmt"
	"github.com/russross/blackfriday"
	"os"
	"path/filepath"
)

//递归创建文件
func CreateFile(src string) (string, error) {
	if FileExists(src) {
		return src, nil
	}

	if err := os.MkdirAll(src, 0777); err != nil {
		if os.IsPermission(err) {
			fmt.Println("你不够权限创建文件")
		}
		return "", err
	}

	return src, nil
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullpath string, err error) {
	for _, path := range paths {
		if fullpath = filepath.Join(path, filename); FileExists(fullpath) {
			return
		}
	}
	err = errors.New(fullpath + " not found in paths")
	return
}

func MarkdownCommon(content string) string {
	return string(blackfriday.MarkdownCommon([]byte(content)))
}
