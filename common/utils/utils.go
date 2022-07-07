package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

var UpperLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var FullLetters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetAllFile 递归获取指定目录下的所有文件名
func GetAllFile(pathname string) ([]string, error) {
	var fileList = make([]string, 0)
	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return fileList, err
	}
	for _, fi := range fis {
		fullName := filepath.Join(pathname, fi.Name())
		if fi.IsDir() {
			temp, err := GetAllFile(fullName)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullName, err)
				return fileList, err
			}
			fileList = append(fileList, temp...)
			continue
		}
		fileList = append(fileList, fullName)
	}
	return fileList, nil
}

func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func GetRandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = FullLetters[rand.Intn(len(FullLetters))]
	}
	return string(b)
}

func GetRandomUpperString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = UpperLetters[rand.Intn(len(UpperLetters))]
	}
	return string(b)
}
