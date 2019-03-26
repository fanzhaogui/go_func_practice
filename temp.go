package main

import (
	"fmt"
	"runtime"
)

func main()  {
	// go Go()
	// time.Sleep(3 * time.Second)
	runtime.GOMAXPROCS(runtime.NumCPU())
	// c := make(chan bool)
	// go func() {
	// 	fmt.Println("gogogo....")
	// 	c <- true // 存入channel的操作
	//
	// 	close(c)
	// }()
	// <- c // 从channel取出

	// for v := range c {
	// 	fmt.Println(v)
	// }

	// for i:=0; i < 10 ; i++ {
	// 	go Gorun(c, i)
	// }
	// <- c

	c := make(chan bool, 10) //
	for i := 0; i < 10; i++ {
		go Gorun(c, i)
	}

	for i := 0; i < 10; i++ {
		<- c
	}

}

// func Go()  {
// 	fmt.Println("go go go...")
// }

func Gorun(c chan bool, index int)  {
	a := 1
	for i := 0; i < 100000000 ; i++ {
		a += i
	}
	fmt.Println(index, a)
	// if index == 9 {
		c <- true
	// }
}
