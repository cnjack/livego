# livego
简单高效的直播服务器：
- 安装和使用非常简单；
- 纯 Golang 编写，性能高，跨平台；
- 支持常用的传输协议、文件格式、编码格式；

#### supported transport protocol
- [x] RTMP
- [x] AMF
- [x] HLS
- [x] HTTP-FLV

#### supported layout
- [x] FLV
- [x] TS

#### supported coding formant 
- [x] H264
- [x] AAC
- [x] MP3

#### build from resource
1. pull the code `git clone https://github.com/cnjack/livego.git`
```
cd livego
export GOPATH=$(pwd):$GOPATH
cd src/app
go build
```

## usage


### thx
fork from [gwuhaolin/livego](https://github.com/gwuhaolin/livego)
