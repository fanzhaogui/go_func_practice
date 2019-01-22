package five

import (
	"fmt"
	"runtime"
	"strconv"
)


//当 if 结构内有 break、continue、goto 或者 return 语句
func IfTest()  {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	// 注意事项： 不要同时在 if-else结果的两个分支里都使用return语句
	/*if bool1 {
		return 2
	} else {
		return 1
	}*/

	// demo

	// 1. 判断一个字符串是否为空
	str := ""
	if str == "" {

	}
	if len(str) == 0 {

	}

	// 2. 判断运行go程序的操作系统类型，这可以通过常量runtime.GOOS来判断
	var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)

	// 3 函数Abs() 返回一个整形数字的绝对值
	var i1 = 6
	fmt.Println(Abs(i1))
	i1 = -2
	fmt.Println(Abs(i1))
}

func Abs(x int) int {
	if x < 0 {
		return -x;
	}
	return x;
}

// 测试多返回值函数的错误
func ErrorsTest() {
	var orig string = "ABC"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}