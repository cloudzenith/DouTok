---
sidebar_position: 2
---

# 快速启动DouTok

## 准备环境

本教程将带领你从零开始，循序渐进搭建并启动 `DouTok` 项目，若读者已具备相关知识，可选择性阅读。

## 环境准备

1. Golang 1.22+
   > - [https://golang.org/dl/](https://golang.org/dl/)
   > - [https://golang.google.cn/dl/](https://golang.google.cn/dl/)
2. Node 14.17+
   > - [https://nodejs.org/en/download/](https://nodejs.org/en/download/)
3. React.js + Next.js
   > - [https://reactjs.org/](https://reactjs.org/)
   > - [https://nextjscn.org/](https://nextjscn.org/)
4. JetBrains GoLand/WebStorem
   > - [https://www.jetbrains.com/](https://www.jetbrains.com/)
5. VSCode
   > - [https://code.visualstudio.com/](https://code.visualstudio.com/)
6. Docker
   > - [https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)

## 必要组件配置及启动

- Consul: 通过`backend/gopkgs/launcher`提供能力，所有后端服务均自动注册到Consul中
- Redis: 缓存
- MySQL: 持久化存储
- MinIO: 对象存储
- RocketMQ: 消息队列（不是必须）

1. 找到`env/basic.yml`文件，通过命令`docker-compose -f ./env/basic.yml up -d`启动Consul, Redis, MySQL, MinIO
（2、3步不是必须）
2. 找到`env/rocketmq/broker.conf`文件，将`brokerIP1`修改为本地局域网IP
3. 找到`env/rocketmq.yml`文件，通过命令`docker-compose -f ./env/rocketmq.yml up -d`启动RocketMQ

## MySQL库表结构同步

1. 进入`sql`目录
2. 检查`sql/Makefile`文件，其中涉及的MySQL连接需注意应与本地环境一致
3. 安装 goose 工具，执行`go install github.com/pressly/goose/v3/cmd/goose@latest`
4. 执行`make up`命令，MySQL库表结构会同步到本地

## 启动后端服务

### 编译运行

1. 进入`backend`目录下除`gopkgs`外的所有服务目录，依次 `go run cmd/main.go` 启动服务

### 镜像运行

1. 进入`backend`目录下除`gopkgs`外的所有服务目录，执行`make build`以编译Docker镜像
2. 进入`env`目录，检查`configs`下各个配置文件，应与本地环境保持一致，特别是 `./baseservice/config.yaml` 中，`minio.default.host` 需要改成本机局域网IP
3. 进入`env`目录，执行`docker-compose -f backend.yml up -d`启动所有后端服务

## 启动前端服务

1. 进入`frontend/doutok`目录，执行`pnpm install`安装依赖
2. 执行`pnpm dev`启动前端服务，通过 [http://localhost:23000](http://localhost:23000) 访问

## 开始开发

### 后端开发过程中需额外安装的工具

在后端开发的过程中，有一些地方会需要安装一些额外的开源工具，具体如下：

|用途|工具|安装方式|备注|
|:---:|:---:|:---:|:---:|
|数据库迁移|github.com/pressly/goose|`go install github.com/pressly/goose/v3/cmd/goose@latest`||
|提供接口给前端|github.com/cloudzenith/DouTok/backend/gopkgs/tools|`backend/shortVideoApiService`目录下 `make init`|安装proto插件，统一处理`code`，`msg`字段|
|提供接口给前端|github.com/favadi/protoc-go-inject-tag@latest|`backend/shortVideoApiService`目录下 `make init`|基于proto生成golang代码时，对结构体的tag做特殊处理|
