# web-go

#### 一些问题

##### 中间件中执行 Next 方法时报错没有处理，这样会不会导致一些错误没有被捕获到而丢失。仅ServeHTTP中对Next方法执行时捕获了错误

#### 一些得到

##### 强制让编译器知道某个类型实现了某个接口类型

```$golang
// 方式一
var _ Container = (*HadeContainer)(nil)

// 方式二
var _ Container = new(HadeContainer)

```