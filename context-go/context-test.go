package context_go

import (
	"context"
	"fmt"
	"time"
	"unsafe"
)

func Test1() {
	ctx := context.Background()
	ctx.Done()
	process(ctx)
	ctx = context.WithValue(ctx, "traceId", "qcrao-2019")
	process(ctx)
}

func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceId)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}

func Gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			fmt.Println("kazhule")
			select {
			case <-ctx.Done():
				fmt.Println("meikazhu, zouren")
				return
			case ch <- n:
				n++
				fmt.Println("n: ", n)
				time.Sleep(time.Second)
			}
			fmt.Println("meikazhu")
		}
	}()
	return ch
}

func Test2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 避免其他地方忘记 cancel，且重复调用不影响

	for n := range Gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			fmt.Println("wait 5")
			fmt.Println("continue")
			cancel()
			break
		}
		time.Sleep(5 * time.Second) //用来测验Gen函数会不会卡在select中，事实证明不会，因为select会阻塞在管道那部分，
		// 但是Test2也会在等待5秒后读取管道，这样就不卡了
	}
	time.Sleep(time.Hour)
}

type animal interface {
	eat()
	sleep()
}

type milk struct {
	animal
	v1 int
	v2 int
	v3 int
}

func Test3() {
	var m milk = milk{v1: 1, v2: 2, v3: 3}
	var a animal = m
	fmt.Println("size: ", unsafe.Sizeof(m))
	fmt.Println("value: ", m)
	fmt.Println("size: ", unsafe.Sizeof(a))
	fmt.Println("value: ", a)
}
