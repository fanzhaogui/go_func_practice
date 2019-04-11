package main

import (
	"fmt"
	"sync"
)

func main() {
	// 第1次测试
	// fmt.Println("before go Go")
	// go Go()
	// fmt.Println("after go Go")
	// time.Sleep(3 * time.Second) // 延迟3秒钟，是为了查看Go()函数被调用后的结果

	// 相关概念
	// Channel
	// chan 是引用类型
	// 通过make创建，close关闭
	// Channel是goroutine沟通的桥梁，大都是阻塞同步的
	// 可以使用for range来迭代不断操作channel
	// 可以设置单项或双向通道
	// 可以设置缓存大小，在未被填满前不会发生阻塞

	// Select
	// 可处理一个或多个channel的发送与接收
	// 同时有多个可用的channel时按随机顺序处理
	// 可用空的select来阻塞main函数
	// 可设置超时

	// 第2次测试
	/*c := make(chan bool)
	go func() {
		fmt.Println("gogogo....")
		c <- true // 存入channel的操作
	}()
	<- c */ // 阻塞，等待从channel取出

	// 第3次测试 利用for range 等待 close
	/*c := make(chan bool)
	go func() {
		fmt.Println("gogogo....")
		c <- true // 存入channel的操作
		close(c) // 关闭
	}()
	for v := range c {
		fmt.Println(v)
	}*/

	// 第4次测试 channel无缓存，和有缓存
	/*c := make(chan bool)
	go func() {
		fmt.Println("gogogo....")
		<- c
	}()
	c <- true */ // 无缓存时，同步阻塞，这里无输出

	/*c := make(chan bool, 1)
	go func() {
		fmt.Println("gogogo....")
		<- c
	}()
	c <- true */ // 有缓存时，异步无阻塞，这里无输出


	// 第5测试
	// runtime.GOMAXPROCS(runtime.NumCPU()) // 多核运行
	/*c := make(chan bool)
	for i:=0; i < 10 ; i++ {
		go Gorun(c, i)
	}
	<- c*/
	// 在启用多核运行时，可能会优先执行index=9，结束
	// 解决 方式一
	/*c := make(chan bool, 10) // 设置10个缓存
	for i := 0; i < 10; i++ {
		go Gorun(c, i)
	}
	for i := 0; i < 10; i++ {
		<- c
	}*/

	// 解决方式2 同步包
	/*wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go GorunWg(&wg, i)
	}
	wg.Wait()*/

	// 第6测试 多个channel的解决方案
	c1, c2 := make(chan int), make(chan string)
	// o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					break
				}
				fmt.Println("c1: ", v)
			case v, ok := <-c2:
				if !ok {
					break
				}
				fmt.Println("c2: ", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hello"
	c1 <- 3
	c2 <- "world"
	close(c1)
	close(c2)
}

// func Go() {
// 	fmt.Println("go go go...")
// }

func Gorun(c chan bool, index int)  {
	a := 1
	for i := 0; i < 100000000 ; i++ {
		a += i
	}
	fmt.Println(index, a)
	// if index == 9 {
	// 	c <- true
	// }
	// time.Sleep(3*time.Second)
	c <- true
}

func GorunWg(wg *sync.WaitGroup, index int)  {
	a := 1
	for i := 0; i < 100000000 ; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}