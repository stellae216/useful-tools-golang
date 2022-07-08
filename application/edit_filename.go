package application

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"useful-tools-golang/common/utils"
)

func EditFileNameServRun() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("please input dir path: ")
	scanner.Scan()
	dirPath := scanner.Text()
	if dirPath == "" {
		fmt.Println("input not allowed to be empty!")
		return
	}
	fileList, err := utils.GetAllFile(dirPath)
	if err != nil {
		fmt.Println("读取文件夹数据失败：", err)
		return
	}
	if err = EditFileNameByModifyTime(dirPath, fileList); err != nil {
		fmt.Println("err：", err)
		return
	}
	fmt.Println("Done.")
}

// EditFileNameByModifyTime 根据文件修改时间批量修改文件名称
func EditFileNameByModifyTime(dirPath string, fileList []string) error {
	// todo 多线程处理任务
	for idx, fPath := range fileList {
		fmt.Printf("[%d] %s\n", idx, fPath)
		fileInfo, err := os.Stat(fPath)
		if err != nil {
			return errors.New(fmt.Sprintf("读取文件%s失败：%s", fPath, err.Error()))
		}
		// 根据文件最后修改时间命名
		//_, lastWriteTime, _ := GetFileTimeAttributeForWindows(fileInfo)
		_, lastWriteTime, _ := GetFileTimeAttribute(fileInfo)
		splitStr := strings.Split(fPath, ".")
		if len(splitStr) <= 1 {
			return errors.New(fmt.Sprintf("不支持文件格式：%s", fPath))
		}
		suffix := splitStr[len(splitStr)-1]
		filename := fmt.Sprintf("%s_%s.%s", lastWriteTime.Format("20060102_150304"), utils.GetRandomUpperString(4), suffix)
		newFilePath := filepath.Join(dirPath, filename)
		if err = os.Rename(fPath, newFilePath); err != nil {
			return errors.New(fmt.Sprintf("修改文件名失败：", err.Error()))
		}
		fmt.Println("=>", newFilePath)
	}
	return nil
}

// GetFileTimeAttribute 获取文件时间属性：创建时间、最后修改时间、最后访问时间
func GetFileTimeAttribute(fileInfo os.FileInfo) (ct, lwt, lat time.Time) {
	//// windows
	//winFileAttr := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	//ct = utils.SecondToTime(winFileAttr.CreationTime.Nanoseconds() / 1e9)
	//lwt = utils.SecondToTime(winFileAttr.LastWriteTime.Nanoseconds() / 1e9)
	//lat = utils.SecondToTime(winFileAttr.LastAccessTime.Nanoseconds() / 1e9)
	linuxFileAttr := fileInfo.Sys().(*syscall.Stat_t)
	ct = utils.SecondToTime(linuxFileAttr.Ctimespec.Sec)
	lwt = utils.SecondToTime(linuxFileAttr.Mtimespec.Sec)
	lat = utils.SecondToTime(linuxFileAttr.Atimespec.Sec)
	return
}
