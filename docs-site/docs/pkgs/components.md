# components

components库针对MySQL，Redis等组件封装了相关初始化的方法，可以更便捷的对这些组件进行初始化和获取，配合`launcher`包，可以进一步降低有关加载、初始化组件的代码重复度。

## 基本原理

components包下的`components.go`中封装了一个“公共父类”，其主要包含3个方法：

- Load: 加载组件配置。结合泛型，定义初始化各个组件的方法入口并在方法中加载配置。
- Start：启动组件，并进行健康检查
- GetConfig：获取组件配置

在components包中，对诸多组件进行了封装，如下以`redisx`为例。

`redisx`中，首先有一个`config.go`文件，这个文件定义了标准的配置文件结构格式，例如Redis地址、密码等。

在`redis.go`中，定义了一个全局变量map，用于存储不同的Redis实例。除此之外，按照`components.Load`中定义的初始化方法的入、出参，定义Redis初始化方法和获取Redis实例的`GetClient`方法。

对于`GetClient`方法，通常需要传入对应的Redis实例的key，key的定义以配置文件中为准，例如如下配置：

```yaml
redis:
  default: # Redis配置1
    dsn: localhost:6379
    password: test
  lock: # Redis配置2
    dsn: localhost:6378
    password: test
```

如上，使用`default`和`lock`作为key，传入`GetClient`方法中，即可获取对应的Redis实例。
