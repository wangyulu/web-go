# web-go

#### 一些问题

##### 中间件中执行 Next 方法时报错没有处理，这样会不会导致一些错误没有被捕获到而丢失。仅ServeHTTP中对Next方法执行时捕获了错误

##### 服务提供者在注册时，里面所需要的参数是如何传递进去的
在初始绑定服务的时候显示传递进去的

```go
container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
```

##### crontab 任务没有存储，每次都是在应用启动的时候注册，这样的话每次添加新的任务都需要升级应用

##### 如果在一个应用下启动多个分布式任务的话，最后启动的进程ID会覆盖掉之前启动任务的进程ID，这样之前启动的任务就因为找不到进程ID而无法通过命令行管理
可参考的方案是：每次启动时为应用指定runtime，通过环境变量传递

#### 一些得到

##### 强制让编译器知道某个类型实现了某个接口类型

```go
// 方式一
var _ Container = (*HadeContainer)(nil)

// 方式二
var _ Container = new(HadeContainer)

```

##### 让某个类型强制实现某个接口的方法
```go
interface Service {
    Get() err
}

type Demo struct {
    // 这是要强制实现Service接口
    Service
}

func (d *Demo) Get() error 
{
    return nil;
}

```
