package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//判断文件是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

//读取文件的内容
func HandFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//按行读取文件
func ReadFileForLine(path string) ([]string, error) {
	if !IsExist(path) {
		return nil, fmt.Errorf(fmt.Sprintf("%s不存在", path))
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	res := make([]string, 0)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}
		res = append(res, str)
	}
	return res, nil
}


