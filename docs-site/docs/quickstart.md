---
sidebar_position: 2
---

# 快速启动DouTok

## 准备环境

本教程将带领你从零开始，循序渐进搭建并启动 `DouTok` 项目，若读者已具备相关知识，可选择性阅读。

## 环境准备

1. Golang 1.22+
   > - https://golang.org/dl/
   > - https://golang.google.cn/dl/
2. Node 14.17+
   > - https://nodejs.org/en/download/
3. JetBrains GoLand/WebStorem
   > - https://www.jetbrains.com/
4. VSCode
   > - https://code.visualstudio.com/
5. Docker
   > - https://www.docker.com/products/docker-desktop

## 必要组件配置及启动

- Consul: 通过`backend/gopkgs/launcher`提供能力，所有后端服务均自动注册到Consul中
- Redis: 缓存
- MySQL: 持久化存储
- MinIO: 对象存储
- RocketMQ: 消息队列

1. 找到`env/basic.yml`文件，通过命令`docker-compose -f ./env/basic.yml up -d`启动Consul, Redis, MySQL, MinIO
2. 找到`env/rocketmq/broker.conf`文件，将`brokerIP1`修改为本地局域网IP
3. 找到`env/rocketmq.yml`文件，通过命令`docker-compose -f ./env/rocketmq.yml up -d`启动RocketMQ

## MySQL库表结构同步

1. 进入`sql`目录
2. 检查`sql/Makefile`文件，其中涉及的MySQL连接需注意应与本地环境一致
3. 执行`make up`命令，MySQL库表结构会同步到本地

## 启动后端服务

1. 进入`backend`目录下除`gopkgs`外的所有服务目录，执行`make build`以编译Docker镜像
2. 进入`env`目录，检查`configs`下各个配置文件，应与本地环境保持一致
3. 进入`env`目录，执行`docker-compose -f backend.yml up -d`启动所有后端服务

## 启动前端服务

1. 进入`frontend/doutok`目录，执行`yarn install`安装依赖
2. 执行`yarn start`启动前端服务
