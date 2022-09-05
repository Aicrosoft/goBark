# goBark

监视 Udp 发送过来的消息，并触发相关事件。

# Getting Started

## 运行goBark
```cmd
goBark.exe -h #显示帮助
goBark.exe -c #指定配置文件

```

## config.json 的相关配置

```json
//UdpServer信息
{
  "udpServer": {
    "host": "0.0.0.0",
    "port": 996,
    "blockSize": 1024
  }
}
```

```json
{
  //是否禁用事件捕获。禁用时，整个服务只是相当于一个UdpServer接受器。
  "disable_capture_message":false,
  //处理消息时优先捕获忽略的消息 ，直接跳过。
  "event_message_ignore_keys": ["dnsmasq-dhcp"],
  //不是直接跳过的消息，只要有一个匹配上就返回，所认优先级高的放前面。
  "event_messages": [
    //捕获消息的正则    新构造的消息标题    新构造的消息的内容(支持正则内容替换)
    //json里的正则要用\\ 来转义原始的\ 符号
    { "captureReg": "`(?P<name>[a-zA-Z]+)\\s+(?P<age>\\d+)\\s+(?P<email>\\w+@\\w+(?:\\.\\w+)+)`", "title": "test0", "content": "content2-$1-$2-$3" },
    { "captureReg": "", "title": "test1", "content": "content2-$1-$2" }
  ]
}  
```

# Build and Test

build for current platform and set version:
```cmd
go build -ldflags="-X 'main.Version=0.1.2'" cmd\goBark.go
```

build for current platform and enable debug:
```cmd
go build -ldflags="-X 'main.Version=0.1.3' -X 'main.DebugMode=true'" cmd\goBark.go
```


# Contribute

## Thanks:

- [Visual Studio Code](https://github.com/Microsoft/vscode)



