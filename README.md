## Sweet Base64

[![Security Status](https://www.murphysec.com/platform3/v3/badge/1617956857323237376.svg?t=1)](https://www.murphysec.com/accept?code=b18b786266a0ffa1ff05d078b6b42b97&type=1&from=2&t=2)

#### 1. 一个简单的Base64编码解码工具

**A Simple Base64 Encode & Decode Tool**

#### 2.易用 轻量 安全 快速部署

**Easy  &  Lightweight  &  Safe  &  FastToDeploy**

#### 3.网页和后台服务分离

**2022/10/18 15:11:32 [Backend] Listening  :6060**
**2022/10/18 15:11:32 [Frontend] Listening  :6062**

#### 4.用户界面展示请点击代码仓库中的 UI_demo.jpg


备注：

> 1. 前端页面使用了vue和axios的cdn资源，因此使用体验可能会受限于实际的网速
> 2. 不建议上传较大的文件进行编码
> 3. 本项目源码不涉及对文本和文件的存储
> 3. 部署时仅需在项目根目录下执行 ./Windows_ver.exe 或 ./Linux_ver



在项目根目录下，生成可执行文件

> windows

~~~
go env -w GOOS=windows
go build -o Windows_ver.exe .
~~~

> linux

~~~
go env -w GOOS=linux
go build -o Linux_ver .
~~~

