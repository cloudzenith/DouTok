# gofer

`gofer`直译为“办事员”，谐音`gopher`。`gofer`是一个对运行goroutine的封装。

## 基本方法

### `gofer.Go`

像运行goroutine一样运行一个函数

```go
func Go(f func())
```

### `gofer.GoWithCtx`

在`gofer.Go`的基础上，传入一个`context.Context`对象，使得可以运行一个依赖于context的函数

```go
func GoWithCtx(ctx context.Context, f func(context.Context))
```

### `gofer.GoWithTimeout`

在gofer.Go`的基础上，传入一个超时时间，使得可以运行一个有超时时间的函数

```go
func GoWithTimeout(f func(), d time.Duration) (isFinish bool)
```

:::tip
使用默认线程池执行以上函数:

通过`gofer.SetUseGlobalPool`可以在项目初始时开启使用默认线程池，在此基础上，使用上述`Go`, `GoWithCtx`和`GoWithTimeout`函数时，会使用默认线程池执行。

```go
func SetUseGlobalPool(value bool)
```

:::

## 任务组

在开发过程中，有时可以通过并行的方式执行多个任务，从而提升整体效率，`gofer`提供了`Group`来支持这种需求。

### `gofer.NewGroup`

用于初始化一个任务组

```go
func NewGroup(ctx context.Context, options ...GroupOption) *Group
```

### `gofer.Group.Run`

用于向任务组中添加一个任务

```go
func (g *Group) Run(f func() error) error
```

### `gofer.Group.Wait`

用于等待所有任务执行完毕

```go
func (g *Group) Wait() error
```
