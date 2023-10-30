# go_study练手
## Internet_coding
### port_scan
全端口扫描(仅IPV4,地址跟v6暂时不行)
#### 使用
```
cd port_scan
go run main.go
```
#### bug
扫不出来就是你的机魂不悦
![](https://cdn.jsdelivr.net/gh/taosu0216/picgo/bacbd5e567cb6ad575d00a06fa165c2d.jpg)
#### 借鉴于
black hat go
### tcp_socket
最基础的socket练手代码,使用tcp协议,本地(可更改)12345端口进行通信
先运行server,再运行client
### http
先运行server.exe,再运行client.exe(本地打开127.0.0.1:1234也可以)

### web_scoket

一个在线聊天室

## Spider
```go
go run main.go
```
简单的爬虫小程序

## Go_delete

一个自动清空对应文件夹的脚本

搭配小雅食用!

在注释标注处(第91和95行)填入账密,如果想看看浏览器自动操作就把69行的true改成false,只想自动操作的话默认false就行

```go run main.go```执行
或者也可以``` go build -ldflags "-s -w -H windowsgui" -o main.exe main.go```编译成exe文件
