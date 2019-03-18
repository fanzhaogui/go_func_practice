package seven

import "fmt"

func Array1()  {
	var a [2]int
	var b [1]int
	fmt.Println("array ....")
	fmt.Println(a,b)
	a = [2]int{1,2} // 1 2
	fmt.Println(a)
	a = [2]int{3}// 3 0
	fmt.Println(a)

	// 指定下标赋值
	var c = [20]int{19:1}
	fmt.Println(c)

	// 不定长:同时可以指定下标赋值
	var d = [...]int{1,23,4,5,6,7,0}
	fmt.Println(d)

	//  多维数组   a := [2][3]int { {}, {}}  包含两个长度为3的一维数组


}