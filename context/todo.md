# context

## golang中Context的使用场景
    https://www.cnblogs.com/yjf512/p/10399190.html


## Go语言实战笔记（二十）| Go Context
    https://www.flysnow.org/2017/05/12/go-in-action-go-context.html

1. cancel取消函数
    CancelFunc调用它就可以取消指令，然后我们的监控goroutine就会收到信号，返回结束
    CancelFunc可以取消一个ctx以及这个节点下所有的ctx，不管有多少层级
    向下取消节点

1. WithTimeout WithDeadline
    表示超时自动取消；后面是多少时间自动取消

1. WittValue函数和取消context无关，它是为了生成一个绑定了一个键值对数据的context，这个绑定的数据可以通过context value方法来范文
    WittValue的方法是附加一个kv对，这里的key必须要是等价性的，具有可比性，value需要线程安全
    每个ctx只能存储一个kv对，value的key查找是从当前ctx节点向上查找

1. context使用原则
    不要把ctx放在结构体中，要以参数的方式传递
    以ctx作为参数的函数方法，应该把ctx作为第一个参数，放在第一位
    给一个函数方法传递ctx的时候，不要传递nil，如果不知道传递什么，使用context.TODO 
       context.BackGround可以作为root ctx
    ctx的value相关方法应该传递必须的数据，不要什么数据都使用这个传递
    ctx是线程安全的，可以放心的在多个goroutine中传递

## Golang Context 原理与实战
    https://segmentfault.com/a/1190000022534841

    ctx的val是链式获取

1. ctx应用范围
    业务需要对访问数据库、rpc、api做超时控制
    trace信息跟踪
    上下文信息的携带

1. withcancel
    形成一个新的上下文
    将ctx挂在到parent上下文中

``` go
type cancelCtx struct {
  Context

  // 互斥锁，保证context协程安全
  mu       sync.Mutex
  // cancel 的时候，close 这个chan
  done     chan struct{}
  // 派生的context
  children map[canceler]struct{}    通过map来存储子ctx相关信息
  err      error
}
```

1. 函数取消
    从上到下取消
    关联关系是从下到上删除

## Golang 之context库用法
    https://www.jianshu.com/p/b3ddeab7370f

    d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
    定义的时间一到自动会取消

    WithTimeout实际就是调用了WithDeadline
``` go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout))
}
```

## 深度解密Go语言之context
    https://zhuanlan.zhihu.com/p/68792989

1. 应用场景
    requestId  通过ctx共享

1. 主流程
``` go
ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
go Perform(ctx)

// ……
// app 端返回页面，调用cancel 函数
cancel()
``` 
    注意一个细节，WithTimeout函数返回的ctx和cancelFunc是分开的
    ctx本省没有取消函数，这样做的原因是取消函数智能由外层函数调用，防止子节点ctx调用取消函数，从而严格控制信息的流向：由父节点ctx流向子节点ctx


1. 避免goroutine泄露
``` go
unc main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // 避免其他地方忘记 cancel，且重复调用不影响

    for n := range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            cancel()
            break
        }
    }
    // ……
}
为何有主动调用cancel，还需要加一个defer cancel函数；避免其他地方忘记，且重复调用不影响

增加一个 context，在 break 前调用 cancel 函数，取消 goroutine。gen 函数在接收到取消信号后，直接退出，系统回收资源
```
