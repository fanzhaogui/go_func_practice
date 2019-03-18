package six

import "fmt"

// 闭包

//func(){
//	sum := 0
//	for i := 1; i <= 1e6; i ++ {
//		sum += i
//	}
//}()

func FuncClourse() (ret int) {
	defer func() {
		ret ++
		fmt.Println(ret)
	}()
	return 1
}