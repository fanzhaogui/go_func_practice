package seven

import "fmt"

// 创建slice
func Test() {
	fmt.Println("slice testt ...")
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} // 数组
	fmt.Println(a)
	s1 := a[5:10] // 包含起始索引，不包含终止索引
	fmt.Println(s1)
	s2 := a[5:len(a)]
	fmt.Println(s2)
	s3 := a[5:]
	fmt.Println(s3)
	s4 := a[:5]
	fmt.Println(s4)
}

func Test2() {
	// 初始化 slice 类型 长度 容量
	s1 := make([]int, 3)
	fmt.Println(s1)
	s2 := make([]int, 3, 10)
	fmt.Println(s2, len(s2), cap(s2))
}

// reslcie
func Test3() {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	s1[0] = 9
	fmt.Println(s1, s2)
}

// append
func Test4() {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	s2 = append(s2, 1, 1, 1, 1, 1, 1, 1, 1, 1) // 超过cap(a) 重新分配内存地址
	s1[0] = 9
	fmt.Println(s1, s2)
}

// copy
func Test5() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8}
	copy(s1, s2)
	fmt.Println(s1)
}

func Test6() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8}
	copy(s2, s1)
	fmt.Println(s2)
}

func Test7() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	fmt.Println(s2)
}
