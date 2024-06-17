---
sidebar_position: 2
---

# 快速启动DouTok

## 准备环境

本教程将带领你从零开始，循序渐进搭建并启动 `DouTok` 项目 ，若读者已具备相关知识，可选择性阅读。



## 后端环境

### 基础需求

1. 下载golang安装
   - 国际: https://golang.org/dl/
   - 国内: https://golang.google.cn/dl/
2. 命令行运行 go 若控制台输出各类提示命令 则安装成功
3. 开发工具推荐 [Goland](https://www.jetbrains.com/go/)
4. 使用 `Docker` 或自行安装所需中间件



### 必要服务及中间件

- Redis
- MySQL
- Etcd
- Zookeeper
- HBase
- Kafka
- MinIO
- Kafka-ui(可选)

找到 `./env/dependencies.yml` 文件，我们将使用 `docker-compose` 来进行安装

```yaml
docker-compose -f ./env/dependencies.yml up -d
```

确认所需各项服务成功安装

你也可以自行安装对应服务，但需在后续修改你的对应配置文件



1. 确保你的主机中的 `Hosts` 文件下存在这样的映射关系

   ```
   localhost(你的HBase部署的IP) hb-master
   ```

   否则会无法找到HBase的主节点信息

2.  进入你的 `HBase` 的 `Docker` 容器内内部，执行下面的命令

   ```
   $ hbase shell # 使用hbase命令行工具
   $ create 'publish','data' # 创建表 publish , 列族为 data
   $ create 'feed','data' # 创建表 feed , 列族为 data
   ```

   这将为你配置好 `HBase` 的初始表格



### 拉取代码

1. 运行

```
git clone https://github.com/TremblingV5/DouTok.git
```

将代码拉取到你的本地

2. 准备数据库表结构

   `sql` 文件在 `./scripts/DouTok.sql` 文件中，将它执行到你的数据库中

3. 修改配置文件

   所有的配置文件都在 `./config` 目录下的 `.yaml` 文件

   将 `MySQL` 、`Redis` 、`HBase` 、`Kafka` 、`MinIO` 、`Etcd` 的相关配置修改为你的配置

至此准备工作完成



### 启动服务

所有的服务都在 `./applications` 文件夹下面

#### API网关

`./applications/api/main.go` 

#### RPC服务

##### 逻辑层

- Entity

`./applications/comment/main.go`

- Feed

`./applications/feed/main.go`

- Message

`./applications/message/main.go`

- Publish

`./applications/publish/main.go`

- Relation

`./applications/relation/main.go`

- User

`./applications/user/main.go`

##### 业务层

- CommentDomain

`./applications/commentDomain/cmd/server/main.go`

- FavoriteDomain

`./applications/favoriteDomain/main.go`

- MessageDomain

`./applications/messageDomain/main.go`

- RelationDomain

`./applications/relationDomain/main.go`

- UserDomain

`./applications/userDomain/main.go`

- VideoDomain

`./applications/videoDomain/main.go`



### 可能遇到的问题

1. 没有正确配置好 `yaml` 文件，无法连接对应服务或中间件崩溃。

   请仔细检查自己的环境配置，包括 `IP` 、端口、密码等

2. 报错没有 `./tmp/DouTok.log` 。

   打印日志的文件，自己添加该文件即可



### Swagger接口文档

在 `API` 网关服务启动后，访问 [http://localhost:8088/swagger/index.html](http://localhost:8088/swagger/index.html) 即可
可以看到目前有的接口相关信息



## 前端环境

We can download `.apk` file from `./ui` of this repo. Now the client is only support Android. After downloading and installing of this app. We can open it firstly. Then we can click `我` on the right bottom to enter configure page. After opening `高级配置`, we can input base url of the backend. An example is `http://localhost:8080/`.









