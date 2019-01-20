package four

import (
	"fmt"
	"strings"
)

func TestFunc()  {
	// 判断字符串s是否以suffix结尾
	var str = "abcdefghijkl"
	fmt.Println(strings.HasPrefix(str,"abc"))
	fmt.Println(strings.HasPrefix(str, "bb"))

	// 字符串包含关系
	fmt.Println(strings.Contains(str, "bcd"))
	fmt.Println(strings.Contains(str, "gggg"))

	// 如需要查询非ASCII编码的支付在父字符串中的位置
	// strings.Index
	// strings.LastIndex
	// strings.IndexRune(s string, r rune) int
	var s1 = "Hi, I'm Marc, Hi. Hi";
	fmt.Println(strings.Index(s1, "Marc"))
	fmt.Println(strings.Index(s1, "Hi")) // 匹配的值从0开始
	fmt.Println(strings.LastIndex(s1,"Hi"))
	fmt.Println(strings.Index(s1, "Lucy")) // 如不存在，则为-1

	// 字符串替换  Replace
	fmt.Println(strings.Replace(s1, "Hi", "NO", 4)) // 改变n查看结果

	// Count 用于统计字符串str在字符串s中出现的非重叠次数
	fmt.Println(strings.Count(s1, "H"))

	// 重复字符串
	var origS string = "Hi there! "
	var newS string
	newS = strings.Repeat(origS, 4)
	fmt.Println(newS)

	// 修改字符串大小写
	// ToLower ToUpper
	origS = "Hey, How are you Lisa?";
	var lower, upper string
	lower = strings.ToLower(origS)
	upper = strings.ToUpper(origS)
	fmt.Println(lower)
	fmt.Println(upper)

	// 修减字符串
	// TrimSpace(s) 剔除字符串开头和结尾的空白符号
	// Trim(s, "cut") // 将开头和结尾的cut去除掉
	// TrimLeft  TrimRight
	origS = " ,and,you, "
	fmt.Println(strings.TrimSpace(origS))
	origS = ",and,you,"
	fmt.Println(strings.Trim(origS, ","))
	fmt.Println(strings.TrimLeft(origS, ","))
	fmt.Println(strings.TrimRight(origS, ","))

	// 分割字符串
	// Fields(s) 将会利用1个或多个空白符号来作为动态长度的分隔符将字符串分隔成若干小块，并返回一个slice，如果字符串只包含空白字符串，则返回长度为0的slice
	origS = "          ";
	fmt.Println(strings.Fields(origS))
	origS = "a bdcdfbfgfgdffgbgb";
	fmt.Println(strings.Fields(origS))
	var fieldSlices = strings.Fields(origS)
	for _,v := range fieldSlices {
		fmt.Println(v)
	}
	// Splite(s, sep) 用于自定义分隔符号来对指定字符串进行分隔，同样返回slice
	fmt.Println(strings.Split(origS, "b"))
	fieldSlices = strings.Split(origS, "b")
	for k,v := range fieldSlices {
		fmt.Println(k,v)
	}

	// 拼接字符串
	// Join 用于将元素类型为string的slice使用分隔符号拼接组成一个字符串
	fmt.Println(strings.Join(fieldSlices, "b"))

	// 从字符串中读取内容
	// strings.NewReader(str) 用于生成一个Reader并读取字符串中的内容，然后返回指向该Reader的指针
	// Read() 从[]byte中读取内容
	// ReadByte() 和 ReadRune() 从字符串中读取下一个byte或者true
	fmt.Println(strings.NewReader(origS))
}
