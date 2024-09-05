---
sidebar_position: 1
---

# IM Service 设计思路

IM Service 用于提供即时通讯服务，主要包括消息发送、消息接收、消息撤回、消息删除等功能。IM Service 通过消息队列进行消息的异步处理，通过消息队列的方式，可以保证消息的可靠性，同时也可以保证消息的顺序性。

## 基本概念

1. 用户：在一组即时通信中，参与通信的实体，对应BaseService中Account的概念
2. 聊天室：对应“一组即时通信”，一个聊天室可以有1~n个用户，具体有多少用户可以由上层服务自行定义
3. 消息：用户之间传递的消息，包括文本消息、图片消息、语音消息、视频消息等类型。文本消息下，消息内容为相应文本；其他情况下为对应的文件id

## 主要接口

### 聊天室服务

- 创建聊天室 Create
- 获取聊天室信息 Get
- 删除聊天室 Remove
- 加入聊天室 AddAccount2Room
- 离开聊天室 RmAccountFromRoom

### 聊天记录服务

- 推送聊天记录接口 Push
- 拉取聊天记录 Pull
- 获取长连接地址用于拉取聊天记录 PullByActiveConnection

### 特殊接口：通过长连接进行消息拉取

优点：

- 低延迟
- 降低连接的数量
- 可以随时增加或减少IM Service服务的数量

## 主要链路

### 发起单聊

```mermaid
sequenceDiagram
    participant fe as 前端
    participant api as 上游业务网关
    participant im as IM Service
    
    fe ->> api: 请求发起单聊，携带双方user id
    api ->> im: 创建聊天室
    im -->> api: 返回
    api ->> im: 添加双方user id到聊天室
    im -->> api: 返回
    api -->> fe: 返回
```

### 发起群聊

```mermaid
sequenceDiagram
    participant fe as 前端
    participant api as 上游业务网关
    participant im as IM Service
    
    fe ->> api: 请求发起群聊
    api ->> im: 创建聊天室
    im -->> api: 返回
    api -->> fe: 返回

    fe ->> api: 增加用户请求
    api ->> im: 增加用户到聊天室
    im -->> api: 返回
    api -->> fe: 返回
```

### 开始聊天

```mermaid
sequenceDiagram
    participant fe as 前端
    participant api as 上游业务网关
    participant im as IM Service
    participant core as 业务核心服务
    participant mq as 消息队列
    participant bs as BaseService
    
    fe ->> api: 请求聊天室
    api ->> im: 获取/创建聊天室
    im -->> api: 返回
    api ->> fe: 返回

    fe ->> api: 获取聊天参与人相关信息
    api ->> im: 获取聊天室所有参与人
    im ->> api: 返回
    api ->> core: 获取聊天参与人用户信息
    core -->> api: 返回
    api -->> fe: 返回

    fe ->> api: 请求建立长连接消息拉取
    api ->> im: 获取需要代理的长连接请求（负载均衡）
    im -->> api: 返回
    api -->> fe: 返回长连接地址

    fe ->> api: 拉取历史消息记录
    api ->> im: 拉取消息
    im -->> api: 返回消息
    api -->> fe: 返回
    fe ->> fe: 结合之前获取的“聊天参与人相关信息”，将历史消息记录进行展示

    fe ->> api: 发送消息
    api ->> bs: 上传文件（针对非文本消息）
    bs -->> api: 文件id
    api ->> mq: 通过MQ发送消息体
    mq -->> api: 发送成功
    api -->> fe: 发送成功

    fe ->> api: 访问对应长连接地址
    api ->> im: 将长连接代理给IM Service
    im ->> mq: 消费，获取被发送的消息
    mq -->> im: 返回消息
    im ->> im: 异步消息持久化
    im ->> api: 代理成功，持续推送消息
    api ->> fe: 连接成功，持续推送消息
```
