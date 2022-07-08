package application

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"useful-tools-golang/common/utils"
)

var (
	dirPath         string
	taskNumber      int
	processTaskChan chan int
	errTaskList     []string
	mutex           sync.Mutex
)

func EditFileNameServRun() {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("please input dir path: ")
	scanner.Scan()
	dirPath = scanner.Text()
	if dirPath == "" {
		fmt.Println("input not allowed to be empty!")
		return
	}
	fmt.Print("please input task number(1-100): ")
	scanner.Scan()
	taskNumber, err = strconv.Atoi(scanner.Text())
	if err != nil || taskNumber > 100 || taskNumber < 1 {
		fmt.Println("task number error!")
		return
	}
	fileList, err := utils.GetAllFile(dirPath)
	if err != nil {
		fmt.Printf("读取文件夹数据失败：%s", err.Error())
		return
	}
	processTaskChan = make(chan int, taskNumber)
	var wg sync.WaitGroup
	wg.Add(len(fileList))
	for idx, fPath := range fileList {
		processTaskChan <- 1
		go EditFileNameByModifyTime(idx, fPath, &wg)
	}
	wg.Wait()
	fmt.Println("Done.")
}

// EditFileNameByModifyTime 根据文件修改时间批量修改文件名称
func EditFileNameByModifyTime(idx int, fPath string, wg *sync.WaitGroup) {
	fileInfo, err := os.Stat(fPath)
	defer func() {
		if err != nil {
			mutex.Lock()
			errTaskList = append(errTaskList, fPath)
			mutex.Unlock()
		}
		<-processTaskChan
		wg.Done()
	}()
	if err != nil {
		fmt.Printf("读取文件%s失败：%s", fPath, err.Error())
		return
	}
	// 根据文件最后修改时间命名
	//_, lastWriteTime, _ := GetFileTimeAttributeForWindows(fileInfo)
	_, lastWriteTime, _ := GetFileTimeAttribute(fileInfo)
	splitStr := strings.Split(fPath, ".")
	if len(splitStr) <= 1 {
		err = errors.New(fmt.Sprintf("不支持文件格式：%s", fPath))
		fmt.Printf(err.Error())
		return
	}
	suffix := splitStr[len(splitStr)-1]
	filename := fmt.Sprintf("%s_%s.%s", lastWriteTime.Format("20060102_150304"), utils.GetRandomUpperString(4), suffix)
	newFilePath := filepath.Join(dirPath, filename)
	if err = os.Rename(fPath, newFilePath); err != nil {
		fmt.Println("修改文件名失败：", err.Error())
		return
	}
	fmt.Printf("[%d] %s\n=>%s\n", idx, fPath, newFilePath)
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
