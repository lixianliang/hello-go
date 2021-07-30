package main

import (
	"fmt"
	"time"
)

/*
	测试map data race问题
	通过go run -race 会有对应的数据竞争报警
	自己在algo-server中并行下载使用mutex，但upload没有使用mutext对mpa[string]error做保护
	uplod同事报错的出现的概率比较小
*/

func main() {
	kv := map[string]string{}

	if true {
	} else {
		for i := 0; i < 10; i++ {
			x := i
			go func() {
				kv[fmt.Sprintf("%d", x)] = fmt.Sprintf("%d", x)
			}()
		}
	}

	time.Sleep(3 * time.Second)
	fmt.Println("%v", kv)
}
