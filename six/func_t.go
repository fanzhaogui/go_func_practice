package six

// Go 语言拥有一些不需要进行导入操作就可以使用的函数
// close 用于管道通信
// len/cap 用于返回某个类型的长度或数量（字符串、数组、切片、map和管道）；cap是容量的意思，用于返回某个类型的最大容量（只能用于切片和map）

// new make 均是用于分配内存；
// new用于值类型和用户定义的类型，如自定义结构，
// make用于内置引用类型（切片、map和管道）
// 他们的用法就像是函数，但是将类型作为参数 new(type)、make(type),
// new(T)分配类型T的零值并返回其地址，也就是指向类型T的指针，它也可以被用于基于类型 v := new(int)
// make(T)返回类型T的初始化之后的值，因此它比new进行更多的工作
// new() 是一个函数，不要忘记它的括号

// copy、append
// 用于复制和链接切片

// panic/recover 两者均用于错误处理机制

// print println 底层打印函数，在部署环境中建议使用fmt包

// complex/real imag 用于创建和操作复数

