# launcher

`launcher`是一个服务启动器，用于加载配置、初始化组件、启动服务。

## 定义一个Launcher

定义启动器需要调用`launcher.New`方法，入参是用于创建加载器的诸多配置项。

### `launcher.WithBeforeConfigInitHandler`

用于定义在加载配置之前的操作。

### `launcher.WithBeforeServerStartHandler`

用于定义在启动服务之前的操作。

### `launcher.WithAfterServerStartHandler`

用于定义在启动服务之后的操作。

### `launcher.WithShutdownHandler`

用于定义在关闭服务时的操作。

### `launcher.WithConfigOptions`

用于传递Kratos Config所需的一些配置项。

### `launcher.WithConfigWatcher`

用于定义一些动态配置行为。

### `launcher.WithHttpServer`

用于定义应用所附带的HTTP服务。

### `launcher.WithGrpcServer`

用于定义应用所附带的gRPC服务。

### `launcher.WithKratosOptions`

用于传递Kratos的一些配置项。

## 启动服务

启动服务需要调用`launcher.Run`方法，将按照配置的流程进行启动。

## 示例代码

```go
func main() {
	launcher.New(
		launcher.WithConfigOptions(
			config.WithSource(
				file.NewSource("configs/"),
			),
			config.WithDecoder(func(keyValue *config.KeyValue, m map[string]interface{}) error {
				return yaml.Unmarshal(keyValue.Value, m)
			}),
		),
		launcher.WithHttpServer(initHttpServer()),
		launcher.WithGrpcServer(initGrpcServer()),
		launcher.WithAfterServerStartHandler(func() {
			redisClient := redisx.GetClient(context.Background())
			if err := redisClient.Ping(context.Background()).Err(); err != nil {
				panic(err)
			} else {
				log.Info("redis connected")
			}
		}),
	).Run()
}
```
