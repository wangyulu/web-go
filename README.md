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

##### 在配置文件章节，启动应用进程时会遇到死锁现象
具体步骤：

在对Env服务进行绑定的时候，由于是修改操作故这里使用了读写锁的写锁（记为M）。

在执行Boot初始化时，由于依赖于App服务，故执行了Make查找App服务

在Make的过程中使用了读写锁中的读锁（记录N）

最终导致M与N的读写冲突

```go

// 没有添加死锁检测时，没有任务输出
/*
$ ./web-go

*/

// 添加死锁检测后的报销
/* 
$ ./web-go
POTENTIAL DEADLOCK: Recursive locking:
current goroutine 1 lock 0xc00048e220
framework/container.go:140 framework.(*HadeContainer).make { hade.lock.RLock() } <<<<<
framework/container.go:139 framework.(*HadeContainer).make { func (hade *HadeContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) { }
framework/container.go:112 framework.(*HadeContainer).MustMake { if err != nil { }
framework/provider/env/provider.go:18 env.(*HadeEnvProvider).Boot { func (provider *HadeEnvProvider) Boot(c framework.Container) error { }
framework/container.go:78 framework.(*HadeContainer).Bind { if err := provider.Boot(hade); err != nil { }
main.go:23 main.main { container.Bind(&env.HadeEnvProvider{}) }
/usr/local/Cellar/go@1.16/1.16.10/libexec/src/runtime/proc.go:233 runtime.main { // Once it does, it will exit. See issues 3934 and 20018. }

Previous place where the lock was grabbed (same goroutine)
framework/container.go:70 framework.(*HadeContainer).Bind { hade.lock.Lock() } <<<<<
framework/container.go:69 framework.(*HadeContainer).Bind { func (hade *HadeContainer) Bind(provider ServiceProvider) error { }
main.go:23 main.main { container.Bind(&env.HadeEnvProvider{}) }
/usr/local/Cellar/go@1.16/1.16.10/libexec/src/runtime/proc.go:233 runtime.main { // Once it does, it will exit. See issues 3934 and 20018. }
*/
```

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

##### Go中同一协程下，读写锁中的读锁也是不可重入的，也是会造成死锁的
