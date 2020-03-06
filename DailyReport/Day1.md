# Day1

## 开发环境

- 操作系统：Ubuntu Server 18.04.1 LTS 64位
- 使用工具：`docker` , `Coolq http api`

## 完成内容

### 环境相关

保持了白嫖的良好传统，从腾讯云又白嫖了一个云服务器用于这次的实习

pull这个api的过程极慢，pull完成之后保存了一个镜像

### docker相关

- 建好了存放`Coolq http apt`的容器，并通过以下代码开始运行，目前还不清楚宿主目录挂载用于持久化程序文件是什么意思，但是按照这个已经能成功运行`Coolq`平台

```bash
docker run -ti --rm --name http-api \
             -v $(pwd)/coolq:/home/ubuntu/coolq \  # 将宿主目录挂载到容器内用于持久化 酷Q 的程序文件
             -p 9000:9000 \  # noVNC 端口，用于从浏览器控制 酷Q
             -p 5700:5700 \  # HTTP API 插件开放的端口
             -e COOLQ_ACCOUNT=1092443987 \ # 要登录的 QQ 账号，可选但建议填
             -e CQHTTP_POST_URL=http://127.0.0.1:8080 \  # 事件上报地址
             -e CQHTTP_SERVE_DATA_FILES=yes \  # 允许通过 HTTP 接口访问 酷Q 数据文件
             richardchien/cqhttp:latest            
```

- 从docker中另外新建一个容器，用于存放自己写的插件，获取前一个容器给出的信息并进行处理然后返回
- 目前止步于golang程序的编写，明天需要找点例子来看看

## 目前想明白的

1. 酷q是个平台，能够登陆qq并且返回信息
2. 酷q上有GitHub上开发的sdk（软件开发工具包）
3. 目前要做的事情就是基于酷q上golang的开发包进行编写程序
4. 需要实现的一阶段功能就是简单能回复就行
5. 需要两个容器，一个用于存放api的东西，一个用于处理数据，两个容器通过某些（还不知道的）的方法进行通信
6. api存放的东西会通过http或者是websock上报信息，http比较了解，但是websock还是不明不白，还是http吧

## 目前已经获取的资源

[Coolq http api 使用说明书](https://github.com/richardchien/coolq-http-api)

[Golang 开发Coolq接口用的SDK](https://github.com/catsworld/qq-bot-api)

[qqbotapi package的使用说明](https://godoc.org/github.com/catsworld/qq-bot-api#pkg-subdirectories)

