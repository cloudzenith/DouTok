# DouTok

## 前言
DouTok 全面拥抱 go-kratos，这是一个微服务框架，它的设计理念是简单、高效、可扩展。它的目标是提供一种简单、可靠的方法来构建微服务。

在开始之前，你需要了解一些 go-kratos 项目的基本概念，以便更好地理解和使用 DouTok。

https://go-kratos.dev/docs/

## 下载 Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## 项目结构
DouTok 的项目划分以 kraos 项目模板为基础。
```
.
├── LICENSE
├── README.md
├── api // proto 文件
│   ├── helloworld
│   │   └── v1
│   └── shortvideo
│       └── v1
├── go.mod
├── go.sum
├── internal
│   ├── app // 微服务
│   │   ├── helloworld
│   │   │   ├── Dockerfile 
│   │   │   ├── Makefile
│   │   │   ├── cmd // 启动文件
│   │   │   ├── configs // 配置文件
│   │   │   └── internal // 业务逻辑
│   │   └── shortvideo
│   │       ├── Dockerfile
│   │       ├── Makefile
│   │       ├── cmd
│   │       ├── configs
│   │       └── internal
│   └── pkg // 内部通用库
├── openapi.yaml
└── third_party // 第三方 proto 文件
    ├── README.md
    ├── errors
    │   └── errors.proto
    ├── google
    │   ├── api
    │   └── protobuf
    ├── openapi
    │   └── v3
    └── validate
        ├── README.md
        └── validate.proto
```

## 新增一个微服务
DouTok 是一个大仓项目，包含了多个微服务，你可以通过 kratos new internal/app/demo 命令来创建一个新 demo 微服务，微服务间共享 go.mod 文件，方便管理依赖。

当前项目已经包含了 helloworld 微服务，作为一个示例，你可以参考它来创建新的微服务。
```
1. kratos new app/demo --nomod（可以手动创建相关目录）
2. 在 api/demo 目录下定义 proto 文件
```

## 依赖注入 (wire)
使用 wire 生成依赖注入代码，确保你已经安装了 wire 工具，且已经了解了 wire 的基本使用方法。

https://github.com/google/wire
```
# install wire
go get github.com/google/wire/cmd/wire
```
