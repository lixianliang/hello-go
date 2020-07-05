# go的内存模型

## 
1. https://www.jianshu.com/p/5e44168f47a3

go org英文文章
1. https://golang.org/ref/mem


## select{} 和 for {} 区别
for{}是忙等待，cpu会100%，一直在执行循环
select{} uses nearly 0%，等待goroutine 的read write close，没有case对应的事件   
``` shell
func main() {
for i := 0; i < 10; i++ {
c := consumer.New()
go c.Start()
}
}

select {} // 阻塞
```

## 限制协程池运行个数的方式
1. 为什么要控制goroutine的数量？     goroutine固然好，但是数量太多了，往往会带来很多麻烦，比如耗尽系统资源导致程序崩溃，或者CPU使用率过高导致系统忙不过来
2. 如果在G里面做远程调用，可能会导致发起很多请求到后端服务

两种方式都可以，方法一可以省略各goroutine个数的同步
方法二更加通用，通过取莫来进行goroutine限制，多次调用wg.Wait()不会有问题，当wg.Wait()判断里面的计数为0则进行下一次
通过buffer channel来进行goroutine个数通过

后面想了下方法一是不安全的的，有可能运行到11，然后一直在wg.Wait()阶段，程序挂起在这里了


### 方法一
``` go

threadPoolNumber := 10
var wg sync.WaitGroup{}

for i := 0; i < 100; i++ {
    wg.Add(1)
    go func(int j) {
        defer wg.Done()
        if j % threadPoolNumber == 0 {
            wg.Wait()
        }
        fmt.Println("abc: ", j)
    }(i)
}

wg.Wait()

```

### 方法二  
``` go
poolLimit := make(chan int, 10)
var wg sync.WaitGroup{}

for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        poolLimit <- 0
        // do something
        fmt.Println("abc: ")
        <-poolLimit
    }()
}

wg.Wait()   // 等待所有G退出


优雅的并使用并空goroutine的数量

```

## Go Modules 终极入门
1.  https://eddycjy.com/posts/go/go-moduels/2020-02-28-go-modules/
