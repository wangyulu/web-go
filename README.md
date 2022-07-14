# web-go

#### 一些问题

##### 中间件中执行 Next 方法时报错没有处理，这样会不会导致一些错误没有被捕获到而丢失。仅ServeHTTP中对Next方法执行时捕获了错误

##### 服务提供者在注册时，里面所需要的参数是如何传递进去的

#### 一些得到

##### 强制让编译器知道某个类型实现了某个接口类型

```$golang
// 方式一
var _ Container = (*HadeContainer)(nil)

// 方式二
var _ Container = new(HadeContainer)

```

##### 让某个类型强制实现某个接口的方法
```$golang
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
